<template>
  <NavBar title="Создание сценария">
    <template v-slot:default>
      <v-btn
        :class="demo ? 'onboarding' : ''"
        :disabled="saveLoading"
        :loading="saveLoading"
        variant="flat"
        @click="CreateFlow">
        Сохранить
      </v-btn>
    </template>
  </NavBar>
  <v-card>
    <v-card-text>
      <v-form validate-on="submit" ref="flow-form">
        <FlowForm @flow-changed="flowChanged" :demo="demo"/>
      </v-form>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {flowService} from "@/generated/services"
import {CreateFlowRequest, Task} from "@/generated/flow";
import {TaskArg, taskTypes} from '@/components/tasks/tasks'
import FlowForm from "@/components/flow/FlowForm.vue";
import NavBar from "@/components/NavBar.vue";

export default defineComponent({
  name: "CreateFlow",
  components: {NavBar, FlowForm},
  data() {
    return {
      demo: false,
      tasks: [] as Task[],
      stepTypes: taskTypes,
      show: this.showProp,
      item: {} as CreateFlowRequest,
      saveLoading: false,
    }
  },
  methods: {
    flowChanged(label: string, tasks: TaskArg[]) {
      this.tasks = []
      tasks.forEach(t => {
        if (t.task) {
          this.tasks.push(t.task)
        }
      })
      this.item.label = label
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['flow-form'].validate()

      return valid
    },
    async CreateFlow() {
      if (!(await this.validateForm())) {
        return
      }

      try {
        this.saveLoading = true
        this.item.tasks = this.tasks

        const res = await flowService.flowServiceCreateFlow({body: this.item})
        this.$router.push({name: "ViewFlow", params: {id: res.flow.id}})
      } finally {
        this.saveLoading = false
      }
    },
  },
  created() {
    const s = this.$route.query.demo
    if (s !== null && !Array.isArray(s)) {
      this.demo = Boolean(s)
    }
  }
})


</script>


<style>


</style>

