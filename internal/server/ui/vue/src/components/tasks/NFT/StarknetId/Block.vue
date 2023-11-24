<template>
  <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <v-select
            density="compact"
            variant="outlined"
            label="network"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="networks"
            v-model="item.network"
            :disabled="disabled"
          />
        </v-col>
      </v-row>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Network, SimpleTask, Task, TaskType} from "@/generated/flow";
import {taskProps} from "@/components/tasks/tasks";
import {Component} from "vue-facing-decorator";
import DefaultSimple from "@/components/tasks/OTHER/DefaultSimple";

@Component({name: 'StarkNetId'})
export default class StarkNetId extends DefaultSimple {

  networks = [Network.StarkNet]

  item: SimpleTask = {
    network: Network.StarkNet,
  }

  created() {
    if (this.task?.starkNetIdMintTask) {
      this.item = this.task.starkNetIdMintTask
    }
  }

  getTask(): Task {
    const taskType = TaskType.StarkNetIdMint
    const task = {
      weight: this.weight.toString(),
      starkNetIdMintTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.dmailTask) {
        this.item = this.task.dmailTask
        this.$emit('taskChanged', this.getTask())
      }
    }
  }
}


</script>

<style scoped>

</style>
