import {
  AmUni,
  Task,
} from "@/generated/flow";
import {ProcessStatus} from "@/generated/process";
import {link} from "@/components/tasks/links";
import {taskProps} from "@/components/tasks/tasks";
import GasOptions from "@/components/tasks/menu/Details.vue";
import MenuTaskSettings from "@/components/tasks/menu/Settings.vue";

import {Component, Vue, Prop,} from 'vue-facing-decorator'
import {getAmountSend} from "../helper";

@Component({
  template: `
    <div>
    <div class="d-inline-flex">
      <MenuTaskSettings :network="item.fromNetwork"/>
    </div>
    <div>From: <b>{{ item.fromNetwork }}</b></div>
    <div>To: <b>{{ item.toNetwork }}</b></div>
    <div>Token: <b>{{ item.token }} </b></div>
    <div>Amount to swap: {{ getAmountSend(item.amount) }}</div>
    <div>Status:
      <span :style="statusColor">{{ getTxStatus }}</span>
      <GasOptions :item="item.tx" :network="item.fromNetwork"/>
    </div>
    </div>`,
  name: 'DefaultLiquidityBridgeMenu',
  components: {MenuTaskSettings, GasOptions},
})
export default class DefaultLiquidityBridgeMenu extends Vue {
  @Prop() task!: Task
  @Prop() status!: ProcessStatus

  item = {} as any

  get taskProps() {
    return taskProps
  }

  get statusColor() {
    return `color: ${this.getTxStatusColor}`
  }

  get link() {
    return link
  }

  get getTxStatusColor(): string {
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
  }

  get getTxStatus(): string {
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
  }


  getAmountSend(a) {
    return getAmountSend(a)
  }

  getBalance(am?: AmUni): string {
    if (!am) {
      return ''
    }

    return `${Number(am.usd).toPrecision(3)} USD`
    // return `${Number(am.eth).toPrecision(3)} ETH ${Number(am.usd).toPrecision(3)} USD`
  }

  created() {
    // if (this.task?.sithSwapTask) {
    //   this.item = this.task.sithSwapTask
    // }
  }
};
