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
import DefaultSimple from "@/components/tasks/OTHER/DefaultSimple";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";

@Component({
  name: 'Dmail',
  components: {NetworkSelector}
})
export default class Dmail extends DefaultSimple {

  networks = [Network.Base, Network.ZKSYNCERA, Network.ARBITRUM, Network.POLIGON, Network.BinanaceBNB]

  item: SimpleTask = {
    network: Network.Base,
  }

  created() {
    if (this.task?.mintMerklyTask) {
      this.item = this.task.mintMerklyTask
    }
  }

  getTask(): Task {
    const taskType = TaskType.MintMerkly
    const task = {
      weight: this.weight.toString(),
      mintMerklyTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  syncTask() {
    if (this.task) {
      if (this.task.mintMerklyTask) {
        this.item = this.task.mintMerklyTask
        this.$emit('taskChanged', this.getTask())
      }
    }
  }
}


</script>

<style scoped>

</style>
