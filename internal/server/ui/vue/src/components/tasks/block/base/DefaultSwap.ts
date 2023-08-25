import {DefaultSwap as TaskSignature, Network, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType, render} from "vue";
import WEIInputField from "@/components/WEIInputField.vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {link} from "@/components/tasks/links";
import {taskProps} from "@/components/tasks/tasks";
import {required} from "@/components/tasks/menu/helper";
import {SwapPair, tokenSwapPair} from "@/components/helper";
import {Component, Prop, Vue, Watch} from "vue-facing-decorator";
import MenuTaskSettings from "@/components/tasks/menu/Settings.vue";
import GasOptions from "@/components/tasks/menu/Details.vue";


@Component({
  inheritAttrs: true,
  template: `
    <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <v-select
            ref="stargate-bridge-form"
            density="compact"
            variant="outlined"
            label="network"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="networks"
            v-model="item.network"
            :disabled="true"
          />
        </v-col>
        <v-col>
          <v-autocomplete
            density="compact"
            variant="outlined"
            label="direction"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="pairs"
            v-model="pair"
            :disabled="disabled"
            item-title="name"
            return-object
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <AmountInput :coin="item.fromToken" :disabled="disabled" v-model="item.amount"/>
        </v-col>
      </v-row>
    </v-container>
    </v-card-actions>
  `,
  name: 'TaskDefaultSwap',
  components: {AmountInput, WEIInputField},
  emits: ['taskChanged'],
})
export default class DefaultSwap extends Vue {
  @Prop weight!: Number
  @Prop disabled!: Boolean
  @Prop task!: Task

  @Watch('pair', {
    deep: true,
  })
  pairHandler(a, b) {
    this.item.fromToken = this.pair.from
    this.item.toToken = this.pair.to
  }

  @Watch('item', {deep: true})
  itemHandler(a, b) {
    this.$emit('taskChanged', this.getTask())
  }

  @Watch('task', {deep: true})
  taskHandler(a, b) {
    if (JSON.stringify(a) !== JSON.stringify(b)) {
      this.syncTask()
    }
  }

  pairs: SwapPair[] = [
    tokenSwapPair(Token.ETH, Token.USDC),
    tokenSwapPair(Token.USDC, Token.ETH),
  ]
  pair = null as SwapPair | null
  networks: Network[] = [Network.StarkNet]
  item: TaskSignature = {
    network: Network.StarkNet,
    amount: {
      sendAll: true,
    },
    toToken: Token.USDC,
    fromToken: Token.ETH,
  }

  get taskProps() {
    return taskProps
  }

  get link() {
    return link
  }

  get Token() {
    return Token
  }

  required

  getTask(): Task {
    const taskType = TaskType.SithSwap
    const task = {
      weight: this.weight.toString(),
      sithSwapTask: this.item,
      taskType: taskType,
      description: "",
    }
    task.description = taskProps[taskType].descFn(task)
    return task
  }

  inputChanged() {
  }

  syncTask() {
    if (this.task) {
      if (this.task.sithSwapTask) {
        this.item = this.task.sithSwapTask
        this.pair = tokenSwapPair(this.item.fromToken, this.item.toToken)
        this.$emit('taskChanged', this.getTask())
      }
    }
  }

  created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
  }
}
