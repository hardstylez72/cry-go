<template>
  <v-menu
    v-model="menu"
    persistent
    width="auto"
    :close-on-back="false"
    :close-on-content-click="false"
  >
    <template v-slot:activator="{ props }">


      <TaskBlock v-bind="props" :task="task.task"
                 :color="GetStatusColor(task.status)" style="cursor: pointer"/>

    </template>

    <template v-slot:default="{ props }">
      <v-card width="auto">
        <v-card-text>

          <span class="d-flex justify-space-between">
            <div>
               <h3 class="d-inline-flex">
                   {{ taskSpec(task.task.taskType).service.name }}
                   <a target="_blank" :href="taskSpec(task.task.taskType).service.link" class="mx-1">
                    <v-img height="22px" :src="taskSpec(task.task.taskType).service.img"/>
                   </a>
               </h3>
              <div>
               <StatusCard :status="task.status"/>
               <v-menu v-model="details" :close-on-back="false" :close-on-content-click="false"
                       location="end" width="400px">
                <template v-slot:activator="{ props }">
                  <v-icon class="ml-3" v-bind="props" size="20" @click="details = true" icon="mdi-information-outline"/>
                </template>

                <v-card>
                  <v-card-text style="white-space: pre-wrap;">
                    <div class="d-flex justify-end">
                         <v-icon icon="mdi-close" @click="details = false"/>
                    </div>
                    <div>{{ getDescription }}</div></v-card-text>
                </v-card>
              </v-menu>
              </div>
            </div>
            <v-icon icon="mdi-close" @click="menu = false"/>
          </span>

          <br/>
          <component :is="getTaskMenuComponent" :task="task.task" :status="task.status"
          />
          <br/>
          <div v-if="getRecords.length > 0">History:</div>

          <div v-for="record in getRecords" :key="record.id">
            <v-divider/>
            {{ `${dayjs(record.startedAt).format('YYYY-MM-DD HH:mm:ss')}` }}
            <StatusCard :status="record.startStatus" :size="'compact'"/>
            ->
            <StatusCard :msg="record.msg" :size="'compact'"
                        :status="record.finishStatus ? record.finishStatus : ProcessStatus.StatusReady"/>
          </div>
          <div v-if="showHistoryDots">
            <v-icon icon="mdi-dots-horizontal" @click="expandHistory = true"/>
          </div>

          <div v-if="getTransactions.length > 0">
            Transactions:
            <div v-for="(tx, i) in getTransactions">
              <div>
                {{ getTransactions.length - i }})
                <a :href="tx.url" target="_blank"><b>{{ tx.code }}</b></a>
                time: {{ formatTime(tx.createdAt) }}
              </div>
            </div>
          </div>

        </v-card-text>
        <v-card-actions class="d-flex">
          <Support :process-id="processId" :task-id="task.id" btn-text="Проблема" :skip-title="true"/>


          <v-btn v-if="errorHappen(task)" @click="retry(task, ppId)" :loading="retryLoading">
            <v-icon icon="mdi-repeat-once"/>
            <v-tooltip activator="parent" location="top">Попробовать снова</v-tooltip>
          </v-btn>


          <v-btn v-if="skipable()" @click="skip(task, ppId)" :loading="skippingTask">
            <v-icon icon="mdi-debug-step-over"></v-icon>
            <v-tooltip activator="parent" location="top">Нажать - если телега не едет</v-tooltip>
          </v-btn>
          <EstimateTask v-if="estimatedTaskMap.get(task.task.taskType)" :task-id="task.id" :profile-id="profileId"
                        :process-id="processId"/>
        </v-card-actions>
      </v-card>
    </template>
  </v-menu>

</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue';
import {processService} from "@/generated/services"
import {ProcessStatus, ProcessTask, ProcessTaskHistoryRecord, TaskType, Transaction} from "@/generated/process";
import ViewFlow from "@/components/flow/OldView.vue";
import StatusCard from "@/components/StatusCard.vue";
import {Delay, formatTime, GetStatusColor, GetStatusText} from "@/components/helper";
import dayjs from "dayjs";
import {estimatedTaskMap, menuTaskComponentMap, taskProps} from '@/components/tasks/tasks'
import EstimateTask from "@/components/tasks/Estimate.vue";
import Support from "@/components/issue/Support.vue";
import {TaskSpec} from "@/components/tasks/utils";
import TaskBlock from "@/components/flow/TaskBlock.vue";

export default defineComponent({
  name: "ProcessTaskMenu",
  components: {TaskBlock, Support, EstimateTask, StatusCard, ViewFlow},
  watch: {
    menu: {
      handler(a, b) {
        if (!a) {
          this.expandHistory = false
        } else {
          this.loadHistory()
          this.loadTransactions()
        }
      }
    },
    task: {
      handler() {
        console.log('task changed')
        this.loadHistory()
        this.loadTransactions()
      },
      deep: true,
    },
  },
  props: {
    task: {
      type: Object as PropType<ProcessTask>,
      required: true,
    },
    ppId: {
      type: String,
      required: true
    },
    processId: {
      type: String,
      required: true
    },
    profileId: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      estimateLoading: false,
      estimateError: '',
      details: false,
      historyLoading: true,
      expandHistory: false,
      menu: false,
      disabled: true,
      retryLoading: false,
      records: [] as ProcessTaskHistoryRecord[],
      skippingTask: false,

      transactionsLoading: false,
      transactions: [] as Transaction[],
    }
  },
  computed: {
    estimatedTaskMap() {
      return estimatedTaskMap
    },
    getTaskMenuComponent(): any {
      return menuTaskComponentMap.get(this.task.task.taskType)
    },
    TaskHasMenu(): boolean {
      return menuTaskComponentMap.has(this.task.task.taskType)
    },
    getTransactions(): Transaction[] {
      return this.transactions
    },
    getRecords(): ProcessTaskHistoryRecord[] {
      if (this.expandHistory) {
        return this.records
      }

      if (this.records.length >= 4) {
        return this.records.slice(0, 4)
      }

      return this.records
    },
    showHistoryDots(): boolean {
      if (this.expandHistory) {
        return false
      }
      if (this.records.length >= 4) {
        return true
      }
      return false
    },
    getDescription(): string {
      return this.task.task.description.replaceAll(`\\`, "")
    },
    TaskType() {
      return TaskType
    },
    ProcessStatus() {
      return ProcessStatus
    }
  },
  methods: {
    formatTime,
    async estimate(taskId: string) {
      try {
        this.estimateError = ''
        this.estimateLoading = true
        const res = await processService.processServiceEstimateCost({
          body: {
            taskId: taskId,
            processId: this.processId,
            profileId: this.profileId
          }
        })
        await Delay(1000)
        if (res.error) {
          this.estimateError = res.error
        }

      } finally {
        this.estimateLoading = false
      }

    },
    GetStatusText,
    dayjs,
    GetStatusColor,
    errorHappen(task: ProcessTask): boolean {
      if (task.status == ProcessStatus.StatusDone) {
        return true
      }
      if (task.status == ProcessStatus.StatusError) {
        return true
      }
      if (task.status == ProcessStatus.StatusRetry) {
        return true
      }
      return false
    },
    taskSpec(item: TaskType): TaskSpec {
      return taskProps[item]
    },
    async loadHistory() {
      if (!this.menu) {
        return
      }
      try {
        this.historyLoading = true
        const res = await processService.processServiceGetProcessTaskHistory({body: {taskId: this.task.id}})
        this.records = res.records
      } finally {
        this.historyLoading = false
      }
    },
    async loadTransactions() {
      if (!this.menu) {
        return
      }
      try {
        this.transactionsLoading = true
        const res = await processService.processServiceGetTaskTransactions({body: {taskId: this.task.id}})
        this.transactions = res.transactions
      } finally {
        this.transactionsLoading = false
      }
    },
    skipable(): boolean {
      if (this.task.status === ProcessStatus.StatusError) {
        return true
      }

      if (this.task.status === ProcessStatus.StatusReady) {
        return true
      }

      if (this.task.status === ProcessStatus.StatusStop) {
        return true
      }

      if (this.task.status === ProcessStatus.StatusRunning) {
        return true
      }

      return false
    },
    async skip(task: ProcessTask, profileId: string) {
      try {
        this.skippingTask = true
        await processService.processServiceSkipProcessTask({
          body: {
            processId: this.processId,
            profileId: profileId,
            taskId: task.id,
          }
        })
      } finally {
        await Delay(100)
        this.skippingTask = false
      }

    },
    async retry(task: ProcessTask, profileId: string) {
      try {
        this.retryLoading = true
        await processService.processServiceRetryProcess({
          body: {
            processId: this.processId,
            profileId: profileId,
            taskId: task.id,
          }
        })
      } finally {
        await Delay(100)
        this.retryLoading = false
      }

    },
  },
})


</script>


<style>


</style>

