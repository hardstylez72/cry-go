<template>
  <MenuTaskSettings :network="item.network"/>
  <div><b>Status: </b> <span :style="'color: ' + getStatusColor(getStatus)">{{ getStatus }}</span></div>
  <div>Network: <b>{{ item.network }}</b></div>
</template>

<script lang="ts">
import {DeployStarkNetAccountTask, OkexDepositTask, SnapshotVoteTask, Task, TaskType} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {formatTime, humanDuration} from "@/components/helper";
import {ProcessStatus} from "@/generated/process";
import {getStatusColor} from "@/components/tasks/menu/helper";
import {link} from "@/components/tasks/links";
import MenuTaskSettings from "@/components/tasks/menu/Settings.vue";
import {TaskSpec} from "@/components/tasks/utils";


interface Proposal {
  status: string,
  link: string
}

export default defineComponent({
  name: "DeployStarkNetAccount",
  components: {MenuTaskSettings},
  props: {
    spec: {
      type: Object as PropType<TaskSpec>,
      required: true,
    },
    task: {
      type: Object as PropType<Task>,
      required: true,
    },
    status: {
      type: String as PropType<ProcessStatus>,
      required: true,
    }
  },
  watch: {
    item: {
      handler() {
        if (this.task?.deployStarkNetAccountTask) {
          this.item = this.task.deployStarkNetAccountTask
        }
      },
      deep: true
    }
  },
  data() {
    return {
      item: {} as DeployStarkNetAccountTask,
    }
  },
  computed: {
    getStatus(): string {
      if (this.status == ProcessStatus.StatusError) {
        return 'error'
      }

      if (this.status == ProcessStatus.StatusDone) {
        return 'completed'
      }

      if (this.status == ProcessStatus.StatusRunning) {
        return 'processing'
      }

      return 'not started'
    }
  },
  methods: {
    getStatusColor,
    humanDuration,
    formatTime,
  },
  async created() {
    if (this.task?.deployStarkNetAccountTask) {
      this.item = this.task.deployStarkNetAccountTask
    }
  }
})
</script>

<style scoped>

</style>
