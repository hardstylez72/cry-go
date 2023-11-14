<template>
  <MenuTaskSettings :network="item.fromNetwork"/>
  <div>From <b>{{ item.fromNetwork }}</b> to <b>{{ item.toNetwork }}</b></div>
  <div>Status:
    <span :style="`color: ${getTxStatusColor}`">{{ getTxStatus }}</span>
  </div>
  <GasOptions :item="item.mintTx" name="mint" :network="item.fromNetwork"/>
  <GasOptions :item="item.bridgeTx" name="bridge" :network="item.fromNetwork"/>
</template>

<script lang="ts">
import {
  AmUni,
  MerklyMintAndBridgeNFTTask,
  Task,
} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {ProcessStatus} from "@/generated/process";
import {getAmountSend} from "../../helper";
import {taskProps} from "@/components/tasks/tasks";
import GasOptions from "@/components/tasks/Details.vue";
import MenuTaskSettings from "@/components/tasks/Settings.vue";

export default defineComponent({
  name: "MenuTaskMerklyNFT",
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
      item: {} as MerklyMintAndBridgeNFTTask,
    }
  },
  watch: {
    task: {
      handler() {
        if (this.task?.merklyMintAndBridgeNFTTask) {
          this.item = this.task.merklyMintAndBridgeNFTTask
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
      if (!this.item.tx) {
        return 'not started'
      }

      if (this.item.tx.txCompleted) {
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
    if (this.task?.merklyMintAndBridgeNFTTask) {
      this.item = this.task.merklyMintAndBridgeNFTTask
    }
  }
})
</script>

<style scoped>

</style>
