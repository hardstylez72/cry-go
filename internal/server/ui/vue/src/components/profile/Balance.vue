<template>
  <div>
    <NetworkChip :network="network"/>
    <span v-if="loading">
      <v-progress-circular
        indeterminate
        color="primary"
      ></v-progress-circular>
    </span>
    <div v-else-if="error">
      error loading balance
      <v-btn @click="loadBalance">retry</v-btn>
    </div>
    <div v-else-if="balance.length === 0">
      no balance
      <v-btn density="compact" @click="loadBalance">retry</v-btn>
    </div>
    <div v-else>
      <div v-for="b in balance">
        {{ b.token }}:
        <span v-if="b.error">{{ b.error }}</span>
        <span v-else>{{ Number(b.amount).toPrecision(3) }}</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {profileService} from "@/generated/services"
import {Balance, Network} from "@/generated/profile";
import NetworkChip from "@/components/tasks/NetworkChip.vue";

export default defineComponent({
  name: "Balance.vue",
  components: {NetworkChip},
  props: {
    profileId: {
      type: String,
      required: true
    },
    network: {
      type: String as PropType<Network>,
      required: true
    }
  },
  data() {
    return {
      loading: false,
      balance: [] as Balance[],
      error: false
    }
  },
  computed: {},
  methods: {
    async loadBalance() {
      try {
        this.error = false
        this.loading = true
        const res = await profileService.profileServiceGetBalance({
          body: {
            profileId: this.profileId,
            network: this.network
          }
        })
        if (res.balance) {
          this.balance = res.balance
        }
      } catch (e) {
        this.error = true
      } finally {
        this.loading = false
      }
    }
  },
  created() {
    this.loadBalance()
  }
})


</script>


<style>


</style>

