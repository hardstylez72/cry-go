import {DefaultSwap as TaskSignature, Network, SimpleTask, Task, TaskType, Token} from "@/generated/flow";
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
      </v-row>
    </v-container>
    </v-card-actions>
  `,
  name: 'DefaultSimple',
  components: {AmountInput, WEIInputField},
  emits: ['taskChanged'],
})
export default class DefaultSimple extends Vue {
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

  networks: Network[] = [Network.StarkNet]
  item: SimpleTask = {
    network: Network.ZKSYNCERA,
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
    const taskType = TaskType.Dmail
    const task = {
      weight: this.weight.toString(),
      dmailTask: this.item,
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
      if (this.task.dmailTask) {
        this.item = this.task.dmailTask
        this.$emit('taskChanged', this.getTask())
      }
    }
  }

  created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
  }
}
