import Vue from "vue"
import Router from "vue-router"
import Home from "./views/Home.vue"
import User from "./views/User.vue"
import Repo from "./views/Repo.vue"
import Login from "./views/Login.vue"

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: "/",
            name: "home",
            component: Home,
        },
        {
            path: "/login",
            name: "login",
            component: Login,
        },
        {
            path: "/user",
            name: "user",
            component: User,
        },
        {
            path: "/repo/:id",
            name: "repo",
            component: Repo,
            props: true,
        },
    ],
})
