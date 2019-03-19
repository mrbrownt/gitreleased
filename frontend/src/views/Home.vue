<template>
    <div class="home">
        <img alt="Vue logo" src="../assets/logo.png" />
    </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator"

@Component({
    components: {},
})
export default class Home extends Vue {}
</script>

<style>
img {
    width: 50%;
    max-height: auto;
}
</style>
