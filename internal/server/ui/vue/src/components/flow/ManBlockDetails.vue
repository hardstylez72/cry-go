<template>

  <v-dialog
    v-model="dialog"
    width="90%"
    height="90%"
    location="center"
    :close-on-content-click="false"
    :close-on-back="false"
    :persistent="true"
  >
    <template v-slot:activator="{ props }">
      <v-chip :disabled="disable" rounded variant="outlined" v-bind="props" @click="dialog=true">Изменить</v-chip>
    </template>


    <v-card width="100%" height="100%">
      <v-card-title class="d-flex justify-space-between">
        <div>{{ block.weight }}) Мужицкий блок</div>
        <div>
          <v-btn @click="save">Сохранить</v-btn>
          <v-icon icon="mdi-close" @click="dialog = false"/>
        </div>

      </v-card-title>
      <v-card-text>

        <v-row class="mx-5">
          <v-col cols="5" v-if="!isMobile"
                 style="float: left; font-size: 18px; width: 30%; overflow: auto; box-sizing: border-box;">
            <v-card class="px-2 py-2">
              <v-card-title class="font-weight-bold text-blue-darken-2 text-center">
                <span>Сценарий</span>
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
                    <a target="_blank" :href="taskSpec(element.taskType).service.link" class="mx-1">
                      <v-img style="background-color: lightgray" height="22px"
                             :src="taskSpec(element.taskType).service.img"/>
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
                        <a target="_blank" :href="taskSpec(element.taskType).service.link" class="mx-1">
                          <v-img style="background-color: lightgray" height="22px"
                                 :src="taskSpec(element.taskType).service.img"/>
                        </a>
                        <b>{{ element.taskType }}</b>
                      </div>
                      <span>{{ taskDescription(element.taskType) }}</span>
                      <v-icon v-if="!disable" icon="mdi-close" @click="taskDeleted(element.weight)"></v-icon>
                    </a>
                  </div>
                </template>
              </draggable>

            </v-card>
          </v-col>


          <v-col style="float: right; font-size: 18px; width: 30%; overflow-y: scroll">
            <v-form ref="form">
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
                          <div v-for="t of taskSpec(element.taskType).profileType.values()"><b>{{ t + " " }}</b></div>
                          ]
                        </div>

                        <v-card-title class="d-flex justify-space-between px-0 align-center">
                          <div class="d-inline-flex align-center">
                            <h4 v-if="!disable">
                              <v-icon icon="mdi-dots-vertical" class="handle" @click="() => {}"/>
                            </h4>
                            <a><h4 class="mx-2">{{ `${element.weight}) ${element.taskType}` }}</h4></a>
                            <a target="_blank" :href="taskSpec(element.taskType).service.link" class="mx-1">
                              <v-img height="22px" width="22px" :src="taskSpec(element.taskType).service.img"/>
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
                          :spec="taskSpec(element.taskType)"
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
                              <div v-for="t of taskSpec(element.taskType).profileType.values()"><b>{{ t + " " }}</b>
                              </div>
                              ]
                            </div>

                            <v-card-title class="d-flex justify-space-between px-0 align-center">
                              <div class="d-inline-flex align-center">
                                <h4 v-if="!disable">
                                  <v-icon icon="mdi-dots-vertical" class="handle" @click="() => {}"/>
                                </h4>
                                <a><h4 :id="element.weight" class="mx-2">{{
                                    `${element.weight}) ${element.taskType}`
                                  }}</h4></a>
                                <a target="_blank" :href="taskSpec(element.taskType).service.link" class="mx-1">
                                  <v-img height="22px" width="22px" :src="taskSpec(element.taskType).service.img"/>
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
                </v-card-text>
              </v-card>
            </v-form>
          </v-col>

        </v-row>

      </v-card-text>
    </v-card>

  </v-dialog>
</template>

<script lang="ts">

import {defineComponent, PropType, ref} from 'vue';
import {FlowBlock, Network, Task, TaskType} from "@/generated/flow";
import {taskComponentMap, TaskArg, taskTypes, taskProps} from '@/components/tasks/tasks'
import draggable from 'vuedraggable'
import TaskChip from "@/components/tasks/TaskChip.vue";
import {networkProps} from "@/components/helper";
import TaskSelector from "@/components/flow/TaskSelector.vue";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskSelectorMenu from "@/components/flow/TaskSelectorBtnMenu.vue";


export default defineComponent({
  name: "ManBlockDetails",
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
    localBlock: {
      handler() {
        // this.sync()
      },
      deep: true
    },
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
  emits: ['flowChanged'],
  data() {
    return {
      dialog: false,
      menu: false,
      selectErrorMessage: "",
      drag: false,
      localBlock: {man: {tasks: [], randomTasks: []}} as FlowBlock,
      randomTasks: [] as TaskArg[],
      tasks: [] as TaskArg[],
    }
  },
  methods: {
    sync() {
      this.localBlock = {
        weight: this.block.weight,
        man: {
          tasks: this.tasks.map(t => t.task),
          randomTasks: this.randomTasks.map(t => t.task),
        }
      }

      this.$emit('flowChanged', this.localBlock)
    },
    addRandomStep(task: TaskType) {
      this.randomTasks.push({
        component: taskComponentMap.get(task),
        weight: this.randomTasks.length + 1,
        taskType: task
      });
      this.resolveRandomOrder()
    },
    taskSpec(item: TaskType): TaskSpec {
      return taskProps[item]
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
    },
    taskDeleted(i: number) {
      this.tasks = this.tasks.filter(t => t.weight !== i)
      this.resolveOrder()
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
      this.updateTask(task)
    },
    randomTaskChanged(task: Task) {
      this.updateRandomTask(task)
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
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
    async save() {
      if (!await this.validate()) {
        return
      }

      this.sync()
      this.dialog = false
    }
  },
  created() {


    if (this.block && this.block.man) {
      this.block.man.tasks.forEach((v) => {
        this.tasks.push({
          component: taskComponentMap.get(v.taskType),
          weight: Number(v.weight),
          task: v,
          taskType: v.taskType,
        })
      })

      this.block.man.randomTasks.forEach((v) => {
        this.randomTasks.push({
          component: taskComponentMap.get(v.taskType),
          weight: Number(v.weight),
          task: v,
          taskType: v.taskType,
        })
      })

    } else {
      this.localBlock = {
        weight: this.block.weight,
        man: {
          randomTasks: [],
          tasks: []
        }
      }
    }

  }
})


</script>


<style>

</style>

