<template>
  <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <NetworkSelector
            label="from network"
            :items="GetFromNetworks"
            :disabled="disabled"
            v-model="network"
          />
        </v-col>
        <v-col>
          <v-autocomplete
            density="compact"
            variant="outlined"
            label="direction"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="getPairs"
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

@Component({name: 'OdosSwap'})
export default class OdosSwap extends DefaultSwapTask {


  pairs: SwapPair[] = [

    tokenSwapPair(Network.Base, Token.ETH, Token.WETH),
    tokenSwapPair(Network.Base, Token.WETH, Token.ETH),

    tokenSwapPair(Network.Base, Token.ETH, Token.USDC),
    tokenSwapPair(Network.Base, Token.USDC, Token.ETH),

    tokenSwapPair(Network.Base, Token.ETH, Token.USDCBridged),
    tokenSwapPair(Network.Base, Token.USDCBridged, Token.ETH),

    tokenSwapPair(Network.Base, Token.USDC, Token.USDCBridged),
    tokenSwapPair(Network.Base, Token.USDCBridged, Token.USDC),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.BUSD, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.BUSD),

    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.USDT),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDT, Token.USDC),
  ]

  created() {
    if (this.task?.odosSwapTask) {
      this.item = this.task.odosSwapTask
    }
  }

  getTask(): Task {
    const taskType = TaskType.OdosSwap
    const task = {
      weight: this.weight.toString(),
      odosSwapTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.odosSwapTask) {
        this.item = this.task.odosSwapTask
        this.network = this.item.network
        if (this.item.network && this.item.fromToken && this.item.toToken) {
          this.pair = tokenSwapPair(this.item.network, this.item.fromToken, this.item.toToken)
        }
        this.$emit('taskChanged', this.getTask())
      }
    }
  }
}


</script>

<style scoped>

</style>
