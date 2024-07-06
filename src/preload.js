const { contextBridge, ipcRenderer } = require('electron')

contextBridge.exposeInMainWorld('windowController', {
    send: (channel, ...args) => {
        ipcRenderer.send(channel, ...args)
    },
    sendSync: (channel, ...args) => {
        return ipcRenderer.sendSync(channel, ...args)
    }
})
