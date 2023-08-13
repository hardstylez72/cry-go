<template>
  <div class="d-inline-flex">
    <a :href="taskProps.SyncSwap.service.link" target="_blank">{{ taskProps.SyncSwap.service.name }}</a>
    <MenuTaskSettings :network="item.network"/>
  </div>
  <div>Network: <b>{{ item.network }}</b></div>
  <div>Pool: <b>{{ `${item.a}/${item.b}` }}</b></div>
  <div>Operation: <b>{{ `${item.add ? 'deposit' : 'withdraw'}` }}</b></div>
  <div>Amount: {{ getAmountSend(item.amount) }}</div>
  <div>Status:
    <span :style="`color: ${getTxStatusColor}`">{{ getTxStatus }}</span>
  </div>
  <GasOptions :item="item.tx" :network="item.network"/>
</template>

<script lang="ts">
import {AmUni, Network, SyncSwapLPTask, SyncSwapTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ProcessStatus} from "@/generated/process";
import {getAmountSend} from "./helper";
import {link} from "@/components/tasks/links";
import GasOptions from "@/components/tasks/menu/Details.vue";
import {taskProps} from "../tasks";
import MenuTaskSettings from "@/components/tasks/menu/Settings.vue";

export default defineComponent({
  name: "MenuTaskSyncSwapLP",
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
      item: {} as SyncSwapLPTask,
    }
  },
  watch: {
    task: {
      handler() {
        if (this.task?.syncSwapLPTask) {
          this.item = this.task.syncSwapLPTask
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
    getTxStatusColor(): string {
      const s = this.getTxStatus
      switch (s) {
        case 'error':
          return 'red'
        case 'not started':
          return 'grey'
        case 'completed':
          return 'green'
        default:
          return 'blue'
      }
    },
    getTxStatus(): string {
      if (this.status == ProcessStatus.StatusError) {
        return 'error'
      }
      if (!this.item.txId) {
        return 'not started'
      }

      if (this.item.txCompleted) {
        return 'completed'
      }

      return 'waiting'
    },
  },
  methods: {
    getAmountSend,
    getBalance(am?: AmUni): string {
      if (!am) {
        return ''
      }

      return `${Number(am.usd).toPrecision(3)} USD`
      // return `${Number(am.eth).toPrecision(3)} ETH ${Number(am.usd).toPrecision(3)} USD`
    },

  },
  async created() {
    if (this.task?.syncSwapLPTask) {
      this.item = this.task.syncSwapLPTask
    }
  }
})
</script>

<style scoped>

</style>
