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

@Component({name: 'SyncSwap'})
export default class SyncSwap extends DefaultSwapTask {

  pairs: SwapPair[] = [
    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.LUSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.LUSD, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.LSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.LSD, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.MUTE),
    tokenSwapPair(Network.ZKSYNCERA, Token.MUTE, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.USDT, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.USDT),
  ]

  created() {
    if (this.task?.syncSwapTask) {
      this.item = this.task.syncSwapTask
    }
  }

  getTask(): Task {
    const taskType = TaskType.SyncSwap
    const task = {
      weight: this.weight.toString(),
      syncSwapTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.syncSwapTask) {
        this.item = this.task.syncSwapTask
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
