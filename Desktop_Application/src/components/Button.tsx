export default function Button({ children, onClick, bColor="grey", type = "button" }: any) {
  return (
    <button
      type={type}
      onClick={onClick}
      className={`
        w-full py-2 px-4 rounded-md
        text-white font-medium transition
        ${bColor === "green" ? "!bg-blue-500 !hover:border-red-500" : ""}
        ${bColor === "violet" ? "!bg-slate-700 !hover:bg-red-500" : ""}
        ${bColor === "grey" ? "!bg-zinc-700 !hover:bg-blue-500" : ""}
      `}
    >
      {children}
    </button>
  )
}
