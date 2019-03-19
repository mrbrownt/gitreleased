<template>
    <div v-bind="subscriptions">
        <h1>{{ subcription.owner }}/{{ subcription.name }}</h1>
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
}
</script>

