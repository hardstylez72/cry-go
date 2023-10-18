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

@Component({name: 'IzumiSwap'})
export default class IzumiSwap extends DefaultSwapTask {
  pairs: SwapPair[] = [
    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.IZI),
    tokenSwapPair(Network.ZKSYNCERA, Token.IZI, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.WETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.WETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.WETH, Token.IZI),
    tokenSwapPair(Network.ZKSYNCERA, Token.IZI, Token.WETH),
  ]

  created() {
    if (this.task?.izumiSwapTask) {
      this.item = this.task.izumiSwapTask
    }
  }

  getTask(): Task {
    const taskType = TaskType.IzumiSwap
    const task = {
      weight: this.weight.toString(),
      izumiSwapTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.izumiSwapTask) {
        this.item = this.task.izumiSwapTask
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
