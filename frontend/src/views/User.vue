<template>
    <div>
        <SubscribeForm />
        <SubscriptionList
            v-for="subscription in subscriptions"
            :subcription="subscription"
            :key="subscription.id"
        />
    </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator"
import SubscribeForm from "@/components/Subscribe.vue"
import SubscriptionList from "@/components/Subscriptions.vue"
import user from "@/store/modules/user"
import { User } from "@/store/models"

@Component({ components: { SubscribeForm, SubscriptionList } })
export default class UserComponent extends Vue {
    get subscriptions() {
        return user.subscriptions
    }

    public async created() {
        await user.loadSubs()
    }
}
</script>
