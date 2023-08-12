<template>
  <v-card-actions>
    <v-container>
      <div class="mb-3">use <a target="_blank"
                               :href="taskProps.MerklyMintAndBridgeNFT.service.link">{{
          taskProps.MerklyMintAndBridgeNFT.service.name
        }}</a>
        to see
        available options
      </div>
      <v-row>
        <v-col cols="6">
          <v-select
            density="compact"
            variant="outlined"
            label="from network"
            :rules="[required]"
            :items="getFromNetworks"
            v-model="item.fromNetwork"
            :disabled="true"
            hide-details
          />
        </v-col>
        <v-col cols="6">
          <v-select
            density="compact"
            variant="outlined"
            label="to network"
            :rules="[required]"
            :items="getToNetworks"
            v-model="item.toNetwork"
            :disabled="disabled"
            hide-details
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

export default defineComponent({
  name: "TaskMerklyNFT",
  components: {AmountInput, WEIInputField},
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
      return this.networks
    },
    getToNetworks(): Network[] {
      return this.networks.filter((n) => n !== this.item.fromNetwork)
    }
  },
  watch: {
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
        Network.OPTIMISM,
        Network.ARBITRUM,
        Network.AVALANCHE,
        Network.POLIGON,
        Network.BinanaceBNB,
      ] as Network[],
      item: {
        fromNetwork: Network.ZKSYNCERA,
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
