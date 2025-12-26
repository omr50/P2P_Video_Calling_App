import { useNavigate } from "react-router-dom"
import Button from "../components/Button"
import Button2 from "../components/Button2"
import UserProfile from "../components/UserProfile"
import { useAuth } from "../context/AuthContext"

export default function Home() {
  const navigate = useNavigate()
  const {token, setToken} = useAuth()
  return (
    <div className="min-h-screen w-full text-white relative">
      {/* Floating user profile */}
      <div className="fixed top-4 right-4 z-50">
        <UserProfile />
      </div>

      {/* Main content */}
      <main className="min-h-screen w-full flex items-center justify-center px-4">
        <div className="w-full max-w-2xl bg-zinc-800 rounded-xl p-8 shadow-lg">
          
          <h1 className="text-3xl font-bold mb-3">
            Welcome {token ? "Username_(change)" : "to DeCent Chat"}
          </h1>

          <p className="text-zinc-400 mb-6">
            Secure peer-to-peer video calling powered by modern networking.
          </p>


          {
            token ? 
            <>
              <div className="m-4"></div>
              <Button onClick={()=> {navigate("/login")}} bColor="green">🌐 Start a Call </Button>
              <div className="m-4"></div>
              <Button bColor="violet">🙋🏻‍♂️ Add Friend</Button>
            </>
              : 
            <>
              
              <div className="grid gap-4 sm:grid-cols-2">
                <div className="bg-zinc-700/40 rounded-lg p-4">
                  <h2 className="font-semibold mb-1">🔒 Secure</h2>
                  <p className="text-sm text-zinc-400">
                    End-to-end encrypted media using modern cryptography.
                  </p>
                </div>

                <div className="bg-zinc-700/40 rounded-lg p-4">
                  <h2 className="font-semibold mb-1">⚡ Fast</h2>
                  <p className="text-sm text-zinc-400">
                    Direct peer connections with minimal latency.
                  </p>
                </div>

                <div className="bg-zinc-700/40 rounded-lg p-4">
                  <h2 className="font-semibold mb-1">🌍 Decentralized</h2>
                  <p className="text-sm text-zinc-400">
                    No central media relay unless required by NAT.
                  </p>
                </div>

                <div className="bg-zinc-700/40 rounded-lg p-4">
                  <h2 className="font-semibold mb-1">🧠 Built to Learn</h2>
                  <p className="text-sm text-zinc-400">
                    Designed to explore real-world networking concepts.
                  </p>
                </div>
              </div>
              <div className="m-4"></div>
              <Button onClick={()=> {navigate("/login")}} bColor="green">Log In</Button>
              <div className="m-4"></div>
              <Button onClick={()=> {navigate("/signup")}}>Sign Up</Button>
            </>

            
          }
        </div>
      </main>
    </div>
  )
}

