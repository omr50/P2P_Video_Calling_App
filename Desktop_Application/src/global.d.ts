export {}

declare global {
  interface Window {
    api: {
      login: (data: { email: string; password: string }) => Promise<any>
      signup: (data: { email: string; username: string, password: string }) => Promise<any>
      logout: () => Promise<any>
    }
  }
}
