<template>
  <v-tooltip
    v-model="show"
    width="300px"
    height="auto"
    location="end"
    attach
    :open-on-click="true"
  >
    <template v-slot:activator="{ props }">
      <v-icon v-bind="props" class="mx-1" @click="show = !show" icon="mdi-information-outline"/>
    </template>


    {{ text }}


  </v-tooltip>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';

import {GetStatusColor, GetStatusText} from "./helper";
import {errorRef} from "./error-ref";

export default defineComponent({
  name: "Explain",
  data() {
    return {
      show: false
    }
  },
  props: {
    text: {
      type: String,
      required: true,
    },
  },

})


</script>


<style>


</style>

