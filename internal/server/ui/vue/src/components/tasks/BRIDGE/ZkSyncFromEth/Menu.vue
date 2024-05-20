<template>
  <div class="d-inline-flex">
    <a :href="taskProps.ZkSyncOfficialBridgeFromEthereum.service.link"
       target="_blank">{{ taskProps.ZkSyncOfficialBridgeFromEthereum.service.name }}</a>
    <MenuTaskSettings :network="Network.Etherium"/>
  </div>
  <div>Network: <b>{{ Network.Etherium }}</b></div>
  <div>Bridge to: <b>{{ Network.ZKSYNCERA }}</b></div>
  <div>Amount to swap: {{ getAmountSend(item.amount) }}</div>
  <div>Status:
    <span :style="`color: ${getTxStatusColor}`">{{ getTxStatus }}</span>
  </div>
  <GasOptions :item="item.tx" :network="Network.Etherium"/>
</template>

<script lang="ts">
import {
  AmUni,
  Network,
  Task,
  Token, ZkSyncOfficialBridgeFromEthereumTask,
} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ProcessStatus} from "@/generated/process";
import {getAmountSend} from "../../helper";
import GasOptions from "@/components/tasks/Details.vue";
import MenuTaskSettings from "@/components/tasks/Settings.vue";
import {taskProps} from "../../tasks";

export default defineComponent({
  name: "MenuZkSyncOfficialBridgeFromEth",
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
      item: {} as ZkSyncOfficialBridgeFromEthereumTask,
    }
  },
  watch: {
    task: {
      handler() {
        if (this.task?.syncSwapTask) {
          this.item = this.task.syncSwapTask
        }
      },
      deep: true
    }
  },
  computed: {
    taskProps() {
      return taskProps
    },

    Network() {
      return Network
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

      if (this.item.txId) {
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
    },
  },
  async created() {
    if (this.task?.zkSyncOfficialBridgeFromEthereumTask) {
      this.item = this.task.zkSyncOfficialBridgeFromEthereumTask
    }
  }
})
</script>

<style scoped>

</style>
