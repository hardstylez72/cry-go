<template>
  <div>
    <div>
      <v-dialog
        v-model="menu"
        height="auto"
        :close-on-content-click="false"
        :close-on-back="false"
        persistent="true"
        location="center"
        location-strategy="static"
        min-width="320px"
      >
        <template v-slot:activator="{ props }">
          <v-btn class="ml-3" variant="flat" height="18px" width="120px" @click="openSettings">Настройки</v-btn>
        </template>

        <v-card class="px-4 py-4">
          <v-card-title class="d-flex justify-space-between">
            <div>Настройки</div>
            <v-icon icon="mdi-close" @click="menu = false"/>
          </v-card-title>
          <v-card-text>
            <v-row>
              <v-col>
                <NetworkGasMultiplier v-model="settings.gasMultiplier"/>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <NetworkMaxGas v-model="settings.maxGas" :network="network"/>
              </v-col>
            </v-row>
            <v-row>
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
                    <TaskSettingsC :network="network" :task-type="task[0]"
                                   :settings-prop="task[1]"
                                   @updated="taskUpdated"/>
                  </v-list-item>
                </div>
              </v-list>
            </v-row>
          </v-card-text>
          <v-card-actions class="d-flex justify-end">
            <v-btn v-if="changed" @click="updateSettings">Update</v-btn>
          </v-card-actions>

        </v-card>
      </v-dialog>
    </div>
  </div>

</template>

<script lang="ts">
import {Amount, AmUni, Network, TaskTx, TaskType, Token} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import NetworkGasMultiplier from "@/components/settings/NetworkGasMultiplier.vue";
import NetworkMaxGas from "@/components/settings/NetworkMaxGas.vue";
import {settingsService} from "@/generated/services";
import {NetworkSettings} from "@/generated/settings";
import deepEqual from "deep-equal";
import {mapStores} from "pinia";
import {useSysStore} from "@/plugins/pinia";
import {castTaskSettingsMap, castTaskSettingsMapObj, TaskSettings} from "@/components/helper";
import {taskProps} from "@/components/tasks/tasks";
import TaskSettingsC from "@/components/settings/TaskSettings.vue";

export default defineComponent({
  name: "MenuTaskSettings",
  components: {TaskSettingsC, NetworkMaxGas, NetworkGasMultiplier},
  props: {
    network: {
      type: String as PropType<Network>,
      required: true,
    },

  },
  watch: {
    settings: {
      handler() {
        if (!deepEqual(this.settings, this.orig)) {
          this.changed = true
        }
      },
      deep: true
    }
  },
  data() {
    return {
      menu: false,
      settings: {} as NetworkSettings,
      orig: {} as NetworkSettings,
      changed: false,
      taskMap: new Map<TaskType, TaskSettings>(),
    }
  },
  computed: {
    ...mapStores(useSysStore),
  },
  methods: {
    slippageAvailable(t: TaskType) {
      if (!t) {
        return false
      }
      return taskProps[t].networks.has(this.network)
    },
    taskUpdated(taskType: TaskType, p: TaskSettings) {
      this.taskMap.set(taskType, p)
      this.settings.taskSettings = this.taskMap
    },
    async updateSettings() {
      try {
        this.settings.taskSettings = castTaskSettingsMapObj(this.taskMap)
        await settingsService.settingsServiceUpdateSettings({body: {settings: this.settings}})
        this.sysStore.showSnackBar("settings updated", false)
        this.menu = false
        this.changed = false
      } catch (e) {
        this.sysStore.showSnackBar("settings update error", true)
      }

    },
    async openSettings() {
      const res = await settingsService.settingsServiceGetSettings({body: {network: this.network}})
      this.settings = res.settings
      this.orig = Object.assign({}, this.settings)
      this.taskMap = castTaskSettingsMap(this.settings.taskSettings)
      this.menu = true
      this.changed = false
    },
  },
})
</script>

<style scoped>

</style>
