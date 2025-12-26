import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import Login from "./pages/Login.jsx"
import Signup from "./pages/Signup.jsx"
import Home from "./pages/Home.jsx"

function App() {

  return  (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home/>} />
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Signup/>} />
      </Routes>    
    </BrowserRouter>
  )
}

export default App
