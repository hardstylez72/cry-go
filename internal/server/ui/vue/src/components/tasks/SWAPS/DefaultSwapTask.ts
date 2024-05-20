import {DefaultSwap as TaskSignature, Network, Task, TaskType, Token} from "@/generated/flow";
import WEIInputField from "@/components/WEIInputField.vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {taskProps} from "@/components/tasks/tasks";
import {SwapPair, tokenSwapPair} from "@/components/helper";
import {Component, Prop, Vue, Watch} from "vue-facing-decorator";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";
import {required} from "@/components/tasks/helper";


@Component({
  inheritAttrs: true,
  template: `
    <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <NetworkSelector
            label="from network"
            :items="GetFromNetworks"
            :disabled="disabled"
            v-model="item.network"
          />
        </v-col>
        <v-col>
          <v-autocomplete
            density="compact"
            variant="outlined"
            label="direction"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="getPairs"
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
  components: {NetworkSelector, AmountInput, WEIInputField},
  emits: ['taskChanged'],
})
export default class DefaultSwapTask extends Vue {
  @Prop weight!: Number
  @Prop disabled!: boolean
  @Prop task!: Task

  @Watch('pair', {
    deep: true,
  })
  pairHandler(a, b) {
    if (!this.pair) {
      this.item.fromToken = undefined
      this.item.toToken = undefined
      return
    }

    this.item.fromToken = this.pair.from
    this.item.toToken = this.pair.to
    this.item.network = this.pair.network
  }

  @Watch('network', {deep: true})
  networkHandler(a: Network, b: Network) {
    this.item.network = this.network
    if (!this.init) {
      this.pair = null
    }
  }

  @Watch('item', {deep: true})
  itemHandler(a: TaskSignature, b: TaskSignature) {
    this.$emit('taskChanged', this.getTask())
  }

  @Watch('task', {deep: true})
  async taskHandler(a, b) {
    if (JSON.stringify(a) !== JSON.stringify(b)) {
      this.syncTask()
      await this.$nextTick(() => {
        this.init = false
      })
    }
  }

  network = null as null | Network
  init = true
  pairs: SwapPair[] = []
  pair = null as SwapPair | null
  item: TaskSignature = {
    network: null,
    amount: {
      sendAll: true,
    },
    toToken: null,
    fromToken: null,
  }

  get taskProps() {
    return taskProps
  }


  get Token() {
    return Token
  }

  get getPairs() {
    const pairs = this.pairs.filter((p) => {
      if (p.network === this.network) {
        return p
      }
    })

    if (pairs.length > 0) {
      return pairs
    }

    return []
  }

  required = required

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

  get GetFromNetworks(): Network[] {
    const m = new Set<Network>()
    this.pairs.forEach((p) => {
      m.add(p.network)
    })

    const out = []
    m.forEach(v => {
      out.push(v)
    })

    return out
  }

  syncTask() {
    if (this.task) {
      if (this.task.sithSwapTask) {
        this.item = this.task.sithSwapTask
        if (this.item.network && this.item.fromToken && this.item.toToken) {
          this.pair = tokenSwapPair(Network.ZKSYNCERA, this.item.fromToken, this.item.toToken)
        }
        this.$emit('taskChanged', this.getTask())
      }
    }
  }

  created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
  }
}
