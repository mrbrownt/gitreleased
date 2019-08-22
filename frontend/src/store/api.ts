import axios from "axios"
import { User, Repo } from "./models"
import PubSub from "pubsub-js"

export const gitReleasedAPI = axios.create({ baseURL: "" })

export async function getRepo(repo: string) {
    const response = await gitReleasedAPI.get("/api/repo/" + repo)
    return response.data as Repo
}

gitReleasedAPI.interceptors.request.use(
    config => {
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
