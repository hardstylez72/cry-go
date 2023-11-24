<template>
  <div>
    <NavBar title="Сценарии">
      <template v-slot:default>
        <div class="d-flex justify-end">
          <v-btn @click="$router.push({name: 'SharedFlowList'})" class="mx-1" variant="flat">Общие</v-btn>
          <v-btn @click=CreateFlow :class="noFlow ? 'onboarding' : ''" variant="flat">Добавить</v-btn>
          <v-btn @click="$router.push({name: 'CreateFlowV2'})" variant="flat">ДобавитьV2</v-btn>
        </div>
      </template>
    </NavBar>
    <Loader v-if="loading"/>
    <div v-else>
      <div v-if="noFlow" class="my-8 mx-8 text-center">
        Для создания сценария вам потребуются профили, если вы уже завели профили - нажмите добавить, для вас будет
        пре-заполнен демонстрационный
        бесплатный сценарий
      </div>
      <v-list v-else max-width="96vw" class="px-5">
        <v-list-item
          density="compact"
          variant="plain"
          class="my-1 my-4"
          v-for="item in getList"
          :key="item.id"
          rounded
          height="auto"
          elevation="1"
          @click="viewFlow(item.id, item.version)"
        >
          <div>
            <div class="mr-2"><b>{{ `${item.label}` }}</b> {{ formatTime(item.createdAt) }}</div>
          </div>
        </v-list-item>
      </v-list>
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService} from "@/generated/services"
import {flow_Flow as Flow, FlowListItem, TaskType} from "@/generated/flow";
import CreateFlow from "@/components/flow/OldCreateFlow.vue";
import CheckBox from "@/components/CheckBox.vue";
import {getFlow} from "@/components/tasks/tasks";
import Loader from "@/components/Loader.vue";
import NavBar from "@/components/NavBar.vue";
import {formatTime} from "../helper";

export default defineComponent({
  name: "Constructor",
  components: {
    NavBar,
    Loader,
    CheckBox,
    CreateFlow
  },
  props: {},
  data() {
    return {
      list: [] as FlowListItem[],
      loading: true,
      loadingError: false,
    }
  },
  computed: {
    noFlow() {
      return this.getList.length === 0
    },
    getList(): FlowListItem[] {
      return this.list
    },
  },
  methods: {
    formatTime,
    getFlow,
    viewFlow(id: string, version: string) {

      if (version == "2") {
        this.$router.push({name: 'FlowViewV2', params: {id}})
      } else {
        this.$router.push({name: 'Flow', params: {id}})
      }
    },
    showStep(flow: Flow): string {
      if (!flow || !flow.tasks) {
        return "-"
      }
      let res = ""
      flow.tasks.forEach((task) => {
        res += ` ${task.weight}) ${task.taskType}`
      })
      return res
    },
    CheckboxChanged(id: string, value: boolean) {
      if (value) {
        this.selected.add(id)
      } else {
        this.selected.delete(id)
      }
    },
    async UpdateList() {
      try {
        this.loadingError = false
        this.loading = true
        const res = await flowService.flowServiceListFlow()
        this.list = res.flows
        this.selected = new Set<string>()
      } catch (e) {
        this.loadingError = true
      } finally {
        this.loading = false
      }

    },
    CreateFlow() {
      if (this.noFlow) {
        this.$router.push({name: 'CreateFlow', query: {demo: true}})
      } else {
        this.$router.push({name: 'CreateFlow'})
      }
    },
  },
  created() {
    this.UpdateList()
  }
})


</script>


<style>


</style>

