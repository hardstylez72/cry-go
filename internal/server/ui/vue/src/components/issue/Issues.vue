<template>

  <v-container>
    <NavBar title="Поддержка">
      <template v-slot:default>
        <Support @created="issueCreated" btn-text="Создать заявку" :skip-title="false"/>
      </template>
    </NavBar>

    <v-radio-group style="float: left" v-model="issueType" inline="true">
      <v-radio class="mx-4" value="мои">Мои</v-radio>
      <v-radio class="mx-4" value="все">Все</v-radio>
    </v-radio-group>
    <div class="d-flex justify-center">


      <div class="text-h6"><b>Список заявок</b></div>
    </div>


    <Loader v-if="listLoading"/>

    <v-list width="100%" v-else>
      <v-list-item v-if="items.length === 0" elevation="1">
        Заявки не найдены
      </v-list-item>
      <v-list-item v-else v-for="item in items" elevation="1" :to='`/issue/${item.id}`'>
        <b class="mx-1">{{ item.title }}</b>
        <i class="mx-1">{{ item.description }}</i>
        <span class="mx-1"> [{{ item.status }}]</span>
        <span class="mx-1"> {{ formatTime(item.createdAt) }}</span>
      </v-list-item>
    </v-list>


  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import NavBar from "@/components/NavBar.vue";
import Login from "@/components/Login.vue";
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";
import General from "@/components/landing/modules/General.vue";
import {Issue} from "@/generated/issue";
import {issueService} from "@/generated/services";
import Support from "@/components/issue/Support.vue";
import {formatTime} from "../helper";
import Loader from "@/components/Loader.vue";

export default defineComponent({
  name: "Issues",
  watch: {
    issueType: {
      handler() {
        this.reset()
        this.loadIssues()
      }
    }
  },
  components: {Loader, Support, General, NavBar},
  methods: {
    formatTime,
    issueCreated() {
      this.loadIssues()
    },
    async loadIssues() {
      try {
        this.listLoading = true
        if (this.issueType === 'мои') {
          const res = await issueService.issueServiceIssuesUser({
            body: {
              limit: this.limit,
              offset: this.offset,
            }
          })
          this.items = res.items
        } else {
          const res = await issueService.issueServiceIssues({
            body: {
              limit: this.limit,
              offset: this.offset,
            }
          })
          this.items = res.items
        }
      } catch (e) {

      } finally {
        this.listLoading = false
      }
    },
    reset() {
      this.offset = 0;
    }
  },
  computed: {},
  data() {
    return {
      limit: 30,
      offset: 0,
      listLoading: false,
      issueType: 'мои' as 'все' | 'мои',
      items: [] as Issue[]
    }
  },
  created() {
    this.loadIssues()
  }
})


</script>


<style>
</style>

