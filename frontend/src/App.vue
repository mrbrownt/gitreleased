<template>
    <v-app id="inspire">
        <v-navigation-drawer v-model="drawer" fixed app disable-resize-watcher>
            <Sidebar />
        </v-navigation-drawer>
        <v-toolbar color="indigo" dark fixed app>
            <v-toolbar-side-icon
                @click.stop="drawer = !drawer"
            ></v-toolbar-side-icon>
            <v-toolbar-title>GitReleased</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items>
                <template v-if="user">
                    <v-btn href="/auth/logout" flat>Logout</v-btn>
                </template>
                <template v-else>
                    <v-btn flat href="/auth?provider=github">Login</v-btn>
                </template>
            </v-toolbar-items>
        </v-toolbar>
        <v-content>
            <v-container fluid fill-height>
                <v-layout justify-center align-top>
                    <router-view />
                </v-layout>
            </v-container>
        </v-content>
    </v-app>
</template>


<script lang="ts">
import { Vue, Component, Prop } from "vue-property-decorator"
import { mapState } from "vuex"
import store from "@/store"
import user from "@/store/modules/user"
import { User } from "@/store/models"
import Sidebar from "@/components/Sidebar.vue"

@Component({ components: { Sidebar } })
export default class extends Vue {
    private drawer: boolean = false

    get user() {
        return user.user
    }

    get loading(): boolean {
        return store.state.loading
    }

    private async created() {
        await user.loadUser()
    }
}
</script>


<style>
</style>
