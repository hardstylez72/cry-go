<template>
  <span class="d-inline-flex">
      <v-img :src="p.img" height="22px" width="20px" class="mx-1"/>
  {{ p.name }}
  </span>


</template>

<script lang="ts">
import {Amount, AmUni, Network, TaskType} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {networkProps, percent} from "@/components/helper";

export default defineComponent({
  name: "NetworkChip",
  computed: {
    p() {
      return networkProps[this.network]
    }
  },
  props: {
    network: {
      type: String as PropType<Network>,
      required: true,
    },
  },
  data() {
    return {}
  },
})
</script>

<style scoped>

</style>
