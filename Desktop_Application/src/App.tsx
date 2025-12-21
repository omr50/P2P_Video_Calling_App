import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import Login from "./pages/Login.jsx"
import Home from "./pages/Home.jsx"

function App() {

  return  (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home/>} />
        <Route path="/login" element={<Login />} />
      </Routes>    
    </BrowserRouter>
  )
}

export default App
