<template>
  <v-container>
    <NavBar title="Статистика по ZkSync">
    </NavBar>
    <span class="d-inline-flex">
             <v-switch class="mx-2" label="подробнее" v-model="more" density="compact" hide-details/>
          <v-switch class="mx-2" v-if="more" label="еще подробнее" v-model="more2" density="compact" hide-details/>
            <span v-if="selectedProfiles.length" class="mx-2 my-2">Скопировать статистику<Copy
              :text="copy.join('\n\r')"/></span>
      </span>


    <ProfileSearch
      :profile-type-filter="[ProfileType.EVM]"
      v-model="selectedProfiles"
      :disabled="loading"
      :shufflable="false"
    />


    <div class="mx-3">
      <div class="d-flex justify-space-between">
        <div><b>Д/M</b>
          <Explain text="Количество активных дней/месяцев"/>
        </div>
        <div><b>Мой последний раз</b></div>
        <div><b>Транзы</b>
          <Explain text="Количество исходящих транзакций"/>
        </div>
        <div><b>У-транзы</b>
          <Explain text="Количество уникальных контрактов, с которыми было взаимодействие"/>
        </div>
        <div><b>Убьем</b>
          <Explain text="Суммарный обьем токенов (ETH/USDT/USDC/LUSD/BUSD) отправленных с адреса"/>
        </div>
        <div v-if="more"><b>Подробности</b></div>
      </div>
    </div>
    <v-list>

      <ZkSyncProfile v-for="p in selectedProfiles" :profile="p" :more="more" :more2="more2" @loaded="loaded"/>
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
import {ProfileType} from "@/generated/helper";
import Explain from "@/components/Explain.vue";
import Copy from "@/components/profile/Copy.vue";

export default defineComponent({
  name: "ZkSyncStat",
  data() {
    return {
      data: [] as Stat[],
      loading: false,
      more: false,
      more2: false,
      selectedProfiles: [] as Profile[],
      statMap: new Map<string, Stat>(),
      copy: [] as string[],
    }
  },
  computed: {
    ProfileType() {
      return ProfileType
    }
  },
  components: {Copy, Explain, ZkSyncProfile, ProfileSearch, NavBar, Loader},
  methods: {
    formatTime,
    loaded(wa: string, d: Stat) {
      this.statMap.set(wa, d)

      this.statMap.forEach((v, k) => {
        if (v.interactions.length) {
          this.copy.push(`${k}; ${v.activeDays}; ${v.activeMonths}; ${formatTime(v.lastActivity)}; ${v.interactions.length}; ${v.uniqAddress}; ${Number(v.totalUsdAmount).toFixed(2)}`)
        } else {
          this.copy.push(`${k};;;;;;`)
        }
      })
    }
  },
  created() {
  }
})


</script>


<style scoped>

</style>

