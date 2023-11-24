<template>
  <div>
    <NavBar :title="`Сценарий: ${flow.label}`">
      <template v-slot:default>

        <v-btn @click="shareFlow" variant="flat" class="mx-1">
          <v-icon v-if="isMobile()" icon="mdi-share-variant"/>
          <span v-else>Поделиться</span>
        </v-btn>

        <v-btn @click="copyFlow" variant="flat" class="mx-1">
          <v-icon v-if="isMobile()" icon="mdi-content-copy"/>
          <span v-else>Копия</span>
        </v-btn>

        <span v-if="!flow.deletedAt && !flow.nextId">
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

        <v-btn variant="flat" :disabled="selectedProfiles.length > 0" v-if="!flow.deletedAt && !flow.nextId"
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
  <v-form validate-on="submit" ref="flow-form">
    <FlowForm v-if="!flowLoading" :disable="disabled" :label-value="flow.label" :tasks-value="tasks"
              @flow-changed="flowChanged" :random-tasks-value="randomTasks"/>
  </v-form>

</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService, processService, profileService} from "@/generated/services"
import {flow_Flow as Flow, Task, TaskType} from "@/generated/flow";
import {Profile} from "@/generated/profile";
import ProfileCard from "@/components/profile/ProfileCard.vue";
import {taskComponentMap, TaskArg, taskTypes} from "@/components/tasks/tasks";
import {Delay, formatTime, getDate, getTime, isMobile, Timer, ts} from "@/components/helper";
import FlowForm from "@/components/flow/OldForm.vue";
import ProfileSearch from "@/components/profile/ProfileSearch.vue";
import NavBar from "@/components/NavBar.vue";
import {required} from "@/components/tasks/helper";
import Schedule from "@/components/flow/Schedule.vue";

export default defineComponent({
  name: "Flow",
  components: {Schedule, NavBar, ProfileSearch, FlowForm, ProfileCard},

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
      flow: {} as Flow,
      disabled: true,
      tasks: [] as Task[],
      randomTasks: [] as Task[],
      showUpdateBtn: false,
      updatingFlow: false,
      timer: new Timer(),
      flowLoading: true
    }
  },
  computed: {
    getProfileList() {
      return this.selectedProfiles.map(e => e)
    }
  },
  methods: {
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
        this.$router.push(`/flow/${res.id}`, {force: true})
        window.location.assign(`/flow/${res.id}`)
      } finally {

      }
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['flow-form'].validate()

      return valid
    },
    flowChanged(label: string, tasks: TaskArg[], randomTasks: TaskArg[]) {
      this.flow.label = label
      this.tasks = []
      tasks.forEach(t => {
        if (t.task) {
          this.tasks.push(t.task)
        }
      })

      this.randomTasks = []
      randomTasks.forEach(t => {
        if (t.task) {
          this.randomTasks.push(t.task)
        }
      })

      this.validateForm()
    },
    async updateFlow() {
      if (!(await this.validateForm())) {
        return
      }
      try {
        this.updatingFlow = true
        this.flow.tasks = this.tasks
        this.flow.randomTasks = this.randomTasks

        const res = await flowService.flowServiceUpdateFlow({body: {flow: this.flow}})
        this.flow = res.flow
        this.flowId = this.flow.id
        this.editModeChanged()
        this.$router.push({name: 'Flow', params: {id: this.flow.id}})

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
        const res = await flowService.flowServiceGetFlow({body: {id: this.flowId}})
        this.flow = res.flow
        this.tasks = this.flow.tasks
        this.randomTasks = this.flow.randomTasks
      } finally {
        this.flowLoading = false
      }
    }
  },
  async mounted() {
    await this.loadFlow()

  }
})


</script>


<style scoped>

</style>

