<template>
  <div class="d-inline-flex">
    <a :href="taskProps.ZkSyncOfficialBridgeFromEthereum.service.link"
       target="_blank">{{ taskProps.ZkSyncOfficialBridgeFromEthereum.service.name }}</a>
    <MenuTaskSettings :network="Network.ZKSYNCERA"/>
  </div>
  <div>Network: <b>{{ item.network }}</b></div>
  <div>Bridge to: <b>{{ Network.Etherium }}</b></div>
  <div>Amount to swap: {{ getAmountSend(item.amount) }}</div>
  <div>Status:
    <span :style="`color: ${getTxStatusColor}`">{{ getTxStatus }}</span>
  </div>
  <GasOptions :item="item.tx" :network="item.network"/>
</template>

<script lang="ts">
import {
  AmUni,
  Network,
  Task,
  ZkSyncOfficialBridgeToEthereumTask
} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ProcessStatus} from "@/generated/process";
import {getAmountSend} from "./helper";
import {link} from "@/components/tasks/links";
import GasOptions from "@/components/tasks/menu/Details.vue";
import {taskProps} from "@/components/tasks/tasks";
import MenuTaskSettings from "@/components/tasks/menu/Settings.vue";

export default defineComponent({
  name: "MenuZkSyncOfficialBridgeToEth",
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
      item: {} as ZkSyncOfficialBridgeToEthereumTask,
      required: (v: any) => !!v || 'required'
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
    link() {
      return link
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
    if (this.task?.zkSyncOfficialBridgeToEthereumTask) {
      this.item = this.task.zkSyncOfficialBridgeToEthereumTask
    }
  }
})
</script>

<style scoped>

</style>
