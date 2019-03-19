<template>
    <div id="app">
        <el-menu></el-menu>
        <div id="nav">
            <router-link to="/">Home</router-link>|
            <router-link v-if="user" to="/user">{{ user.github_user_name }}</router-link>
            <router-link v-else to="/login">Login</router-link>
        </div>
        <router-view/>
    </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator"
import { mapState } from "vuex"
import store from "@/store/store"
import user from "@/store/modules/user"
import { User } from "@/store/models"

@Component
export default class extends Vue {
    get user() {
        return user.user
    }

    public async created() {
        await user.loadUser()
    }
}
</script>


<style>
#body {
    background-color: #c9e9f1;
}

#app {
    font-family: "Avenir", Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #1723a2;
}

#nav {
    padding: 30px;
}

#nav a {
    font-weight: bold;
    color: #59667e;
}

#nav a.router-link-exact-active {
    color: #2e2268;
}
</style>
