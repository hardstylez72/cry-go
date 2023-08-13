<template>
  <div>
    <div v-if="item && item.txCompleted" class="pl-2">
      <b>{{ name }}</b>
      <div v-if="item.gasLimit">
        <i>Max:</i>
        {{ Number(item.gasLimit?.usd).toFixed(2) }} USD
      </div>
      <div v-if="item.multiplier">
        <i>Multiplier:</i>
        {{ Number(item.multiplier).toFixed(2) }}
      </div>
      <div v-if="item.gasEstimated">
        <i>Estimated:</i>
        {{ Number(item.gasEstimated?.usd).toFixed(2) }} USD
      </div>
      <div v-if="item.gasResult">
        <i>Paid:</i>
        {{ Number(item?.gasResult?.usd).toFixed(2) }} USD
      </div>
      <div v-for="d in item.details">
        <i>{{ d.key }}</i> {{ Number(d.value.split(" ")[0]).toPrecision(2) + " " + d.value.split(" ")[1] }}
      </div>
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
  name: "TaskDetails",
  components: {TaskSettingsC, NetworkMaxGas, NetworkGasMultiplier},
  props: {
    name: {
      type: String,
      required: false,
      default: 'operation'
    },
    network: {
      type: String as PropType<Network>,
      required: true,
    },
    item: {
      default: {} as TaskTx,
      type: Object as PropType<TaskTx>,
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
