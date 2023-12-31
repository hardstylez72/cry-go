<template>
  <span class="d-inline-flex">
      <v-img :src="p.img" height="22px" width="20px" class="mx-1"/>
  {{ p.name }}
  </span>


</template>

<script lang="ts">
import {Amount, AmUni, Network, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {networkProps, percent} from "@/components/helper";
import {tokenProps} from "@/components/tokens";

export default defineComponent({
  name: "TokenChip",
  computed: {
    p() {
      return tokenProps[this.token]
    }
  },
  props: {
    token: {
      type: String as PropType<Token>,
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
