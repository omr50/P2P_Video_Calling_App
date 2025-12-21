import { useState } from "react"
import { useAuth } from "../context/AuthContext"
import { useNavigate } from "react-router-dom"
import Button from "./Button"

export default function UserProfile() {
  const { token, setToken } = useAuth()
  const [open, setOpen] = useState(false)
  const navigate = useNavigate()

  if (!token) return null // not logged in → show nothing

  // Simple initials for now
  const initials = "U" // later: derive from username/email
console.log("UserProfile token:", token)

  return (
    <div className="relative">
      {/* Avatar */}
      <Button
        onClick={() => setOpen(!open)}
        className="
          w-10 h-10
          rounded-full
          bg-blue-600
          text-white
          font-semibold
          flex items-center justify-center
          hover:bg-blue-500
        "
      >
        {initials}
      </Button>

      {/* Dropdown */}
      {open && (
        <div
          className="
            absolute right-0 mt-2
            w-48
            bg-zinc-800
            border border-zinc-700
            rounded-lg
            shadow-lg
            p-3
            z-50
          "
        >
          <div className="text-sm text-zinc-300 mb-2">
            Logged in
          </div>

          <Button
            onClick={() => {
              setToken("")
              setOpen(false)
              navigate("/")
            }}
            className="
              w-full
              text-left
              text-sm
              text-red-400
              hover:text-red-300
            "
          >
            Log out
          </Button>
        </div>
      )}
    </div>
  )
}
