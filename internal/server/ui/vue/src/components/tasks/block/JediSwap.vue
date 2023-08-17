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
import {DefaultSwap, MaverickSwapTask, Network, SpaceFiSwapTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import WEIInputField from "@/components/WEIInputField.vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {link} from "@/components/tasks/links";
import {taskProps} from "@/components/tasks/tasks";
import {required} from "@/components/tasks/menu/helper";
import {SwapPair, tokenSwapPair} from "@/components/helper";
import DefaultSwapMenu from "@/components/tasks/menu/DefaultSwapMenu.js";
import TaskDefaultSwap from "@/components/tasks/block/TaskDefaultSwap.js";
import {Component, toNative} from "vue-facing-decorator";

@Component({name: 'JediSwap'})
export default class JediSwap extends TaskDefaultSwap {

  pairs: SwapPair[] = [
    tokenSwapPair(Token.ETH, Token.USDC),
    tokenSwapPair(Token.USDC, Token.ETH),
  ]

  getTask(): Task {
    const taskType = TaskType.JediSwap
    const task = {
      weight: this.weight.toString(),
      jediSwapTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.jediSwapTask) {
        this.item = this.task.jediSwapTask
        this.pair = tokenSwapPair(this.item.fromToken, this.item.toToken)
        this.$emit('taskChanged', this.getTask())
      }
    }
  }

  created() {
    if (this.task?.jediSwapTask) {
      this.item = this.task.jediSwapTask
    }
  }

}


</script>

<style scoped>

</style>
