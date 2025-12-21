

export default function Button2({ children, onClick, bColor="blue", type = "button" }: any) {
  return (
    <button
      type={type}
      onClick={onClick}
      className={`
        w-full py-2 px-4 rounded-md
        text-white font-medium transition
        ${bColor === "green" && "bg-green-600 hover:bg-green-500"}
        ${bColor === "blue" && "bg-blue-600 hover:bg-blue-500"}
        ${bColor === "red" && "bg-red-600 hover:bg-red-500"}
      `}
    >
      {children}
    </button>
  )
}
