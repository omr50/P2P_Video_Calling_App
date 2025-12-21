import { useState } from "react"
import Button from "../components/Button"
import Input from "../components/Input"
import {useAuth} from "../context/AuthContext"
import { useNavigate } from "react-router-dom"

export default function Login() {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [wrongCredentials, setWrongCredentials] = useState(false)
  const {token, setToken} = useAuth()
  const navigate = useNavigate()

  const onSubmit = async (e) => {
    e.preventDefault()
    console.log("test?")
    const result = await window.api.login({
      email,
      password,
    })

    console.log("http result:", result)
    if (result.token === ""){
      setWrongCredentials(true);
      setTimeout(()=> {setWrongCredentials(false);}, 3000);
    } else {
      setToken(result.token)
      navigate("/")
    }
  }
  return (
    <div className="max-h-screen flex items-center justify-center bg-zinc-900">
      <div className="w-full max-w-sm bg-zinc-800 p-6 rounded-lg shadow-lg">
        <h1 className="text-2xl font-bold text-white mb-1">
          
        </h1>
        <p className="text-zinc-400 mb-6">
          Sign in to continue
        </p>

        <form onSubmit={onSubmit} className="flex flex-col gap-4">
          {wrongCredentials ? <div className="bg-rose-400 text-white rounded text-sm py-1">Incorrect Credentials!</div> : ""}
          <Input label="Email" type="email" value={email} onChange={e=>setEmail(e.target.value)} error={wrongCredentials} />
          <Input label="Password" type="password" value={password} onChange={e=>setPassword(e.target.value)} error={wrongCredentials}/>

          <Button type="submit">
            Sign In
          </Button>
        </form>
      </div>
    </div>
  )
}
