const {contextBridge, ipcRenderer} = require('electron')

contextBridge.exposeInMainWorld('appObject', {
    send: (channel, ...args) => {
        ipcRenderer.send(channel, ...args)
    },
    getAllBillbooks() {
        return ipcRenderer.invoke('billbooks.all-billbooks');
    },
})
