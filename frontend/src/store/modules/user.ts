import {
    Module,
    VuexModule,
    getModule,
    MutationAction,
} from "vuex-module-decorators"
import store from "@/store"
import { User } from "../models"
import { getUser } from "../api"

@Module({
    dynamic: true,
    namespaced: true,
    name: "user",
    store,
})
class UserModule extends VuexModule {
    public user: User | null = null

    @MutationAction
    public async loadUser() {
        const user = await getUser()
        return { user }
    }
}

export default getModule(UserModule)
