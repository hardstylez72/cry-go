<template>
  <v-row class="mx-5">
    <v-col cols="5" v-if="!isMobile"
           style="float: left; font-size: 18px; width: 30%; overflow: auto; box-sizing: border-box;">
      <v-card class="px-2 py-2">
        <v-card-title class="font-weight-bold text-blue-darken-2 text-center">
          Сценарий
        </v-card-title>

        <v-card-subtitle class="font-weight-bold text-blue-darken-2 text-center">
          Случайные задачи
        </v-card-subtitle>
        <div v-if="randomTasks.length === 0" class="text-center">-</div>

        <div v-else class="handle my-2 elevation-1 py-1 px-1" style="cursor: pointer; height: 50px;"
             v-for="element in randomTasks">
          <a style="text-decoration: none; color: black" :href="`#${element.weight}`">
            <div class="d-inline-flex">
              {{ element.weight }})
              <a target="_blank" :href="taskSpec(element).service.link" class="mx-1">
                <v-img style="background-color: lightgray" height="22px" :src="taskSpec(element).service.img"/>
              </a>
              <b>{{ element.taskType }}</b>
            </div>
            <span>{{ taskDescription(element) }}</span>
            <v-icon v-if="!disable" icon="mdi-close" @click="randomTaskDeleted(element.weight)"></v-icon>
          </a>
        </div>

        <v-card-subtitle class="font-weight-bold text-blue-darken-2 text-center">
          Основные задачи
        </v-card-subtitle>
        <div v-if="tasks.length === 0" class="text-center">-</div>
        <draggable
          v-else
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
            <div class="handle my-2 elevation-1 py-1 px-1" style="cursor: pointer; height: 50px;">
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

      </v-card>
    </v-col>

    <v-col style="float: right; font-size: 18px; width: 30%; overflow-y: scroll">
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

      <div class="my-2">
        <TaskSelectorMenu label="Добавить случайную задачу" :disable="disable" @addStep="addRandomStep"/>
      </div>
      <div class="my-2">
        <TaskSelectorMenu label="Добавить основную задачу" :disable="disable" @addStep="addStep"/>
      </div>


      <v-card v-if="randomTasks.length !== 0">
        <v-card-title class="font-weight-bold text-blue-darken-2 text-center">
          Случайные задачи
        </v-card-title>
        <v-card-text>


          <v-row class="d-flex justify-space-between" v-for="element in randomTasks">
            <v-col>
              <v-card>
                <div class="d-inline-flex my-1 mx-1">
                  Доступна для типов профилей: [
                  <div v-for="t of taskSpec(element).profileType.values()"><b>{{ t + " " }}</b></div>
                  ]
                </div>

                <v-card-title class="d-flex justify-space-between px-0 align-center">
                  <div class="d-inline-flex align-center">
                    <h4 v-if="!disable">
                      <v-icon icon="mdi-dots-vertical" class="handle" @click="() => {}"/>
                    </h4>
                    <a><h4 class="mx-2">{{ `${element.weight}) ${element.taskType}` }}</h4></a>
                    <a target="_blank" :href="taskSpec(element).service.link" class="mx-1">
                      <v-img height="22px" :src="taskSpec(element).service.img"/>
                    </a>
                  </div>

                  <v-icon v-if="!disable" icon="mdi-close" @click="randomTaskDeleted(element.weight)"></v-icon>
                </v-card-title>
                <component
                  @taskChanged="randomTaskChanged"
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

        </v-card-text>
      </v-card>


      <v-card v-if="tasks.length !== 0">
        <v-card-title class="font-weight-bold text-blue-darken-2 text-center">
          Основные задачи
        </v-card-title>
        <v-card-text>
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
                      <div v-for="t of taskSpec(element).profileType.values()"><b>{{ t + " " }}</b></div>
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
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>

</template>

<script lang="ts">

import {defineComponent, PropType, ref} from 'vue';
import {Network, Task, TaskType} from "@/generated/flow";
import {taskComponentMap, TaskArg, taskTypes, taskProps} from '@/components/tasks/tasks'
import draggable from 'vuedraggable'
import TaskChip from "@/components/tasks/TaskChip.vue";
import {networkProps} from "@/components/helper";
import TaskSelector from "@/components/flow/TaskSelector.vue";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskSelectorMenu from "@/components/flow/TaskSelectorBtnMenu.vue";


export default defineComponent({
  name: "OldForm",
  computed: {
    isMobile() {
      return window.innerWidth < 1280
    },
    taskProps() {
      return taskProps
    },
    taskTypes() {
      return taskTypes
    },
  },
  components: {
    TaskSelectorMenu,
    TaskSelector,
    TaskChip,
    draggable,
  },
  watch: {
    label: {
      handler() {
        this.sync()
      }
    },
    tasks: {
      handler() {
        this.sync()
      }
    },
    randomTasks: {
      handler() {
        this.sync()
      }
    },
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
    },
    randomTasksValue: {
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
      randomTasks: [] as TaskArg[],
    }
  },
  methods: {
    sync() {
      console.log('taskChanged')
      this.$emit('flowChanged', this.label, this.tasks, this.randomTasks)
    },
    addRandomStep(task: TaskType) {
      this.randomTasks.push({
        component: taskComponentMap.get(task),
        weight: this.randomTasks.length + 1,
        taskType: task
      });
      this.resolveRandomOrder()
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
    randomTaskDeleted(i: number) {
      this.randomTasks = this.randomTasks.filter(t => t.weight !== i)
      this.resolveRandomOrder()
      this.sync()
    },
    taskDeleted(i: number) {
      this.tasks = this.tasks.filter(t => t.weight !== i)
      this.resolveOrder()
      this.sync()
    },
    orderChanged() {
      this.drag = false
      this.resolveOrder()
    },
    resolveRandomOrder() {
      this.randomTasks = this.randomTasks.map((t, i) => {
        t.weight = ++i
        return t
      })
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
      this.sync()
    },
    randomTaskChanged(task: Task) {
      console.log('randomTaskChanged')
      this.updateRandomTask(task)
      this.sync()
    },
    updateTask(task: Task) {
      for (let i = 0; i < this.tasks.length; i++) {
        if (this.tasks[i].weight.toString() === task.weight) {
          this.tasks[i].task = task
        }
      }
    },
    updateRandomTask(task: Task) {
      for (let i = 0; i < this.randomTasks.length; i++) {
        if (this.randomTasks[i].weight.toString() === task.weight) {
          this.randomTasks[i].task = task
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

    if (this.randomTasksValue) {
      this.randomTasksValue.forEach((v) => {
        this.randomTasks.push({
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

</style>

