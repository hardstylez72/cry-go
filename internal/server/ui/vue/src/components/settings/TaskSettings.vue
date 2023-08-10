<template>
  <div class="d-inline-flex text-center">
    <div class="mt-2"><b>{{ taskType }}</b> <i>slippage</i></div>

    <v-radio-group direction="horizontal" inline hide-details v-model="settings.slippage">
      <v-radio value="2" label="2%"/>
      <v-radio value="1" label="1%"/>
      <v-radio value="0.5" label="0.5%"/>
      <v-radio value="0.1" label="0.1%"/>
      <v-radio value="0" label="0%"/>
    </v-radio-group>
  </div>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {Network, TaskType} from "@/generated/flow";
import {TaskSettings, Timer} from "@/components/helper";
import deepEqual from "deep-equal";
import {taskProps} from "@/components/tasks/tasks";

export default defineComponent({
  name: "TaskSettings",
  watch: {
    settings: {
      handler() {
        if (!deepEqual(this.settings, this.orig)) {
          this.$emit('updated', this.taskType, this.settings)
        }
      },
      deep: true
    }
  },
  props: {
    settingsProp: {
      type: Object as PropType<TaskSettings>,
      required: true,
    },
    network: {
      type: Object as PropType<Network>,
      required: true,
    },
    taskType: {
      type: String as PropType<TaskType>,
      required: true
    }
  },
  emits: ['updated'],
  data() {
    return {
      priceUSD: "-1",
      timer: new Timer(),
      settings: null as TaskSettings,
      orig: null as TaskSettings,
    }
  },

  created() {
    this.orig = Object.assign({}, this.settingsProp)
    this.settings = Object.assign({}, this.settingsProp)
  }
})


</script>


<style>


</style>

