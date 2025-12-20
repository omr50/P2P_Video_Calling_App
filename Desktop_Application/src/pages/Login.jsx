import { useState } from "react"
import Button from "../components/Button"
import Input from "../components/Input"

export default function Login() {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const onSubmit = async (e) => {
    e.preventDefault()

    const result = await window.api.login({
      email,
      password,
    })

    console.log(result)
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
          <Input label="Email" type="email" value={email} onChange={e=>setEmail(e.target.value)} />
          <Input label="Password" type="password" value={password} onChange={e=>setPassword(e.target.value)}/>

          <Button type="submit">
            Sign In
          </Button>
        </form>
      </div>
    </div>
  )
}
