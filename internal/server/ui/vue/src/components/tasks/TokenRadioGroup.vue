<template>
  <v-radio-group density="compact" :inline="true" :rules="[required]" v-model="selected" :disabled="disabled"
                 :label="label">
    <v-radio v-for="item in items" :value="item" class="mx-10">
      <template v-slot:default>

        <v-img :src="tokenProps[item].img" width="20px" class="mx-1"/>
        {{ tokenProps[item].name }}
      </template>
    </v-radio>
  </v-radio-group>
</template>

<script lang="ts">
import {Amount, AmUni, Network, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import {processService} from "@/generated/services";
import {EstimationTx} from "@/generated/process";
import {networkProps, percent} from "@/components/helper";
import {ca} from "vuetify/locale";
import {taskProps} from "@/components/tasks/tasks";
import {required} from "@/components/tasks/helper";
import {tokenProps} from "@/components/tokens";

export default defineComponent({
  name: "TokenRadioGroup",
  methods: {required},
  emits: ['update:modelValue'],
  computed: {
    tokenProps() {
      return tokenProps
    },
    selected: {
      get(): Token {
        if (!this.modelValue) {
          return null
        }
        return this.modelValue
      },
      set(n: Token) {
        this.$emit('update:modelValue', n)
      }
    }
  },
  props: {
    modelValue: {
      type: String as PropType<Token | null>,
      required: true
    },
    items: {
      type: Array as PropType<Token[]>,
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
