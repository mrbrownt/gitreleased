import {
    Module,
    VuexModule,
    getModule,
    MutationAction,
} from "vuex-module-decorators"
import store from "@/store"
import { Repo } from "../models"
import { gitReleasedAPI } from "../api"
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
        let foundRepo: Repo | undefined

        foundRepo = user.subscriptions.find(data => data.id === repoID)
        if (foundRepo) {
            return { repo: foundRepo }
        }

        const response = await gitReleasedAPI.get("/api/repo" + repoID)
        foundRepo = response.data as Repo
        return { repo: foundRepo }
    }
}

export default getModule(RepoModule)
