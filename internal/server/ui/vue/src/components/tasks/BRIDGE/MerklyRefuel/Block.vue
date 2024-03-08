<template>
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
          <!--          <AmountInput :coin="item.fromToken" :disabled="disabled" v-model="item.amount"/>-->
        </v-col>
      </v-row>

    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {Network, Task, TaskType, Token} from "@/generated/flow";
import {tokenBridgePair} from "@/components/helper";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";
import {Component} from "vue-facing-decorator";
import TaskDefaultBridge from "@/components/tasks/BRIDGE/DefaultBridgeTask";

@Component({
  name: 'TaskAcrossBridgeBlock',
  components: {NetworkSelector}
})

export default class TaskAcrossBridgeBlock extends TaskDefaultBridge {
  pairs = [
    tokenBridgePair(Network.ARBITRUM, Network.POLIGON, Token.ETH, Token.MATIC),
    tokenBridgePair(Network.POLIGON, Network.Conflux, Token.MATIC, Token.CFX),
    tokenBridgePair(Network.BinanaceBNB, Network.DFK, Token.BNB, Token.JEWEL),
    tokenBridgePair(Network.BinanaceBNB, Network.Celo, Token.BNB, Token.CELO),
  ]

  getTask(): Task {
    return {
      weight: this.weight.toString(),
      merklyRefuel: this.item,
      taskType: TaskType.MerklyRefuel,
      description: JSON.stringify(this.item),
    }
  }

  syncTask() {
    if (this.task) {
      if (this.task.merklyRefuel) {
        this.item = this.task.merklyRefuel

        if (this.item.fromNetwork && this.item.toNetwork && this.item.fromToken && this.item.toToken) {
          this.pair = tokenBridgePair(this.item.fromNetwork, this.item.toNetwork, this.item.fromToken, this.item.toToken)
        }

        this.$emit('taskChanged', this.getTask())
      }
    }
  }
}
</script>

<style scoped>

</style>
