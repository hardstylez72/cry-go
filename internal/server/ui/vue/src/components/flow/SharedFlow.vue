<template>
  <div>
    <NavBar :title="`Общий Сценарий: ${flow.label}`">
      <template v-slot:default>

        <v-btn v-if="allowedRemoval" @click="removeFlow" variant="flat" class="mx-1">
          <v-icon v-if="isMobile()" icon="mdi-delete"/>
          <span v-else>Удалить</span>
        </v-btn>
        <v-btn @click="useSharedFlow" variant="flat" class="mx-1">
          <v-icon v-if="isMobile()" icon="mdi-handshake-outline"/>
          <span v-else>Использовать</span>
        </v-btn>
      </template>
    </NavBar>
  </div>

  <v-spacer class="my-6"/>
  <v-form validate-on="submit" ref="flow-form">
    <FlowForm v-if="!flowLoading" :disable="disabled" :label-value="flow.label" :tasks-value="tasks"/>
  </v-form>

</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService, processService, profileService} from "@/generated/services"
import {flow_Flow as Flow, FlowShared, Task, TaskType} from "@/generated/flow";
import TaskStargateBridge from "@/components/tasks/BRIDGE/Stargate/Block.vue";
import {Profile} from "@/generated/profile";
import TaskDelay from "@/components/tasks/OTHER/Delay/Block.vue";
import ProfileCard from "@/components/profile/ProfileCard.vue";
import {taskComponentMap, TaskArg, taskTypes} from "@/components/tasks/tasks";
import {Delay, isMobile, Timer} from "@/components/helper";
import FlowForm from "@/components/flow/FlowForm.vue";
import ProfileSearch from "@/components/profile/ProfileSearch.vue";
import NavBar from "@/components/NavBar.vue";
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";

export default defineComponent({
  name: "SharedFlow",
  components: {NavBar, ProfileSearch, FlowForm, ProfileCard},
  props: {
    propId: {
      type: String,
      required: false,
    }
  },
  computed: {
    ...mapStores(useUserStore),
    allowedRemoval(): boolean {
      return this.userStore.publisher && this.userStore.id === this.flow.creatorId
    }
  },
  data() {
    return {
      flowId: "" as string,
      flow: {} as FlowShared,
      disabled: true,
      tasks: [] as Task[],
      flowLoading: true
    }
  },
  methods: {
    async removeFlow() {
      try {
        await flowService.flowServiceHideFlow({body: {id: this.flowId}})
        this.$router.push({name: 'SharedFlowList'})
      } finally {

      }
    },
    async useSharedFlow() {
      try {
        const res = await flowService.flowServiceUseSharedFlow({body: {id: this.flowId}})
        this.$router.push({name: 'Flow', params: {id: res.id}})
      } finally {

      }
    },
    isMobile,
    async loadFlow() {
      if (this.propId) {
        this.flowId = this.propId
      } else if (this.$route.params.id) {
        this.flowId = this.$route.params.id.toString()
      }

      try {
        this.flowLoading = true
        const res = await flowService.flowServiceSharedFlow({body: {id: this.flowId}})
        this.flow = res.flow
        this.tasks = this.flow.tasks
      } finally {
        this.flowLoading = false
      }
    }
  },
  async mounted() {
    await this.loadFlow()

  }
})


</script>


<style scoped>

.header {
  margin: auto;
  height: auto;
  z-index: 100;
  right: 2px;
  left: 56px;
  background-color: #FFF3E0;
  box-shadow: rgba(0, 0, 0, 0.16) 0px 3px 6px, rgba(0, 0, 0, 0.23) 0px 3px 6px;
}
</style>

