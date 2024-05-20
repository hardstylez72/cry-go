<template>
  <v-menu
    v-model="show"
    location="end"
    :open-on-click="true"
    :open-on-hover="true"
    width="350px"
    z-index="2"

  >
    <template v-slot:activator="{ props }">
      <v-icon v-bind="props" class="mx-1" @click="show = !show" icon="mdi-information-outline"/>
    </template>

    <template v-slot:default>
      <v-card width="300px" style="white-space: pre-wrap; line-break: auto" class="px-2 py-2">
        <v-card-text>
          <slot name="default"/>
        </v-card-text>
      </v-card>
    </template>


  </v-menu>
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

})


</script>


<style>


</style>

