<template>
  <v-container>
    <Loader v-if="loading"/>
    <div v-else>
      <v-list-item v-if="!data.lastActivity" elevation="1">
        <v-list-item-title>
          <div class="d-inline-flex">
            <span>{{ profile.num }}<i>({{ profile.label }})</i></span>
            <ProfileAddress class="mx-1" :addr="profile.mmskId" v-model="addrShow"/>
          </div>
        </v-list-item-title>
        Пусто
      </v-list-item>
      <div v-else>
        <v-list-item
          elevation="1"
          width="100%"
          rounded
        >
          <v-list-item-title>
            <div class="d-inline-flex">
              <span>{{ profile.num }}<i>({{ profile.label }})</i></span>
              <ProfileAddress class="mx-1" :addr="profile.mmskId" v-model="addrShow"/>
            </div>
          </v-list-item-title>
          <div class="d-flex justify-space-between">
            <div class="mx-1">{{ data.activeDays }}/{{ data.activeMonths }}</div>
            <div class="mx-1">{{ formatTime((data.lastActivity)) }}</div>
            <div class="mx-1">{{ data.txCount }}</div>
            <div class="mx-1">{{ data.uniqAddress }}</div>
            <div class="mx-1">{{ data.totalUsdAmount.toFixed(2) }} USD</div>

            <div class="mx-1" v-if="more" style="width: 350px">
              <div v-for="item in interactions" class="mx-5">
                <a :href="item.toUrl">{{ item.serviceName }}</a>
                Txs: ({{ item.txs }}) V: {{ item.amountUsd.toFixed(2) }} USD
              </div>
              <br/>
            </div>

          </div>

        </v-list-item>

      </div>


    </div>
  </v-container>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import Loader from "@/components/Loader.vue";
import NavBar from "@/components/NavBar.vue";
import {formatTime, networkProps} from "@/components/helper";
import {publicService, statService} from "@/generated/services";
import {SwapStatItem, SwapStatRes} from "@/generated/public";
import {TaskType} from "@/generated/flow";
import {taskProps} from "@/components/tasks/tasks";
import {Interaction, Stat} from "@/generated/stat";
import boba from './dd'
import {Profile} from "@/generated/profile";
import ProfileAddress from "@/components/profile/ProfileAddress.vue";

export default defineComponent({
  name: "ZkSyncProfile",
  emits: ['loaded'],
  props: {
    addrShowP: {
      type: Boolean,
      required: false,
      default: false,
    },
    profile: {
      type: Object as PropType<Profile>,
      required: true
    },
    more: {
      type: Boolean,
      required: false
    },
    more2: {
      type: Boolean,
      required: false
    }
  },
  data() {
    return {
      addrShow: false,
      data: {} as Stat,
      loading: false,
    }
  },
  computed: {
    interactions(): Interaction[] {
      if (this.more2) {
        return this.data.interactions
      }
      return this.data.interactions.filter(i => i.serviceName !== 'незвестный сервис' && i.serviceName !== 'noname')
    }
  },
  components: {ProfileAddress, NavBar, Loader},
  methods: {
    formatTime,
    async loadData() {
      try {
        this.loading = true
        const res = await statService.statServiceZkSyncStat({body: {profileId: this.profile.id}})


        this.data = res.stat

        this.data.interactions.sort((a, b) => {
          if (a.serviceName !== 'незвестный сервис' && a.serviceName !== 'noname') {
            return -1
          }
          return 1
        })

        this.$emit('loaded', this.profile.mmskId, this.data)
      } catch (e) {

      } finally {
        this.loading = false
      }
    },
  },
  created() {
    this.loadData()
  }
})


</script>


<style scoped>

</style>

