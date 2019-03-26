import Vue from "vue"
import Router from "vue-router"
const Home = () => import("./views/Home.vue")
const User = () => import("./views/User.vue")
const Repo = () => import("./views/Repo.vue")
const Login = () => import("./views/Login.vue")

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
