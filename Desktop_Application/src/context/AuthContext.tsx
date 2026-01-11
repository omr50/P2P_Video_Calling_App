import { createContext, useContext, useState } from "react"

type AuthContextType = {
    token: string | null
    setToken: any | null
    email: string | null
    setEmail: any | null
    username: string | null
    setUsername: any | null
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: any) {
    const [token, setToken] = useState(null)
    const [email, setEmail] = useState(null)
    const [username, setUsername] = useState(null)

    return (
        <AuthContext.Provider value={{ token, setToken, email, setEmail, username, setUsername }}>
            {children}
        </AuthContext.Provider>
    )
}

export function useAuth() {
    const ctx = useContext(AuthContext)
    if (!ctx) {
        throw new Error("useAuth must be used inside AuthProvider")
    }

    return ctx;
}

export function authFetch() {
    
}