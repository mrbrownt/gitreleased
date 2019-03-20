import {
    Module,
    VuexModule,
    getModule,
    MutationAction,
} from "vuex-module-decorators"
import store from "@/store"
import { Repo } from "../models"
import { getRepo } from "../api"
import user from "./user"

@Module({
    dynamic: true,
    namespaced: true,
    name: "repo",
    store,
})
class RepoModule extends VuexModule {
    public repo: Repo = {
        id: "loading",
        owner: "loading",
        name: "loading",
        description: "loading",
        url: "loading",
    } as Repo

    @MutationAction({ mutate: ["repo"] })
    public async loadRepo(repoID: string) {
        const localStore = user.subscriptions.find(repo => repo.id === repoID)
        if (localStore) {
            return { repo: localStore }
        }
        const result = await getRepo(repoID)
        return { repo: result }
    }
}

export default getModule(RepoModule)
