<template>
  <v-chip size="auto" density="compact">
    <template v-slot:prepend>
      <div class="d-inline-flex mx-1">
        <v-img height="22px" v-if="service.img" :src="service.img"/>
        {{ service.name }}
      </div>
    </template>

  </v-chip>
</template>

<script lang="ts">
import {Amount, AmUni, TaskType} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {processService} from "@/generated/services";
import {EstimationTx} from "@/generated/process";
import {percent} from "@/components/helper";
import {ca} from "vuetify/locale";
import {taskProps} from "@/components/tasks/tasks";

export default defineComponent({
  name: "TaskChip",
  computed: {
    taskProps() {
      return taskProps
    },
    service() {
      const s = this.taskProps[this.taskType].service
      if (!s) {
        console.log('this.taskType', this.taskType)
      }
      return this.taskProps[this.taskType].service
    }
  },
  props: {
    taskType: {
      type: String as PropType<TaskType>,
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
