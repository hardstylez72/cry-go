<template>
  <v-menu
    location-strategy="connected"
    max-width="700px"
    v-model="menu"
    persistent="true"
    no-click-animation
    transition="false"
    :close-on-content-click="false"
    :close-on-back="false">
    <template v-slot:activator="{ props }">
      <v-btn rounded="false" width="100%" v-bind="props" :disabled="disable">{{ label }}</v-btn>
    </template>
    <v-card>
      <div class="d-flex justify-end">
        <v-icon icon="mdi-close" @click="menu = false"></v-icon>
      </div>
      <TaskSelector v-if="!disable" @task-selected="addStep"/>
    </v-card>
  </v-menu>
</template>

<script lang="ts">

import {defineComponent, PropType, ref} from 'vue';
import {Network, Task, TaskType} from "@/generated/flow";
import {taskComponentMap, TaskArg, taskTypes, taskProps} from '@/components/tasks/tasks'
import draggable from 'vuedraggable'
import TaskChip from "@/components/tasks/TaskChip.vue";
import {networkProps} from "@/components/helper";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskSelector from "@/components/flow/TaskSelector.vue";


export default defineComponent({
  name: "TaskSelectorMenu",
  emits: ['addStep'],
  components: {
    TaskSelector,
    TaskChip,
    draggable,
  },
  watch: {},
  props: {
    disable: {
      type: Boolean as PropType<boolean>,
      required: false,
    },
    label: {
      type: String as PropType<string>,
      required: true,
    },
  },
  data() {
    return {
      menu: false
    }
  },
  methods: {
    addStep(t: TaskType) {
      this.$emit("addStep", t)
    }
  },
  created() {
  }
})


</script>


<style>


</style>

