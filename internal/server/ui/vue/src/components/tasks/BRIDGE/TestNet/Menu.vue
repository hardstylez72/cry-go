<template>
  <div>Network: <b>{{ `${item.network}` }}</b></div>
  <div>Settings: <b>{{ `[ ${item.minAmount} : ${item.maxAmount} ] ETH` }}</b></div>
  <div v-if="item.amount">Amount ETH: <b>{{ `${item.amount}` }}</b></div>
  <div>Status:
    <span :style="`color: ${getStatusColor(getStatus)}`">{{ getStatus }}</span>
  </div>
</template>

<script lang="ts">
import {
  Task,
  TestNetBridgeSwapTask,
} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ProcessStatus} from "@/generated/process";
import {getStatusColor} from "@/components/tasks/helper";

export default defineComponent({
  name: "MenuTaskTestNetBridge",
  methods: {getStatusColor},
  components: {},
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
      item: {} as TestNetBridgeSwapTask,
    }
  },
  watch: {
    task: {
      handler() {
        if (this.task?.testNetBridgeSwapTask) {
          this.item = this.task.testNetBridgeSwapTask
        }
      },
      deep: true
    }
  },
  computed: {
    getStatus(): string {
      if (this.status == ProcessStatus.StatusDone) {
        return 'completed'
      }
      if (this.status == ProcessStatus.StatusError) {
        return 'error'
      }

      if (this.status == ProcessStatus.StatusReady) {
        return 'not started'
      }


      return 'waiting'
    },
  },
  async created() {
    if (this.task?.testNetBridgeSwapTask) {
      this.item = this.task.testNetBridgeSwapTask
    }
  }
})
</script>

<style scoped>

</style>
