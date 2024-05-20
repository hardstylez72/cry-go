<template>
  <NavBar title="Создание сценария">
    <template v-slot:default>
      <v-btn
        :disabled="saveLoading"
        :loading="saveLoading"
        variant="flat"
        @click="CreateFlow">
        Сохранить
      </v-btn>
    </template>
  </NavBar>

  <v-card>
    <div>
      <v-btn @click="addManBlock" class="mx-1 font-weight-bold">Мужцкий блок 💪💪💪</v-btn>
      <v-btn @click="addRandBlock" class="mx-1 font-weight-bold">Рандомизированный блок</v-btn>
    </div>

    <v-card-text>
      <v-form ref="form">
        <v-text-field
          v-model="label"
          label="Название сценария"
          density="comfortable"
          variant="outlined"
          :rules="[required]"
          :disabled="disable"
        ></v-text-field>

        <div v-for="block in blocks">
          <component
            class="my-2"
            v-if="block.man"
            is="ManBlock"
            :disable="disable"
            :block="block"
            @flow-changed="flowChanged"
            @removeBlock="removeBlock"
          />
          <component
            class="my-2"
            v-if="block.rand"
            is="RandomBlock"
            :disable="disable"
            :block="block"
            @block-changed="flowChanged"
            @removeBlock="removeBlock"
          />
        </div>
      </v-form>
      <Preview :data="preview"/>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService} from "@/generated/services"
import {
  FlowBlock, FlowPreviewRes,
} from "@/generated/flow";
import NavBar from "@/components/NavBar.vue";
import RandomBlock from "@/components/flow/RandomBlock.vue";
import ManBlock from "@/components/flow/ManBlock.vue";
import Preview from "@/components/flow/Preview.vue";
import {Timer} from "@/components/helper";
import {required} from "@/components/tasks/helper";

export default defineComponent({
  name: "CreateFlowV2",
  components: {Preview, RandomBlock, ManBlock, NavBar},
  data() {
    return {
      blockMap: new Map<number, FlowBlock>(),
      disable: false,
      saveLoading: false,
      preview: null as null | FlowPreviewRes,
      previewError: '',
      timer: new Timer(),
      label: '',
      parallel: false,
    }
  },
  computed: {
    blocks(): FlowBlock[] {
      const blocks = [] as FlowBlock[]
      this.blockMap.forEach((v, k) => {
        blocks.push(v)
      })
      return blocks
    }
  },
  methods: {
    removeBlock(i: number) {
      this.blockMap.delete(i)
      const tmp = new Map<number, FlowBlock>()
      let order = []
      this.blockMap.forEach((k, v) => {
        order.push(k.weight)
      })
      order = order.sort()
      order.forEach((w, index) => {
        index++
        const t = this.blockMap.get(w)
        t.weight = index
        tmp.set(index, t)
      })
      this.blockMap = tmp
    },
    addManBlock() {
      const weight = this.blockMap.size + 1
      this.blockMap.set(weight, {man: {tasks: [], randomTasks: []}, weight: weight})
    },
    addRandBlock() {
      const weight = this.blockMap.size + 1
      this.blockMap.set(weight, {rand: {tasks: []}, weight: weight})
    },
    required,
    async flowChanged(block: FlowBlock) {

      if (!block || !block.weight) {
        return
      }

      this.blockMap.set(Number(block.weight), block)

      await this.validateForm()

      this.loadPreview()
    },
    async loadPreview() {

      this.timer.add(100)
      this.timer.cb(() => {

        this.previewError = ''
        flowService.flowServiceFlowPreview({body: {blocks: this.blocks, label: ''}})
          .then((data) => {
            this.preview = data
          }).catch(() => {
          this.previewError = 'Ошибка'
        })

      })
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
    async CreateFlow() {
      if (!(await this.validateForm())) {
        return
      }

      try {
        this.saveLoading = true

        const res = await flowService.flowServiceCreateFlowV2({
          body: {
            blocks: this.blocks,
            label: this.label,

          }
        })
        this.$router.push({name: "FlowViewV2", params: {id: res.id}})
      } finally {
        this.saveLoading = false
      }
    },
  },
  created() {
    this.blockMap.set(1, {man: {tasks: [], randomTasks: []}, weight: 1})
    this.blockMap.set(2, {rand: {tasks: []}, weight: 2})
  }
})


</script>


<style>


</style>

