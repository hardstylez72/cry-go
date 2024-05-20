<template>
  <NavBar title="Связи профилей">
  </NavBar>
  <div>

    <Loader v-if="loading"/>
    <div v-else>
      <div class="d-flex justify-start mx-8">
        <div class="mx-8">
          Тип профиля 1:
          <v-radio-group v-model="pType" hide-details density="compact">
            <v-radio class="mx-8" :value="ProfileType.EVM">EVM</v-radio>
          </v-radio-group>
          Подтип 1:
          <v-radio-group v-model="pSubType" hide-details density="compact" center-affix>
            <v-radio class="mx-8" :value="ProfileSubType.Metamask">{{ ProfileSubType.Metamask }}</v-radio>
          </v-radio-group>
        </div>

        <div class="mx-8">
          Тип профиля 2:
          <v-radio-group v-model="relPType" hide-details density="compact">
            <v-radio class="mx-8" :value="ProfileType.StarkNet">StarkNet</v-radio>
          </v-radio-group>
          Подтип 2:
          <v-radio-group v-model="relPSubType" hide-details density="compact" center-affix>
            <v-radio class="mx-8" :value="ProfileSubType.Braavos">{{ ProfileSubType.Braavos }}</v-radio>
            <v-radio class="mx-8" :value="ProfileSubType.UrgentX">{{ ProfileSubType.UrgentX }}</v-radio>
          </v-radio-group>
        </div>

      </div>
      <div>

        <v-list max-width="96vw" class="px-5" :nav="true">
          <v-list-item
            density="compact"
            variant="plain"
            class="my-0"
            v-for="item in list"
            :key="item.id"
            rounded
            height="auto"
            width="auto"
            elevation="1"
            style="border: 0 solid "
          >
            <RelationItem
              :item="item"
              :rel-p-sub-type="relPSubType"
              :rel-p-type="relPType"
            />
          </v-list-item>
        </v-list>
      </div>
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {profileService} from "@/generated/services"
import {Profile, ProfileRelation, ProfileType, ProfileSubType, UnrelatedProfile} from "@/generated/profile";
import CreateProfile from "@/components/profile/CreateProfile.vue";
import CheckBox from "@/components/CheckBox.vue";
import ProfileCard from "@/components/profile/ProfileCard.vue";
import BtnCheckProxy from "@/components/BtnCheckProxy.vue";
import EditProfile from "@/components/profile/EditProfile.vue";
import ProfilesDocs from "@/components/profile/ProfilesDocs.vue";
import ProfilesExport from "@/components/profile/ProfilesExport.vue";
import Loader from "@/components/Loader.vue";
import NavBar from "@/components/NavBar.vue";
import ProfileRelationSearch from "@/components/profile/ProfileRelationSearch.vue";
import RelationItem from "@/components/profile/RelationItem.vue";

export default defineComponent({
  name: "ProfileRelations",
  watch: {
    relPSubType: {
      handler() {
        this.loadList()
      }
    }
  },
  computed: {
    ProfileSubType() {
      return ProfileSubType
    },
    ProfileType() {
      return ProfileType
    }
  },
  components: {
    RelationItem,
    ProfileRelationSearch,
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
      pType: ProfileType.EVM,
      relPType: ProfileType.StarkNet,

      pSubType: ProfileSubType.Metamask,
      relPSubType: ProfileSubType.UrgentX,

      loading: false,

      list: [] as ProfileRelation[],
    }
  },

  methods: {

    async loadList() {
      this.loading = true
      try {
        const res = await profileService.profileServiceGetProfileRelations({
          body: {
            p1Type: this.pType,
            p1SubType: this.pSubType,
            p2Type: this.relPType,
            p2SubType: this.relPSubType,
          }
        })
        this.list = res.items
        this.list = this.list.sort((a, b) => a.num - b.num)
      } finally {
        this.loading = false
      }

    }
  },
  created() {
    this.loadList()
  }
})
</script>

<style>

</style>

