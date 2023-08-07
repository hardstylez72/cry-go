<template>
  <v-container>
    <NavBar title="Биллинг"/>
    <div class="mb-5">
      <div class="my-1 text-h6"> Пользователь: {{ store.email }}</div>
      <div class="d-flex flex-row my-1 text-h5 ">
        <div class="mr-1">Баланс: <b>{{ store.ass }}</b>
          <TopUp/>
        </div>

      </div>
      <div class="my-1 text-h5">Task execution price: <b>{{ store.taskPrice }}</b></div>

      <v-dialog width="300px">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" @click="dialog=true" density="compact">Типы задач</v-btn>
        </template>

        <v-card class="px-5 py-5">
          <b>Платные</b>
          <br/>
          <ul>
            <li v-for="task in store.payableTasks">
              {{ task }}
            </li>
          </ul>
          <b>Бесплатные</b>
          <br/>
          <ul>
            <li v-for="task in store.nonpayableTasks">
              {{ task }}
            </li>
          </ul>
        </v-card>
      </v-dialog>
    </div>
    <TaskHistory class="mt-5"/>
    <PaymentHistory class="mt-5"/>
  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";
import TaskHistory from "@/components/billing/TaskHistory.vue";
import TopUp from "@/components/billing/TopUp.vue";
import PaymentHistory from "@/components/billing/PaymentHistory.vue";
import NavBar from "@/components/NavBar.vue";

export default defineComponent({
  name: "Billing",
  components: {NavBar, PaymentHistory, TopUp, TaskHistory},
  created() {
  },
  data() {
    return {
      dialog: false
    }
  },
  computed: {
    ...mapStores(useUserStore),
    store() {
      return this.userStore
    }
  }
})


</script>


<style>
</style>

