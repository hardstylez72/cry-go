<template>
  <v-menu
    location-strategy="connected"
    offset="10"
    location="bottom"
    max-width="700px"
    v-model="menu"
    persistent="true"
    no-click-animation
    transition="false"
    :close-on-content-click="false"
    :close-on-back="false">
    <template v-slot:activator="{ props }">
      <v-chip rounded variant="outlined" v-bind="props" width="100%"
              :disabled="disable">{{
          label
        }}
      </v-chip>
    </template>
    <v-card>
      <div class="d-flex justify-end">
        <v-icon icon="mdi-close" @click="menu = false"></v-icon>
      </div>
      <TaskSelector v-if="!disable" @task-selected="addStep" :base-network-filter="baseNetworkFilter"
                    :exclude-jobs="excludeJobs"/>
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
  name: "TaskSelectorChip",
  emits: ['addStep'],
  components: {
    TaskSelector,
    TaskChip,
    draggable,
  },
  watch: {},
  props: {
    excludeJobs: {
      type: Array as PropType<TaskJob[]>,
      required: false,
    },
    baseNetworkFilter: {
      type: String as PropType<Network>,
      required: false,
    },
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

