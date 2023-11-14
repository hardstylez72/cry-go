<template>
  <v-select
    density="compact"
    variant="outlined"
    :label="label"
    :rules="[required]"
    :items="items"
    v-model="selected"
    :disabled="disabled"
  >
    <template v-slot:item="{item, props}">
      <v-list-item v-bind="props" title="">
        <v-list-item-title class="d-inline-flex">
          <v-img :src="networkProps[item.title].img" width="20px" class="mx-1"/>
          {{ networkProps[item.title].name }}
        </v-list-item-title>
      </v-list-item>
    </template>

    <template v-slot:selection="{item}">
      <div class="d-inline-flex">
        <v-img :src="networkProps[item.title].img" width="20px" class="mx-1"/>
        {{ networkProps[item.title].name }}
      </div>
    </template>
  </v-select>
</template>

<script lang="ts">
import {Amount, AmUni, Network, TaskType} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {processService} from "@/generated/services";
import {EstimationTx} from "@/generated/process";
import {networkProps, percent} from "@/components/helper";
import {ca} from "vuetify/locale";
import {taskProps} from "@/components/tasks/tasks";
import {required} from "@/components/tasks/helper";

export default defineComponent({
  name: "NetworkSelector",
  methods: {required},
  emits: ['update:modelValue'],
  computed: {
    networkProps() {
      return networkProps
    },
    selected: {
      get(): Network {
        if (!this.modelValue) {
          return null
        }
        return this.modelValue
      },
      set(n: Network) {
        this.$emit('update:modelValue', n)
      }
    }
  },
  props: {
    modelValue: {
      type: String as PropType<Network | null>,
      required: true
    },
    items: {
      type: Array as PropType<Network[]>,
      required: true,
    },
    label: {
      type: String,
      required: true,
    },
    disabled: {
      type: Boolean,
      required: false
    },
  },
  data() {
    return {}
  },
})
</script>

<style scoped>

</style>
