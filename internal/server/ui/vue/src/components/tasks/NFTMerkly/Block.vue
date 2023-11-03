<template>
  <v-card-actions>
    <v-container>
      <v-row>
        <v-col cols="6">
          <NetworkSelector
            label="from network"
            :items="getFromNetworks"
            :disabled="disabled"
            v-model="item.fromNetwork"
          />
        </v-col>
        <v-col cols="6">
          <NetworkSelector
            label="to network"
            :items="getToNetworks"
            :disabled="disabled"
            v-model="item.toNetwork"
          />
        </v-col>
      </v-row>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {MerklyMintAndBridgeNFTTask, Network, Task, TaskType} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import WEIInputField from "@/components/WEIInputField.vue";
import AmountInput from "@/components/tasks/AmountInput.vue";
import {required} from "@/components/tasks/menu/helper";
import {taskProps} from "@/components/tasks/tasks";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";

export default defineComponent({
  name: "TaskMerklyNFT",
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
    taskProps() {
      return taskProps
    },
    getFromNetworks(): Network[] {
      const out = []
      this.networks.forEach((n, v) => {
        out.push(v)
      })
      return out
    },
    getToNetworks(): Network[] {
      const out = []
      this.networks.get(this.item.fromNetwork).forEach(n => {
        out.push(n)
      })
      return out
    }
  },
  watch: {
    "item.fromNetwork": {
      handler() {
        this.item.toNetwork = null
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
      networks: new Map<Network, Set<Network>>([
        [Network.ZKSYNCERA, new Set([
          Network.OPTIMISM,
          Network.ARBITRUM,
          Network.AVALANCHE,
          Network.POLIGON,
          Network.BinanaceBNB,
          Network.ArbitrumNova,
          Network.Base,
          Network.Canto,
          Network.Fantom,
          Network.Meter,
          Network.opBNB,
          Network.PolygonZKEVM,
          Network.Tenet,
        ])],
        [Network.BinanaceBNB, new Set([
          Network.ZKSYNCERA,
          Network.OPTIMISM,
          Network.ARBITRUM,
          Network.AVALANCHE,
          Network.POLIGON,
          Network.ArbitrumNova,
          Network.Base,
          Network.Canto,
          Network.Fantom,
          Network.Meter,
          Network.opBNB,
          Network.PolygonZKEVM,
          Network.Tenet,
        ])],
        [Network.ARBITRUM, new Set([
          Network.BinanaceBNB,
          Network.ZKSYNCERA,
          Network.OPTIMISM,
          Network.AVALANCHE,
          Network.POLIGON,
          Network.ArbitrumNova,
          Network.Base,
          Network.Canto,
          Network.Fantom,
          Network.Meter,
          Network.opBNB,
          Network.PolygonZKEVM,
          Network.Tenet,
        ])],
        [Network.POLIGON, new Set([
          Network.BinanaceBNB,
          Network.ZKSYNCERA,
          Network.OPTIMISM,
          Network.AVALANCHE,
          Network.ARBITRUM,
          Network.ArbitrumNova,
          Network.Base,
          Network.Canto,
          Network.Fantom,
          Network.Meter,
          Network.opBNB,
          Network.PolygonZKEVM,
          Network.Tenet,
        ])],
        [Network.Base, new Set([
          Network.BinanaceBNB,
          Network.ZKSYNCERA,
          Network.OPTIMISM,
          Network.AVALANCHE,
          Network.ARBITRUM,
          Network.ArbitrumNova,
          Network.Base,
          Network.Canto,
          Network.Fantom,
          Network.Meter,
          Network.opBNB,
          Network.PolygonZKEVM,
          Network.Tenet,
        ])]
      ]),
      item: {
        fromNetwork: Network.ZKSYNCERA,
        toNetwork: Network.Fantom,
      } as MerklyMintAndBridgeNFTTask,
    }
  },
  methods: {
    required,
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        merklyMintAndBridgeNFTTask: this.item,
        taskType: TaskType.MerklyMintAndBridgeNFT,
        description: "",
      }
    },
    syncTask() {
      if (this.task) {
        if (this.task.merklyMintAndBridgeNFTTask) {
          this.item = this.task.merklyMintAndBridgeNFTTask
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
