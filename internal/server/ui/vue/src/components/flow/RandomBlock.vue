<template>
  <v-card elevation="5" density="compact" variant="elevated">
    <v-card-title>
      {{ block.weight }}) Рандомизированный блок
    </v-card-title>
    <v-card-text>
      <v-form validate-on="lazy" ref="form">

        <div class="flex-wrap d-flex align-center">
          <div>
            <NetworkSelector style="width: 200px" label="Сеть"
                             :items="[
                             Network.Base,
                              Network.ZKSYNCERA,
                               Network.StarkNet,
                               Network.ARBITRUM]"
                             v-model="network"
                             :disabled="disable"
                             class="mx-2"/>
          </div>

          <div>
            <v-text-field label="Кол-во" type="number" v-model="taskCount" variant="outlined" density="compact"
                          style="width: 120px" :disabled="disable" :rules="[required, taskCountRule]"/>
          </div>
          <div style="width: 500px">
            <v-range-slider
              v-model="randomRange"
              density="compact"
              :min="1"
              step="10"
              :max="600"
              persistent-hint
              label=""
              center-affix
              hide-details
              :disabled="disable"
            >
              <template v-slot:prepend>
                <span style="width: 60px">{{ humanDuration(randomRange[0]) }}</span>
              </template>
              <template v-slot:append>
                <span style="width: 60px">{{ humanDuration(randomRange[1]) }}</span>
              </template>
            </v-range-slider>
          </div>

        </div>
        <div v-if="tasks.length">
          <TokenRadioGroup :disabled="disable" label="Монета на вход" :items="fromTokens" v-model="startToken"
                           class="mx-3"/>
        </div>

        <div v-if="tasks.length">
          <TokenRadioGroup :disabled="disable" label="Монета на выход" :items="toTokens" v-model="finishToken"
                           class="mx-3"/>
        </div>


        <div>
          <h4>Основные</h4>
          <v-chip-group :rules="[required]">
            <TaskSelectorChip v-if="network" :base-network-filter="network" label="Добавить"
                              :disable="disable"
                              @addStep="addStep"
                              :exclude-jobs="[TaskJob.LP, TaskJob.Bridge, TaskJob.Exchange, TaskJob.Other]"
            />
            <v-chip v-for="(task, index) in getTasks" rounded variant="outlined"
                    class="mx-1" :disabled="disable">
              <a target="_blank" :href="taskSpec(task.taskType).service.link" class="mx-1">
                <v-img height="22px" :src="taskSpec(task.taskType).service.img"/>
              </a>
              {{ task.taskType }}

              <template v-slot:append>
                <v-icon icon="mdi-close" @click="taskDeleted(index)"/>
              </template>


            </v-chip>
          </v-chip-group>

        </div>

        <div>
          <h4>Случайные</h4>

          <v-chip-group>
            <TaskSelectorChip v-if="network" :base-network-filter="network" label="Добавить"
                              @addStep="addRandomStep" :disable="disable"
                              :exclude-jobs="[TaskJob.Swap, TaskJob.LP, TaskJob.Bridge, TaskJob.Exchange, TaskJob.Other]"/>

            <v-chip v-for="(task, index) in randomTasks" rounded closable="true" @click:close="randomTaskDeleted(index)"
                    variant="outlined" class="mx-1" :disabled="disable">
              <a target="_blank" :href="taskSpec(task.taskType).service.link" class="mx-1">
                <v-img height="22px" :src="taskSpec(task.taskType).service.img"/>
              </a>
              {{ task.taskType }}
            </v-chip>
          </v-chip-group>
        </div>

        <div v-if="previewError">
          Ошибка формирования комбинаций
        </div>
        <RandomFlowPreview v-else :data="preview" v-if="preview" :combo="preview.uniquePercent"/>
      </v-form>
    </v-card-text>
  </v-card>

</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {
  FlowBlock,
  Network,
  OnlyRandomFlowPreviewReq,
  RandomFlowPreviewRes,
  RandomTask,
  TaskType,
  Token
} from "@/generated/flow";
import {taskProps, taskTypes} from '@/components/tasks/tasks'
import TaskChip from "@/components/tasks/TaskChip.vue";
import TaskSelector from "@/components/flow/TaskSelector.vue";
import {TaskJob, TaskSpec} from "@/components/tasks/utils";
import NetworkSelector from "@/components/tasks/NetworkSelector.vue";
import TokenSelector from "@/components/tasks/TokenSelector.vue";
import TaskSelectorChip from "@/components/flow/TaskSelectorChipMenu.vue";
import RandomFlowPreview from "@/components/flow/Preview.vue";
import {flowService} from "@/generated/services";
import {humanDuration, Timer} from "@/components/helper";
import TokenRadioGroup from "@/components/tasks/TokenRadioGroup.vue";
import {required} from "@/components/tasks/helper";


export default defineComponent({
  name: "CreateRandomFlowBlock",
  computed: {
    TaskJob() {
      return TaskJob
    },
    Network() {
      return Network
    },
    Token() {
      return Token
    },
    taskProps() {
      return taskProps
    },
    taskTypes() {
      return taskTypes
    },
    getTasks() {
      return this.tasks
    },
    fromTokens(): Token[] {
      return this.fromTokens.sort()
    },
    async toTokens() {
      return this.toTokens.sort()
    }
  },
  components: {
    TokenRadioGroup,
    RandomFlowPreview,
    TaskSelectorChip,
    TokenSelector,
    NetworkSelector,
    TaskSelector,
    TaskChip,
  },
  watch: {
    taskCount: {
      handler() {
        this.localBlock.rand.taskCount = this.taskCount
      }
    },
    randomRange: {
      handler() {
        this.localBlock.rand.minDelay = this.randomRange[0]
        this.localBlock.rand.maxDelay = this.randomRange[1]
      },
      deep: true,
    },
    tasks: {
      async handler() {
        const tasks = []
        tasks.push(...this.tasks)
        tasks.push(...this.randomTasks)
        this.localBlock.rand.tasks = tasks
        await this.loadTokenFrom()
      },
      deep: true,
    },
    randomTasks: {
      async handler() {
        const tasks = []
        tasks.push(...this.tasks)
        tasks.push(...this.randomTasks)
        this.localBlock.rand.tasks = tasks
        await this.loadTokenFrom()
      },
      deep: true,
    },
    startToken: {
      handler() {
        this.loadToTokens()
        this.localBlock.rand.startToken = this.startToken
      },
    },
    finishToken: {
      handler() {
        this.loadFlowPreview()
        this.localBlock.rand.finishToken = this.finishToken
      }
    },
    localBlock: {
      async handler() {
        await this.loadFlowPreview()
        await this.sync()
      },
      deep: true,
    },
    network: {
      handler() {
        this.localBlock.rand.startNetwork = this.network
      }
    }
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
  emits: ['blockChanged'],
  data() {
    return {
      randomRange: [180, 300] as Array<Number>,
      network: Network.ZKSYNCERA as null | Network,
      startToken: null as Token | null,
      finishToken: null as Token | null,
      tasks: [] as RandomTask[],
      randomTasks: [] as RandomTask[],
      preview: null as null | RandomFlowPreviewRes,
      previewError: '',
      taskCount: 1,
      timerTokenCombo: new Timer(),
      timerFlowPreview: new Timer(),
      timerSync: new Timer(),
      localBlock: {
        rand: {
          tasks: []
        }
      } as FlowBlock,
      fromTokens: [] as Token[],
      toTokens: [] as Token[],
    }
  },
  methods: {
    required,
    humanDuration,

    async sync() {

      if (!await this.validateForm()) {
        return
      }

      this.$emit('blockChanged', {
        startNetwork: this.network,
        maxDelay: this.randomRange[1],
        minDelay: this.randomRange[0],
        tasks: [...this.randomTasks, ...this.tasks],
        taskCount: this.taskCount,
        startToken: this.startToken,
        finishToken: this.finishToken,
      })
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
    async loadTokenFrom() {
      this.previewError = ''
      const tasks = []
      tasks.push(...this.tasks)
      tasks.push(...this.randomTasks)

      const b: OnlyRandomFlowPreviewReq = {
        startToken: this.startToken,
        finishToken: this.finishToken,
        ignoreFinishToken: true,
        ignoreStartToken: true,
        startNetwork: this.network,
        taskCount: this.taskCount,
        tasks: tasks,
        minDelay: "0",
        maxDelay: "0",
      }

      await flowService.flowServiceOnlyRandomFlowFromTokens({
        body: b
      }).then((data) => {
        this.fromTokens = data.tokens
      }).catch((err) => {
        this.previewError = 'Ошибка'
      })

    },
    async loadFlowPreview() {

      this.timerFlowPreview.add(500)
      this.timerFlowPreview.cb(async () => {
        this.previewError = ''

        const b: OnlyRandomFlowPreviewReq = {
          startToken: this.startToken,
          finishToken: this.finishToken,
          ignoreFinishToken: false,
          ignoreStartToken: false,
          startNetwork: this.network,
          taskCount: this.taskCount,
          tasks: [...this.tasks, ...this.randomTasks],
          minDelay: "0",
          maxDelay: "0",
        }

        await flowService.flowServiceOnlyRandomFlowPreview({
          body: b
        }).then((data) => {
          this.preview = data
        }).catch((err) => {
          this.previewError = 'Ошибка'
        })

      })

    },
    taskCountRule(v: any) {
      if (this.tasks.length >= this.taskCount) {
        return true
      }

      return `> ${this.tasks.length} !`
    },
    async loadToTokens() {
      this.previewError = ''

      try {
        const res = await flowService.flowServiceOnlyRandomFlowPreview({
          body: {
            startToken: this.startToken,
            finishToken: this.finishToken,
            ignoreFinishToken: true,
            ignoreStartToken: false,
            startNetwork: this.network,
            taskCount: this.taskCount,
            tasks: [...this.tasks, ...this.randomTasks],
            minDelay: "0",
            maxDelay: "0",
          }
        })

        this.toTokens = res.tokens[0].to

      } catch (e) {
        this.previewError = 'Ошибка'
      }

    },
    addRandomStep(task: TaskType) {

      switch (taskProps[task].job) {
        case TaskJob.Swap:
          this.randomTasks.push({
            taskType: task,
            optional: true,
            swap: {
              items: taskProps[task].swapParis,
            }
          });
          break
        case TaskJob.NFT, TaskJob.Simple:
          this.randomTasks.push({
            taskType: task,
            optional: true,
            simple: {
              network: this.network
            }
          });
          break
      }

      this.resolveRandomOrder()
    },
    taskSpec(item: TaskType): TaskSpec {
      return taskProps[item]
    },
    addStep(task: TaskType) {
      switch (taskProps[task].job) {
        case TaskJob.Swap:
          this.tasks.push({
            taskType: task,
            optional: false,
            swap: {
              items: taskProps[task].swapParis,
            }
          });
          break
        case TaskJob.NFT, TaskJob.Simple:
          this.randomTasks.push({
            taskType: task,
            optional: true,
            simple: {
              network: this.network
            }
          });
          break
      }

      this.resolveOrder()
    },
    taskDeleted(i: number) {
      this.tasks.splice(i, 1)
      this.resolveOrder()
      this.sync()
    },
    randomTaskDeleted(i: number) {
      this.randomTasks.splice(i, 1)
      this.resolveRandomOrder()
      this.sync()
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
    tasksRules() {
      return [
        (v: string) => this.tasks.length ? true : 'add at least one task'
      ]
    },
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['label-input'].validate()

      return valid
    },
  },
  created() {
    if (this.block && this.block.rand) {

      this.localBlock = this.block

      if (this.block.rand.taskCount) {
        this.taskCount = this.block.rand.taskCount
      }

      if (this.block.rand.startNetwork) {
        this.network = this.block.rand.startNetwork
      }
      if (this.block.rand.startToken) {
        this.startToken = this.block.rand.startToken
      }

      if (this.block.rand.finishToken) {
        this.finishToken = this.block.rand.finishToken
      }
      if (this.block.rand.minDelay) {
        this.randomRange = [this.block.rand.minDelay, this.block.rand.maxDelay]
      }


      this.block.rand.tasks.forEach((v) => {
        if (v.optional) {

          switch (taskProps[v.taskType].job) {
            case TaskJob.Swap:
              this.randomTasks.push({
                taskType: v.taskType,
                optional: v.optional,
                swap: v.swap
              });
              break
            case TaskJob.NFT, TaskJob.Simple:
              this.randomTasks.push({
                taskType: v.taskType,
                optional: v.optional,
                simple: v.simple
              });
              break
          }

        } else {
          switch (taskProps[v.taskType].job) {
            case TaskJob.Swap:
              this.tasks.push({
                taskType: v.taskType,
                optional: v.optional,
                swap: v.swap
              });
              break
            case TaskJob.NFT, TaskJob.Simple:
              this.tasks.push({
                taskType: v.taskType,
                optional: v.optional,
                simple: v.simple
              });
              break
          }
        }

      })

      this.sync()
    }
  }
})


</script>


<style>

</style>

