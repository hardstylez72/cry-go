<template>
  <v-card rounded variant="elevated" density="compact" elevation="5">
    <v-card-title>
      {{ block.weight }}) Мужицкий блок
      <ManBlockDetails :disable="disable" :block="block" @flowChanged="flowChanged"/>
    </v-card-title>
    <v-card-text>
      <Preview :data="preview"/>
    </v-card-text>
  </v-card>

</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {FlowBlock, FlowPreviewRes} from "@/generated/flow";
import draggable from 'vuedraggable'
import TaskChip from "@/components/tasks/TaskChip.vue";
import TaskSelector from "@/components/flow/TaskSelector.vue";
import TaskSelectorMenu from "@/components/flow/TaskSelectorBtnMenu.vue";
import TaskSelectorChip from "@/components/flow/TaskSelectorChipMenu.vue";
import ManBlockDetails from "@/components/flow/ManBlockDetails.vue";
import Preview from "@/components/flow/Preview.vue";
import {flowService} from "@/generated/services";
import {Timer} from "@/components/helper";

export default defineComponent({
  name: "FormV2",
  components: {
    Preview,
    ManBlockDetails,
    TaskSelectorChip,
    TaskSelectorMenu,
    TaskSelector,
    TaskChip,
    draggable,
  },
  props: {
    disable: {
      type: Boolean as PropType<boolean>,
      required: false,
    },
    block: {
      type: Object as PropType<FlowBlock>,
      required: true
    },
  },
  watch: {
    block: {
      handler() {
        this.localBlock = this.block
        this.loadPreview()
      },
      deep: true,
      immediate: true
    }
  },
  methods: {
    flowChanged(block: FlowBlock) {
      this.$emit('flowChanged', block)
      this.localBlock = block
    },
    async loadPreview() {

      this.timer.add(100)
      this.timer.cb(() => {

        this.previewError = ''
        flowService.flowServiceFlowPreview({
          body: {
            label: '',
            blocks: [
              this.localBlock,
            ]
          }
        }).then((data) => {
          this.preview = data
        }).catch(() => {
          this.previewError = 'Ошибка'
        })

      })
    },
  },
  emits: ['flowChanged'],
  data() {
    return {
      localBlock: {man: {tasks: [], randomTasks: []}} as FlowBlock,
      previewError: '',
      preview: {} as FlowPreviewRes,
      timer: new Timer(),
    }
  }
})


</script>


<style>

</style>

