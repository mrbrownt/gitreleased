<template>
    <form @submit="subscribe">
        <div class="field">
            <label class="label"
                >Repo: <input v-model="sub.name" placeholder="test"
            /></label>
        </div>
        <div><button type="submit">Subscribe</button></div>
    </form>
</template>

<script lang="ts">
import { Vue, Component, Prop } from "vue-property-decorator"
import user from "@/store/modules/user"

@Component
export default class SubscribeForm extends Vue {
    public sub: { name: string } = { name: "github/repo" }

    private async subscribe() {
        await user.addSub(this.sub.name)
    }
}
</script>

<style>
</style>
