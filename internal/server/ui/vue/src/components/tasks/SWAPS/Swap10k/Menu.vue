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
import {Component, toNative, mixins} from "vue-facing-decorator";

@Component({
  name: 'MenuSwap10k',
})
export default class MenuSwap10k extends DefaultSwapMenu {
  created() {
    if (this.task?.swap10k) {
      this.item = this.task.swap10k
    }
  }
}

</script>

<style scoped>

</style>
