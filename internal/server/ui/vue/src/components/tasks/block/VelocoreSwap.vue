<template>
  <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <v-select
            ref="stargate-bridge-form"
            density="compact"
            variant="outlined"
            label="network"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="networks"
            v-model="item.network"
            :disabled="true"
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

@Component({name: 'VelocoreSwap'})
export default class VelocoreSwap extends DefaultSwapTask {

  networks = [Network.ZKSYNCERA]

  item: DefaultSwap = {
    network: Network.ZKSYNCERA,
    amount: {
      sendAll: true,
    },
    toToken: Token.USDC,
    fromToken: Token.ETH,
  }
  pairs: SwapPair[] = [
    tokenSwapPair(Token.ETH, Token.USDC),
    tokenSwapPair(Token.USDC, Token.ETH),

    tokenSwapPair(Token.ETH, Token.VC),
    tokenSwapPair(Token.VC, Token.ETH),
  ]

  created() {
    if (this.task?.velocoreSwapTask) {
      this.item = this.task.velocoreSwapTask
    }
  }

  getTask(): Task {
    const taskType = TaskType.VelocoreSwap
    const task = {
      weight: this.weight.toString(),
      velocoreSwapTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.velocoreSwapTask) {
        this.item = this.task.velocoreSwapTask
        this.pair = tokenSwapPair(this.item.fromToken, this.item.toToken)
        this.$emit('taskChanged', this.getTask())
      }
    }
  }
}


</script>

<style scoped>

</style>
