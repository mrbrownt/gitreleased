<template>
    <div>
        <SubscribeForm />
        <v-data-table
            :items="subscriptions"
            hide-actions
            hide-headers
            class="elevation-1"
        >
            <template v-slot:items="props">
                <td>{{ props.item.owner }}/{{ props.item.name }}</td>
                <td>{{ props.item.description }}</td>
            </template>
        </v-data-table>
    </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator"
import SubscribeForm from "@/components/Subscribe.vue"
import user from "@/store/modules/user"
import { User } from "@/store/models"

@Component({ components: { SubscribeForm } })
export default class UserComponent extends Vue {
    get subscriptions() {
        return user.subscriptions
    }

    public async created() {
        await user.loadSubs()
    }
}
</script>
