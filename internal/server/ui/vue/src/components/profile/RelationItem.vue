<template>
  <div class="d-inline-flex align-center">
    <v-chip class="mx-2">
      <b class="mx-1">{{ item.num }}</b>
      <i>{{ item.type }} ({{ item.subType }})</i>
    </v-chip>
    <div v-if="relationExist">
      <v-chip class="mx-2" :closable="true" @click:close="close">
        <b class="mx-1">{{ rel.num }}</b> <i>{{ rel.type }} ({{ rel.subType }})</i>
      </v-chip>
    </div>
    <span v-else>
        <ProfileRelationSearch
          class="mx-2"
          style="width: 200px"
          :disabled="false"
          :p1-type="item.type"
          :p1-sub-type="item.subType"
          :p2-type="relPType"
          :p2-sub-type="relPSubType"
          @selected="selected"
          :p-id="item.id"
          :num="item.num"
        />
    </span>
  </div>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import ProfileRelationSearch from "@/components/profile/ProfileRelationSearch.vue";
import {ProfileRelation, ProfileSubType, ProfileType, UnrelatedProfile} from "@/generated/profile";
import {profileService} from "@/generated/services";

export default defineComponent({
  name: "RelationItem",
  components: {ProfileRelationSearch},
  created() {
    this.rel = this.item.rel
  },
  props: {
    item: {
      type: Object as PropType<ProfileRelation>,
      required: true,
    },
    relPType: {
      type: String as PropType<ProfileType>,
      required: true,
    },
    relPSubType: {
      type: String as PropType<ProfileSubType>,
      required: true,
    }
  },
  computed: {
    relationExist(): Boolean {
      if (this.rel) {
        return true
      }
      return false
    }
  },
  data: function () {
    return {
      rel: {} as UnrelatedProfile
    }
  },
  methods: {

    async close() {
      try {
        await profileService.profileServiceDeleteProfileRelation({
          body: {
            p1SubType: this.item.subType,
            p1Id: this.item.id,
            p1Type: this.item.type,
            p2Id: this.rel.id,
            p2Type: this.rel.type,
            p2SubType: this.rel.subType
          }
        })
        this.rel = undefined
      } catch (e) {

      }

    },
    async selected(p: ProfileRelation) {
      try {
        await profileService.profileServiceAddProfileRelation({
          body: {
            p1SubType: p.subType,
            p1Id: p.id,
            p1Type: p.type,
            p2Id: p.rel.id,
            p2Type: p.rel.type,
            p2SubType: p.rel.subType
          }
        })

        this.rel = p.rel
      } catch (e) {

      }
    },
  }
})
</script>

<style>
</style>

