<template>
  <div>
    <div class="mt-2"><b>{{ taskType }}</b></div>
    <div>
      <div class="d-inline-flex text-center align-center">
        <i>Проскальзывание</i>
        <v-text-field
          class="mx-1"
          suffix="%"
          style="width: 100px; height: 50px"
          type="number"
          density="compact"
          variant="outlined"
          :rules="[required, slippage]"
          :error-messages="slippageErrorMessage"
          v-model="settings.slippage"/>

      </div>
    </div>
    <div>
      <div class="d-inline-flex text-center align-center">
        <i class="mr-2">Допустимая разница с курсом бинанса</i>
        <v-text-field
          class="mx-1"
          suffix="%"
          style="width: 100px"
          type="number"
          density="compact"
          :rules="[required]"
          variant="outlined"
          v-model="settings.swapRateRatio"/>
      </div>
    </div>


  </div>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {Network, TaskType} from "@/generated/flow";
import {TaskSettings, Timer} from "@/components/helper";
import deepEqual from "deep-equal";
import {taskProps} from "@/components/tasks/tasks";
import {required} from "@/components/tasks/helper";

export default defineComponent({
  name: "TaskSettings",
  methods: {
    required,
    slippage: (v: any) => {
      if (!v) {
        return 'required'
      }

      if (Number.isNaN(Number(v))) {
        return 'invalid number'
      }

      if (Number(v) > 100) {
        return 'must be less than 100%'
      }

      if (Number(v) < 0.001) {
        return 'must be more than 0.001%'
      }

      return true
    }
  },
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
      slippageErrorMessage: '',
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

