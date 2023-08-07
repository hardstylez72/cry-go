<template>
  <v-dialog
    v-model="menu"
    :close-on-content-click="false"
    min-width="300px"
    max-width="400px"
  >
    <template v-slot:activator="{ props }">
      <v-btn width="85px" :loading="loading" variant="text" @click="menu=true" density="compact">Support</v-btn>
    </template>

    <template v-slot:default="{ props }">
      <v-card class="text-break text-pre-wrap">
        <v-card-title>
          Обращение в поддержку
        </v-card-title>
        <v-card-text>
          <v-textarea v-model="text" variant="outlined" auto-grow
                      placeholder="Добавление описания проблемы/предложения - уменьшает время на реализацию"/>
        </v-card-text>
        <v-card-actions class="d-flex justify-end">
          <v-btn @click="menu = false">Отмена</v-btn>
          <v-btn @click="sendMessage">Отправить</v-btn>
        </v-card-actions>
      </v-card>
    </template>
  </v-dialog>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {helperService} from "@/generated/services"
import {mapStores} from "pinia";
import {useSysStore} from "@/plugins/pinia";

export default defineComponent({
  name: "Support",
  computed: {
    ...mapStores(useSysStore),
  },
  props: {
    processId: {
      type: String,
      required: false
    },
    taskId: {
      type: String,
      required: false
    }
  },
  methods: {
    async sendMessage() {
      try {
        this.loading = true
        await helperService.helperServiceSupportMessage({
          body: {
            processId: this.processId,
            taskId: this.taskId,
            text: this.text
          }
        })
        this.menu = false
        this.sysStore.showSnackBar("Обращение отправлено", false)
      } catch (e) {
        this.sysStore.showSnackBar("Ошибка при отправке обращения", true)
      } finally {
        this.loading = false
      }
    },
  },
  data() {
    return {
      menu: false,
      text: '',
      loading: false
    }
  }

})


</script>
<style>
</style>

