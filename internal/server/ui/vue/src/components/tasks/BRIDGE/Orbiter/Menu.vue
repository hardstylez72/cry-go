<template>
  <div class="d-inline-flex">
    <a :href="taskProps.OrbiterBridge.service.link" target="_blank">{{ taskProps.OrbiterBridge.service.name }}</a>
    <MenuTaskSettings :network="item.fromNetwork"/>
  </div>
  <div>Network: <b>{{ `${item.fromNetwork} to ${item.toNetwork}` }}</b></div>
  <div>Token: <b>{{ `${item.fromToken} to ${item.toToken}` }}</b></div>
  <div>Amount to swap: {{ getAmountSend(item.amount) }}</div>
  <div>Status:
    <span :style="`color: ${getStatusColor(getStatus)}`">{{ getStatus }}</span>
  </div>
  <GasOptions :item="item.tx" :network="item.fromNetwork"/>
</template>

<script lang="ts">
import {
  OrbiterBridgeTask,
  Task,
} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ProcessStatus} from "@/generated/process";
import {getAmountSend, getStatusColor} from "@/components/tasks/helper";
import GasOptions from "@/components/tasks/Details.vue";
import MenuTaskSettings from "@/components/tasks/Settings.vue";
import {taskProps} from "../../tasks";

export default defineComponent({
  name: "MenuTaskStargateBridge",
  methods: {getAmountSend, getStatusColor},
  components: {MenuTaskSettings, GasOptions},
  props: {
    task: {
      type: Object as PropType<Task>,
      required: true,
    },
    status: {
      type: String as PropType<ProcessStatus>,
      required: true,
    }
  },
  data() {
    return {
      item: {} as OrbiterBridgeTask,
    }
  },
  watch: {
    task: {
      handler() {
        if (this.task?.stargateBridgeTask) {
          this.item = this.task.stargateBridgeTask
        }
      },
      deep: true
    }
  },
  computed: {
    taskProps() {
      return taskProps
    },
    link() {
      return link
    },
    getStatus(): string {
      if (this.status == ProcessStatus.StatusDone) {
        return 'completed'
      }
      if (this.status == ProcessStatus.StatusError) {
        return 'error'
      }

      if (!this.item.txId) {
        return 'not started'
      }

      if (this.item.txCompleted) {
        return 'transaction send'
      }

      return 'waiting'
    },
  },
  async created() {
    if (this.task?.orbiterBridgeTask) {
      this.item = this.task.orbiterBridgeTask
    }
  }
})
</script>

<style scoped>

</style>
