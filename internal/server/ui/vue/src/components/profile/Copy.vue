<template>
  <v-icon class="mx-1" icon="mdi-content-copy" @click="click"/>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {processService, profileService} from "@/generated/services"
import Balance from "@/components/profile/Balance.vue";
import {Network, Profile, ProfileType} from "@/generated/profile";
import TopUp from "@/components/billing/TopUp.vue";
import Withdraw from "@/components/exchange.acc/ExchangeAccounts.vue";
import TopUpProfile from "@/components/exchange.acc/TopUpProfile.vue";
import {Transaction} from "@/generated/process";
import {copyToClipboard, formatTime} from "../helper";
import {useSysStore} from "@/plugins/pinia";
import {mapStores} from "pinia";

export default defineComponent({
  name: "Copy",
  props: {
    text: {
      type: String,
      required: true
    }
  },
  computed: {
    ...mapStores(useSysStore),
  },
  methods: {
    async click() {
      await copyToClipboard(this.text)
      this.sysStore.showSnackBar("Скопировано в буфер", false)
    },
  },
})


</script>


<style>


</style>

