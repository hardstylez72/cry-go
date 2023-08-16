<template>
  <div cols="12" class="my-3"
       style="box-shadow: rgba(0, 0, 0, 0.16) 0 3px 6px, rgba(0, 0, 0, 0.23) 0 3px 6px;">
    <v-col cols="12">
      <div class="d-inline-flex align-center">{{ `Тип фильтра: ` }}
        <v-radio-group inline="true" v-model="filterType" hide-details>
          <v-radio v-for="item in Filters" :value="item" :label="item"/>
        </v-radio-group>
      </div>
      <div>Значения фильтра:
        <v-radio-group v-if="filterType === 'Сеть'" v-model="filterNetwork" inline="true">
          <v-chip density="compact" size="5px" value="Все сети" @click="filterNetwork = 'Все сети'">{{
              "Все сети"
            }}
          </v-chip>
          <v-chip v-for="network in networks" @click="filterNetwork = network" :value="network"
                  density="compact" class="mx-1 my-1">
            <v-img height="22px" :src="networkProps[network].img"/>
            {{ networkProps[network].name }}
          </v-chip>
        </v-radio-group>
        <v-radio-group v-if="filterType === 'Airdrop'" v-model="filterAirdrop" inline="true">
          <v-radio value="Все Airdrop" label="Все Airdrop"/>
          <v-radio v-for="airdrop in airdrops" :value="airdrop" :label="airdrop"></v-radio>
        </v-radio-group>

      </div>
      <div v-for="job in taskJobs" class="d-flex flex-wrap">
        <v-divider class="my-1"/>
        <b>{{ job }}</b>
        <div v-for="tt in filterTasksByJob(job)">
          <TaskChip class="mx-1 my-1" @click="addStep(tt)" :task-type="tt"/>
        </div>
      </div>
    </v-col>
  </div>

</template>

<script lang="ts">

import {defineComponent, PropType, ref} from 'vue';
import {Network, Task, TaskType} from "@/generated/flow";
import {taskComponentMap, TaskArg, taskTypes, taskProps, TaskJob, TaskSpec, Airdrop} from '@/components/tasks/tasks'
import draggable from 'vuedraggable'
import TaskChip from "@/components/tasks/TaskChip.vue";
import {networkProps} from "@/components/helper";


export default defineComponent({
  name: "TaskSelector",
  computed: {
    networkProps() {
      return networkProps
    },
    taskProps() {
      return taskProps
    },
    taskTypes() {
      return taskTypes
    },
    Filters(): string[] {
      return ['Сеть', 'Airdrop']
    },
  },
  components: {
    TaskChip,
    draggable,
  },
  watch: {
    filterNetwork: {
      handler() {
        this.getTaskList()
      },
    },
    filterAirdrop: {
      handler() {
        this.getTaskList()
      },
    },
    filter: {
      handler() {
        this.getTaskList()
      },
    }
  },
  props: {
    disable: {
      type: Boolean as PropType<boolean>,
      required: false,
    },
  },
  emits: ['taskSelected'],
  data() {
    return {
      selectErrorMessage: "",
      selectedTask: TaskType.Delay,

      filterType: 'Сеть' as 'Сеть' | 'Airdrop',


      taskJobs: [] as TaskJob[],
      jobSet: new Set<TaskJob>(),

      filterNetwork: 'Все сети' as 'Все сети' | Network,
      networks: [] as Network[],
      networkSet: new Set<Network>(),

      filterAirdrop: 'Все Airdrop' as 'Все Airdrop' | Airdrop,
      airdrops: [] as Airdrop[],
      airdropSet: new Set<Airdrop>()
    }
  },
  methods: {
    filterTasksByJob(job: TaskJob): TaskType[] {
      const out: TaskType[] = []

      this.taskTypes.forEach((taskType) => {
        if (this.taskProps[taskType].job !== job) {
          return
        }

        switch (true) {
          case this.filterType === 'Сеть' :
            if (this.filterNetwork !== 'Все сети' && !this.taskProps[taskType].networks.has(this.filterNetwork)) {
              return;
            }
            break
          case this.filterType === 'Airdrop' :
            if (this.filterAirdrop !== 'Все Airdrop' && !this.taskProps[taskType].airdrops.has(this.filterAirdrop)) {
              return;
            }
            break
        }

        out.push(taskType)
      })
      return out
    },
    getTaskList() {

      this.networks = []
      this.airdrops = []
      this.taskJobs = []

      this.jobSet = new Set<TaskJob>()
      this.airdropSet = new Set<Airdrop>()
      this.networkSet = new Set<Network>()

      this.taskTypes.forEach((taskType) => {

        this.jobSet.add(this.taskProps[taskType].job)
        this.taskProps[taskType].airdrops.forEach((airdrop) => {
          this.airdropSet.add(airdrop)
        })

        taskProps[taskType].networks.forEach((network) => {
          this.networkSet.add(network)
        })
      })

      for (const airdrop of this.airdropSet.values()) {
        this.airdrops.push(airdrop)
      }

      for (const network of this.networkSet.values()) {
        this.networks.push(network)
      }

      for (const job of this.jobSet.values()) {
        this.taskJobs.push(job)
      }
    },
    taskSpec(item: TaskArg): TaskSpec {
      return taskProps[item.taskType]
    },
    async addStep(task: TaskType) {
      this.$emit('taskSelected', task)
    },
    orderChanged() {
      this.drag = false
      this.resolveOrder()
    },
    tasksRules() {
      return [
        (v: string) => this.tasks.length ? true : 'add at least one task'
      ]
    },
    async labelChanged(e: InputEvent) {
      await this.validate()
    },
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['label-input'].validate()

      return valid
    },
  },
  created() {
    this.getTaskList()
  }
})


</script>


<style>


</style>

