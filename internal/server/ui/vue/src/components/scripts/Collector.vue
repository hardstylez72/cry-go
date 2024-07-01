<template>
  <v-container>
    Куда скидывать
    <div class="mx-8 my-3" style="width: 80vw">
      <ProfileSearch
          v-model="sendTo"
          :disabled="disabled"
      />
      <div class="ml-3 mb-2 d-inline-flex flex-wrap" v-for="profile in sendTo">
        <ProfileCard :label="profile.num" :profile-id="profile.id"/>
      </div>
    </div>

    Откуда скидывать
    <div class="mx-8 my-3" style="width: 80vw">
      <ProfileSearch
          v-model="sendFrom"
          :disabled="disabled"
      />
    </div>
    <div class="mx-8 my-3 d-inline-flex">
      <NetworkSelector class="mx-2" label="Сеть"
                       :items="[Network.StarkNet, Network.ARBITRUM, Network.ZKSYNCERA, Network.POLIGON]"
                       v-model="network"
                       :disabled="disabled"/>
      <TokenSelector class="mx-2" label="Монетка" :items="[Token.ETH, Token.ZRO, Token.MATIC]" v-model="token"
                     :disabled="disabled"/>
      <v-btn @click="collect" :disabled="disabled">Собирать дэньги</v-btn>
    </div>
    <div class="ml-3 mb-2" v-for="profile in sendFrom">
      <ProfileCard :label="profile.num" :profile-id="profile.id"/>
      {{ profile.meta }}
      <div></div>
    </div>


  </v-container>


</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {profileService} from "@/generated/services"
import {Network, Token} from "@/generated/flow";
import {Profile} from "@/generated/profile";
import ProfileCard from "@/components/profile/ProfileCard.vue";
import ProfileSearch from "@/components/profile/ProfileSearch.vue";
import TokenSelector from "@/components/tasks/TokenSelector.vue";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";

export default defineComponent({
  name: "Collector",
  components: {NetworkSelector, TokenSelector, ProfileCard, ProfileSearch},

  data() {
    return {
      sendTo: [] as Profile[],
      sendFrom: [] as Profile[],
      token: Token.ETH,
      network: Network.StarkNet,
      disabled: false,
    }
  },
  computed: {
    Token() {
      return Token
    },
    Network() {
      return Network
    },
    getProfileList() {
      return this.selectedProfiles.map(e => e)
    },
  },
  methods: {
    async collect() {
      try {
        this.disabled = true

        for (let i = 0; i < this.sendFrom.length; i++) {

          try {
            this.sendFrom[i].meta = "в процессе"
            const res = await profileService.profileServiceTransferP2P({
              body: {
                from: this.sendFrom[i].id,
                network: this.network,
                token: this.token,
                to: this.sendTo[0].id
              }
            })
            this.sendFrom[i].meta = res.result
          } catch (e) {

          }
        }
      } finally {
        this.disabled = false
      }
    }
  },
  async mounted() {

  }
})


</script>


<style scoped>

</style>

