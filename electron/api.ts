import {BrowserWindow} from 'electron';
import BilladmDao from "./billadmDao";

export interface WorkspaceState {
    last: string,
    workspaces: string[],
}

export interface Workspace {
    billadmDao: BilladmDao;
    workspaceDir: string;
    mainWindow: BrowserWindow;
}