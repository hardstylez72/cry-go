<template>
  <Exchange/>
  <Zksync/>
  <LayerZero/>
  <Other/>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import Zksync from "@/components/modules/Zksync.vue";
import Orbiter from "@/components/modules/Other.vue";
import LayerZero from "@/components/modules/LayerZero.vue";
import Exchange from "@/components/modules/Exchange.vue";
import Other from "@/components/modules/Other.vue";


export default defineComponent({
  name: "General",
  components: {Other, Exchange, LayerZero, Orbiter, Zksync},
  created() {
  },
})


</script>


<style>
</style>

