export default function Input({ label, type = "text", error, ...props } : any) {
  return (
    <div className="flex flex-col gap-1">
      <label className="text-sm text-zinc-300">
        {label}
      </label>
      <input
        type={type}
        className={`
          bg-zinc-800
          border 
          rounded-md
          px-3 py-2
          text-white
          focus:outline-none
          focus:ring-2
          focus:ring-blue-600
         ${error ? "border-red-400" : "border-zinc-700"}`}
        {...props}
      />
    </div>
  )
}
