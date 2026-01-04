import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import Login from "./pages/Login.jsx"
import Signup from "./pages/Signup.jsx"
import Home from "./pages/Home.jsx"
import SearchBar from './components/SearchBar.js'
import UserSearch from './pages/Search.js'

function App() {

  return  (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home/>} />
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Signup/>} />
        <Route path="/search" element={<UserSearch/>} />
      </Routes>    
    </BrowserRouter>
  )
}

export default App
