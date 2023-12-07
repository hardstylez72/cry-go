<template>
  <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <NetworkSelector
            label="from network"
            :items="GetFromNetworks"
            :disabled="disabled"
            v-model="fromNetwork"
          />
        </v-col>
        <v-col>
          <NetworkSelector
            label="to network"
            :items="GetToNetworks"
            :disabled="disabled"
            v-model="toNetwork"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-autocomplete
            return-object
            density="compact"
            variant="outlined"
            v-on:change="inputChanged"
            label="direction"
            :items="getPairs"
            :rules="[required]"
            v-model="pair"
            :disabled="disabled"
            item-title="name"
          />
        </v-col>
      </v-row>
      <div class="my-2" v-if="pair && item.fromToken && item.toToken">{{
          `Swap possible - ${swapOpt.error ? "NO" : `YES [fee ${swapOpt.fee} USD min: ${swapOpt.min} max: ${swapOpt.max}]`}`
        }}
      </div>
      <AmountInput :coin="item.fromToken" :disabled="disabled" v-model="item.amount"/>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Network, Task, TaskType, Token} from "@/generated/flow";
import {orbiterService} from "@/generated/services";
import {GetSwapOptionsRes} from "@/generated/orbiter";
import {BridgePair, tokenBridgePair} from "@/components/helper";
import TaskDefaultBridge from "@/components/tasks/BRIDGE/DefaultBridgeTask";
import {Component, Watch} from "vue-facing-decorator";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";

@Component({
  name: 'TaskOrbiterBridgeBlock',
  components: {NetworkSelector}
})
export default class TaskOrbiterBridgeBlock extends TaskDefaultBridge {
  pairs = [
    tokenBridgePair(Network.ZKSYNCERA, Network.ARBITRUM, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ZKSYNCERA, Network.OPTIMISM, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ZKSYNCERA, Network.Base, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ZKSYNCERA, Network.Etherium, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ZKSYNCERA, Network.ArbitrumNova, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ZKSYNCERA, Network.ZKSYNCLITE, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ZKSYNCERA, Network.Linea, Token.ETH, Token.ETH),

    tokenBridgePair(Network.Base, Network.ARBITRUM, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Base, Network.OPTIMISM, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Base, Network.ZKSYNCERA, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Base, Network.Etherium, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Base, Network.ArbitrumNova, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Base, Network.ZKSYNCLITE, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Base, Network.Linea, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Base, Network.Zora, Token.ETH, Token.ETH),

    tokenBridgePair(Network.ARBITRUM, Network.Base, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ARBITRUM, Network.OPTIMISM, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ARBITRUM, Network.ZKSYNCERA, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ARBITRUM, Network.Etherium, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ARBITRUM, Network.ArbitrumNova, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ARBITRUM, Network.ZKSYNCLITE, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ARBITRUM, Network.Linea, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ARBITRUM, Network.Zora, Token.ETH, Token.ETH),

    tokenBridgePair(Network.OPTIMISM, Network.Base, Token.ETH, Token.ETH),
    tokenBridgePair(Network.OPTIMISM, Network.ARBITRUM, Token.ETH, Token.ETH),
    tokenBridgePair(Network.OPTIMISM, Network.ZKSYNCERA, Token.ETH, Token.ETH),
    tokenBridgePair(Network.OPTIMISM, Network.Etherium, Token.ETH, Token.ETH),
    tokenBridgePair(Network.OPTIMISM, Network.ArbitrumNova, Token.ETH, Token.ETH),
    tokenBridgePair(Network.OPTIMISM, Network.ZKSYNCLITE, Token.ETH, Token.ETH),
    tokenBridgePair(Network.OPTIMISM, Network.Linea, Token.ETH, Token.ETH),

    tokenBridgePair(Network.Linea, Network.Base, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Linea, Network.ARBITRUM, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Linea, Network.ZKSYNCERA, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Linea, Network.Etherium, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Linea, Network.OPTIMISM, Token.ETH, Token.ETH),


    tokenBridgePair(Network.Zora, Network.Base, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Zora, Network.ARBITRUM, Token.ETH, Token.ETH),

  ] as BridgePair[]

  swapOpt = {min: "", max: "", fee: '', error: 'fill up the form above'} as GetSwapOptionsRes
  swapOptLoading = false


  @Watch('pair', {
    deep: true,
  })
  pairHandler(a, b) {
    if (!this.pair) {
      this.item.fromToken = undefined
      this.item.toToken = undefined
      return
    }
    this.item.fromToken = this.pair.from
    this.item.toToken = this.pair.to
    this.item.fromNetwork = this.pair.fromNetwork
    this.item.toNetwork = this.pair.toNetwork
    this.fromNetwork = this.pair.fromNetwork
    this.toNetwork = this.pair.toNetwork
    this.getSwapOpt()
  }

  async getSwapOpt() {
    if (!this.pair) {
      return
    }
    if (!this.item.fromToken) {
      return
    }
    if (!this.item.toToken) {
      return
    }

    try {
      this.swapOptLoading = true
      const res = await orbiterService.orbiterServiceGetSwapOptions({
        body: {
          fromNetwork: this.item.fromNetwork,
          toNetwork: this.item.toNetwork,
          fromToken: this.item.fromToken,
          toToken: this.item.toToken,
        }
      })

      this.swapOpt.max = res.max
      this.swapOpt.min = res.min
      this.swapOpt.fee = res.fee
      this.swapOpt.error = res.error

    } finally {
      this.swapOptLoading = false
    }
  }

  getTask(): Task {
    return {
      weight: this.weight.toString(),
      orbiterBridgeTask: this.item,
      taskType: TaskType.OrbiterBridge,
      description: JSON.stringify(this.item),
    }
  }

  syncTask() {
    if (this.task && this.task.orbiterBridgeTask) {
      const t = this.task.orbiterBridgeTask
      this.item = {
        fromNetwork: t.fromNetwork,
        received: true,
        toNetwork: t.toNetwork,
        amount: t.amount,
        toToken: t.toToken,
        fromToken: t.fromToken,
      }

      if (this.item.fromNetwork && this.item.toNetwork && this.item.fromToken && this.item.toToken) {
        this.pair = tokenBridgePair(this.item.fromNetwork, this.item.toNetwork, this.item.fromToken, this.item.toToken)
      }

      this.$emit('taskChanged', this.getTask())
    }
  }
}
</script>

<style scoped>

</style>

