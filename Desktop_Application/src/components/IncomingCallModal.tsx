import { useCall } from "../context/CallProvider"

export default function IncomingCallModal() {
  const { incomingCall, acceptCall, declineCall } = useCall()

  if (!incomingCall) return null

  return (
    <div className="modal">
      <h2>Incoming Call</h2>
      <p>{incomingCall.callerEmail}</p>
      <button onClick={acceptCall}>Accept</button>
      <button onClick={declineCall}>Decline</button>
    </div>
  )
}
