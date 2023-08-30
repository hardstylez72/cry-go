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
import DefaultSimple from "@/components/tasks/block/base/DefaultSimple";

@Component({name: 'Dmail'})
export default class Dmail extends DefaultSimple {

  networks = [Network.ZKSYNCERA, Network.StarkNet]

  item: SimpleTask = {
    network: Network.ZKSYNCERA,
  }

  created() {
    if (this.task?.dmailTask) {
      this.item = this.task.dmailTask
    }
  }

  getTask(): Task {
    const taskType = TaskType.Dmail
    const task = {
      weight: this.weight.toString(),
      dmailTask: this.item,
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
