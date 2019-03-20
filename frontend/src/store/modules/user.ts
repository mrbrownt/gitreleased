import {
    Module,
    VuexModule,
    getModule,
    MutationAction,
} from "vuex-module-decorators"
import store from "@/store"
import { User, Repo } from "../models"
import { getUser, getSubscriptions, subscribe } from "../api"

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

export default getModule(UserModule)
