<template>
  <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <NetworkSelector label="сеть" :items="networks" v-model="item.network" :disabled="disabled"/>
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
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";

@Component({
  name: 'Dmail',
  components: {NetworkSelector}
})
export default class Dmail extends DefaultSimple {

  networks = [Network.Base]

  item: SimpleTask = {
    network: Network.Base,
  }

  created() {
    if (this.task?.mintZeriusTask) {
      this.item = this.task.mintZeriusTask
    }
  }

  getTask(): Task {
    const taskType = TaskType.MintZerius
    const task = {
      weight: this.weight.toString(),
      mintZeriusTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.mintZeriusTask) {
        this.item = this.task.mintZeriusTask
        this.$emit('taskChanged', this.getTask())
      }
    }
  }
}


</script>

<style scoped>

</style>
