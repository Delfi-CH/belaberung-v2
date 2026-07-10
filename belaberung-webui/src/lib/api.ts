import axios from "axios";

const backendURL = import.meta.env.DEV ? "http://localhost:8081" : "/api"

export const api = axios.create({
    baseURL: backendURL,
    withCredentials: true
})

export async function isLoggedIn() {
    try {
        await api.get("/auth/status")
        return true
    } catch {
        return false
    }
}