<template>
    <div v-bind="subscriptions">
        <a @click="navigateToRepo" href="#">
            <h2>{{ subcription.owner }}/{{ subcription.name }}</h2>
        </a>
        {{ subcription.description }}
    </div>
</template>

<script lang="ts">
import { Vue, Component, Prop, Watch } from "vue-property-decorator"
import { Repo } from "@/store/models"
import user from "@/store/modules/user"
import { mapState } from "vuex"

@Component
export default class SubscriptionList extends Vue {
    @Prop({ default: {} as Repo }) public subcription?: Repo

    private navigateToRepo() {
        if (this.subcription) {
            this.$router.push({
                name: "repo",
                params: { id: this.subcription.id },
            })
        }
    }
}
</script>
