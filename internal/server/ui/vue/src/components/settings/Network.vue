<template>

  <v-expansion-panel @group:selected="loadSettings">
    <v-expansion-panel-title>
      <div class="d-flex align-center ">
        <v-img height="22px" width="22px" :src="network.img" class="mx-4"/>
        <div>{{ network.name }}</div>
      </div>

    </v-expansion-panel-title>
    <v-expansion-panel-text>
      <Loader v-if="loading"/>
      <v-card v-else>
        <v-form ref="form" validate-on="lazy input">
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
                  hide-details
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
                  hide-details
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <NetworkMaxGas v-model="settings.maxGas" :network="network.value"/>
              </v-col>
            </v-row>
            <v-row class="align-center my-2 px-2">
              <v-checkbox
                density="compact"
                hide-details
                label="Авто-пополнение с биржи"
                v-model="settings.autoRefuel.enabled"
              />

              <div v-if="settings.autoRefuel.enabled" class="d-flex flex-wrap">
                <v-text-field
                  :label="`Минимум`"
                  density="compact"
                  v-model="settings.autoRefuel.min"
                  variant="outlined"
                  style="width: 130px; height: 40px;"
                  class="mx-2 mr-3"
                  :rules="[required, moreThanZero, gte]"
                  :suffix="nativeToken"
                />
                <v-text-field
                  :label="`Maксимум`"
                  density="compact"
                  v-model="settings.autoRefuel.max"
                  variant="outlined"
                  style="width: 130px; height: 40px;"
                  class="mx-2 mr-6"
                  :rules="[required, moreThanZero, gte]"
                  :suffix="nativeToken"
                />
                <div style="width: 200px">
                  <WithdrawerSelect v-model="settings.autoRefuel.withdrawerId"/>
                </div>

              </div>

            </v-row>

            <v-row class="my-1">
              <v-expansion-panels :multiple="false" variant="popout">
                <v-expansion-panel rounded title="Продвинутые">
                  <v-expansion-panel-text>
                    <h3 class="my-3 text-center bg-red-accent-1">При изменении продвинутых настроек возможно появление
                      ошибок. Действуйте осознанно!</h3>
                    <NetworkGasMultiplier v-model="settings.gasMultiplier"/>
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
                  </v-expansion-panel-text>
                </v-expansion-panel>
              </v-expansion-panels>
            </v-row>
          </v-card-text>


        </v-form>
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
import WithdrawerSelect from "@/components/exchange.acc/WithdrawerSelect.vue";
import {moreThanZero, required} from "@/components/tasks/helper";

export default defineComponent({
  name: "SettingsNetwork",
  components: {WithdrawerSelect, TaskSettingsC, Loader, NetworkMaxGas, NetworkGasMultiplier},
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
        console.log('this.settingsChanged ', this.settingsChanged)
      },
      deep: true,
    }
  },
  data() {
    return {
      timer: new Timer(),
      loading: false,
      suggestedGasPrice: 0,
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
    nativeToken(): string {
      const n = this.network.value
      switch (n) {
        case Network.ARBITRUM:
          return 'ETH'
        case Network.Linea:
          return 'ETH'
        case Network.Base:
          return 'ETH'
        case Network.ZKSYNCERA:
          return 'ETH'
        case Network.Etherium:
          return 'ETH'
        case Network.OPTIMISM:
          return "ETH"
        case Network.StarkNet:
          return "ETH"
        case Network.BinanaceBNB:
          return "BNB"
        case Network.AVALANCHE:
          return 'AVAX'
        case Network.POLIGON:
          return "MATIC"
      }
    },
  },
  methods: {
    gte(a: any) {

      if (Number.isNaN(this.settings.autoRefuel.min) || Number.isNaN(this.settings.autoRefuel.max)) {
        return "Должно быть числом"
      }
      if (Number(this.settings.autoRefuel.min) > Number(this.settings.autoRefuel.max)) {
        return "Минимус > Максимус"
      }

      return true
    },
    moreThanZero,
    required,
    slippageAvailable(t: TaskType) {
      return taskProps[t].networks.has(this.network.value)
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

      let j = JSON.stringify(this.settings)
      this.orig = JSON.parse(j)
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
        await this.loadSettings()
      } finally {
        this.reseting = false
      }
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
    async Update() {
      try {
        if (!await this.validateForm()) {
          return
        }

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

        if (!res.settings.autoRefuel) {
          res.settings.autoRefuel = {}
        }
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

