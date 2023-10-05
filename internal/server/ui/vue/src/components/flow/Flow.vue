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
    <div>
      <v-checkbox class="mx-8" v-model="runAfter" label="Запланировать запуск процесса"/>

      <v-list v-if="runAfter" variant="plain" rounded>
        <v-list-item elevation="1" width="95%" v-for="(item, i) in runAfterSchedule" height="auto">
          <div class="d-inline-flex align-center">
            <div style="max-width: 200px" class="mx-1 my-2">
              <v-text-field :rules="[required]" type="date" v-model="item.date" label="Дата запуска" density="compact"
                            variant="outlined" hide-details/>
            </div>
            <div style="max-width: 220px; min-width: 140px" class="mx-1 my-2">
              <v-text-field :rules="[required]" type="time" v-model="item.time" label="Время запуска" density="compact"
                            variant="outlined" hide-details/>
            </div>
            <v-icon v-if="runAfterSchedule.length -1 === i" icon="mdi-plus" color="green" @click="addRunAfterTime"/>
            <v-icon v-if="runAfterSchedule.length -1 === i && i !== 0" icon="mdi-close" color="red"
                    @click="removeRunAfterTime"/>
            <i class="mx-1">{{ formatTime(runAfterList[i]) }}</i>

          </div>


        </v-list-item>
      </v-list>

      <v-btn
        class="mx-8"
        style="width: 80vw"
        :disabled="selectedProfiles.length === 0"
        @click="startProcess"
        v-if="!runAfter"
      >
        Создать процесс
      </v-btn>
      <v-btn
        class="mx-8"
        style="width: 80vw"
        :disabled="selectedProfiles.length === 0"
        @click="startProcess"
        v-if="runAfter"
      >
        Запланировать процесс
      </v-btn>
    </div>
  </div>


  <v-spacer class="my-6"/>
  <v-form validate-on="submit" ref="flow-form">
    <FlowForm v-if="!flowLoading" :disable="disabled" :label-value="flow.label" :tasks-value="tasks"
              @flow-changed="flowChanged"/>
  </v-form>

</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService, processService, profileService} from "@/generated/services"
import {flow_Flow as Flow, Task, TaskType} from "@/generated/flow";
import TaskStargateBridge from "@/components/tasks/BridgeStargate/Block.vue";
import {Profile} from "@/generated/profile";
import TaskDelay from "@/components/tasks/block/Delay.vue";
import ProfileCard from "@/components/profile/ProfileCard.vue";
import {taskComponentMap, TaskArg, taskTypes} from "@/components/tasks/tasks";
import {Delay, formatTime, getDate, getTime, isMobile, Timer, ts} from "@/components/helper";
import FlowForm from "@/components/flow/FlowForm.vue";
import ProfileSearch from "@/components/profile/ProfileSearch.vue";
import NavBar from "@/components/NavBar.vue";
import {required} from "@/components/tasks/menu/helper";

export default defineComponent({
  name: "Flow",
  components: {NavBar, ProfileSearch, FlowForm, ProfileCard},
  watch: {
    runAfter: {
      handler() {
        if (this.runAfter) {
          this.addRunAfterTime()
        }
      }
    },
    runAfterSchedule: {
      handler() {
        this.runAfterList = []
        this.runAfterSchedule.forEach(e => {
          this.runAfterList.push(ts(e.date, e.time))
        })
      },
      deep: true,
    }
  },
  props: {
    propId: {
      type: String,
      required: false,
    }
  },
  data() {
    return {
      runAfterSchedule: [],
      runAfterList: [],
      runAfter: false,
      selectedProfiles: [] as Profile[],
      editMode: false,
      flowId: "" as string,
      flow: {} as Flow,
      disabled: true,
      tasks: [] as Task[],
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
    removeRunAfterTime() {
      this.runAfterSchedule.pop()
    },
    addRunAfterTime() {
      const date = getDate()
      const time = getTime()
      this.runAfterSchedule.push({date: date, time: time})
    },
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
    flowChanged(label: string, tasks: TaskArg[]) {
      this.flow.label = label
      this.tasks = []
      tasks.forEach(t => {
        if (t.task) {
          this.tasks.push(t.task)
        }
      })
    },
    async updateFlow() {
      if (!(await this.validateForm())) {
        return
      }
      try {
        this.updatingFlow = true
        this.flow.tasks = this.tasks

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
    async startProcess() {


      if (!this.runAfter) {
        const res = await processService.processServiceCreateProcess({
          body: {
            flowId: this.flowId,
            profileIds: this.selectedProfiles.map((p) => p.id),
          }
        })
        this.$router.push({name: "ViewProcess", params: {id: res.process.id}})
      } else {
        this.runAfterList.forEach(runAfter => {
          const res = processService.processServiceCreateProcess({
            body: {
              runAfter: runAfter,
              flowId: this.flowId,
              profileIds: this.selectedProfiles.map((p) => p.id),
            }
          })
        })

        this.$router.push({
          name: "Processes",
          query: {ready: 'true', done: false, error: false, running: false, stop: false}
        })
      }


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

.header {
  margin: auto;
  height: auto;
  z-index: 100;
  right: 2px;
  left: 56px;
  background-color: #FFF3E0;
  box-shadow: rgba(0, 0, 0, 0.16) 0px 3px 6px, rgba(0, 0, 0, 0.23) 0px 3px 6px;
}
</style>

