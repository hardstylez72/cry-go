<template>
  <v-container>
    <NavBar title="Статистика по ZkSync">
    </NavBar>
    <span class="d-inline-flex">
             <v-switch class="mx-2" label="подробнее" v-model="more" density="compact" hide-details/>
          <v-switch class="mx-2" v-if="more" label="еще подробнее" v-model="more2" density="compact" hide-details/>
      </span>


    <ProfileSearch
      v-model="selectedProfiles"
      :disabled="loading"
    />

    <v-list>
      <v-list-item>
        <div class="d-flex justify-space-between">
          <div><b>Дни</b></div>
          <div><b>Месяцы</b></div>
          <div><b>Мой последний раз</b></div>
          <div><b>Транзы</b></div>
          <div><b>У-транзы</b></div>
          <div><b>Убьем</b></div>
          <div v-if="more"><b>Подробности</b></div>
        </div>
      </v-list-item>
      <ZkSyncProfile v-for="p in selectedProfiles" :profile="p" :more="more" :more2="more2"/>
    </v-list>
  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import Loader from "@/components/Loader.vue";
import NavBar from "@/components/NavBar.vue";
import {formatTime, networkProps} from "@/components/helper";
import {publicService, statService} from "@/generated/services";
import {SwapStatItem, SwapStatRes} from "@/generated/public";
import {TaskType} from "@/generated/flow";
import {taskProps} from "@/components/tasks/tasks";
import {Interaction, Stat} from "@/generated/stat";
import boba from './dd'
import ProfileSearch from "@/components/profile/ProfileSearch.vue";
import {Profile} from "@/generated/profile";
import ZkSyncProfile from "@/components/stats/checker/ZkSyncProfile.vue";

export default defineComponent({
  name: "ZkSyncStat",
  data() {
    return {
      data: [] as Stat[],
      loading: false,
      more: false,
      more2: false,
      selectedProfiles: [] as Profile[]
    }
  },
  computed: {},
  components: {ZkSyncProfile, ProfileSearch, NavBar, Loader},
  methods: {
    formatTime,
  },
  created() {
  }
})


</script>


<style scoped>

</style>

