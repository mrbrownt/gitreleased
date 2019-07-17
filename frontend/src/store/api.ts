import axios from "axios"
import { User, Repo } from "./models"
import PubSub from "pubsub-js"

export const gitReleasedAPI = axios.create({
    baseURL: process.env.production
        ? "https://api.gitreleased.app"
        : "http://localhost:8080",
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

let lastURL

gitReleasedAPI.interceptors.request.use(
    config => {
        lastURL = config.url
        return config
    },

    error => {
        return Promise.reject(error)
    },
)

gitReleasedAPI.interceptors.response.use(
    response => {
        return response
    },

    error => {
        PubSub.publish("API_ERROR", error)
        return Promise.reject(error)
    },
)
