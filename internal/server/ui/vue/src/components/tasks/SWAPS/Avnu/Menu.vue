<template>
  <div>
    <div class="d-inline-flex">
      <MenuTaskSettings :network="item.network"/>
    </div>
    <div>Network: <b>{{ item.network }}</b></div>
    <div>Swap: <b>{{ item.fromToken }} to {{ item.toToken }}</b></div>
    <div>Amount to swap: {{ getAmountSend(item.amount) }}</div>
    <div>Status:
      <span :style="statusColor">{{ getTxStatus }}</span>
      <GasOptions :item="item.tx" :network="item.network"/>
    </div>
  </div>
</template>

<script lang="ts">

import DefaultSwapMenu from "@/components/tasks/SWAPS/DefaultSwapMenu.js";
import {Component} from "vue-facing-decorator";

@Component({name: 'MenuAvnuSwap'})
export default class MenuAvnuSwap extends DefaultSwapMenu {
  created() {
    if (this.task?.avnuSwapTask) {
      this.item = this.task.avnuSwapTask
    }
  }
}

</script>

<style scoped>

</style>
