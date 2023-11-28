<template>
  <v-dialog
    v-model="dialog"
    width="800"
    :close-on-content-click="false"
    :close-on-back="false"
    :persistent="true"
  >
    <template v-slot:activator="{ props }">
      <v-btn density="comfortable" @click="dialog=true">Пополнить</v-btn>
    </template>

    <v-card>
      <v-card-title>Пополнение</v-card-title>
      <v-card-text>
        <v-form ref="form" class="d-flex justify-space-between my-4">
          <v-text-field
            class="mr-3"
            type="number"
            v-model="amInput"
            density="compact"
            variant="outlined"
            placeholder="введите количество USDT"
            :disabled="amInputDisabled"
            :rules="[onlyInteger,required]"
            suffix="USDT"
          />
          <v-btn v-if="!amInputDisabled" rounded="false" @click="createOrder">Next</v-btn>
        </v-form>
        <div v-if="amInputDisabled">
          <span>Send exact amount of <b>{{ amToSend }} <a :href="coinAddrUrl"
                                                          target="_blank">USDT</a></b>
            <br/>
            to address: <b>{{
                walletToSend
              }}</b> using Arbitrum Network</span>
          <div class="text-center my-10">
            <div v-if="finished" style="color: green" class="text-h6 font-weight-bold">Payment confirmed</div>
            <div v-else>
              Waiting for transaction confirmation
              <v-progress-linear
                color="deep-purple-accent-4"
                indeterminate
                rounded
                height="6"
              ></v-progress-linear>
            </div>
          </div>
        </div>

      </v-card-text>
      <v-card-actions class="d-flex justify-end">
        <v-btn v-if="finished" @click="cancel">Close</v-btn>
        <v-btn v-else @click="cancel">Cancel</v-btn>
      </v-card-actions>
    </v-card>

  </v-dialog>
</template>


<script lang="ts">

import {defineComponent} from 'vue';
import {helperService} from "@/generated/services";
import {onlyInteger, required} from "@/components/tasks/helper";
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";

export default defineComponent({
  name: "TopUp",
  created() {
  },
  computed: {
    ...mapStores(useUserStore),
  },
  data() {
    return {
      amInput: "",
      promo: "",
      dialog: false,
      amInputDisabled: false,
      orderCreating: false,

      coinAddrUrl: "",
      orderId: "",
      walletToSend: "",
      amToSend: 0,
      errorMsg: "",
      finished: false,
      pollerTimer: null as NodeJS.Timer | null
    }
  },
  methods: {
    required,
    onlyInteger,
    async createOrder() {
      try {

        if (!await this.validateForm()) {
          return
        }

        this.orderCreating = true
        const res = await helperService.helperServiceCreateOrder({body: {am: this.amInput}})
        this.amInputDisabled = true
        this.orderId = res.id
        this.walletToSend = res.toWallet
        this.coinAddrUrl = res.coinAddrUrl
        this.amToSend = res.am
        this.startPolling()
      } finally {
        this.orderCreating = false
      }
    },
    cancel() {
      this.stopPolling()
      this.amInput = ""
      this.amInputDisabled = false
      this.orderCreating = false
      this.errorMsg = ""
      this.finished = false
      this.coinAddrUrl = ''
      this.dialog = false
      this.pollerTimer = null
    },
    startPolling() {
      this.pollerTimer = setInterval(() => {
        this.checkOrder()
      }, 1000)

    },
    stopPolling() {
      if (this.pollerTimer) {
        clearInterval(this.pollerTimer)
      }
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
    async checkOrder() {
      try {
        const res = await helperService.helperServiceGetOrderStatus({body: {orderId: this.orderId}})
        if (res.status === 'Error') {
          this.errorMsg = 'error happened'
          this.stopPolling()

        } else if (res.status === 'Processed') {
          this.finished = true
          await this.userStore.syncUser()
          this.stopPolling()
        }
      } finally {

      }

    }
  }

})


</script>


<style>
</style>

