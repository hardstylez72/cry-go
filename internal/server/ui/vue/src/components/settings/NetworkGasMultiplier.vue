<template>
  <div>Множитель газа</div>
  <v-slider
    v-model="multiplier"
    density="compact"
    label=""
    min="0.1"
    max="1"
    persistent-hint
    thumb-label="always"
    hide-details
  >
    <template v-slot:append>

      <v-tooltip v-model="hint">
        <template v-slot:activator="{ props }">
          <v-icon v-bind="props" @click="hint = true" icon="mdi-information-outline"></v-icon>
        </template>
        <div>
          Мультипликатор газа, позволяет уменьшать газ за странзакцию.
          <br/>
          GAS = ESTIMATED_GAS * GAS_MULTIPLIER
          <br/>
          Exmaple:
          <br/>
          ESTIMATED_GAS = 1$
          <br/>
          GAS_MULTIPLIER = 0.6
          <br/>
          GAS = 1$ * 0.6 = 0.6$
        </div>
      </v-tooltip>


    </template>
    <template v-slot:thumb-label="{ modelValue }">
      {{ Number(modelValue).toPrecision(2) }}
    </template>
  </v-slider>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {Network} from "@/generated/flow";
import {ratio} from "@/components/helper";

export default defineComponent({
  name: "NetworkGasMultiplier",
  props: {
    modelValue: {
      type: Number,
      required: false,
    },
  },
  emits: ['update:modelValue'],
  data() {
    return {
      hint: false
    }
  },
  computed: {
    Network() {
      return Network
    },
    multiplier: {
      get(): number {
        return this.modelValue ? this.modelValue : 1
      },
      async set(value: number) {
        this.$emit('update:modelValue', Number(value))
      }
    },
  },
  methods: {ratio},
})


</script>


<style>


</style>

