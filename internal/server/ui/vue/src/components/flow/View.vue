<template>
  <div>
    <NavBar :title="`Сценарий: ${label}`">
      <template v-slot:default>

        <v-btn @click="shareFlow" variant="flat" class="mx-1">
          <v-icon v-if="isMobile()" icon="mdi-share-variant"/>
          <span v-else>Поделиться</span>
        </v-btn>

        <v-btn @click="copyFlow" variant="flat" class="mx-1">
          <v-icon v-if="isMobile()" icon="mdi-content-copy"/>
          <span v-else>Копия</span>
        </v-btn>

        <span v-if="!deletedAt && !nextId">
        <v-btn :disabled="selectedProfiles.length > 0" v-if="!editMode" variant="flat" class="mx-1" focused
               @click="editModeChanged">
          <v-icon v-if="isMobile()" icon="mdi-note-edit-outline"/>
          <span v-else>Изменить</span>
        </v-btn>

        <v-btn v-else @click="updateFlow" variant="flat" class="mx-1">
          <v-icon v-if="isMobile()" icon="mdi-content-save"/>
          <span v-else>Сохранить</span>
        </v-btn>
        </span>

        <v-btn variant="flat" :disabled="selectedProfiles.length > 0" v-if="!deletedAt && !nextId"
               color="red" @click="DeleteFlow">
          <v-icon v-if="isMobile()" icon="mdi-delete"/>
          <span v-else>Удалить</span>
        </v-btn>
      </template>
    </NavBar>
    <div class="mx-8 my-3" style="width: 80vw">
      <ProfileSearch
        v-model="selectedProfiles"
        :disabled="editMode"
      />
      <div class="ml-3 mb-2 d-inline-flex flex-wrap" v-for="profile in getProfileList">
        <ProfileCard :label="profile.num" :profile-id="profile.id"/>
      </div>
    </div>


    <div class="d-flex flex-wrap mx-8">
      <!--      <v-checkbox label="запустить параллельно" v-model="parallel"/>-->
      <v-btn
        width="100%"
        :disabled="selectedProfiles.length === 0" @click="startProcess">
        Запустить сейчас
      </v-btn>
      <div class="my-3" style="width: 100%;">
        <Schedule :disabled="selectedProfiles.length === 0" @bobaSatisfied="bobaSatisfied"/>
      </div>
    </div>
  </div>

  <v-spacer class="my-6"/>

  <div class="px-4 py-4">

    <v-form ref="form">
      <div class="my-3 ml-4">
        <v-btn v-if="!disabled" @click="addManBlock" class="mx-1 font-weight-bold">Мужцкий блок 💪💪💪
        </v-btn>
        <v-btn v-if="!disabled" @click="addRandBlock" class="mx-1 font-weight-bold">Рандомизированный
          блок
        </v-btn>
      </div>


      <v-text-field
        v-model="label"
        label="Название сценария"
        density="comfortable"
        variant="outlined"
        :rules="[required]"
        :disabled="disabled"
      ></v-text-field>

      <div v-for="block in blocks">
        <component
          class="my-2"
          v-if="block.man"
          is="ManBlock"
          :disable="disabled"
          :block="block"
          @flow-changed="flowChanged"
          @removeBlock="removeBlock"
        />
        <component
          class="my-2"
          v-if="block.rand"
          is="RandomBlock"
          :disable="disabled"
          :block="block"
          @block-changed="flowChanged"
          @removeBlock="removeBlock"
        />
      </div>

    </v-form>

    <Preview :data="preview">
      <template v-slot:default>
        <h3 class="my-1">Возможные финальные комбинации</h3>
      </template>
    </Preview>

  </div>

</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService, processService, profileService} from "@/generated/services"
import {flow_Flow as Flow, FlowBlock, FlowPreviewRes, Task, TaskType} from "@/generated/flow";
import {Profile} from "@/generated/profile";
import ProfileCard from "@/components/profile/ProfileCard.vue";
import {Delay, formatTime, getDate, getTime, isMobile, Timer, ts} from "@/components/helper";
import FlowForm from "@/components/flow/OldForm.vue";
import ProfileSearch from "@/components/profile/ProfileSearch.vue";
import NavBar from "@/components/NavBar.vue";
import {required} from "@/components/tasks/helper";
import Schedule from "@/components/flow/Schedule.vue";
import Preview from "@/components/flow/Preview.vue";

import RandomBlock from "@/components/flow/RandomBlock.vue";
import ManBlock from "@/components/flow/ManBlock.vue";

export default defineComponent({
  name: "FlowViewV2",
  components: {Preview, Schedule, NavBar, ProfileSearch, ProfileCard, RandomBlock, ManBlock},

  props: {
    propId: {
      type: String,
      required: false,
    }
  },
  data() {
    return {
      selectedProfiles: [] as Profile[],
      editMode: false,
      flowId: "" as string,
      disabled: true,
      deletedAt: null as null | Date,
      nextId: null as null | string,
      showUpdateBtn: false,
      updatingFlow: false,
      timer: new Timer(),
      flowLoading: true,

      blockMap: new Map<number, FlowBlock>(),
      preview: null as null | FlowPreviewRes,
      previewError: '',
      label: '',
      parallel: false,
    }
  },
  computed: {
    getProfileList() {
      return this.selectedProfiles.map(e => e)
    },
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
    formatTime,
    required,
    isMobile,
    async shareFlow() {
      try {
        const res = await flowService.flowServiceShareFlow({body: {id: this.flowId, description: ''}})
        this.$router.push(`/shared-flow/${res.id}`, {force: true})
      } finally {

      }
    },
    async copyFlow() {
      try {
        const res = await flowService.flowServiceCopyFlow({body: {id: this.flowId}})
        this.$router.push(`/flow/random/${res.id}`, {force: true})
        window.location.assign(`/flow/random/${res.id}`)
      } finally {

      }
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
    async loadPreview() {

      this.timer.add(100)
      this.timer.cb(() => {

        this.previewError = ''
        flowService.flowServiceFlowPreview({
          body: {
            label: '',
            blocks: this.blocks,
          }
        }).then((data) => {
          this.preview = data
        }).catch((err) => {
          this.previewError = 'Ошибка'
        })

      })
    },
    async flowChanged(block: FlowBlock) {
      await this.validateForm()
      this.blockMap.set(block.weight, block)
      await this.loadPreview()
    },
    async updateFlow() {
      if (!(await this.validateForm())) {
        return
      }
      try {
        this.updatingFlow = true

        const res = await flowService.flowServiceUpdateFlowV2({
          body: {
            blocks: this.blocks,
            label: this.label,
            id: this.flowId,
          }
        })
        this.flowId = res.id
        this.editModeChanged()
        this.$router.push({name: 'FlowViewV2', params: {id: this.flowId}})

      } finally {
        this.updatingFlow = false
      }
    },
    editModeChanged() {
      this.editMode = !this.editMode
      this.disabled = !this.editMode
      this.showUpdateBtn = this.editMode
    },
    getCheckboxLabel(): string {
      return this.editMode ? "sw" : 'switch tp editing mode'
    },
    async bobaSatisfied(runAfterList: Date[]) {

      for (let i = 0; i < runAfterList.length; i++) {
        const runAfter = runAfterList[i]
        await processService.processServiceCreateProcess({
          body: {
            runAfter: runAfter,
            flowId: this.flowId,
            profileIds: this.selectedProfiles.map((p) => p.id),
            parallel: this.parallel,
          }
        })
      }

      this.$router.push({
        name: "Processes",
        query: {ready: 'true', done: false, error: false, running: false, stop: false}
      })
    },
    async startProcess() {
      const res = await processService.processServiceCreateProcess({
        body: {
          flowId: this.flowId,
          profileIds: this.selectedProfiles.map((p) => p.id),
          parallel: this.parallel,
        }
      })
      this.$router.push({name: "ViewProcess", params: {id: res.process.id}})
    },
    async DeleteFlow() {
      try {
        await flowService.flowServiceDeleteFlow({body: {id: this.flowId}})
        this.$router.push({name: 'Constructor'})
      } catch (e) {

      }
    },
    async loadFlow() {
      if (this.propId) {
        this.flowId = this.propId
      } else if (this.$route.params.id) {
        this.flowId = this.$route.params.id.toString()
      }

      try {
        this.flowLoading = true
        const res = await flowService.flowServiceGetFlowV2({body: {id: this.flowId}})

        this.blockMap.clear()
        res.blocks.forEach((block) => {
          this.blockMap.set(block.weight, block)
        })
        this.label = res.label

      } finally {
        this.flowLoading = false
      }
    }
  },
  async mounted() {
    await this.loadFlow()
    await this.loadPreview()

  }
})


</script>


<style scoped>

</style>

