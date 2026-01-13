import { createContext, useContext, useEffect, useState } from "react"

const CallContext = createContext<CallContextType | null>(null)


type CallOffer = {
  callerEmail: string
}
type CallContextType = {
  incomingCall: CallOffer | null 
  acceptCall: () => void
  declineCall: () => void
}

export const useCall = () => useContext(CallContext)!

export function CallProvider({ children }: { children: React.ReactNode }) {
  const [incomingCall, setIncomingCall] = useState<CallOffer | null>(null)

useEffect(() => {
  console.log("SSE effect mount")

  const es = new EventSource("http://localhost:5000/sse")

  es.onmessage = (e) => {
    console.log("SSE DATA:", e.data)
  }

  return () => {
    console.log("SSE effect cleanup")
    es.close()
  }
}, [])

  const acceptCall = async () => {
    // call backend
    setIncomingCall(null)
  }

  const declineCall = async () => {
    // call backend
    setIncomingCall(null)
  }

  return (
    <CallContext.Provider value={{ incomingCall, acceptCall, declineCall }}>
      {children}
    </CallContext.Provider>
  )
}
