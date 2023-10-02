<template>
  <v-icon class="mx-1" icon="mdi-eye" @click="click" v-if="!show"/>
  <div v-if="show" class="mx-1"><i>{{ addr }}</i>
    <Copy :text="addr"/>
  </div>
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
import {formatTime} from "../helper";
import Copy from "@/components/profile/Copy.vue";

export default defineComponent({
  name: "ProfileAddress",
  components: {Copy, TopUpProfile, Withdraw, TopUp, Balance},
  props: {
    modelValue: {
      type: Boolean,
      required: true
    },
    addr: {
      type: String,
      required: true
    }
  },
  methods: {
    async click() {
      this.show = true
    },
  },
  emits: ['update:modelValue'],
  computed: {
    show: {
      get() {
        return this.modelValue
      },
      set(v: boolean) {
        this.$emit('update:modelValue', v)
      }
    }
  },
})


</script>


<style>


</style>

