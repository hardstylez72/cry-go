<template>

  <v-container>
    <NavBar :title="`Заявка: ${item.id}`">
    </NavBar>

    <Loader v-if="listLoading"/>
    <div v-else>
      <v-card elevation="2" class="mx-1 my-1">
        <v-card-title class="d-flex justify-space-between text-break text-pre-wrap">
          <div>
            <span class="mx-1">  <StatusChip :status="item.status"/></span>
            <b>{{ item.title }}</b>
          </div>
          <span class="mx-1"> {{ formatTime(item.createdAt) }}</span>
        </v-card-title>
        <v-card-text>{{ item.description }}</v-card-text>
        <v-card-actions class="d-flex justify-end align-center">
          <div style="width: 200px; height: 50px">
            <v-select v-if="support" :items="statuses" :loading="updating" v-model="status" variant="outlined"
                      density="compact" label="изменить статус"/>
          </div>

          <v-btn @click="remove" :loading="removing" v-if="item.my">Удалить</v-btn>
        </v-card-actions>
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
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";
import General from "@/components/landing/modules/General.vue";
import {Issue, IssueComment, IssueStatus} from "@/generated/issue";
import {issueService} from "@/generated/services";
import Support from "@/components/issue/Support.vue";
import {formatTime} from "../helper";
import Loader from "@/components/Loader.vue";
import {required} from "@/components/tasks/menu/helper";
import StatusChip from "@/components/issue/StatusChip.vue";

export default defineComponent({
  name: "Issues",
  watch: {
    status: {
      async handler() {
        try {
          this.updating = true
          await issueService.issueServiceIssueUpdateStatus({body: {status: this.status, issueId: this.issueId}})
          await this.loadIssues()
        } catch (e) {

        } finally {
          this.updating = false
        }
      }
    },
    issueType: {
      handler() {
        this.reset()
        this.loadIssues()
      }
    }
  },
  data() {
    return {
      status: null as IssueStatus,
      statuses: [
        IssueStatus.Created,
        IssueStatus.Done,
        IssueStatus.Processing,
        IssueStatus.Rejected,
        IssueStatus.Stop,
      ],
      listLoading: false,
      item: {} as Issue,
      comments: [] as IssueComment[],
      commentText: '',
      commentLoading: false,
      removing: false,
      updating: false,
    }
  },
  components: {Loader, Support, General, NavBar, StatusChip},
  methods: {
    async remove() {
      try {
        this.removing = true
        await issueService.issueServiceDeleteIssue({body: {issueId: this.issueId}})
        this.$router.push({name: 'Issues'})
      } catch (e) {

      } finally {
        this.removing = false
      }
    },
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
    IssueStatus() {
      return IssueStatus
    },
    support(): boolean {
      return this.userStore.support
    },
    ...mapStores(useUserStore),
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
