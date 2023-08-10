<template>
  <div class="d-inline-flex">
    <MenuTaskSettings :network="item.network"/>
  </div>
  <div>Network: <b>{{ item.network }}</b></div>
  <div>Swap: <b>{{ `${item.wrap ? 'ETH -> WETH' : 'WETH -> ETH'}` }}</b></div>
  <div>Amount to swap: {{ getAmountSend(item.amount) }}</div>
  <div>Status:
    <span :style="`color: ${getTxStatusColor}`">{{ getTxStatus }}</span>
  </div>
  <GasOptions :item="item.tx" :network="item.network"/>
</template>

<script lang="ts">
import {AmUni, Network, SyncSwapTask, Task, TaskType, Token, WETHTask} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ProcessStatus} from "@/generated/process";
import {getAmountSend} from "./helper";
import {link} from "@/components/tasks/links";
import GasOptions from "@/components/tasks/menu/TaskDetails.vue";
import MenuTaskSettings from "@/components/tasks/menu/MenuTaskSettings.vue";

export default defineComponent({
  name: "MenuTaskWethSwap",
  components: {MenuTaskSettings, GasOptions,},
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
      item: {} as WETHTask,
    }
  },
  watch: {
    task: {
      handler() {
        if (this.task?.wETHTask) {
          this.item = this.task.wETHTask
        }
      },
      deep: true
    }
  },
  computed: {
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
      if (this.status == ProcessStatus.StatusDone) {
        return 'completed'
      }
      if (!this.item.tx) {
        return 'not started'
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
    if (this.task?.wETHTask) {
      this.item = this.task.wETHTask
    }
  }
})
</script>

<style scoped>

</style>
