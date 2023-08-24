<template>
  <div class="d-inline-flex">
    <a :href="taskProps.StargateBridge.service.link" target="_blank">{{ taskProps.StargateBridge.service.name }}</a>
    <MenuTaskSettings :network="item.fromNetwork"/>
  </div>
  <div>Network: <b>{{ `${item.fromNetwork} to ${item.toNetwork}` }}</b></div>
  <div>Token: <b>{{ `${item.fromToken} to ${item.toToken}` }}</b></div>
  <div>Amount to swap: {{ getAmountSend(item.amount) }}</div>
  <div>Status:
    <span :style="`color: ${getStatusColor(getStatus)}`">{{ getStatus }}</span>
  </div>
  <div v-if="item.lzscanUrl"><a :href="item.lzscanUrl">LayerZero scan</a></div>
  <GasOptions :item="item.tx" :network="item.fromNetwork"/>
</template>

<script lang="ts">
import {AmUni, Network, StargateBridgeTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ProcessStatus} from "@/generated/process";
import {getAmountSend, getStatusColor} from "@/components/tasks/menu/helper";
import {link} from "@/components/tasks/links";
import GasOptions from "@/components/tasks/menu/Details.vue";
import {taskProps} from "../tasks";
import MenuTaskSettings from "@/components/tasks/menu/Settings.vue";

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
      item: {} as StargateBridgeTask,
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

      if (this.item.layerZeroStatus) {
        return 'waiting stargate swap complete'
      }

      if (!this.item.tx) {
        return 'not started'
      }

      if (this.item.txCompleted) {
        return 'transaction send'
      }

      return 'waiting'
    },
  },
  async created() {
    if (this.task?.stargateBridgeTask) {
      this.item = this.task.stargateBridgeTask
    }
  }
})
</script>

<style scoped>

</style>
