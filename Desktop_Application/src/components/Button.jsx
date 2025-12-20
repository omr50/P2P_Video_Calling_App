export default function Button({ children, onClick, type = "button" }) {
  return (
    <button
      type={type}
      onClick={onClick}
      className="
        w-full
        py-2 px-4
        rounded-md
        bg-blue-600
        hover:bg-blue-500
        text-white
        font-medium
        transition
      "
    >
      {children}
    </button>
  )
}
