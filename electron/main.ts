import * as path from 'path';
import * as fs from 'fs';
import Logger from "./logger";
import BilladmDao from "./billadmDao";

const {
    app,
    BrowserWindow,
    ipcMain,
    screen,
    dialog,
} = require('electron');

if (require('electron-squirrel-startup')) {
    app.quit();
}

const confDir = path.join(app.getPath("home"), ".config", "billadm");
const windowStatePath = path.join(confDir, "windowState.json");
const workspaceStatePath = path.join(confDir, "workspace.json")
const logger = new Logger(path.join(confDir, 'app.log'));
let firstOpen = false;

try {
    firstOpen = !fs.existsSync(workspaceStatePath);
    if (!fs.existsSync(confDir)) {
        fs.mkdirSync(confDir, {mode: 0o755, recursive: true});
    }
} catch (e) {
    console.error(e);
    dialog.showErrorBox("创建配置目录失败", "需要在用户家目录下创建配置文件夹（~/.config/billadm），请确保该路径具有写入权限");
    app.exit();
}

logger.info(`start to launch billadm, firstOpen: ${firstOpen}`);


interface WorkspaceState {
    last: string,
    workspaces: string[],
}

interface Workspace {
    billadmDao: BilladmDao;
    workspaceDir: string;
}

let currentWindow: any;
let workspace: Workspace;

const initWorkspace = () => {
    let oldWorkspace: WorkspaceState = JSON.parse(fs.readFileSync(workspaceStatePath, "utf8"));
    if (!oldWorkspace.last) {
        oldWorkspace.last = oldWorkspace.workspaces[0];
    }
    if (fs.existsSync(oldWorkspace.last) && fs.statSync(oldWorkspace.last).isFile()) {
        dialog.showErrorBox("工作目录错误", `工作目录【${oldWorkspace.last}】已存在同名文件，请检查`);
        app.exit();
    }
    logger.info(`当前工作目录[${oldWorkspace.last}]`)
    let needInit = false;
    try {
        if (!fs.existsSync(oldWorkspace.last)) {
            fs.mkdirSync(oldWorkspace.last, {mode: 0o755, recursive: true});
            needInit = true;
        }
    } catch (e) {
        console.error(e);
        dialog.showErrorBox("创建工作目录失败", `工作目录【${oldWorkspace.last}】创建失败，请确保该路径具有写入权限`);
        app.exit();
    }
    if (needInit) {
        fs.mkdirSync(path.join(oldWorkspace.last, 'data'), {mode: 0o755, recursive: true});
    }
    let dbFile = path.join(oldWorkspace.last, 'data', 'billadm.db');
    let logFile = path.join(oldWorkspace.last, 'billadm.db.log');
    workspace = {
        billadmDao: new BilladmDao(dbFile, logFile),
        workspaceDir: oldWorkspace.last,
    };
    workspace.billadmDao.init().then(
        () => logger.info(`init billadmDao for ${workspace.workspaceDir} success`)
    ).catch(
        (err: Error) => {
            logger.info(`init billadmDao for ${workspace.workspaceDir} failed`);
            dialog.showErrorBox("初始化billadmDao失败", `工作目录【${workspace.workspaceDir}】错误信息【${err.message}】`);
            app.exit();
        }
    )
    if (needInit) workspace.billadmDao.initDB().then(
        () => logger.info(`init db for ${workspace.workspaceDir} success`)
    ).catch(
        (err: Error) => {
            logger.info(`init db for ${workspace.workspaceDir} failed`);
            dialog.showErrorBox("初始化数据库失败", `工作目录【${workspace.workspaceDir}】错误信息【${err.message}】`);
            app.exit();
        }
    );
};

const createWindow = () => {
    // 恢复主窗体状态
    let oldWindowState = {};
    try {
        oldWindowState = JSON.parse(fs.readFileSync(windowStatePath, "utf8"));
    } catch (e) {
        fs.writeFileSync(windowStatePath, "{}");
    }
    let defaultWidth;
    let defaultHeight;
    let workArea;
    try {
        defaultWidth = Math.floor(screen.getPrimaryDisplay().size.width * 0.6);
        defaultHeight = Math.floor(screen.getPrimaryDisplay().workAreaSize.height * 0.8);
        workArea = screen.getPrimaryDisplay().workArea;
    } catch (e) {
        console.error(e);
    }
    const windowState = Object.assign({}, {
        isMaximized: false,
        fullscreen: false,
        isDevToolsOpened: false,
        x: 0,
        y: 0,
        width: defaultWidth,
        height: defaultHeight,
    }, oldWindowState);

    logger.info("windowStat [x=" + windowState.x + ", y=" + windowState.y + ", width=" + windowState.width + ", height=" + windowState.height + "], default [width=" + defaultWidth + ", height=" + defaultHeight + "], workArea [width=" + workArea.width + ", height=" + workArea.height + "]");

    let x = windowState.x;
    let y = windowState.y;

    const mainWindow = new BrowserWindow({
        width: windowState.width,
        height: windowState.height,
        minWidth: Math.floor(defaultWidth * 0.3),
        minHeight: Math.floor(defaultHeight * 0.3),
        frame: false,
        fullscreenable: true,
        fullscreen: windowState.fullscreen,
        webPreferences: {
            nodeIntegration: false,
            contextIsolation: true,
            preload: path.join(__dirname, 'preload.js'),
        },
    });

    mainWindow.setPosition(x, y);

    currentWindow = mainWindow;

    if (FRONTEND_VITE_DEV_SERVER_URL) {
        mainWindow.loadURL(FRONTEND_VITE_DEV_SERVER_URL);
    } else {
        mainWindow.loadFile(path.join(__dirname, `../${FRONTEND_VITE_NAME}/index.html`));
    }

    ipcMain.on("window-close", () => {
        mainWindow.close();
    })

    ipcMain.on("window-maximize", () => {
        if (mainWindow.isMaximized()) {
            mainWindow.unmaximize();
        } else {
            mainWindow.maximize();
        }
    })

    ipcMain.on("window-minimize", () => {
        mainWindow.minimize();
    })

    if (windowState.isDevToolsOpened) {
        mainWindow.webContents.openDevTools({mode: "bottom"});
    }

    // 主界面事件监听
    mainWindow.once("ready-to-show", () => {
        mainWindow.show();
        if (windowState.isMaximized) {
            mainWindow.maximize();
        } else {
            mainWindow.unmaximize();
        }
    });

    mainWindow.on('close', () => {
        exitApp();
    })
};

const exitApp = () => {
    logger.info('start to exit app');

    const bounds = currentWindow.getBounds();
    fs.writeFileSync(windowStatePath, JSON.stringify({
        isMaximized: currentWindow.isMaximized(),
        fullscreen: currentWindow.isFullScreen(),
        isDevToolsOpened: currentWindow.webContents.isDevToolsOpened(),
        x: bounds.x,
        y: bounds.y,
        width: bounds.width,
        height: bounds.height,
    }));
    logger.info('end to exit app');
}

app.whenReady().then(() => {
    initWorkspace();
    createWindow();
    app.on('activate', () => {
        if (BrowserWindow.getAllWindows().length === 0) {
            createWindow();
        }
    });
});

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit();
    }
});
