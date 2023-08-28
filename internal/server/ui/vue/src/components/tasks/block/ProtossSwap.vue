<template>
  <v-card-actions>
    <v-container>
      <b style="color:red;">Не вздумайте гонять обьемы, проскальзывание 50%. это означает что половину от обьема вы
        подарите этой платформе</b>
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
import {Task, TaskType, Token} from "@/generated/flow";
import {taskProps} from "@/components/tasks/tasks";
import {SwapPair, tokenSwapPair} from "@/components/helper";
import DefaultSwapTask from "@/components/tasks/block/base/DefaultSwapTask.js";
import {Component} from "vue-facing-decorator";

@Component({name: 'ProtossSwap'})
export default class ProtossSwap extends DefaultSwapTask {

  pairs: SwapPair[] = [
    tokenSwapPair(Token.ETH, Token.USDC),
    tokenSwapPair(Token.USDC, Token.ETH),
  ]

  getTask(): Task {
    const taskType = TaskType.ProtossSwap
    const task = {
      weight: this.weight.toString(),
      protosSwapTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.protosSwapTask) {
        this.item = this.task.protosSwapTask
        this.pair = tokenSwapPair(this.item.fromToken, this.item.toToken)
        this.$emit('taskChanged', this.getTask())
      }
    }
  }

  created() {
    if (this.task?.protosSwapTask) {
      this.item = this.task.protosSwapTask
    }
  }
}


</script>

<style scoped>

</style>
