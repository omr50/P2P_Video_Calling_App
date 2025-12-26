"use strict";
const electron = require("electron");
electron.contextBridge.exposeInMainWorld("ipcRenderer", {
  on(...args) {
    const [channel, listener] = args;
    return electron.ipcRenderer.on(channel, (event, ...args2) => listener(event, ...args2));
  },
  off(...args) {
    const [channel, ...omit] = args;
    return electron.ipcRenderer.off(channel, ...omit);
  },
  send(...args) {
    const [channel, ...omit] = args;
    return electron.ipcRenderer.send(channel, ...omit);
  },
  invoke(...args) {
    const [channel, ...omit] = args;
    return electron.ipcRenderer.invoke(channel, ...omit);
  }
  // login: (credentials: any) => ipcRenderer.invoke('login', credentials),
  // You can expose other APTs you need here.
  // ...
});
electron.contextBridge.exposeInMainWorld("api", {
  login: (credentials) => electron.ipcRenderer.invoke("login", credentials),
  signup: (credentials) => electron.ipcRenderer.invoke("signup", credentials)
});
