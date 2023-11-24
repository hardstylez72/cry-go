<template>
  <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <NetworkSelector
            label="from network"
            :items="networks"
            v-model="item.network"
            :disabled="true"
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-autocomplete
            density="compact"
            variant="outlined"
            label="token"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="tokens"
            item-value="token"
            v-model="selectedToken"
            :disabled="disabled"
            item-title="token"
            return-object
            hide-details
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          Operation:
          {{ item.add ? 'deposit' : 'withdraw' }}
          <v-btn
            :disabled="disabled"
            @click="item.add = !item.add"
          >Change
          </v-btn>
        </v-col>
      </v-row>
      <AmountInput v-if="item.add" :coin="item.a" :disabled="disabled" v-model="item.amount"/>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {DefaultLP, Network, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import WEIInputField from "@/components/WEIInputField.vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {required} from "@/components/tasks/helper";
import {taskProps} from "@/components/tasks/tasks";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";

export default defineComponent({
  name: "TaskZkLendLP",
  components: {NetworkSelector, AmountInput, WEIInputField},
  emits: ['taskChanged'],
  props: {
    weight: {
      type: Number,
      required: true,
    },
    disabled: {
      type: Boolean,
      required: true,
    },
    task: {
      type: Object as PropType<Task>,
      required: false,
    }
  },
  computed: {

    Token() {
      return Token
    },
    getToTokens(): Token[] {
      if (this.item.fromToken === Token.ETH) {
        return this.tokens.filter((n) => n !== this.item.toToken)
      }

      return [Token.ETH]
    }
  },
  watch: {
    selectedToken: {
      handler() {
        this.item.tokens = [{token: this.selectedToken}]
        this.$emit('taskChanged', this.getTask())
      },
      deep: true
    },
    item: {
      handler() {
        this.$emit('taskChanged', this.getTask())
      },
      deep: true
    },
    task: {
      handler(b, a) {
        if (JSON.stringify(a) !== JSON.stringify(b)) {
          this.syncTask()
        }
      },
      deep: true
    }
  },

  data() {
    return {
      selectedToken: Token.USDCBridged,
      networks: [Network.Base] as Network[],
      tokens: [Token.USDCBridged] as Token[],
      item: {
        network: Network.Base,
        amount: {
          sendAll: true,
        },
        tokens: [{token: Token.USDCBridged}]
      } as DefaultLP,
    }
  },
  methods: {
    required,
    getTask(): Task {
      const task = {
        weight: this.weight.toString(),
        aaveLPTask: this.item,
        taskType: TaskType.AaveLP,
        description: ''
      }
      task.description = taskProps[TaskType.AaveLP].descFn(task)

      return task
    },
    inputChanged() {
    },
    syncTask() {
      if (this.task) {
        if (this.task.aaveLPTask) {
          this.item = this.task.aaveLPTask
          this.selectedToken = this.item.tokens[0].token
          this.$emit('taskChanged', this.getTask())
        }
      }
    }
  },
  created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
  }
})
</script>

<style scoped>

</style>
