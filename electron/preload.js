const {contextBridge, ipcRenderer} = require('electron')

contextBridge.exposeInMainWorld('appObject', {
    send: (channel, ...args) => {
        ipcRenderer.send(channel, ...args)
    },
    sendSync: (channel, ...args) => {
        return ipcRenderer.sendSync(channel, ...args)
    }
})
