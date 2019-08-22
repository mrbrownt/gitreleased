<template>
    <v-container class="elevation-1">
        <v-layout row>
            <v-flex pa1 grow>
                <v-text-field
                    v-model="newSubscription"
                    label="Github Repo"
                    required
                ></v-text-field>
            </v-flex>
            <v-flex pa1 shrink>
                <v-btn @click="subscribe" color="info">Subscribe</v-btn>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script lang="ts">
import { Vue, Component, Prop } from "vue-property-decorator"
import user from "@/store/modules/user"

@Component
export default class SubscribeForm extends Vue {
    private newSubscription: string = ""

    private async subscribe() {
        await user.addSub(this.newSubscription)
        this.newSubscription = ""
    }
}
</script>

<style>
</style>
