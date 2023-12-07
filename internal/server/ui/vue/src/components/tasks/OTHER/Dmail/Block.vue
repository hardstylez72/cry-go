<template>
  <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <NetworkSelector label="Сеть" :items="networks" v-model="item.network"/>
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
