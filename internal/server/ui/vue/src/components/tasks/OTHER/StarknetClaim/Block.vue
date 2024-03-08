<template>
  <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          Заклеймить токены и пожелать старкнету добра
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
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";

@Component({
  name: 'Dmail',
  components: {NetworkSelector}
})
export default class Dmail extends DefaultSimple {

  networks = [Network.StarkNet]

  item: SimpleTask = {
    network: Network.StarkNet,
  }

  created() {
    if (this.task?.starknetClaim) {
      this.item = this.task.starknetClaim
    }
  }

  getTask(): Task {
    const taskType = TaskType.StarknetClaim
    const task = {
      weight: this.weight.toString(),
      starknetClaim: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.starknetClaim) {
        this.item = this.task.starknetClaim
        this.$emit('taskChanged', this.getTask())
      }
    }
  }
}


</script>

<style scoped>

</style>
