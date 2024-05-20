<template>
  <div>
    <div class="d-inline-flex">
      <MenuTaskSettings :network="item.fromNetwork"/>
    </div>
    <div>From: <b>{{ item.fromNetwork }}</b></div>
    <div>To: <b>{{ item.toNetwork }}</b></div>
    <div>Token: <b>{{ item.fromToken }} -> {{ item.toToken }} </b></div>
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

@Component({name: 'BridgeAcrossMenu'})
export default class BridgeAcrossMenu extends DefaultBridgeMenu {
  created() {
    if (this.task?.acrossBridgeTask) {
      this.item = this.task.acrossBridgeTask
    }
  }
}

</script>

<style scoped>

</style>
