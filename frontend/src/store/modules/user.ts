import {
    Module,
    VuexModule,
    getModule,
    MutationAction,
} from "vuex-module-decorators"
import { User, Repo } from "../models"
import { getUser, getSubscriptions, subscribe } from "../api"
import store from "@/store/store"

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
        store.state.loading = true
        const user = await getUser()
        store.state.loading = false
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
