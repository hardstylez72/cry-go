<template>
  <v-chip :color="color">{{ this.status }}</v-chip>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {IssueStatus} from "@/generated/issue";

export default defineComponent({
  name: "StatusChip.vue",
  computed: {
    color() {
      switch (this.status) {
        case  IssueStatus.Rejected:
          return 'red'
        case  IssueStatus.Done:
          return 'green'
        case  IssueStatus.Processing:
          return 'blue'
        case  IssueStatus.Created:
          return 'gray'
        case  IssueStatus.Deleted:
          return 'black'
      }
    }
  },
  props: {
    status: {
      type: String as PropType<IssueStatus>,
      required: true
    }
  },

})


</script>
<style>
</style>

