const {contextBridge, ipcRenderer} = require('electron')

contextBridge.exposeInMainWorld('appObject', {
    send: (channel, ...args) => {
        ipcRenderer.send(channel, ...args)
    },
    getBilladmInfo() {
        return ipcRenderer.invoke('billadm.info');
    },
    getAllBillbooks() {
        return ipcRenderer.invoke('billbooks.all-billbooks');
    },
    addOneBillbook(item) {
        return ipcRenderer.invoke('billbooks.add-one-billbook', item);
    },
    deleteOneBillbook(item) {
        return ipcRenderer.invoke('billbooks.delete-one-billbook', item);
    },
    getAllBillsByBookID(bookId) {
        return ipcRenderer.invoke('bills.all-bills', bookId);
    },
    getAllBillsByBookIDWithFilters(bookId, filters) {
        return ipcRenderer.invoke('bills.all-bills-with-filters', bookId, filters);
    },
    addOneBill(item) {
        return ipcRenderer.invoke('bills.add-one-bill', item);
    },
    editOneBill(item) {
        return ipcRenderer.invoke('bills.edit-one-bill', item);
    },
    deleteBills(idList) {
        return ipcRenderer.invoke('bills.delete-bills', idList);
    },
    getAllBillTypes() {
        return ipcRenderer.invoke('billtypes.all-billtypes');
    },
    chooseWorkspaceDirectory() {
        return ipcRenderer.invoke('init.choose-workspace-directory');
    },
    initWorkspace(workspaceDir) {
        return ipcRenderer.invoke('init.init-workspace', workspaceDir);
    },
    removeWorkspace(workspaceDir) {
        return ipcRenderer.invoke('init.remove-workspace', workspaceDir);
    },
    isFirstOpen() {
        return ipcRenderer.invoke("init.first-open");
    },
    workspaceState() {
        return ipcRenderer.invoke("init.workspaceState");
    },
})
