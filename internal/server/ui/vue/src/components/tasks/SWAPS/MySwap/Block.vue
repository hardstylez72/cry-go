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
import {Task, TaskType, Token} from "@/generated/flow";
import {taskProps} from "@/components/tasks/tasks";
import {SwapPair, tokenSwapPair} from "@/components/helper";
import DefaultSwapTask from "@/components/tasks/SWAPS/DefaultSwapTask.js";
import {Component} from "vue-facing-decorator";
import {Network} from "@/generated/process";

@Component({name: 'MySwap'})
export default class MySwap extends DefaultSwapTask {

  getTask(): Task {
    const taskType = TaskType.MySwap
    const task = {
      weight: this.weight.toString(),
      mySwapTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.mySwapTask) {
        this.pairs = this.taskProps[this.task.taskType].swapParis
        this.item = this.task.mySwapTask
        this.network = this.item.network
        if (this.item.network && this.item.fromToken && this.item.toToken) {
          this.pair = tokenSwapPair(this.item.network, this.item.fromToken, this.item.toToken)
        }
        this.$emit('taskChanged', this.getTask())
      }
    }
  }

  created() {
    if (this.task?.mySwapTask) {
      this.item = this.task.mySwapTask
    }
  }

}


</script>

<style scoped>

</style>
