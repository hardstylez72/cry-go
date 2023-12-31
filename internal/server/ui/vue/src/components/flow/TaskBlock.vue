<template>
  <div class="task-container" :style="`background-color: ${color}`">

    <div class="d-inline-flex align-center" v-if="task.taskType !== TaskType.Delay">

      <TaskChip :task-type="task.taskType" text-style/>
      <span class="font-weight-thin" style="font-weight: 100;">({{ p.job }})</span>
    </div>
    <div class="text-center justify-center align-center" v-else>
      <div>
        <v-icon icon="mdi-sleep" size="15"/>
      </div>
      <div>
        <v-icon icon="mdi-sleep" size="15"/>
      </div>
    </div>
    <div class="font-italic text-center">

      <div v-if="p.job === TaskJob.Swap">
        <span>{{ `${p.swap(task).from}->${p.swap(task).to}` }}</span>
      </div>
      <div v-else-if="p.job === TaskJob.NFT">
        <span>mint</span>
        <span v-if="p.nft(task) && p.nft(task).toNetwork">& bridge</span>
      </div>
      <div v-else-if="p.job === TaskJob.Bridge">
        <span class="d-inline-flex align-center">
        <NetworkChip no-label :network="p.bridge(task).fromNetwork"/>
        ->
        <NetworkChip no-label :network="p.bridge(task).toNetwork"/>
        </span>
      </div>
      <div v-else-if="task.taskType === TaskType.Delay">
      </div>
      <div v-else>
        {{ p.service.op }}
      </div>

    </div>

  </div>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {flowService} from "@/generated/services"
import {flow_Flow as Flow, FlowPreviewRes, Task, TaskType} from "@/generated/flow";
import CreateFlow from "@/components/flow/OldCreateFlow.vue";
import CheckBox from "@/components/CheckBox.vue";
import {getFlow, taskProps} from "@/components/tasks/tasks";
import Loader from "@/components/Loader.vue";
import NavBar from "@/components/NavBar.vue";
import {TaskJob} from "@/components/tasks/utils";
import TaskChip from "@/components/tasks/TaskChip.vue";
import NetworkChip from "@/components/tasks/NetworkChip.vue";

export default defineComponent({
  name: "TaskBlock",
  components: {NetworkChip, TaskChip},
  props: {
    task: {
      required: true,
      type: Object as PropType<Task>
    },
    color: {
      required: false,
      type: String,
      default: 'none'
    }
  },
  data() {
    return {}
  },
  computed: {
    TaskType() {
      return TaskType
    },
    TaskJob() {
      return TaskJob
    },
    p() {
      return taskProps[this.task.taskType]
    },
    taskProps() {
      return taskProps
    }
  },
  methods: {},
  created() {

  }
})


</script>


<style>
.task-container {
  font-size: 13px;
  padding: 1px 1px 1px 1px;
  max-height: 60px;
  max-width: 150px;
  /*min-height: 45px;*/
  /*min-width: 70px;*/
  border: 1px solid gray;
  border-radius: 5px;
}

</style>

