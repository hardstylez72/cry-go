<template>
  <v-card-actions>
    <br/>
    <v-container>
      <v-row>
        <v-col>
          <NetworkSelector
              label="network"
              :items="networks"
              :disabled="disabled"
              v-model="network"
          />
        </v-col>
        <v-col>

          <TokenSelector :disabled="disabled" label="Монета на выход" :items="tokens" v-model="item.token"
          />
        </v-col>
      </v-row>
      <AmountInput :coin="item.token" :disabled="disabled" v-model="item.amount"/>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Network, OkexDepositTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {required} from "@/components/tasks/helper";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";
import TokenSelector from "@/components/tasks/TokenSelector.vue";

export default defineComponent({
  name: "TaskOkexDeposit",
  components: {TokenSelector, NetworkSelector, AmountInput},
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
      deep: true,
    }
  },

  watch: {
    network: {
      handler(b, a) {
        this.item.network = this.network
        this.item.token = null
      },
    },
    item: {
      deep: true,
      handler() {
        this.$emit('taskChanged', this.getTask())
      }
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
      network: Network.ARBITRUM,
      opt: new Map<Network, Token[]>([
        [Network.ARBITRUM, [Token.USDT, Token.USDC, Token.USDCBridged, Token.ETH, Token.ZRO]],
        [Network.StarkNet, [Token.ETH, Token.STRK]],
        [Network.ZKSYNCERA, [Token.ETH]],
        [Network.Etherium, [Token.ETH, Token.USDC, Token.USDT]],
        [Network.AVALANCHE, [Token.USDC, Token.USDT, Token.AVAX]],
        [Network.POLIGON, [Token.USDC, Token.USDT]],
        [Network.BinanaceBNB, [Token.USDC, Token.USDT, Token.ZK, Token.BNB]],
        [Network.OPTIMISM, [Token.ETH, Token.USDC, Token.USDT]],
        [Network.Linea, [Token.ETH]],
        [Network.Base, [Token.ETH]],
      ]),
      item: {
        network: Network.ARBITRUM,
        token: Token.USDT
      } as OkexDepositTask,
    }
  },
  computed: {
    tokens(): Token[] {
      const tokens = this.opt.get(this.item.network)
      if (!tokens) {
        return []
      }
      return tokens
    },
    networks(): Network[] {
      return [
        Network.AVALANCHE,
        Network.Etherium,
        Network.POLIGON,
        Network.BinanaceBNB,
        Network.OPTIMISM,
        Network.ARBITRUM,
        Network.ZKSYNCERA,
        Network.StarkNet,
        Network.Linea,
        Network.Base,
      ]
    }
  },
  methods: {
    required,
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        okexDepositTask: this.item,
        taskType: TaskType.OkexDeposit,
        description: ""
      }
    },
    syncTask() {
      if (this.task) {
        if (this.task.okexDepositTask) {
          this.item = this.task.okexDepositTask
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
