import {LiquidityBridgeTask as TaskSignature, Network, Task, TaskType, Token} from "@/generated/flow";
import WEIInputField from "@/components/WEIInputField.vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {link} from "@/components/tasks/links";
import {taskProps} from "@/components/tasks/tasks";
import {tokenSwapPair} from "@/components/helper";
import {Component, Prop, Vue, Watch} from "vue-facing-decorator";


@Component({
  template: `
    <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <v-select
            density="compact"
            variant="outlined"
            label="from"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="networks"
            v-model="item.fromNetwork"
            :disabled="true"
          />
          <v-select
            density="compact"
            variant="outlined"
            label="to"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="networks"
            v-model="item.toNetwork"
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
            :items="tokens"
            v-model="item.token"
            :disabled="disabled"
            item-title="name"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <AmountInput :coin="item.token" :disabled="disabled" v-model="item.amount"/>
        </v-col>
      </v-row>
    </v-container>
    </v-card-actions>
  `,
  name: 'DefaultLiquidityBridge',
  components: {AmountInput, WEIInputField},
  emits: ['taskChanged'],
})
export default class DefaultLiquidityBridge extends Vue {
  @Prop weight!: Number
  @Prop disabled!: Boolean
  @Prop task!: Task

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

  tokens: Token[] = [
    Token.ETH
  ]
  networks: Network[] = [Network.StarkNet, Network.Etherium]
  item: TaskSignature = {
    fromNetwork: Network.Etherium,
    toNetwork: Network.StarkNet,
    amount: {
      sendAll: true,
    },
    token: Token.ETH,
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
    const taskType = TaskType.StarkNetBridge
    const task = {
      weight: this.weight.toString(),
      starkNetBridgeTask: this.item,
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
      if (this.task.starkNetBridgeTask) {
        this.item = this.task.starkNetBridgeTask
        this.$emit('taskChanged', this.getTask())
      }
    }
  }

  created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
  }
}
