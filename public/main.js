const {
  app,
  BrowserWindow,
  ipcMain,
  screen,
  dialog,
} = require('electron');
const path = require('node:path');
const fs = require('fs');
const logger = require('./logger');

if (require('electron-squirrel-startup')) {
  app.quit();
}

const confDir = path.join(app.getPath("home"), ".config", "billadm");
const windowStatePath = path.join(confDir, "windowState.json");
let firstOpen = false;

logger.init(path.join(confDir, 'app.log'));

try {
  firstOpen = !fs.existsSync(path.join(confDir, "workspace.json"));
  if (!fs.existsSync(confDir)) {
    fs.mkdirSync(confDir, { mode: 0o755, recursive: true });
  }
} catch (e) {
  console.error(e);
  dialog.showErrorBox("创建配置目录失败", "需要在用户家目录下创建配置文件夹（~/.config/billadm），请确保该路径具有写入权限");
  app.exit();
}

logger.info(`start to launch billadm, firstOpen: ${firstOpen}`);

let currentWindow;

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
    mainWindow.webContents.openDevTools({ mode: "bottom" });
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
