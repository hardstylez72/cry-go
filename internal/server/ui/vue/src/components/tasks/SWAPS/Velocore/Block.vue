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
import {DefaultSwap, Task, TaskType, Token} from "@/generated/flow";
import {taskProps} from "@/components/tasks/tasks";
import {SwapPair, tokenSwapPair} from "@/components/helper";
import DefaultSwapTask from "@/components/tasks/SWAPS/DefaultSwapTask.js";
import {Component} from "vue-facing-decorator";
import {Network} from "@/generated/process";

@Component({name: 'VelocoreSwap'})
export default class VelocoreSwap extends DefaultSwapTask {

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
        this.pairs = this.taskProps[this.task.taskType].swapParis
        this.item = this.task.velocoreSwapTask
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
