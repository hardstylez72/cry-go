<template>
  <div>
    <NavBar title="Общие сценарии"/>
    <Loader v-if="loading"/>
    <div v-else>

      <v-list max-width="96vw" class="px-5">
        <v-list-item
          density="compact"
          variant="plain"
          class="my-1 my-4"
          v-for="item in getList"
          :key="item.id"
          rounded
          height="auto"
          elevation="1"
          @click="viewFlow(item.id)"
        >
          <v-list-item-title>
            <div class="mr-2"><b>{{ `${item.label}` }}</b></div>
          </v-list-item-title>
          <v-list-item-subtitle>
            {{ item.description }}
          </v-list-item-subtitle>
          <div>
            <div class="mr-2" v-for="(d, i) in getFlow(item)">
              <b>{{ i + 1 }})</b> {{ d }}
            </div>
          </div>
        </v-list-item>
      </v-list>
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService} from "@/generated/services"
import {flow_Flow as Flow, FlowShared, TaskType} from "@/generated/flow";
import CreateFlow from "@/components/flow/OldCreateFlow.vue";
import CheckBox from "@/components/CheckBox.vue";
import {getFlow} from "@/components/tasks/tasks";
import Loader from "@/components/Loader.vue";
import NavBar from "@/components/NavBar.vue";

export default defineComponent({
  name: "SharedFlowList",
  components: {
    NavBar,
    Loader,
    CheckBox,
    CreateFlow
  },
  props: {},
  data() {
    return {
      list: [] as FlowShared[],
      loading: true,
      loadingError: false,
    }
  },
  computed: {
    noFlow() {
      return this.getList.length === 0
    },
    getList(): FlowShared[] {
      return this.list
    },
  },
  methods: {
    getFlow,
    viewFlow(id: string) {
      this.$router.push({name: 'SharedFlow', params: {id}})
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
        const res = await flowService.flowServiceSharedFlows()
        this.list = res.items
        this.selected = new Set<string>()
      } catch (e) {
        this.loadingError = true
      } finally {
        this.loading = false
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

