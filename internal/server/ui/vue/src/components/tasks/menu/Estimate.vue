<template>
  <v-menu
    v-model="showTooltip"
    width="350px"
    persistent
    :close-on-back="false"
    :close-on-content-click="false"
  >
    <template v-slot:activator="{ props }">
      <v-btn
        class="ml-1"
        :rounded="false"
        density="compact"
        v-bind="props"
      >
        Estimate
      </v-btn>
    </template>

    <v-card>
      <v-card-title class="d-flex justify-space-between">
        <div>Market Estimation</div>
        <v-icon icon="mdi-close" @click="showTooltip = false"/>
      </v-card-title>
      <Loader v-if="loading"/>
      <v-card-text v-else style="white-space: pre-wrap">
        <div v-if="errorMessage" style="color: red">{{ errorMessage }}</div>
        <div v-else v-for="estimation in estimations" class="d-flex justify-space-between elevation-1 px-3 py-3">
          <div><b>{{ estimation.name }}</b><br/></div>
          <div>
            <div><b>gas</b>: {{ getBalance(estimation.gas, true) }}</div>
            <div v-for="d in estimation.details">
              <b>{{ d.key }}</b> {{ Number(d.value.split(" ")[0]).toPrecision(3) + " " + d.value.split(" ")[1] }}
            </div>
            <div><b>gas limit</b>: {{ getBalance(estimation.gasLimit, false) }}</div>
            <div><b>gas price</b>: {{ getBalance(estimation.gasPrice, false) }}</div>
          </div>
        </div>
      </v-card-text>
    </v-card>
  </v-menu>
</template>

<script lang="ts">
import {AmUni, Network, Token} from "@/generated/flow";
import {defineComponent} from "vue";
import {processService} from "@/generated/services";
import {EstimationTx} from "@/generated/process";
import Loader from "@/components/Loader.vue";

export default defineComponent({
  name: "EstimateTask",
  components: {Loader},
  props: {
    taskId: {
      type: String,
      required: true,
    },
    profileId: {
      type: String,
      required: true,
    },
    processId: {
      type: String,
      required: true,
    },
  },
  watch: {
    showTooltip: {
      handler() {
        if (this.showTooltip) {
          this.estimate()
        }
      }
    }
  },
  data() {
    return {
      errorMessage: '',
      loading: false,
      showTooltip: false,
      estimations: {} as EstimationTx[]
    }
  },
  methods: {
    getBalance(am?: AmUni, gas?: boolean): string {
      if (!am) {
        return ''
      }

      if (gas) {
        let token = Token.ETH
        switch (am.network) {
          case Network.BinanaceBNB:
            token = Token.BNB
            break
          case Network.POLIGON:
            token = Token.MATIC
            break
          case  Network.AVALANCHE:
            token = Token.AVAX
            break
        }
        return `${Number(am.usd).toPrecision(3)} USD`
      }
      return `${Number(am.gwei).toPrecision(3)} GWEI`
    },
    async estimate() {
      try {
        this.loading = true
        this.errorMessage = ''
        const res = await processService.processServiceEstimateCost({
          body: {
            profileId: this.profileId,
            taskId: this.taskId,
            processId: this.processId,
          }
        })
        if (res.error) {
          this.errorMessage = res.error
        } else if (res.estimations) {
          this.estimations = res.estimations
        }
      } catch (e) {
        this.errorMessage = 'error estimating'
      } finally {
        this.loading = false
      }
    }
  },
})
</script>

<style scoped>

</style>
