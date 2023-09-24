<template>
  <div>Swap: <b>{{ item.fromToken }} to {{ item.toToken }}</b></div>
  <div>Amount to swap: {{ getAmountSend(item.amount) }}</div>
</template>

<script lang="ts">
import {ExchangeSwapTask, Task, WithdrawExchangeTask} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ProcessStatus} from "@/generated/process";
import {getAmountSend, getStatusColor} from "@/components/tasks/menu/helper";

export default defineComponent({
  name: "SwapMenu",
  methods: {getAmountSend, getStatusColor},
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
  watch: {
    item: {
      handler() {
        if (this.task?.exchangeSwapTask) {
          this.item = this.task.exchangeSwapTask
        }
      },
      deep: true
    }
  },
  data() {
    return {
      item: {} as ExchangeSwapTask,
    }
  },
  computed: {
    getStatus(): string {
      if (this.status == ProcessStatus.StatusError) {
        return 'error'
      }
      if (this.item.txId) {
        return 'completed'
      }

      if (this.item.withdrawOrderId) {
        return 'awaiting approval confirmation'
      }

      if (this.status === ProcessStatus.StatusReady) {
        return 'not started'
      }

      return 'sending transaction'
    }
  },
  async created() {
    if (this.task?.exchangeSwapTask) {
      this.item = this.task.exchangeSwapTask
    }
  }
})
</script>

<style scoped>

</style>
