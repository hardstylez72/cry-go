<template>
  <v-btn variant="flat" :loading="loading" v-if="processCanBeStopped() && !errorHappen()"
         @click="Stop(item.id)"
         color="orange">Стоп
  </v-btn>
  <v-btn variant="flat" v-else-if="item.status !== ProcessStatus.StatusDone && !errorHappen()" :loading="loading"

         @click="Resume(item.id)" color="green">
    {{ processNotStarted() ? "Старт" : "Возобновить" }}
  </v-btn>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {processService} from "@/generated/services"
import {Process, ProcessStatus} from "@/generated/process";
import {Delay} from "@/components/helper";

export default defineComponent({
  name: "BtnProcessStopResume",
  props: {
    item: {
      type: Object as PropType<Process>,
      required: true,
    },
    size: {
      type: String as PropType<'compact' | 'comfortable'>,
      required: true,
    },
  },

  data() {
    return {
      loading: false
    }
  },
  computed: {
    ProcessStatus() {
      return ProcessStatus
    },
  },
  methods: {
    errorHappen(): boolean {
      if (this.item.status === ProcessStatus.StatusError) {
        return true
      }
      return false
    },
    processCanBeStopped(): boolean {

      if (this.item.status === ProcessStatus.StatusReady) {
        return false
      }

      if (this.item.status === ProcessStatus.StatusStop) {
        return false
      }

      if (this.item.status === ProcessStatus.StatusDone) {
        return false
      }
      return true
    },

    processNotStarted(): boolean {

      if (this.item.status === ProcessStatus.StatusReady) {
        return true
      }

      if (this.item.status !== ProcessStatus.StatusStop) {
        return false
      }

      return !this.item.profiles.some(p => p.status !== ProcessStatus.StatusReady)
    },

    async Stop(id: string) {
      try {
        this.loading = true
        await processService.processServiceStopProcess({body: {processId: id}})
        this.$emit("updated")
      } finally {
        this.loading = false
      }
    },
    async Resume(id: string) {
      try {
        this.loading = true
        await processService.processServiceResumeProcess({body: {processId: id}})
        this.$emit("updated")
      } finally {
        this.loading = false
      }
    },
    async mounted() {
    },
  }
})


</script>


<style>


</style>

