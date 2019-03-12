import axios from "axios"
import { User } from "./models"

export const gitReleasedAPI = axios.create({
    baseURL: "http://localhost:8081",
})

export async function getUser(): Promise<User> {
    const response = await gitReleasedAPI.get("/user/")
    return response.data as User
}
