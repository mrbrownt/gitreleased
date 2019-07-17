<template>
    <v-list dense>
        <v-list-tile href="/#/">
            <v-list-tile-action>
                <v-icon>home</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
                <v-list-tile-title>Home</v-list-tile-title>
            </v-list-tile-content>
        </v-list-tile>
        <v-list-tile href="/#/user">
            <v-list-tile-action>
                <v-icon>contact_mail</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
                <v-list-tile-title>Subscriptions</v-list-tile-title>
            </v-list-tile-content>
        </v-list-tile>
    </v-list>
</template>


<script lang="ts">
import { Vue, Component } from "vue-property-decorator"

@Component
export default class SidebarComponent extends Vue {}
</script>
