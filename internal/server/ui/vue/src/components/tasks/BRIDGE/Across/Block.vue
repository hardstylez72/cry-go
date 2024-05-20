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
      <v-row>
        <v-col>
          <AmountInput :coin="item.fromToken" :disabled="disabled" v-model="item.amount"/>
        </v-col>
      </v-row>

    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Network, Task, TaskType, Token} from "@/generated/flow";
import {tokenBridgePair} from "@/components/helper";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";
import {Component} from "vue-facing-decorator";
import TaskDefaultBridge from "@/components/tasks/BRIDGE/DefaultBridgeTask";

@Component({
  name: 'TaskAcrossBridgeBlock',
  components: {NetworkSelector}
})

export default class TaskAcrossBridgeBlock extends TaskDefaultBridge {
  pairs = [
    tokenBridgePair(Network.ZKSYNCERA, Network.ARBITRUM, Token.USDC, Token.USDCBridged),
    tokenBridgePair(Network.ZKSYNCERA, Network.ARBITRUM, Token.USDT, Token.USDT),
    tokenBridgePair(Network.ZKSYNCERA, Network.ARBITRUM, Token.ETH, Token.ETH),

    tokenBridgePair(Network.ARBITRUM, Network.ZKSYNCERA, Token.USDCBridged, Token.USDC),
    tokenBridgePair(Network.ARBITRUM, Network.ZKSYNCERA, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ARBITRUM, Network.Base, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ARBITRUM, Network.Base, Token.USDCBridged, Token.USDCBridged),
    tokenBridgePair(Network.ARBITRUM, Network.ZKSYNCERA, Token.USDT, Token.USDT),

    tokenBridgePair(Network.OPTIMISM, Network.Base, Token.USDCBridged, Token.USDCBridged),
    tokenBridgePair(Network.OPTIMISM, Network.Base, Token.ETH, Token.ETH),
    tokenBridgePair(Network.OPTIMISM, Network.ARBITRUM, Token.USDCBridged, Token.USDCBridged),
    tokenBridgePair(Network.OPTIMISM, Network.ARBITRUM, Token.ETH, Token.ETH),


    tokenBridgePair(Network.Base, Network.ARBITRUM, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Base, Network.ARBITRUM, Token.USDCBridged, Token.USDCBridged),
  ]

  getTask(): Task {
    return {
      weight: this.weight.toString(),
      acrossBridgeTask: this.item,
      taskType: TaskType.AcrossBridge,
      description: JSON.stringify(this.item),
    }
  }

  syncTask() {
    if (this.task) {
      if (this.task.acrossBridgeTask) {
        this.item = this.task.acrossBridgeTask

        if (this.item.fromNetwork && this.item.toNetwork && this.item.fromToken && this.item.toToken) {
          this.pair = tokenBridgePair(this.item.fromNetwork, this.item.toNetwork, this.item.fromToken, this.item.toToken)
        }

        this.$emit('taskChanged', this.getTask())
      }
    }
  }
}
</script>

<style scoped>

</style>
