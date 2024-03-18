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
  name: 'TaskACoreDaoBridgeBlock',
  components: {NetworkSelector}
})

export default class TaskACoreDaoBridgeBlock extends TaskDefaultBridge {
  pairs = [

    // from ARBITRUM
    tokenBridgePair(Network.ARBITRUM, Network.OPTIMISM, Token.ETH, Token.ETH),
    tokenBridgePair(Network.ARBITRUM, Network.OPTIMISM, Token.STG, Token.STG),
    tokenBridgePair(Network.ARBITRUM, Network.OPTIMISM, Token.USDCBridged, Token.USDCBridged),
    tokenBridgePair(Network.ARBITRUM, Network.Base, Token.ETH, Token.ETH),

    tokenBridgePair(Network.ARBITRUM, Network.POLIGON, Token.STG, Token.STG),
    tokenBridgePair(Network.ARBITRUM, Network.POLIGON, Token.USDCBridged, Token.USDC),
    tokenBridgePair(Network.ARBITRUM, Network.POLIGON, Token.USDT, Token.USDT),

    tokenBridgePair(Network.ARBITRUM, Network.BinanaceBNB, Token.STG, Token.STG),
    tokenBridgePair(Network.ARBITRUM, Network.BinanaceBNB, Token.USDT, Token.USDT),

    tokenBridgePair(Network.ARBITRUM, Network.AVALANCHE, Token.STG, Token.STG),
    tokenBridgePair(Network.ARBITRUM, Network.AVALANCHE, Token.USDT, Token.USDT),
    tokenBridgePair(Network.ARBITRUM, Network.AVALANCHE, Token.USDCBridged, Token.USDC),

    // from BinanaceBNB
    tokenBridgePair(Network.BinanaceBNB, Network.AVALANCHE, Token.STG, Token.STG),
    tokenBridgePair(Network.BinanaceBNB, Network.AVALANCHE, Token.USDT, Token.USDT),

    tokenBridgePair(Network.BinanaceBNB, Network.POLIGON, Token.STG, Token.STG),
    tokenBridgePair(Network.BinanaceBNB, Network.POLIGON, Token.USDT, Token.USDT),

    tokenBridgePair(Network.BinanaceBNB, Network.ARBITRUM, Token.STG, Token.STG),
    tokenBridgePair(Network.BinanaceBNB, Network.ARBITRUM, Token.USDT, Token.USDT),

    tokenBridgePair(Network.BinanaceBNB, Network.OPTIMISM, Token.STG, Token.STG),

    //from OPTIMISM

    tokenBridgePair(Network.OPTIMISM, Network.BinanaceBNB, Token.STG, Token.STG),

    tokenBridgePair(Network.OPTIMISM, Network.AVALANCHE, Token.STG, Token.STG),
    tokenBridgePair(Network.OPTIMISM, Network.AVALANCHE, Token.USDCBridged, Token.USDC),

    tokenBridgePair(Network.OPTIMISM, Network.POLIGON, Token.STG, Token.STG),
    tokenBridgePair(Network.OPTIMISM, Network.POLIGON, Token.USDCBridged, Token.USDC),

    tokenBridgePair(Network.OPTIMISM, Network.ARBITRUM, Token.STG, Token.STG),
    tokenBridgePair(Network.OPTIMISM, Network.ARBITRUM, Token.USDCBridged, Token.USDCBridged),
    tokenBridgePair(Network.OPTIMISM, Network.ARBITRUM, Token.ETH, Token.ETH),
    tokenBridgePair(Network.OPTIMISM, Network.Base, Token.ETH, Token.ETH),

    // from AVALANCHE
    tokenBridgePair(Network.AVALANCHE, Network.BinanaceBNB, Token.STG, Token.STG),
    tokenBridgePair(Network.AVALANCHE, Network.BinanaceBNB, Token.USDT, Token.USDT),

    tokenBridgePair(Network.AVALANCHE, Network.POLIGON, Token.USDT, Token.USDT),
    tokenBridgePair(Network.AVALANCHE, Network.POLIGON, Token.USDC, Token.USDC),
    tokenBridgePair(Network.AVALANCHE, Network.POLIGON, Token.STG, Token.STG),

    tokenBridgePair(Network.AVALANCHE, Network.ARBITRUM, Token.USDT, Token.USDT),
    tokenBridgePair(Network.AVALANCHE, Network.ARBITRUM, Token.USDC, Token.USDCBridged),
    tokenBridgePair(Network.AVALANCHE, Network.ARBITRUM, Token.STG, Token.STG),

    tokenBridgePair(Network.AVALANCHE, Network.OPTIMISM, Token.USDC, Token.USDCBridged),
    tokenBridgePair(Network.AVALANCHE, Network.OPTIMISM, Token.STG, Token.STG),

    // from POLIGON
    tokenBridgePair(Network.POLIGON, Network.BinanaceBNB, Token.STG, Token.STG),
    tokenBridgePair(Network.POLIGON, Network.BinanaceBNB, Token.USDC, Token.USDC),
    tokenBridgePair(Network.POLIGON, Network.BinanaceBNB, Token.USDT, Token.USDT),

    tokenBridgePair(Network.POLIGON, Network.AVALANCHE, Token.STG, Token.STG),
    tokenBridgePair(Network.POLIGON, Network.AVALANCHE, Token.USDC, Token.USDC),
    tokenBridgePair(Network.POLIGON, Network.AVALANCHE, Token.USDT, Token.USDT),

    tokenBridgePair(Network.POLIGON, Network.ARBITRUM, Token.STG, Token.STG),
    tokenBridgePair(Network.POLIGON, Network.ARBITRUM, Token.USDC, Token.USDCBridged),
    tokenBridgePair(Network.POLIGON, Network.ARBITRUM, Token.USDT, Token.USDT),

    tokenBridgePair(Network.POLIGON, Network.OPTIMISM, Token.STG, Token.STG),
    tokenBridgePair(Network.POLIGON, Network.OPTIMISM, Token.USDC, Token.USDCBridged),


    tokenBridgePair(Network.POLIGON, Network.Fantom, Token.STG, Token.STG),
    tokenBridgePair(Network.POLIGON, Network.Base, Token.STG, Token.STG),

    tokenBridgePair(Network.Fantom, Network.Base, Token.STG, Token.STG),
    tokenBridgePair(Network.Fantom, Network.POLIGON, Token.STG, Token.STG),

    tokenBridgePair(Network.Base, Network.Fantom, Token.STG, Token.STG),
    tokenBridgePair(Network.Base, Network.POLIGON, Token.STG, Token.STG),
    tokenBridgePair(Network.Base, Network.ARBITRUM, Token.ETH, Token.ETH),
    tokenBridgePair(Network.Base, Network.OPTIMISM, Token.ETH, Token.ETH),

  ]

  getTask(): Task {
    return {
      weight: this.weight.toString(),
      stargateBridge: this.item,
      taskType: TaskType.StargateBridge,
      description: JSON.stringify(this.item),
    }
  }

  syncTask() {
    if (this.task) {
      if (this.task.stargateBridge) {
        this.item = this.task.stargateBridge

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
