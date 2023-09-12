<template>

  <v-container>
    <NavBar :title="`Заявка: ${item.id}`">
    </NavBar>

    <Loader v-if="listLoading"/>
    <div v-else>
      <v-card elevation="2" class="mx-1 my-1">
        <v-card-title class="d-flex justify-space-between text-break text-pre-wrap">
          <div>
            <b>{{ item.title }}</b>
            <span class="mx-1"> [{{ item.status }}]</span>
          </div>
          <span class="mx-1"> {{ formatTime(item.createdAt) }}</span>
        </v-card-title>
        <v-card-text>{{ item.description }}</v-card-text>
      </v-card>

      <v-list class="mx-3 my-3">
        <v-list-item v-if="comments.length === 0" elevation="1" variant="flat">
          <div style="width: 100%">Нет комментариев</div>
        </v-list-item>
        <v-list-item v-for="comment in comments" elevation="1" v-else>
          <span>{{ comment.description }}</span>
          <span class="mx-1 float-right"> <i>{{ formatTime(comment.createdAt) }}</i></span>
        </v-list-item>
        <v-list-item class="my-3">
          <v-form ref="form">
            <v-textarea v-model="commentText" :rules="[required]"/>
            <v-btn width="100%" @click="addComment">Добавить комментарий</v-btn>
          </v-form>

        </v-list-item>
      </v-list>
    </div>


  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import NavBar from "@/components/NavBar.vue";
import Login from "@/components/Login.vue";
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";
import General from "@/components/landing/modules/General.vue";
import {Issue, IssueComment} from "@/generated/issue";
import {issueService} from "@/generated/services";
import Support from "@/components/issue/Support.vue";
import {formatTime} from "../helper";
import Loader from "@/components/Loader.vue";
import {required} from "@/components/tasks/menu/helper";

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
  data() {
    return {
      listLoading: false,
      item: {} as Issue,
      comments: [] as IssueComment[],
      commentText: '',
      commentLoading: false,
    }
  },
  components: {Loader, Support, General, NavBar},
  methods: {
    required,
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
    async addComment() {
      try {
        if (!await this.validate()) {
          return
        }
        this.commentLoading = true
        await issueService.issueServiceAddComment({
          body: {
            issueId: this.issueId,
            text: this.commentText,
          }
        })
        this.commentText = ''
        await this.loadIssues()
      } catch (e) {
        console.error(e)
      } finally {
        this.commentLoading = false
      }
    },
    formatTime,
    async loadIssues() {
      try {
        this.listLoading = true
        const res = await issueService.issueServiceIssue({
          body: {
            id: this.issueId,
          }
        })
        this.item = res.issue
        this.comments = res.comments

      } catch (e) {

      } finally {
        this.listLoading = false
      }
    },
    reset() {
      this.offset = 0;
    }
  },
  computed: {
    issueId: {
      get(): string {
        const pId = this.$route.params.id.toString()
        if (pId) {
          return pId
        }
      },
      set() {

      }
    }
  },

  created() {
    this.loadIssues()
  }
})


</script>


<style>
</style>
