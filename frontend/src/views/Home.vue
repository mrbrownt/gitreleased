<template>
    <v-layout>
        <v-flex xs12 sm6 offset-sm3>
            <v-card>
                <v-img :src="logo" alt="GitReleased Logo" />
            </v-card>
        </v-flex>
    </v-layout>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator"
import SubscribeForm from "@/components/Subscribe.vue"
import SubscriptionList from "@/components/Subscriptions.vue"
import user from "@/store/modules/user"
import { User } from "@/store/models"

@Component({ components: { SubscribeForm, SubscriptionList } })
export default class Home extends Vue {
    private logo: null = require("@/assets/logo.png")
}
</script>

<style>
img {
    width: 50%;
    max-height: auto;
}
</style>
