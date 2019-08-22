<template>
    <div>
        <router-link
            v-bind:to="{ name: 'repo', params: { id: subcription.id } }"
            class="item"
            ><h3>
                {{ subcription.owner }}/{{ subcription.name }}
            </h3></router-link
        >
        {{ subcription.description }}
    </div>
</template>

<script lang="ts">
import { Vue, Component, Prop, Watch } from "vue-property-decorator"
import { Repo } from "@/store/models"

@Component
export default class SubscriptionList extends Vue {
    @Prop({ default: {} as Repo }) public subcription?: Repo
}
</script>
