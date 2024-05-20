<template>
  <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <v-select
            density="compact"
            variant="outlined"
            label="from"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="networks"
            v-model="item.fromNetwork"
            :disabled="true"
          />
          <v-select
            density="compact"
            variant="outlined"
            label="to"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="networks"
            v-model="item.toNetwork"
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
            :items="tokens"
            v-model="item.token"
            :disabled="disabled"
            item-title="name"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <AmountInput :coin="item.token" :disabled="disabled" v-model="item.amount"/>
        </v-col>
      </v-row>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Task, TaskType, Token} from "@/generated/flow";
import {taskProps} from "@/components/tasks/tasks";
import StrarkNetBridgeTask from "@/components/tasks/BRIDGE/DefaultLiquidityBridge.js";
import {Component} from "vue-facing-decorator";

@Component({name: 'StrarkNetBridge'})
export default class StrarkNetBridge extends StrarkNetBridgeTask {

  created() {
    if (this.task?.starkNetBridgeTask) {
      this.item = this.task.starkNetBridgeTask
    }
  }

  getTask(): Task {
    const taskType = TaskType.StarkNetBridge
    const task = {
      weight: this.weight.toString(),
      starkNetBridgeTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.starkNetBridgeTask) {
        this.item = this.task.starkNetBridgeTask
        this.$emit('taskChanged', this.getTask())
      }
    }
  }
}


</script>

<style scoped>

</style>
