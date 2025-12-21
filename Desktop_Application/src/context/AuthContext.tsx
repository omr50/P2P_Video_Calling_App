import { createContext, useContext, useState } from "react"

type AuthContextType = {
    token: string | null
    setToken: any
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: any) {
    const [token, setToken] = useState(null)

    return (
        <AuthContext.Provider value={{ token, setToken }}>
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