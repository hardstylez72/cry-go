<template>
  <v-card-actions>
    <v-container>
      <v-row>
        <v-col>
          <v-select
            ref="stargate-bridge-form"
            density="compact"
            variant="outlined"
            label="from network"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="networks"
            v-model="fromNetwork"
            :disabled="true"
          />
        </v-col>
        <v-col>
          <v-select
            density="compact"
            variant="outlined"
            label="to network"
            v-on:change="inputChanged"
            :rules="[required]"
            :items="GetToNetworks"
            v-model="toNetwork"
            :disabled="disabled"
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
import {DefaultBridge, Network, StargateBridgeTask, Task, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {link} from "@/components/tasks/links";
import {taskProps} from "@/components/tasks/tasks";
import {BridgePair, tokenBridgePair} from "@/components/helper";
import {required} from "@/components/tasks/menu/helper";

export default defineComponent({
  name: "TaskAcrossBridgeBlock",
  components: {AmountInput},
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
        this.toNetwork = null
        this.pair = null
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
        Network.ZKSYNCERA,
        Network.ARBITRUM,
      ] as Network[],
      pairs: [
        tokenBridgePair(Network.ZKSYNCERA, Network.ARBITRUM, Token.USDC, Token.USDCBridged),
      ] as BridgePair[],
      fromNetwork: Network.ZKSYNCERA as null | Network,
      toNetwork: Network.ARBITRUM as null | Network,
      pair: null as BridgePair | null,
      item: {
        fromNetwork: Network.ZKSYNCERA,
        toNetwork: Network.ARBITRUM,
        toToken: Token.USDCBridged,
        fromToken: Token.USDC,
        amount: {
          sendAll: true,
        }
      } as DefaultBridge,
    }
  },
  methods: {
    required,
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        acrossBridgeTask: this.item,
        taskType: TaskType.AcrossBridge,
        description: JSON.stringify(this.item),
      }
    },
    inputChanged() {
      return this.validateForm()
    },
    async syncTask() {
      if (this.task) {
        if (this.task.acrossBridgeTask) {
          this.item = this.task.acrossBridgeTask

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
