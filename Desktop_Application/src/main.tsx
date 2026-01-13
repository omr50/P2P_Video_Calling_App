import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import {AuthProvider} from './context/AuthContext.js'
import { CallProvider } from './context/CallProvider.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
    <AuthProvider>
      <CallProvider>
        <App />
      </CallProvider>
    </AuthProvider>
)

// Use contextBridge
window.ipcRenderer.on('main-process-message', (_event, message) => {
  console.log(message)
})
