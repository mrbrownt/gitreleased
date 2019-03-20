import axios from "axios"
import { User, Repo } from "./models"

export const gitReleasedAPI = axios.create({
    baseURL: "http://localhost:8081",
})

export async function getUser(): Promise<User> {
    const response = await gitReleasedAPI.get("/api/user/")
    return response.data as User
}

export async function getSubscriptions(): Promise<Repo[]> {
    const response = await gitReleasedAPI.get("/api/user/subscriptions")
    return response.data as Repo[]
}

export async function subscribe(repo: string) {
    await gitReleasedAPI.post("/api/user/subcribe?repo=" + repo)
}

export async function getRepo(repo: string) {
    const response = await gitReleasedAPI.get("/api/repo/" + repo)
    return response.data as Repo
}
