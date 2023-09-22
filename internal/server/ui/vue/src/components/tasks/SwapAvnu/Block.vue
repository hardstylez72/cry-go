<template>
  <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <NetworkSelector label="сеть" :items="networks" v-model="item.network" :disabled="true"/>
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
import {Task, TaskType, Token} from "@/generated/flow";
import {taskProps} from "@/components/tasks/tasks";
import {SwapPair, tokenSwapPair} from "@/components/helper";
import DefaultSwapTask from "@/components/tasks/block/base/DefaultSwapTask.js";
import {Component, toNative} from "vue-facing-decorator";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";

@Component({
  name: 'AvnuSwap',
  components: {NetworkSelector}
})
export default class AvnuSwap extends DefaultSwapTask {

  pairs: SwapPair[] = [
    tokenSwapPair(Token.ETH, Token.USDC),
    tokenSwapPair(Token.USDC, Token.ETH),

    tokenSwapPair(Token.ETH, Token.USDT),
    tokenSwapPair(Token.USDT, Token.ETH),

    tokenSwapPair(Token.USDC, Token.USDT),
    tokenSwapPair(Token.USDT, Token.USDC),
  ]

  getTask(): Task {
    const taskType = TaskType.AvnuSwap
    const task = {
      weight: this.weight.toString(),
      avnuSwapTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.avnuSwapTask) {
        this.item = this.task.avnuSwapTask
        this.pair = tokenSwapPair(this.item.fromToken, this.item.toToken)
        this.$emit('taskChanged', this.getTask())
      }
    }
  }

  created() {
    if (this.task?.avnuSwapTask) {
      this.item = this.task.avnuSwapTask
    }
  }

}


</script>

<style scoped>

</style>
