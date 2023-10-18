<template>
  <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <NetworkSelector
            label="network"
            :items="networks"
            :disabled="disabled"
            v-model="item.network"
          />
        </v-col>
        <v-col>
          <v-autocomplete
            density="compact"
            variant="outlined"
            label="direction"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="pairs"
            v-model="pair"
            :disabled="disabled"
            item-title="name"
            return-object
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
import {DefaultSwap, Network, Task, TaskType, Token} from "@/generated/flow";
import {taskProps} from "@/components/tasks/tasks";
import {SwapPair, tokenSwapPair} from "@/components/helper";
import DefaultSwapTask from "@/components/tasks/block/base/DefaultSwapTask.js";
import {Component} from "vue-facing-decorator";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";

@Component({
  name: 'TraderJoeSwap',
  components: {NetworkSelector}
})
export default class TraderJoeSwap extends DefaultSwapTask {

  networks = [Network.ARBITRUM]

  item: DefaultSwap = {
    network: Network.ARBITRUM,
    amount: {
      sendAll: true,
    },
    toToken: Token.USDT,
    fromToken: Token.USDC,
  }
  pairs: SwapPair[] = [
    tokenSwapPair(Network.ARBITRUM, Token.STG, Token.ETH),

    tokenSwapPair(Network.ARBITRUM, Token.ETH, Token.USDT),
    tokenSwapPair(Network.ARBITRUM, Token.USDT, Token.ETH),

    tokenSwapPair(Network.ARBITRUM, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ARBITRUM, Token.USDC, Token.ETH),

    tokenSwapPair(Network.ARBITRUM, Token.USDT, Token.USDC),
    tokenSwapPair(Network.ARBITRUM, Token.USDC, Token.USDT),

    tokenSwapPair(Network.ARBITRUM, Token.USDC, Token.USDCBridged),
    tokenSwapPair(Network.ARBITRUM, Token.USDCBridged, Token.USDC),

    tokenSwapPair(Network.ARBITRUM, Token.ETH, Token.USDCBridged),
    tokenSwapPair(Network.ARBITRUM, Token.USDCBridged, Token.ETH),
  ]

  created() {
    if (this.task?.traderJoeSwapTask) {
      this.item = this.task.traderJoeSwapTask
    }
  }

  getTask(): Task {
    const taskType = TaskType.TraderJoeSwap
    const task = {
      weight: this.weight.toString(),
      traderJoeSwapTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.traderJoeSwapTask) {
        this.item = this.task.traderJoeSwapTask
        this.pair = tokenSwapPair(this.item.network, this.item.fromToken, this.item.toToken)
        this.$emit('taskChanged', this.getTask())
      }
    }
  }
}


</script>

<style scoped>

</style>
