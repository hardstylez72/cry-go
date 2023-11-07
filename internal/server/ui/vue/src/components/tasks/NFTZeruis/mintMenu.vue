<template>
  <div>
    <div class="d-inline-flex">
      <MenuTaskSettings :network="item.network"/>
    </div>
    <div>Network: <b>{{ item.network }}</b></div>
    <div>Status:
      <span :style="statusColor">{{ getTxStatus }}</span>
      <GasOptions :item="item.tx" :network="item.network"/>
    </div>
  </div>
</template>

<script lang="ts">

import DefaultSwapMenu from "@/components/tasks/menu/base/DefaultSwapMenu.js";
import {Component} from "vue-facing-decorator";
import DefaultSimpleMenu from "@/components/tasks/menu/base/DefaultSimpleMenu";

@Component({name: 'MenuDmail'})
export default class MenuDmail extends DefaultSimpleMenu {
  created() {
    if (this.task?.mintZeriusTask) {
      this.item = this.task.mintZeriusTask
    }
  }
}

</script>

<style scoped>

</style>
