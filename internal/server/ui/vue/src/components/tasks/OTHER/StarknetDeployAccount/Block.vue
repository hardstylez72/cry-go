<template>
  <v-card-actions>
    <br/>
    <v-container>
      <v-row>
        <div>Для работы со StarkNet, в отличие от EVM сетей, необходимо задеплоить аккаунт.
          Эта функция позволяет произвести деплой аккаунта
        </div>
      </v-row>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {DeployStarkNetAccountTask, Network, Task, TaskType} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {TaskSpec} from "@/components/tasks/utils";

export default defineComponent({
  name: "TaskDeployStarkNetAccount",
  emits: ['taskChanged'],
  props: {
    spec: {
      type: Object as PropType<TaskSpec>,
      required: true,
    },
    weight: {
      type: Number,
      required: true,
    },
    disabled: {
      type: Boolean,
      required: true,
    },
    task: {
      type: Object as PropType<Task>,
      required: false,
      deep: true,
    }
  },
  watch: {
    local: {
      handler(b, a) {
        if (JSON.stringify(a) !== JSON.stringify(b)) {
        }
      },
      deep: true
    },
    task: {
      handler(b, a) {
        if (JSON.stringify(a) !== JSON.stringify(b)) {
          this.syncTask()
        }
      },
      deep: true
    }
  },

  data() {
    return {
      local: {network: Network.StarkNet} as DeployStarkNetAccountTask,
    }
  },
  methods: {
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        deployStarkNetAccountTask: this.local,
        taskType: TaskType.DeployStarkNetAccount,
        description: ""
      }
    },
    inputChanged() {
      this.$emit('taskChanged', this.getTask())
      return this.validateForm()
    },
    syncTask() {
      if (this.task) {
        if (this.task.deployStarkNetAccountTask) {
          this.local = this.task.deployStarkNetAccountTask
          this.$emit('taskChanged', this.getTask())
        }
      }
    }
  },
  created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
  }
})
</script>

<style scoped>

</style>
