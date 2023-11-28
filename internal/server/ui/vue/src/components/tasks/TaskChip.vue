<template>
  <v-chip size="auto" density="compact" variant="outlined" rounded>
    <template v-slot:prepend>
      <v-img height="22px" width="20px" class="mx-1" v-if="service.img" :src="service.img" position="left"/>
    </template>
    <div class="px-1">{{ service.name }}</div>
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
      console.log('this.taskType', this.taskType)
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
