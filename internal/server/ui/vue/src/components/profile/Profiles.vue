<template>
  <NavBar :title="`Профили: ${selectProfileType}`">
    <template v-slot:default>
      <v-btn :class="noProfiles ? 'onboarding mx-2 my-2' : `mx-2 my-2`" v-if="selectedSome" variant="flat" color="black"
             @click=DeleteProfiles>Удалить
      </v-btn>
      <v-btn v-else :class="noProfiles ? 'onboarding' : `mx-2 my-2`" variant="flat"
             @click="$router.push({name: 'CreateProfileBatch'})">
        Добавить
      </v-btn>
    </template>
  </NavBar>
  <div>

    <Loader v-if="loading && offset === 0"/>
    <div v-else>
      <div class="mr-6 d-flex justify-start">
        <v-radio-group v-model="selectProfileType" :inline="true" hide-details density="compact">
          <v-radio class="mx-8" :value="ProfileType.EVM">EVM</v-radio>
          <v-radio class="mx-8" :value="ProfileType.StarkNet">StarkNet</v-radio>
          <v-btn class="mx-3 my-2" density="default" @click="$router.push({name: 'ProfileRelations'})">Связи</v-btn>
        </v-radio-group>

        <div style="width: 100px">
          <v-checkbox
            density="compact"
            label="Адреса"
            v-model="showWalletAddrs"
            hide-details
          />
        </div>
      </div>
      <div v-if="noProfiles" class="my-8 mx-8 text-center">
        Профиль - это абстракция над "кошельком" в блокчейне, она помогает обезопасить приватные ключи.
        <br/>
        Без профилей - не будет возможности насладиться автоматизацией. Создание не займет много времени, обещаю:)
        <br/>
        Нажмите кнопку <a href="/profiles/batch/create">добавить</a>
      </div>
      <div v-else>

        <v-list max-width="96vw" class="px-5" :nav="true">
          <v-list-item
            density="compact"
            variant="plain"
            class="my-0"
            v-for="item in getList"
            :key="item.id"
            rounded
            height="auto"
            width="auto"
            elevation="1"
            style="border: 0 solid "
          >

            <v-row align="center">
              <v-col cols="1">
                <CheckBox style="height: 30px" density="compact" :id="item.id" focused
                          @CheckboxChanged="CheckboxChanged"/>
              </v-col>

              <v-col>
                <ProfileCard :label="item.num" :profile-id="item.id"/>
              </v-col>
              <v-col class="font-weight-bold" style="font-size: 12px" v-if="showWalletAddrs">{{
                  item.mmskId
                }}
              </v-col>
              <v-col class="font-italic" style="font-size: 14px">{{
                  item.subType
                }}
              </v-col>
              <v-col>
                <i v-if="item.label">({{ item.label }})</i>
              </v-col>
              <v-col>
                <BtnCheckProxy :proxy="item.proxy"/>
              </v-col>
              <v-col>
                <EditProfile :profile-id="item.id" @updated="UpdateList"/>
              </v-col>

            </v-row>
          </v-list-item>
        </v-list>
        <v-btn width="100%" @click="loadMore" :loading="loading" v-if="!loadedAll">Загрузить еще
        </v-btn>

      </div>
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {profileService} from "@/generated/services"
import {Profile, ProfileType} from "@/generated/profile";
import CreateProfile from "@/components/profile/CreateProfile.vue";
import CheckBox from "@/components/CheckBox.vue";
import ProfileCard from "@/components/profile/ProfileCard.vue";
import BtnCheckProxy from "@/components/BtnCheckProxy.vue";
import EditProfile from "@/components/profile/EditProfile.vue";
import ProfilesDocs from "@/components/profile/ProfilesDocs.vue";
import ProfilesExport from "@/components/profile/ProfilesExport.vue";
import Loader from "@/components/Loader.vue";
import {formatTime} from "../helper";
import NavBar from "@/components/NavBar.vue";

export default defineComponent({
  name: "Profiles",
  components: {
    NavBar,
    Loader,
    ProfilesExport,
    ProfilesDocs,
    EditProfile,
    BtnCheckProxy,
    CheckBox,
    ProfileCard
  },
  props: {},
  data() {
    return {
      showWalletAddrs: false,
      list: [] as Profile[],
      showCreateProfileDialog: false,
      selected: new Set<string>(),
      loading: true,
      loadingError: false,
      showLabels: false,
      selectProfileType: ProfileType.EVM,
      offset: 0,
      loadedAll: false,
    }
  },
  watch: {
    selectProfileType: {
      handler() {
        this.reset()
        this.UpdateList()
      }
    }
  },
  computed: {
    ProfileType() {
      return ProfileType
    },
    noProfiles() {
      return this.getList.length === 0
    },
    getList(): Profile[] {
      return this.list
    },
    selectedSome(): boolean {
      return this.selected.size > 0
    },
  },
  methods: {
    reset() {
      this.loadedAll = false
      this.list = []
      this.offset = 0
    },
    formatTime,
    CheckboxChanged(id: string, value: boolean) {
      if (value) {
        this.selected.add(id)
      } else {
        this.selected.delete(id)
      }
    },
    async UpdateList() {
      try {
        this.loadingError = false
        this.loading = true
        const res = await profileService.profileServiceListProfile({
          body: {
            type: this.selectProfileType,
            offset: this.offset.toString()
          }
        })
        if (res.profiles.length === 0) {
          this.loadedAll = true
        }
        this.list.push(...res.profiles)
        this.selected = new Set<string>()
      } catch (e) {
        this.loadingError = true
      } finally {
        this.loading = false
      }

    },
    async DeleteProfiles() {
      for (let id of this.selected.keys()) {
        await profileService.profileServiceDeleteProfile({body: {id: id}})
      }
      await this.UpdateList()
    },
    async loadMore() {
      this.offset += 30
      await this.UpdateList()
    }
  },
  created() {
    this.UpdateList()
  }
})
</script>

<style>

</style>

