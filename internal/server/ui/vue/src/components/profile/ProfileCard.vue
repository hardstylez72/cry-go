<template>
  <v-menu
    v-model="menu"
    :close-on-content-click="false"
    location="end"
  >

    <template v-slot:activator="{ props }">
      <v-btn variant="outlined" @click="load" density="comfortable" v-bind="props">{{ label }}
      </v-btn>
    </template>

    <template v-slot:default="{ props }">
      <v-card class="px-4 py-4">

        <div class="d-flex justify-space-between">
          <div>
            <i class="mx-2">{{ profile.subType }}</i>
            <i class="mx-2" v-if="profile.type === ProfileType.StarkNet &&!deployed && !loading">not deployed</i>
            <v-icon icon="mdi-eye" @click="addrShow = true" v-if="!addrShow"/>
          </div>
          <v-icon icon="mdi-close" @click="menu = false"/>
        </div>

        <div v-if="addrShow" class="my-2">{{ profile.mmskId }}</div>

        <div class="d-flex justify-space-between" v-if="this.profile.type === ProfileType.StarkNet">
          <div>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.StarkNet"/>
            </v-card-item>
          </div>
        </div>
        <div class="d-flex justify-space-between" v-if="this.profile.type === ProfileType.EVM">
          <div>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.ARBITRUM"/>
            </v-card-item>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.Etherium"/>
            </v-card-item>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.OPTIMISM"/>
            </v-card-item>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.BinanaceBNB"/>
            </v-card-item>
          </div>
          <div>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.POLIGON"/>
            </v-card-item>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.AVALANCHE"/>
            </v-card-item>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.ZKSYNCLITE"/>
            </v-card-item>
            <v-card-item>
              <Balance :profile-id="profileId" :network="Network.ZKSYNCERA"/>
            </v-card-item>
          </div>
        </div>
        <div>
          <div v-if="transactions.length > 0">
            Recent Transactions:
            <div v-for="(tx, i) in transactions">
              <div>
                <a class="mr-1" :href="tx.url" target="_blank"> <b> {{ tx.code }}</b></a>
                <a :href="`/process/${tx.processId}`" target="_blank"> {{ formatTime(tx.createdAt) }}</a>
              </div>
            </div>
          </div>
        </div>
        <v-card-actions class="d-flex justify-end">
          <a class="pr-5" :href="`https://debank.com/profile/${profile.mmskId}`" target="_blank">Debank</a>
          <TopUpProfile :profile-id="profileId"/>
        </v-card-actions>
      </v-card>
    </template>
  </v-menu>

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

export default defineComponent({
  name: "ProfileCard",
  components: {TopUpProfile, Withdraw, TopUp, Balance},
  props: {
    profileId: {
      type: String,
      required: true
    },
    label: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      loading: false,
      deployed: false,
      addrShow: false,
      profile: {} as Profile,
      menu: false,
      transactionsLoading: false,
      transactions: [] as Transaction[]
    }
  },
  computed: {
    ProfileType() {
      return ProfileType
    },
    Network() {
      return Network
    }
  },
  methods: {
    async load() {
      this.loadProfile()
      this.loadTransactions()

    },
    formatTime,
    async loadTransactions() {
      try {
        this.transactionsLoading = true
        const res = await processService.processServiceGetProfileTransactions({body: {profileId: this.profileId}})
        this.transactions = res.transactions
      } finally {
        this.transactionsLoading = false
      }
    },
    async loadProfile() {
      this.loading = true
      try {
        const res = await profileService.profileServiceGetProfile({body: {profileId: this.profileId}})
        this.profile = res.profile

        const deployed = await profileService.profileServiceStarkNetAccountDeployed({body: {profileId: this.profileId}})
        this.deployed = deployed.deployed
      } finally {
        this.loading = false
      }

    }
  },
})


</script>


<style>


</style>

