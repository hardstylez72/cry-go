<template>
  <Exchange/>
  <Zksync/>
  <LayerZero/>
  <Other/>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import Zksync from "@/components/landing/modules/Zksync.vue";
import Orbiter from "@/components/landing/modules/Other.vue";
import LayerZero from "@/components/landing/modules/LayerZero.vue";
import Exchange from "@/components/landing/modules/Exchange.vue";
import Other from "@/components/landing/modules/Other.vue";


export default defineComponent({
  name: "General",
  components: {Other, Exchange, LayerZero, Orbiter, Zksync},
  created() {
  },
})


</script>


<style>
</style>

