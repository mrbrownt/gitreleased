<template>
    <div>
        <h1>Hello {{ loggedInUser.github_user_name }}!</h1>
    </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator"
import user from "@/store/modules/user"
import { User } from "@/store/models"

@Component
export default class UserComponent extends Vue {
    public loggedInUser: Partial<User> = {}

    public async created() {
        await user.loadUser()
        this.loggedInUser = user.user || {}
    }
}
</script>

<style>
</style>
