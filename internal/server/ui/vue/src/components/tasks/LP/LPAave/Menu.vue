<template>
  <MenuTaskSettings :network="item.network"/>
  <div>Network: <b>{{ item.network }}</b></div>
  <div>Pool: <b>{{ `${item.tokens.map(t => t.token).join(' - ')}` }}</b></div>
  <div>Operation: <b>{{ `${item.add ? 'deposit' : 'withdraw'}` }}</b></div>
  <div>Amount: {{ getAmountSend(item.amount) }}</div>
  <div>Status:
    <span :style="`color: ${getTxStatusColor}`">{{ getTxStatus }}</span>
  </div>
  <GasOptions :item="item.tx" :network="item.network"/>
</template>

<script lang="ts">
import {AmUni, DefaultLP, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ProcessStatus} from "@/generated/process";
import GasOptions from "@/components/tasks/Details.vue";
import {taskProps} from "../../tasks";
import MenuTaskSettings from "@/components/tasks/Settings.vue";
import {getAmountSend} from "@/components/tasks/helper";

export default defineComponent({
  name: "MenuTaskZkLendLP",
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
      item: {} as DefaultLP,
    }
  },
  watch: {
    task: {
      handler() {
        if (this.task?.zkLendLPTask) {
          this.item = this.task.zkLendLPTask
        }
      },
      deep: true
    }
  },
  computed: {
    taskProps() {
      return taskProps
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
    if (this.task?.aaveLPTask) {
      this.item = this.task.aaveLPTask
    }
  }
})
</script>

<style scoped>

</style>
