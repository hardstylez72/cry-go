import {DefaultBridge as TaskSignature, Network, Task, TaskType, Token} from "@/generated/flow";
import WEIInputField from "@/components/WEIInputField.vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {link} from "@/components/tasks/links";
import {taskProps} from "@/components/tasks/tasks";
import {BridgePair, tokenBridgePair, tokenSwapPair} from "@/components/helper";
import {Component, Prop, Vue, Watch} from "vue-facing-decorator";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";


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
            v-model="fromNetwork"
          />
        </v-col>
        <v-col>
          <NetworkSelector
            label="to network"
            :items="GetToNetworks"
            :disabled="disabled"
            v-model="toNetwork"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-autocomplete
            return-object
            density="compact"
            variant="outlined"
            v-on:change="inputChanged"
            label="direction"
            :items="getPairs"
            :rules="[required]"
            v-model="pair"
            :disabled="disabled"
            item-title="name"
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
  name: 'TaskDefaultBridge',
  components: {NetworkSelector, AmountInput, WEIInputField},
  emits: ['taskChanged'],
})
export default class TaskDefaultBridge extends Vue {
  @Prop weight!: Number
  @Prop disabled!: Boolean
  @Prop task!: Task

  init = true

  fromNetwork = null as null | Network
  toNetwork = null as null | Network

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
    this.item.fromNetwork = this.pair.fromNetwork
    this.item.toNetwork = this.pair.toNetwork
    this.fromNetwork = this.pair.fromNetwork
    this.toNetwork = this.pair.toNetwork
  }

  @Watch('item', {deep: true})
  itemHandler(a, b) {
    this.$emit('taskChanged', this.getTask())
  }

  @Watch('task', {deep: true})
  taskHandler(a, b) {
    if (JSON.stringify(a) !== JSON.stringify(b)) {
      this.syncTask()
      this.$nextTick(() => {
        this.init = false
      })
    }
  }

  @Watch('toNetwork', {deep: true})
  toNetworkHandler(a, b) {
    this.item.toNetwork = this.toNetwork
    if (!this.init) {
      this.pair = null
    }
  }

  @Watch('fromNetwork', {deep: true})
  fromNetworkHandler(a, b) {
    this.item.fromNetwork = this.fromNetwork
    if (!this.init) {
      this.toNetwork = null
      this.pair = null
    }
  }

  pairs: BridgePair[] = []
  pair = null as BridgePair | null
  item: TaskSignature = {
    fromNetwork: null,
    toNetwork: null,
    received: false,
    amount: {
      sendAll: true,
    },
    toToken: null,
    fromToken: null,
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
    return {
      weight: this.weight.toString(),
      acrossBridgeTask: this.item,
      taskType: TaskType.AcrossBridge,
      description: JSON.stringify(this.item),
    }
  }

  inputChanged() {
  }

  syncTask() {
    if (this.task) {
      if (this.task.acrossBridgeTask) {
        this.item = this.task.acrossBridgeTask

        if (this.item.fromNetwork && this.item.toNetwork && this.item.fromToken && this.item.toToken) {
          this.pair = tokenBridgePair(this.item.fromNetwork, this.item.toNetwork, this.item.fromToken, this.item.toToken)
        }

        this.$emit('taskChanged', this.getTask())
      }
    }
  }

  get getPairs(): BridgePair[] {
    const pairs = this.pairs.filter((p) => {
      if (p.fromNetwork === this.item.fromNetwork && p.toNetwork === this.item.toNetwork) {
        return p
      }
    })

    if (pairs.length > 0) {
      return pairs
    }

    return []
  }

  get GetFromNetworks(): Network[] {
    const m = new Set<Network>()
    this.pairs.forEach((p) => {
      m.add(p.fromNetwork)
    })

    const out = []
    m.forEach(v => {
      out.push(v)
    })

    return out
  }

  get GetToNetworks(): Network[] {
    const m = new Set<Network>()

    this.pairs.forEach((p) => {
      if (p.fromNetwork == this.fromNetwork) {
        m.add(p.toNetwork)
      }
    })

    const out = []
    m.forEach(v => {
      out.push(v)
    })

    return out
  }

  created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
  }
}
