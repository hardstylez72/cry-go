<template>
  <v-card>
    <v-card-text>

      <slot name="default"/>
      <div v-if="combo"> Количество уникальных комбинаций: {{ combo }}</div>
      <div v-if="!data">Ошибка</div>
      <div v-else>
        <div v-for="flow in boba.flow">
          <div class="d-flex flex-wrap">
            <div v-for="task in flow.tasks" class="mx-1 my-1">
              <TaskBlock :task="task"/>
            </div>
          </div>
        </div>
      </div>
    </v-card-text>
  </v-card>

</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {flow_Flow as Flow, FlowPreviewRes, TaskType} from "@/generated/flow";
import {getFlow, taskProps} from "@/components/tasks/tasks";
import TaskBlock from "@/components/flow/TaskBlock.vue";

export default defineComponent({
  name: "Preview",
  components: {TaskBlock},
  props: {
    data: {
      required: true,
      type: Object as PropType<FlowPreviewRes>
    },
    combo: {
      required: false,
      type: Number,
    }
  },
  computed: {
    taskProps() {
      return taskProps
    },
    boba() {
      return this.data
    }
  },
  methods: {},
  created() {

  }
})


</script>


<style>
.task-container {
  width: auto;
  height: 50px;
  border: 1px solid gray;
}

</style>

