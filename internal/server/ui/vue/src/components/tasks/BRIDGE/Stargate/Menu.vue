<template>
  <div>
    <div class="d-inline-flex pb-3">
      <MenuTaskSettings :network="item.fromNetwork"/>
    </div>
    <div>
      <div class="d-inline-flex align-center">
        <NetworkChip :network="item.fromNetwork"/>
        ->
        <NetworkChip :network="item.toNetwork"/>
      </div>
    </div>

    <div>
      <div class="d-inline-flex align-center">
        <TokenChip :token="item.fromToken"/>
        ->
        <TokenChip :token="item.toToken"/>
      </div>
    </div>

    <div>Amount to swap: {{ getAmountSend(item.amount) }}</div>
    <div>Status:
      <span :style="statusColor">{{ getTxStatus }}</span>
      <GasOptions :item="item.tx" :network="item.fromNetwork"/>
    </div>
  </div>
</template>

<script lang="ts">

import {Component} from "vue-facing-decorator";
import DefaultBridgeMenu from "@/components/tasks/BRIDGE/DefaultBridgeMenu";
import NetworkChip from "@/components/tasks/NetworkChip.vue";
import TokenChip from "@/components/tasks/TokenChip.vue";

@Component({
  name: 'BridgeAcrossMenu',
  components: {TokenChip, NetworkChip}
})
export default class BridgeAcrossMenu extends DefaultBridgeMenu {
  created() {
    if (this.task?.stargateBridge) {
      this.item = this.task.stargateBridge
    }
  }
}

</script>

<style scoped>

</style>
