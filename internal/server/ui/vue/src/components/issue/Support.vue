<template>
  <v-dialog
    v-model="menu"
    :close-on-content-click="false"
    min-width="300px"
    max-width="400px"
  >
    <template v-slot:activator="{ props }">
      <v-btn :loading="loading" variant="text" @click="menu=true">{{ btnText }}</v-btn>
    </template>

    <template v-slot:default="{ props }">
      <v-card class="text-break text-pre-wrap">
        <v-card-title>
          Обращение в поддержку
        </v-card-title>
        <v-form ref="form">
          <v-card-text>
            <v-text-field v-if="!skipTitle" v-model="title" variant="outlined" auto-grow
                          placeholder="Заголовок" :rules="[required]"/>
          </v-card-text>
          <v-card-text>
            <v-textarea v-model="text" variant="outlined" auto-grow
                        placeholder="Добавление описания проблемы/предложения - уменьшает время на реализацию"
                        :rules="[required]"/>
          </v-card-text>
          <v-card-actions class="d-flex justify-end">
            <v-btn @click="menu = false">Отмена</v-btn>
            <v-btn @click="sendMessage">Отправить</v-btn>
          </v-card-actions>
        </v-form>

      </v-card>
    </template>
  </v-dialog>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {helperService, issueService} from "@/generated/services"
import {mapStores} from "pinia";
import {useSysStore} from "@/plugins/pinia";
import {required} from "@/components/tasks/menu/helper";

export default defineComponent({
  name: "Support",
  emits: ['created'],
  computed: {
    ...mapStores(useSysStore),
  },
  props: {
    skipTitle: {
      type: Boolean,
      required: true
    },
    btnText: {
      type: String,
      required: true
    },
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
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
    required,
    async sendMessage() {
      try {
        if (!await this.validate()) {
          return
        }
        this.loading = true
        await issueService.issueServiceCreateIssue({
          body: {
            description: this.text,
            title: this.title,
            processId: this.processId,
            taskid: this.taskId,
          }
        })
        this.$emit('created')
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
      title: '',
      loading: false
    }
  }

})


</script>
<style>
</style>

