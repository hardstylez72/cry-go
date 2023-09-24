<template>
  <v-card-actions class="mt-3">
    <v-container>
      <v-row>
        <v-col>
          <WithdrawerSelect v-model="item.withdrawerId" :disabled="disabled"/>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-autocomplete
            return-object
            density="compact"
            variant="outlined"
            label="direction"
            :items="pairs"
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
</template>

<script lang="ts">

import WithdrawerSelect from "@/components/exchange.acc/WithdrawerSelect.vue";
import {ExchangeSwapTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {withdrawerService} from "@/generated/services";
import {Withdrawer} from "@/generated/withdraw";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {required} from "@/components/tasks/menu/helper";
import {SwapPair, tokenSwapPair} from "@/components/helper";


export default defineComponent({
  name: "TaskExchangeSwap",
  components: {AmountInput, WithdrawerSelect},
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
      deep: true
    }
  },

  watch: {
    pair: {
      handler() {
        this.item.fromToken = this.pair.from
        this.item.toToken = this.pair.to
      },
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
      pair: tokenSwapPair(Token.ETH, Token.USDC) as SwapPair,
      pairs: [
        tokenSwapPair(Token.ETH, Token.USDC),
        tokenSwapPair(Token.USDC, Token.ETH),

        tokenSwapPair(Token.ETH, Token.USDT),
        tokenSwapPair(Token.USDT, Token.ETH),

      ] as SwapPair[],
      accounts: [] as Withdrawer[],
      item: {
        amount: {sendAll: true},
        fromToken: Token.ETH,
        toToken: Token.USDC,
      } as ExchangeSwapTask,
    }
  },
  methods: {
    required,
    async withdrawerChanged() {
    },
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        exchangeSwapTask: this.item,
        taskType: TaskType.ExchangeSwap,
        description: "",
      }
    },
    syncTask() {
      if (this.task && this.task.exchangeSwapTask) {
        this.item = this.task.exchangeSwapTask
        this.$emit('taskChanged', this.getTask())
      }
    }
  },
  async created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
    try {
      const res = await withdrawerService.withdrawerServiceListWithdrawer()
      this.accounts = res.withdrawers.map((w) => {
        w.label = `${w.label} [${w.exchangeType}]`
        return w
      })
    } finally {

    }
  }
})
</script>

<style scoped>

</style>
