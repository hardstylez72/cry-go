<template>
  <div>
    <NavBar title="Процессы">
      <template v-slot:default>
        <v-btn @click="stopAll" variant="elevated" class="mx-1" :loading="stopLoading" color="red-lighten-3">
          <v-icon v-if="isMobile()" size="30" icon="mdi-stop-circle-outline" color="white"/>
          <span v-else>Остановить всех</span>
        </v-btn>
        <v-btn @click="resumeAll" variant="elevated" class="mx-1" :loading="resumeLoading" color="green">
          <v-icon v-if="isMobile()" size="30" icon="mdi-play-circle-outline"/>
          <span v-else>Возобновить всех</span>
        </v-btn>
      </template>
    </NavBar>
    <Loader v-if="loading"/>
    <div v-else>
      <div class="d-flex justify-start flex-wrap">
        <div class="mx-2" style="width: 80px">
          <v-checkbox color="blue" v-model="StatusReadyW" direction="horizontal" hide-details
                      density="compact">
            <template v-slot:label>
              <StatusCard :status="ProcessStatus.StatusReady"/>
            </template>
          </v-checkbox>
        </div>
        <div class="mx-2" style="width: 80px">
          <v-checkbox color="blue" v-model="StatusStopW" direction="horizontal" hide-details
                      density="compact">
            <template v-slot:label>
              <StatusCard :status="ProcessStatus.StatusStop"/>
            </template>
          </v-checkbox>
        </div>
        <div class="mx-2" style="width: 80px">
          <v-checkbox color="blue" v-model="StatusErrorW" direction="horizontal" hide-details
                      density="compact">
            <template v-slot:label>
              <StatusCard :status="ProcessStatus.StatusError"/>
            </template>
          </v-checkbox>
        </div>
        <div class="mx-2" style="width: 90px">
          <v-checkbox color="blue" v-model="StatusRunningW" direction="horizontal" hide-details
                      density="compact">
            <template v-slot:label>
              <StatusCard :status="ProcessStatus.StatusRunning"/>
            </template>
          </v-checkbox>
        </div>
        <v-checkbox class="mx-2" color="blue" v-model="StatusDoneW" direction="horizontal" hide-details
                    density="compact" style="width: 50px">
          <template v-slot:label>
            <StatusCard :status="ProcessStatus.StatusDone"/>
          </template>
        </v-checkbox>
      </div>

      <div v-if="noProcesses" class="my-8 mx-8 text-center">
        У вас нет активных процессов, можете создать новый в разделе "Конструктор"
      </div>
      <v-list v-else max-width="96vw" class="px-5">
        <v-list-item
          density="compact"
          variant="outlined"
          class="my-1 my-4"
          v-for="item in getList"
          :key="item.id"
          rounded
          height="auto"
          elevation="1"
          style="border: 0px solid "
        >
          <div class="d-flex justify-space-between align-center">
            <div class="d-inline-flex" @click="viewProcess(item.id)" style="cursor: pointer">
              <StatusCard :status="item.status" class="mx-2"/>
              <div class="mr-2">
                <v-progress-linear
                  :model-value="item.progress"
                  color="blue"
                  height="25"
                  style="width: 100px"
                >
                  <template v-slot:default="{ value }">
                    <strong>{{ Math.ceil(value) }}%</strong>
                  </template>
                </v-progress-linear>
              </div>
              <div class="mr-2 text-h6 text-blue-darken-1"><b>{{ `${item.flowLabel}` }}</b></div>
              <div class="mr-2" v-if="item.runAfter && item.status === ProcessStatus.StatusReady"> Запуск запланирован
                {{ formatTime(item.runAfter) }}
              </div>
              <div class="mr-2">Created: {{ dayjs(item.createdAt).format('YYYY-MM-DD HH:mm:ss') }}</div>
            </div>
            <DeleteProcess class="mx-1" :process-id="item.id" @processRemoved="processRemoved"/>

          </div>

          <div class="mr-2">Profiles: <b>{{ getProfiles(item) }}</b></div>
        </v-list-item>
        <v-btn v-if="showNext" @click="next" :loading="nextLoading" width="100%">More</v-btn>
      </v-list>

    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {processService} from "@/generated/services"
import CreateFlow from "@/components/flow/OldCreateFlow.vue";
import CheckBox from "@/components/CheckBox.vue";
import {Process, ProcessStatus} from "@/generated/process";
import StatusCard from "@/components/StatusCard.vue";
import dayjs from "dayjs";
import BtnProcessStopResume from "@/components/BtnProcessStopResume.vue";
import {getFlow} from "@/components/tasks/tasks";
import Loader from "@/components/Loader.vue";
import NavBar from "@/components/NavBar.vue";
import {Delay, formatTime, isMobile, Timer} from "@/components/helper";
import {fi} from "vuetify/locale";
import DeleteProcess from "@/components/process/DeleteProcess.vue";

export default defineComponent({
  name: "Processes",
  watch: {
    StatusRetryW: {
      handler() {
        this.statusChanged()
      }
    },
    StatusStopW: {
      handler() {
        this.statusChanged()
      }
    },
    StatusErrorW: {
      handler() {
        this.statusChanged()
      }
    },
    StatusRunningW: {
      handler() {
        this.statusChanged()
      }
    },
    StatusDoneW: {
      handler() {
        this.statusChanged()
      }
    },
    StatusReadyW: {
      handler() {
        this.statusChanged()
      }
    },
  },
  components: {
    DeleteProcess,
    NavBar,
    Loader,
    BtnProcessStopResume,
    StatusCard,
    CheckBox,
    CreateFlow
  },
  props: {},
  data() {
    return {

      offset: 0,
      loading: false,
      loadingError: false,
      list: [] as Process[],
      showCreateFlowDialog: false,
      selected: new Set<string>(),
      nextLoading: false,

      resumeLoading: false,
      stopLoading: false,
      timer: new Timer()
    }
  },
  computed: {
    StatusStopW: {
      async set(id: string) {
        this.$router.push({query: {...this.$route.query, stop: id}})
      },
      get(): boolean {
        const s = this.$route.query.stop
        if (s === null || Array.isArray(s)) {
          return false
        }
        if (s === 'true') {
          return true
        }

        if (s === undefined) {
          this.$router.push({query: {...this.$route.query, stop: 'true'}})
        }

        return false
      }
    },
    StatusErrorW: {
      async set(id: string) {
        this.$router.push({query: {...this.$route.query, error: id}})
      },
      get(): boolean {
        const s = this.$route.query.error
        if (s === null || Array.isArray(s)) {
          return false
        }

        if (s === undefined) {
          this.$router.push({query: {...this.$route.query, error: 'true'}})
        }

        if (s === 'true') {
          return true
        }
        return false
      }
    },
    StatusRunningW: {
      async set(id: string) {
        this.$router.push({query: {...this.$route.query, running: id}})
      },
      get(): boolean {
        const s = this.$route.query.running
        if (s === null || Array.isArray(s)) {
          return false
        }

        if (s === undefined) {
          this.$router.push({query: {...this.$route.query, running: 'true'}})
        }

        if (s === 'true') {
          return true
        }
        return false
      }
    },
    StatusDoneW: {
      async set(id: string) {
        this.$router.push({query: {...this.$route.query, done: id}})
      },
      get(): boolean {
        const s = this.$route.query.done
        if (s === null || Array.isArray(s)) {
          return false
        }

        if (s === undefined) {
          this.$router.push({query: {...this.$route.query, done: 'false'}})
        }

        if (s === 'true') {
          return true
        }
        return false
      }
    },
    StatusReadyW: {
      async set(id: string) {
        this.$router.push({query: {...this.$route.query, ready: id}})
      },
      get(): boolean {
        const s = this.$route.query.ready
        if (s === null || Array.isArray(s)) {
          return false
        }

        if (s === undefined) {
          this.$router.push({query: {...this.$route.query, ready: 'false'}})
        }

        if (s === 'true') {
          return true
        }
        return false
      }
    },
    noProcesses() {
      return this.getList.length === 0
    },
    showNext(): boolean {
      if (this.list.length === 0) {
        return false
      }
      return this.list.length % 15 === 0
    },
    ProcessStatus() {
      return ProcessStatus
    },
    dayjs() {
      return dayjs
    },
    getList(): Process[] {
      return this.list
    },
    selectedSome(): boolean {
      return this.selected.size > 0
    }
  },
  methods: {
    processRemoved() {
      this.statusChanged()
    },
    formatTime,
    statusChanged() {
      this.timer.add(100)
      this.timer.cb(() => {
        this.reset()
        this.UpdateList()
      })
    },
    isMobile,
    getFlow,
    viewProcess(id: string) {
      this.$router.push({name: 'ViewProcess', params: {id: id}});
    },
    getProfiles(p: Process): string {
      let labels: number[] = []
      p.profiles.forEach((pp) => {
        labels.push(Number(pp.profileLabel))
      })

      const ranges = this.mergeIntervals(labels)

      let out = ""
      ranges.forEach((range) => {
        if (range.length === 1) {
          out += ", " + range[0]
        } else {
          out += `, [${range[0]}-${range[1]}]`
        }
      })
      out = out.substring(2)

      return out
    },
    mergeIntervals(array: number[]): number[][] {

      if (array.length === 1) {
        return [array]
      }

      array = array.sort((a, b) => a - b)

      const out: number[][] = []

      let start = 0
      let end = 0
      let prev = 0

      for (let i = 0; i <= array.length; i++) {

        if (i === 0) {
          prev = i
          start = prev
          continue
        }

        while (true) {
          if (array[i] - array[prev] === 1) {
            end = i
            prev = i
            i++
          } else {
            break
          }
        }

        if (start === end) {
          out.push([array[start]])
          if (end + 1 < array.length) {
            end += 1
            start = end
            prev = end
          }
        } else {
          out.push([array[start], array[end]])
          if (end + 1 < array.length) {
            start = end + 1
            end = i
            prev = start
          }
        }
      }
      return out
    },
    next() {
      this.offset += 15
      this.UpdateList()
    },
    reset() {
      this.offset = 0
      this.list = []
    },
    processProgress(p: Process): number {
      let totalProfiles = p.profiles.length
      if (p.profiles[0]) {
        let totalTask = p.profiles[0].tasks.length
        let total = totalProfiles * totalTask
        let finished = 0

        p.profiles.forEach((p) => {
          p.tasks.forEach((t) => {
            if (t.status == ProcessStatus.StatusDone) {
              finished++
            }
          })
        })

        if (finished == 0) {
          return 0
        }

        return Math.ceil(finished / total)
      }
      return 0
    },
    viewRoute(id: string) {
      window.open("/process/" + id, "_blank")
    },
    async resumeAll() {
      try {
        this.resumeLoading = true
        await processService.processServiceResumeAllProcess()
        this.StatusStopW = true
        await Delay(10)
        this.StatusErrorW = false
        await Delay(10)
        this.StatusRunningW = true
        await Delay(10)


        this.StatusDoneW = false
        await Delay(10)
        this.StatusReadyW = false
        await Delay(10)

        this.statusChanged()
      } catch (e) {

      } finally {
        this.resumeLoading = false
      }
    },
    async stopAll() {
      try {
        this.stopLoading = true
        await processService.processServiceStopAllProcess()
        this.StatusStopW = true
        await Delay(10)
        this.StatusRunningW = true
        await Delay(10)
        this.StatusErrorW = false
        await Delay(10)

        this.StatusDoneW = false
        await Delay(10)
        this.StatusReadyW = false
        await Delay(10)

        this.statusChanged()
      } catch (e) {

      } finally {
        this.stopLoading = false
      }
    },
    getStatuses() {
      const statuses: ProcessStatus[] = []
      if (this.StatusDoneW) {
        statuses.push(ProcessStatus.StatusDone)
      }
      if (this.StatusReadyW) {
        statuses.push(ProcessStatus.StatusReady)
      }
      if (this.StatusErrorW) {
        statuses.push(ProcessStatus.StatusError)
      }
      if (this.StatusRetryW) {
        statuses.push(ProcessStatus.StatusRetry)
      }
      if (this.StatusStopW) {
        statuses.push(ProcessStatus.StatusStop)
      }
      if (this.StatusRunningW) {
        statuses.push(ProcessStatus.StatusRunning)
      }
      return statuses
    },
    async UpdateList() {
      try {

        this.list.length ? this.nextLoading = true : this.loading = true
        this.loadingError = false

        const res = await processService.processServiceListProcess({
          body: {
            offset: this.offset.toString(),
            statuses: this.getStatuses(),
          }
        })
        this.list.push(...res.processes)
        this.selected = new Set<string>()
      } catch (e) {
        this.loadingError = true
      } finally {
        this.nextLoading = false
        this.loading = false
      }
    },
    async cancel(id: string) {
      try {
        await processService.processServiceCancelProcess({body: {processId: id}})
      } finally {
        await this.UpdateList()
      }
    },
    processUpdated() {
      this.UpdateList()
    }
  },
  async created() {
    await this.statusChanged()
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

