<template>
  <v-dialog
      v-model="menu"
      :close-on-content-click="false"
      width="200px"
  >
    <template v-slot:activator="{ props }">
      <v-btn variant="flat" density="compact" color="red" @click="remove">Удалить</v-btn>
    </template>

    <template v-slot:default>
      <v-card>
        <v-card-title>
          Process delete
        </v-card-title>
        <v-card-text>
          Are you sure?
        </v-card-text>
        <v-card-actions class="d-flex justify-end">
          <v-btn @click="menu = false">No</v-btn>
          <v-btn @click="remove">Yes</v-btn>
        </v-card-actions>
      </v-card>
    </template>
  </v-dialog>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {processService} from "@/generated/services"
import {Process, ProcessProfile, ProcessStatus, ProcessTask} from "@/generated/process";
import ViewFlow from "@/components/flow/OldView.vue";
import StatusCard from "@/components/StatusCard.vue";
import {formatTime, GetStatusColor, Timer} from "@/components/helper";
import ProcessTaskMenu from "@/components/process/ProcessTaskMenu.vue";
import ProfileCard from "@/components/profile/ProfileCard.vue";
import BtnProcessStopResume from "@/components/BtnProcessStopResume.vue";
import dayjs from "dayjs";

export default defineComponent({
  name: "DeleteProcess",
  emits: ['processRemoved'],
  props: {
    processId: {
      type: String,
      required: true
    }
  },
  methods: {
    async remove() {
      try {
        await processService.processServiceCancelProcess({body: {processId: this.processId}})
        this.$emit('processRemoved')
      } catch (e) {

      }
    },
  },
  data() {
    return {
      menu: false,
    }
  }

})


</script>
<style>
</style>

