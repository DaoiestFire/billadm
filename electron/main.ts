import * as path from 'path';
import * as fs from 'fs';
import {app, BrowserWindow, ipcMain, screen, dialog, OpenDialogOptions} from 'electron';
import Logger from './logger';
import BilladmDao from './billadmDao';
import {Workspace, WorkspaceState} from './api';

if (require('electron-squirrel-startup')) {
    app.quit();
}

const appVersion = app.getVersion();
const confDir = path.join(app.getPath("home"), ".config", "billadm");
const windowStatePath = path.join(confDir, "windowState.json");
const workspaceStatePath = path.join(confDir, "workspace.json")
const logger = new Logger(path.join(confDir, 'app.log'));
let firstOpen = false;
let workspaceState: WorkspaceState;
let workspace: Workspace = {billadmDao: null, workspaceDir: '', mainWindow: null};


try {
    if (!fs.existsSync(confDir)) {
        fs.mkdirSync(confDir, {mode: 0o755, recursive: true});
    }
} catch (e) {
    console.error(e);
    dialog.showErrorBox("创建配置目录失败", "需要在用户家目录下创建配置文件夹（~/.config/billadm），请确保该路径具有写入权限");
    app.exit();
}

let oldWorkspace = {};
try {
    oldWorkspace = JSON.parse(fs.readFileSync(workspaceStatePath, "utf8"));
} catch (e) {
    fs.writeFileSync(workspaceStatePath, "{}");
}
workspaceState = Object.assign({}, {
    last: '',
    workspaces: [],
}, oldWorkspace);

if (workspaceState.workspaces.length == 0) {
    firstOpen = true;
} else if (!workspaceState.last) {
    workspaceState.last = workspaceState.workspaces[0];
    if (fs.existsSync(workspaceState.last) && fs.statSync(workspaceState.last).isFile()) {
        dialog.showErrorBox("工作目录错误", `工作目录【${workspaceState.last}】已存在同名文件，请检查`);
        app.exit();
    }
    logger.info(`当前工作目录[${workspaceState.last}]`)
}

logger.info(`start to launch billadm, firstOpen: ${firstOpen}`);

const assignWorkspace = async (dbFile: string, logFile: string, workspaceDir: string) => {
    workspace.billadmDao = new BilladmDao(dbFile, logFile);
    workspace.workspaceDir = workspaceDir;
    try {
        await workspace.billadmDao.init();
        logger.info(`init billadmDao for ${workspace.workspaceDir} success`);
    } catch (e) {
        logger.info(`init billadmDao for ${workspace.workspaceDir} failed`);
        dialog.showErrorBox("初始化billadmDao失败", `工作目录【${workspace.workspaceDir}】错误信息【${e.message}】`);
        app.exit();
    }
}

const connectWorkspace = async () => {
    logger.info(`connect workspace: ${workspaceState.last}`);
    let dbFile = path.join(workspaceState.last, 'data', 'billadm.db');
    let logFile = path.join(workspaceState.last, 'billadm.db.log');
    await assignWorkspace(dbFile, logFile, workspaceState.last);
};

const initWorkspace = async (workspaceDir: string) => {
    try {
        let items = fs.readdirSync(workspaceDir);
        if (items.length != 0) {
            dialog.showErrorBox(" 工作空间路径错误", `工作空间不为空：${workspaceDir}`);
            return false;
        }
    } catch (e) {
        console.error(e);
        dialog.showErrorBox("读取目录失败", `读取工作空间失败：${workspaceDir}`);
        return false;
    }
    let dataDir = path.join(workspaceDir, 'data');
    try {
        fs.mkdirSync(dataDir, {mode: 0o755, recursive: true});
    } catch (e) {
        console.error(e);
        dialog.showErrorBox("创建目录失败", `在工作空间下创建数据目录失败：${dataDir}`);
        return false;
    }
    logger.info(`init workspace: ${workspaceDir}`);
    let dbFile = path.join(dataDir, 'billadm.db');
    let logFile = path.join(workspaceDir, 'billadm.db.log');
    await assignWorkspace(dbFile, logFile, workspaceDir);
    try {
        await workspace.billadmDao.initDB(appVersion);
        logger.info(`init db for ${workspaceDir} success`);
    } catch (e) {
        logger.info(`init db for ${workspaceDir} failed`);
        dialog.showErrorBox("初始化数据库失败", `工作目录【${workspaceDir}】错误信息【${e.message}】`);
        app.exit();
    }
    workspaceState.last = workspaceDir;
    workspaceState.workspaces.push(workspaceDir);
    return true;
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
    workspace.mainWindow = mainWindow;

    mainWindow.setPosition(x, y);

    if (FRONTEND_VITE_DEV_SERVER_URL) {
        mainWindow.loadURL(FRONTEND_VITE_DEV_SERVER_URL);
    } else {
        mainWindow.loadFile(path.join(__dirname, `../${FRONTEND_VITE_NAME}/index.html`));
    }

    ipcMain.on("window.close", () => {
        mainWindow.close();
    });

    ipcMain.on("window.maximize", () => {
        if (mainWindow.isMaximized()) {
            mainWindow.unmaximize();
        } else {
            mainWindow.maximize();
        }
    });

    ipcMain.on("window.minimize", () => {
        mainWindow.minimize();
    });

    // billbooks
    ipcMain.handle("billbooks.all-billbooks", async () => {
        return await workspace.billadmDao.queryAllBillbook();
    });
    ipcMain.handle("billbooks.add-one-billbook", async (event, item) => {
        return await workspace.billadmDao.insertOneBillbook(item.name, item.description);
    });
    ipcMain.handle("billbooks.delete-one-billbook", async (event, bookId) => {
        return await workspace.billadmDao.deleteOneBillbookByID(bookId);
    });

    // bills
    ipcMain.handle("bills.all-bills", async (event, bookId: string) => {
        return await workspace.billadmDao.queryAllBillByBookID(bookId);
    });
    ipcMain.handle("bills.all-bills-with-filters", async (event, bookId: string, filters: object) => {
        return await workspace.billadmDao.queryAllBillByBookIDWithFilters(bookId, filters);
    });
    ipcMain.handle("bills.add-one-bill", async (event, item) => {
        return await workspace.billadmDao.insertOneBill(item.money, item.type, item.income, item.bookId, item.description, item.tags, item.creationTime);
    });
    ipcMain.handle("bills.delete-bills", async (event, idList: string[]) => {
        for (let id of idList) {
            await workspace.billadmDao.deleteOneBillByID(id);
        }
    });

    // billtypes
    ipcMain.handle("billtypes.all-billtypes", async () => {
        return await workspace.billadmDao.queryAllType();
    });

    // 用于初始化工作间
    ipcMain.handle('init.first-open', async (event, item) => {
        return firstOpen;
    });
    ipcMain.handle('init.choose-workspace-directory', async (event, item) => {
        const options: OpenDialogOptions = {
            title: '选择目录',
            properties: ['openDirectory'],
        };
        const {canceled, filePaths} = await dialog.showOpenDialog(options);
        if (canceled) {
            return '';
        } else {
            return filePaths[0];
        }
    });
    ipcMain.handle('init.init-workspace', async (event, workspaceDir) => {
        return await initWorkspace(workspaceDir);
    });

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
    // 保存退出前的窗口状态
    const bounds = workspace.mainWindow.getBounds();
    fs.writeFileSync(windowStatePath, JSON.stringify({
        isMaximized: workspace.mainWindow.isMaximized(),
        fullscreen: workspace.mainWindow.isFullScreen(),
        isDevToolsOpened: workspace.mainWindow.webContents.isDevToolsOpened(),
        x: bounds.x,
        y: bounds.y,
        width: bounds.width,
        height: bounds.height,
    }));
    // 保存全部的工作空间路径和上次打开的工作空间路径
    fs.writeFileSync(workspaceStatePath, JSON.stringify(workspaceState));
    logger.info('end to exit app');
}

app.whenReady().then(async () => {
    if (!firstOpen) {
        await connectWorkspace();
    }
    createWindow();
});

app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) {
        createWindow();
    }
});

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit();
    }
});
