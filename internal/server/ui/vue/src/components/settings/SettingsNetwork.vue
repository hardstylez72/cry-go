<template>
  <v-expansion-panel @group:selected="loadSettings">
    <v-expansion-panel-title>
      <div class="d-flex align-center ">
        <v-img height="22px" :src="network.img" class="mx-4"/>
        <div>{{ network.name }}</div>
      </div>

    </v-expansion-panel-title>
    <v-expansion-panel-text>
      <Loader v-if="loading"/>
      <v-card v-else>
        <div class="d-flex justify-end">
          <v-btn v-if="settingsChanged" variant="flat" @click=Update :loading="updating" class="ma-3">Обновить</v-btn>
          <v-btn @click=Reset variant="flat" :loading="reseting" class="ma-3">Сброс</v-btn>
        </div>

        <v-card-text>
          <v-row>
            <v-col>
              <v-text-field
                ref="network-input"
                v-model="settings.rpcEndpoint"
                label="rpc endpoint"
                density="compact"
                variant="outlined"
                :disabled="true"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <v-text-field
                v-model="settings.id"
                label="network id"
                density="compact"
                variant="outlined"
                :loading="loading"
                :disabled="true"
              ></v-text-field>
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <NetworkGasMultiplier v-model="settings.gasMultiplier"/>
            </v-col>
          </v-row>
          <v-row>
            <v-col>
              <NetworkMaxGas v-model="settings.maxGas" :network="network.value"/>
            </v-col>
          </v-row>
        </v-card-text>

        <v-list>
          <div v-for="(task) in taskMap">
            <v-list-item
              v-if="slippageAvailable(task[0])"
              density="compact"
              variant="plain"
              class="my-1 my-4"
              rounded
              height="auto"
              elevation="1"
            >
              <TaskSettingsC :network="network.value" :task-type="task[0]"
                             :settings-prop="task[1]"
                             @updated="taskUpdated"/>
            </v-list-item>
          </div>
        </v-list>

      </v-card>
    </v-expansion-panel-text>
  </v-expansion-panel>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {settingsService} from "@/generated/services"
import {castTaskSettingsMap, castTaskSettingsMapObj, network, TaskSettings, Timer} from "@/components/helper";
import {NetworkSettings} from "@/generated/settings";
import {Network, TaskType} from "@/generated/flow";
import NetworkGasMultiplier from "@/components/settings/NetworkGasMultiplier.vue";
import deepEqual from "deep-equal";
import NetworkMaxGas from "@/components/settings/NetworkMaxGas.vue";
import Loader from "@/components/Loader.vue";
import TaskSettingsC from "@/components/settings/TaskSettings.vue";
import {taskProps} from "@/components/tasks/tasks";

export default defineComponent({
  name: "SettingsNetwork",
  components: {TaskSettingsC, Loader, NetworkMaxGas, NetworkGasMultiplier},
  props: {
    network: {
      type: Object as PropType<network>,
      required: true,
    }
  },
  watch: {
    settings: {
      handler() {
        this.settingsChanged = !deepEqual(this.settings, this.orig)
      },
      deep: true,
    }
  },
  data() {
    return {
      timer: new Timer(),
      loading: false,
      suggestedGasPrice: 0,
      ETHPrice: 2000,
      settings: {} as NetworkSettings,
      orig: {} as NetworkSettings,
      settingsChanged: false,
      updating: false,
      reseting: false,
      taskMap: new Map<TaskType, TaskSettings>(),
    }
  },
  computed: {
    TaskType() {
      return TaskType
    },
    Network() {
      return Network
    },
  },
  methods: {
    slippageAvailable(t: TaskType) {
      return taskProps[t].networks.some(n => n === this.network.value)
    },
    taskUpdated(taskType: TaskType, p: TaskSettings) {
      this.taskMap.set(taskType, p)
      this.settings.taskSettings = this.taskMap
    },
    async endpointChanged(init: boolean) {
      if (this.network.value == Network.ZKSYNCLITE) {
        return
      }
      if (!init) {
        this.timer.add(1000)
        this.timer.cb(async () => {
          this.loading = true
          if (!await this.validate()) {
            this.loading = false
            return
          }

          this.loading = false
        })
      }
    },

    sync() {
      this.orig = Object.assign({}, this.settings)
      this.taskMap = castTaskSettingsMap(this.settings.taskSettings)
    },
    async Reset() {
      try {
        this.reseting = true
        const res = await settingsService.settingsServiceResetSettings({
          body: {
            network: this.network.value
          }
        })
        this.settings = res.settings
        this.sync()
      } finally {
        this.reseting = false
      }
    },
    async Update() {
      try {
        this.updating = true
        this.settings.taskSettings = castTaskSettingsMapObj(this.taskMap)
        const res = await settingsService.settingsServiceUpdateSettings({
          body: {
            settings: this.settings
          }
        })
        this.settings = res.settings
        this.sync()
      } finally {
        this.updating = false
      }
    },
    async loadSettings() {
      try {
        this.loading = true
        const res = await settingsService.settingsServiceGetSettings({body: {network: this.network.value}})
        this.settings = res.settings
        this.sync()
      } finally {
        this.loading = false
      }

    },
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['network-input'].validate()

      return valid
    },
  },
  async created() {

  }
})


</script>


<style>


</style>

