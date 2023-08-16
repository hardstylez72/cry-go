<template>
  <v-row class="mx-5">
    <v-col cols="5" v-if="!isMobile">
      <v-card style="position: fixed; font-size: 18px; width: 30%" class="px-1 py-1">
        Сценарий
        <draggable
          v-model="tasks"
          @start="drag=true"
          @end="orderChanged"
          item-key="weight"
          ghost-class="ghost"
          :animation="500"
          handle=".handle"
          class="list-group"
        >
          <template #item="{element}">
            <div class="handle my-2 elevation-1 py-1 px-1" style="cursor: pointer">
              <a style="text-decoration: none; color: black" :href="`#${element.weight}`">
                <div class="d-inline-flex">
                  {{ element.weight }})
                  <a target="_blank" :href="taskSpec(element).service.link" class="mx-1">
                    <v-img style="background-color: lightgray" height="22px" :src="taskSpec(element).service.img"/>
                  </a>
                  <b>{{ element.taskType }}</b>
                </div>
                <span>{{ taskDescription(element) }}</span>
                <v-icon v-if="!disable" icon="mdi-close" @click="taskDeleted(element.weight)"></v-icon>
              </a>
            </div>
          </template>
        </draggable>
        <v-menu
          location-strategy="connected"
          max-width="700px"
          location="left"
          v-model="menu"
          persistent="true"
          no-click-animation
          transition="false"
          :close-on-content-click="false"
          :close-on-back="false">
          <template v-slot:activator="{ props }">
            <v-btn rounded="false" width="100%" v-bind="props" :disabled="disable">Добавить задачу</v-btn>
          </template>
          <v-card>
            <div class="d-flex justify-end">
              <v-icon icon="mdi-close" @click="menu = false"></v-icon>
            </div>
            <TaskSelector v-if="!disable" @task-selected="addStep"/>
          </v-card>
        </v-menu>


      </v-card>
    </v-col>

    <v-col>
      <v-row>
        <v-col cols="12" sm="6" md="4">
          <v-text-field
            ref="label-input"
            v-model="label"
            label="label"
            density="comfortable"
            variant="outlined"
            :rules="labelRules()"
            @input="labelChanged"
            :disabled="disable"
          ></v-text-field>
        </v-col>
      </v-row>


      <draggable
        v-model="tasks"
        @start="drag=true"
        @end="orderChanged"
        item-key="weight"
        :animation="100"
        handle=".handle"
      >
        <template #item="{element}">
          <v-row class="d-flex justify-space-between">
            <v-col>
              <v-card>
                <div class="d-inline-flex my-1 mx-1">
                  Доступна для типов профилей: [
                  <div v-for="t of taskSpec(element).profileType.values()"><b>{{ t }}</b></div>
                  ]
                </div>

                <v-card-title class="d-flex justify-space-between px-0 align-center">
                  <div class="d-inline-flex align-center">
                    <h4 v-if="!disable">
                      <v-icon icon="mdi-dots-vertical" class="handle" @click="() => {}"/>
                    </h4>
                    <a><h4 :id="element.weight" class="mx-2">{{ `${element.weight}) ${element.taskType}` }}</h4></a>
                    <a target="_blank" :href="taskSpec(element).service.link" class="mx-1">
                      <v-img height="22px" :src="taskSpec(element).service.img"/>
                    </a>
                  </div>

                  <v-icon v-if="!disable" icon="mdi-close" @click="taskDeleted(element.weight)"></v-icon>
                </v-card-title>
                <component
                  @taskChanged="taskChanged"
                  style="box-shadow: rgba(0, 0, 0, 0.16) 0 3px 6px, rgba(0, 0, 0, 0.23) 0 3px 6px;"
                  :is="element.component"
                  :weight="element.weight"
                  :task="element.task"
                  :disabled="disable"
                  :spec="taskSpec(element)"
                />
              </v-card>
            </v-col>
          </v-row>
        </template>
      </draggable>
      <TaskSelector v-if="!disable && isMobile" @task-selected="addStep"/>

    </v-col>
  </v-row>

</template>

<script lang="ts">

import {defineComponent, PropType, ref} from 'vue';
import {Network, Task, TaskType} from "@/generated/flow";
import {taskComponentMap, TaskArg, taskTypes, taskProps, TaskJob, TaskSpec, Airdrop} from '@/components/tasks/tasks'
import draggable from 'vuedraggable'
import TaskChip from "@/components/tasks/TaskChip.vue";
import {networkProps} from "@/components/helper";
import TaskSelector from "@/components/flow/TaskSelector.vue";


export default defineComponent({
  name: "FlowForm",
  computed: {
    isMobile() {
      return window.innerWidth < 1280
    },
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
    TaskSelector,
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
    label: {
      handler() {
        this.$emit('flowChanged', this.label, this.tasks)
      }
    },
    tasks: {
      handler() {
        this.$emit('flowChanged', this.label, this.tasks)
      }
    },
    filter: {
      handler() {
        this.getTaskList()
      },
    }
  },
  props: {
    demo: {
      type: Boolean,
      required: false,
    },
    disable: {
      type: Boolean as PropType<boolean>,
      required: false,
    },
    labelValue: {
      type: String as PropType<string>,
      required: false,
    },
    tasksValue: {
      type: Array as PropType<Task[]>,
      required: false
    }
  },
  emits: ['flowChanged'],
  data() {
    return {
      menu: false,
      selectErrorMessage: "",
      drag: false,
      label: "",
      selectedTask: TaskType.Delay,
      tasks: [] as TaskArg[],

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
    taskDescription(item: TaskArg) {
      let base = ''
      if (!item.task) {
        return base
      }
      base += taskProps[item.taskType].descFn(item.task)
      return base
    },
    async addStep(task: TaskType) {
      this.tasks.push({
        component: taskComponentMap.get(task),
        weight: this.tasks.length + 1,
        taskType: task
      });
      this.resolveOrder()
    },
    taskDeleted(i: number) {
      this.tasks = this.tasks.filter(t => t.weight !== i)
      this.resolveOrder()
      this.$emit('flowChanged', this.label, this.tasks)
    },
    orderChanged() {
      this.drag = false
      this.resolveOrder()
    },
    resolveOrder() {
      this.tasks = this.tasks.map((t, i) => {
        t.weight = ++i
        return t
      })
    },
    taskChanged(task: Task) {
      console.log('taskChanged')
      this.updateTask(task)
      this.$emit('flowChanged', this.label, this.tasks)
    },
    updateTask(task: Task) {
      for (let i = 0; i < this.tasks.length; i++) {
        if (this.tasks[i].weight.toString() === task.weight) {
          this.tasks[i].task = task
        }
      }
    },
    labelRules() {
      return [
        (v: string) => this.label ? true : 'label is required'
      ]
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

    if (this.labelValue) {
      this.label = this.labelValue
    }
    if (this.tasksValue) {
      this.tasksValue.forEach((v) => {
        this.tasks.push({
          component: taskComponentMap.get(v.taskType),
          weight: Number(v.weight),
          task: v,
          taskType: v.taskType,
        })
      })
    }

    if (this.demo) {
      this.label = 'demo'

      this.tasks.push({
        component: taskComponentMap.get(TaskType.Delay),
        weight: 1,
        taskType: TaskType.Delay,
        task: {
          taskType: TaskType.Delay,
          weight: 1,
          delayTask: {
            duration: 2,
          }
        }
      });
      this.tasks.push({
        component: taskComponentMap.get(TaskType.Delay),
        weight: 2,
        taskType: TaskType.Delay,
        task: {
          taskType: TaskType.Delay,
          weight: 2,
          delayTask: {
            duration: 2,
          }
        }
      });
      this.tasks.push({
        component: taskComponentMap.get(TaskType.Delay),
        weight: 3,
        taskType: TaskType.Delay,
        task: {
          taskType: TaskType.Delay,
          weight: 1,
          delayTask: {
            duration: 2,
          }
        }
      });
    }
  }
})


</script>


<style>
.ghost {
  opacity: 0.7;
  background: #c8ebfb;
}

</style>

