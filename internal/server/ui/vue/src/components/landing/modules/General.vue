<template>
  <Exchange/>
  <Base/>
  <StarkNet/>
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
import StarkNet from "@/components/landing/modules/StarkNet.vue";
import Base from "@/components/landing/modules/Base.vue";


export default defineComponent({
  name: "General",
  components: {StarkNet, Other, Exchange, LayerZero, Orbiter, Zksync, Base},
  created() {
  },
})


</script>


<style>
</style>

