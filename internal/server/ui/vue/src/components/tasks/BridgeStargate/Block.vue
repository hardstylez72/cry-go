<template>
  <v-card-actions>
    <v-container>
      <div class="mb-3">use <a :href="taskProps.StargateBridge.service.link"
                               target="_blank">{{ taskProps.StargateBridge.service.name }}</a> to see
        available swap
        options
      </div>
      <v-row>
        <v-col>
          <NetworkSelector
            label="from network"
            :items="networks"
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
</template>

<script lang="ts">
import {Network, StargateBridgeTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {link} from "@/components/tasks/links";
import {taskProps} from "@/components/tasks/tasks";
import {BridgePair, tokenBridgePair} from "@/components/helper";
import {required} from "@/components/tasks/menu/helper";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";

export default defineComponent({
  name: "TaskStargateBridge",
  components: {NetworkSelector, AmountInput},
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

  watch: {
    "toNetwork": {
      handler() {
        this.item.toNetwork = this.toNetwork
        if (this.toNetwork && this.fromNetwork) {
          this.pair = null
        }
      }
    },
    "fromNetwork": {
      handler() {
        this.item.fromNetwork = this.fromNetwork
        if (this.toNetwork && this.fromNetwork) {
          this.pair = null
        }
      }
    },
    pair: {
      handler() {
        if (!this.pair) {
          return
        }
        this.item.fromToken = this.pair.from
        this.item.toToken = this.pair.to
        this.toNetwork = this.pair.toNetwork
        this.fromNetwork = this.pair.fromNetwork

      },
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
      networks: [
        Network.BinanaceBNB,
        Network.ARBITRUM,
        Network.OPTIMISM,
        Network.POLIGON,
        Network.AVALANCHE,
        Network.ZKSYNCERA,
      ] as Network[],
      pairs: [

        // from ARBITRUM
        tokenBridgePair(Network.ARBITRUM, Network.OPTIMISM, Token.ETH, Token.ETH),
        tokenBridgePair(Network.ARBITRUM, Network.OPTIMISM, Token.STG, Token.STG),
        tokenBridgePair(Network.ARBITRUM, Network.OPTIMISM, Token.USDCBridged, Token.USDCBridged),

        tokenBridgePair(Network.ARBITRUM, Network.POLIGON, Token.STG, Token.STG),
        tokenBridgePair(Network.ARBITRUM, Network.POLIGON, Token.USDCBridged, Token.USDC),
        tokenBridgePair(Network.ARBITRUM, Network.POLIGON, Token.USDT, Token.USDT),

        tokenBridgePair(Network.ARBITRUM, Network.BinanaceBNB, Token.STG, Token.STG),
        tokenBridgePair(Network.ARBITRUM, Network.BinanaceBNB, Token.USDT, Token.USDT),

        tokenBridgePair(Network.ARBITRUM, Network.AVALANCHE, Token.STG, Token.STG),
        tokenBridgePair(Network.ARBITRUM, Network.AVALANCHE, Token.USDT, Token.USDT),
        tokenBridgePair(Network.ARBITRUM, Network.AVALANCHE, Token.USDCBridged, Token.USDC),

        // from BinanaceBNB
        tokenBridgePair(Network.BinanaceBNB, Network.AVALANCHE, Token.STG, Token.STG),
        tokenBridgePair(Network.BinanaceBNB, Network.AVALANCHE, Token.USDT, Token.USDT),

        tokenBridgePair(Network.BinanaceBNB, Network.POLIGON, Token.STG, Token.STG),
        tokenBridgePair(Network.BinanaceBNB, Network.POLIGON, Token.USDT, Token.USDT),

        tokenBridgePair(Network.BinanaceBNB, Network.ARBITRUM, Token.STG, Token.STG),
        tokenBridgePair(Network.BinanaceBNB, Network.ARBITRUM, Token.USDT, Token.USDT),

        tokenBridgePair(Network.BinanaceBNB, Network.OPTIMISM, Token.STG, Token.STG),


        tokenBridgePair(Network.BinanaceBNB, Network.ZKSYNCERA, Token.MAV, Token.MAV),

        //from OPTIMISM

        tokenBridgePair(Network.OPTIMISM, Network.BinanaceBNB, Token.STG, Token.STG),

        tokenBridgePair(Network.OPTIMISM, Network.AVALANCHE, Token.STG, Token.STG),
        tokenBridgePair(Network.OPTIMISM, Network.AVALANCHE, Token.USDCBridged, Token.USDC),

        tokenBridgePair(Network.OPTIMISM, Network.POLIGON, Token.STG, Token.STG),
        tokenBridgePair(Network.OPTIMISM, Network.POLIGON, Token.USDCBridged, Token.USDC),

        tokenBridgePair(Network.OPTIMISM, Network.ARBITRUM, Token.STG, Token.STG),
        tokenBridgePair(Network.OPTIMISM, Network.ARBITRUM, Token.USDCBridged, Token.USDCBridged),
        tokenBridgePair(Network.OPTIMISM, Network.ARBITRUM, Token.ETH, Token.ETH),

        // from AVALANCHE
        tokenBridgePair(Network.AVALANCHE, Network.BinanaceBNB, Token.STG, Token.STG),
        tokenBridgePair(Network.AVALANCHE, Network.BinanaceBNB, Token.USDT, Token.USDT),

        tokenBridgePair(Network.AVALANCHE, Network.POLIGON, Token.USDT, Token.USDT),
        tokenBridgePair(Network.AVALANCHE, Network.POLIGON, Token.USDC, Token.USDC),
        tokenBridgePair(Network.AVALANCHE, Network.POLIGON, Token.STG, Token.STG),

        tokenBridgePair(Network.AVALANCHE, Network.ARBITRUM, Token.USDT, Token.USDT),
        tokenBridgePair(Network.AVALANCHE, Network.ARBITRUM, Token.USDC, Token.USDCBridged),
        tokenBridgePair(Network.AVALANCHE, Network.ARBITRUM, Token.STG, Token.STG),

        tokenBridgePair(Network.AVALANCHE, Network.OPTIMISM, Token.USDC, Token.USDCBridged),
        tokenBridgePair(Network.AVALANCHE, Network.OPTIMISM, Token.STG, Token.STG),

        // from POLIGON
        tokenBridgePair(Network.POLIGON, Network.BinanaceBNB, Token.STG, Token.STG),
        tokenBridgePair(Network.POLIGON, Network.BinanaceBNB, Token.USDC, Token.USDC),
        tokenBridgePair(Network.POLIGON, Network.BinanaceBNB, Token.USDT, Token.USDT),

        tokenBridgePair(Network.POLIGON, Network.AVALANCHE, Token.STG, Token.STG),
        tokenBridgePair(Network.POLIGON, Network.AVALANCHE, Token.USDC, Token.USDC),
        tokenBridgePair(Network.POLIGON, Network.AVALANCHE, Token.USDT, Token.USDT),

        tokenBridgePair(Network.POLIGON, Network.ARBITRUM, Token.STG, Token.STG),
        tokenBridgePair(Network.POLIGON, Network.ARBITRUM, Token.USDC, Token.USDCBridged),
        tokenBridgePair(Network.POLIGON, Network.ARBITRUM, Token.USDT, Token.USDT),

        tokenBridgePair(Network.POLIGON, Network.OPTIMISM, Token.STG, Token.STG),
        tokenBridgePair(Network.POLIGON, Network.OPTIMISM, Token.USDC, Token.USDCBridged),

        // from ZKSYNCERA
        tokenBridgePair(Network.ZKSYNCERA, Network.BinanaceBNB, Token.MAV, Token.MAV),

      ] as BridgePair[],
      fromNetwork: Network.OPTIMISM as null | Network,
      toNetwork: Network.ARBITRUM as null | Network,
      pair: null as BridgePair | null,
      item: {
        fromNetwork: Network.OPTIMISM,
        toNetwork: Network.ARBITRUM,
        toToken: Token.ETH,
        fromToken: Token.ETH,
        amount: {
          sendAll: true,
        }
      } as StargateBridgeTask,
    }
  },
  methods: {
    required,
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        stargateBridgeTask: this.item,
        taskType: TaskType.StargateBridge,
        description: JSON.stringify(this.item),
      }
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['stargate-bridge-form'].validate()

      return valid
    },
    inputChanged() {
      return this.validateForm()
    },
    async syncTask() {
      if (this.task) {
        if (this.task.stargateBridgeTask) {
          this.item = this.task.stargateBridgeTask

          if (this.item.fromNetwork && this.item.toNetwork && this.item.fromToken && this.item.toToken) {
            this.pair = tokenBridgePair(this.item.fromNetwork, this.item.toNetwork, this.item.fromToken, this.item.toToken)
          }

          console.log('this.pair', this.pair)
          this.$emit('taskChanged', this.getTask())
        }
      }
    }
  },
  computed: {
    getPairs(): BridgePair[] {
      const pairs = this.pairs.filter((p) => {
        if (p.fromNetwork === this.item.fromNetwork && p.toNetwork === this.item.toNetwork) {
          return p
        }
      })

      if (pairs.length > 0) {
        return pairs
      }

      return []
    },
    taskProps() {
      return taskProps
    },
    link() {
      return link
    },
    GetToNetworks(): Network[] {
      return this.networks.filter((n) => n !== this.item.fromNetwork);
    },
  },
  created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
  }
})
</script>

<style scoped>

</style>
