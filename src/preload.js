const { contextBridge, ipcRenderer } = require('electron')

contextBridge.exposeInMainWorld('windowController', {
    send: (channel) => {
        ipcRenderer.send(channel)
    }
})
