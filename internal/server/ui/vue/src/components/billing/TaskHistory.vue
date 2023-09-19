<template>
  <div class="text-h5">История задач</div>
  <div class="text-center" v-if="items.length === 0">Нет истории задач</div>
  <v-list v-else>
    <v-list-item
      density="compact"
      variant="plain"
      class="my-1"
      elevation="1"
      v-for="item in items"
      :key="item.taskId"
    >
      <div class="d-flex justify-space-between">
        <div>{{ item.taskType }}
          <span>{{ `Price: ${item.taskPrice} USD` }}</span>
        </div>
        <a :href="`/process/${item.processId}`" target="_blank">show process</a>
      </div>

    </v-list-item>
  </v-list>
  <v-btn v-if="showNext" @click="next" width="100%">Еще</v-btn>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";
import {TaskType} from "@/generated/flow";
import {helperService} from "@/generated/services";
import {TaskHistoryRecord} from "@/generated/helper";

export default defineComponent({
  name: "TaskHistory",
  created() {
    this.getHistory()
  },
  computed: {
    showNext(): boolean {
      if (this.items.length === 0) {
        return false
      }
      return this.items.length % 10 === 0
    }
  },
  data() {
    return {
      loading: false,
      offset: 0,
      items: [] as TaskHistoryRecord[]
    }
  },
  methods: {
    next() {
      this.offset += 10
      this.getHistory()
    },
    async getHistory() {
      try {
        this.loading = true
        const res = await helperService.helperServiceGetBillingHistory({body: {offset: this.offset.toString()}})
        this.items.push(...res.records)
      } catch (e) {
      } finally {
        this.loading = false
      }
    }
  }

})


</script>


<style>
</style>

