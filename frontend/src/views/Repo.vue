<template>
    <div>
        <h1>
            <a :href="repo.url" target="_blank" rel="noopener noreferrer">
                {{ repo.name }}/{{ repo.owner }}</a
            >
        </h1>
        {{ repo.description }}
    </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from "vue-property-decorator"
import repo from "@/store/modules/repo"

@Component
export default class RepoComponent extends Vue {
    get repo() {
        return repo.repo
    }

    public async created() {
        await repo.loadRepo(this.$route.params.id)
    }
}
</script>

<style>
</style>
