import { useCall } from "../context/CallProvider"

export default function IncomingCallModal() {
  const { incomingCall, acceptCall, declineCall } = useCall()

  if (!incomingCall) return null

  return (
    <div className="fixed bottom-5 right-5 z-[9999] w-80 rounded-xl bg-gray-900 text-white shadow-2xl p-4 animate-slide-up">
      <h3 className="text-lg font-semibold">Incoming Call from {incomingCall.callerEmail}</h3>
      <p className="mt-1 text-sm text-gray-300">
        {incomingCall.callerEmail}
      </p>

      <div className="mt-4 flex justify-end gap-2">
        <button
          onClick={declineCall}
          className="rounded-md bg-red-600 px-4 py-2 text-sm font-medium hover:bg-red-700"
        >
          Decline
        </button>

        <button
          onClick={acceptCall}
          className="rounded-md bg-green-600 px-4 py-2 text-sm font-medium hover:bg-green-700"
        >
          Accept
        </button>
      </div>
    </div>
  )
}
