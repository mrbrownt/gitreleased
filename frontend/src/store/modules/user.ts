import {
    Module,
    VuexModule,
    getModule,
    MutationAction,
} from "vuex-module-decorators"
import { User, Repo } from "../models"
import { gitReleasedAPI } from "../api"
import store from "@/store"

@Module({
    dynamic: true,
    namespaced: true,
    name: "user",
    store,
})
class UserModule extends VuexModule {
    public user: User | null = null
    public subscriptions: Repo[] = []

    @MutationAction
    public async loadUser() {
        const user = await getUser()
        return { user }
    }

    @MutationAction({ mutate: ["subscriptions"] })
    public async loadSubs() {
        const subscriptions = await getSubscriptions()
        return { subscriptions }
    }

    @MutationAction({ mutate: ["subscriptions"] })
    public async addSub(repo: string) {
        await subscribe(repo)
        const subscriptions = await getSubscriptions()
        return { subscriptions }
    }
}

async function getUser() {
    const response = await gitReleasedAPI.get("/api/user")
    return response.data
}

async function getSubscriptions() {
    const response = await gitReleasedAPI.get("/api/user/subscriptions")
    return response.data
}

async function subscribe(repo: string) {
    await gitReleasedAPI.post("/api/user/subscribe?repo=" + repo)
}

export default getModule(UserModule)
