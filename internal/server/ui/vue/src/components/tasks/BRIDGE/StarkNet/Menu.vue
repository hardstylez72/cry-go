<template>
  <div>
    <div class="d-inline-flex">
      <MenuTaskSettings :network="item.fromNetwork"/>
    </div>
    <div>From: <b>{{ item.fromNetwork }}</b></div>
    <div>To: <b>{{ item.toNetwork }}</b></div>
    <div>Token: <b>{{ item.token }} </b></div>
    <div>Amount to swap: {{ getAmountSend(item.amount) }}</div>
    <div>Status:
      <span :style="statusColor">{{ getTxStatus }}</span>
      <GasOptions :item="item.tx" :network="item.fromNetwork"/>
    </div>
  </div>
</template>

<script lang="ts">

import DefaultSwapMenu from "@/components/tasks/SWAPS/DefaultSwapMenu.js";
import {Component, toNative} from "vue-facing-decorator";
import DefaultLiquidityBridgeMenu from "@/components/tasks/BRIDGE/DefaultLiquidityBridgeMenu";

@Component({name: 'MenuStarkNetBridge'})
export default class MenuStarkNetBridge extends DefaultLiquidityBridgeMenu {
  created() {
    if (this.task?.starkNetBridgeTask) {
      this.item = this.task.starkNetBridgeTask
    }
  }
}

</script>

<style scoped>

</style>
