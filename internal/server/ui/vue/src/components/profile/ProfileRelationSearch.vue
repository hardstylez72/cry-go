<template>
  <div class="my-2" @focusin="click">
    <v-autocomplete
      chips="true"
      closable-chips
      return-object
      v-model="selectedProfiles"
      :loading="profilesLoading"
      @update:search="searchProfiles"
      :items="suggestedProfiles"
      item-title="num"
      item-value="id"
      :multiple="false"
      ref="label-input"
      label="Поиск профилей"
      density="compact"
      variant="outlined"
      clearable="true"
      :disabled="disabled"
      no-data-text="Петушки не нашлись"
      hide-no-data="true"
      hide-details="true"
    >
    </v-autocomplete>
  </div>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {profileService} from "@/generated/services"
import {
  Profile, ProfileRelation, ProfileSubType,
  ProfileType,
  SearchProfileRelationReq,
  SearchProfileRelationRes,
  UnrelatedProfile
} from "@/generated/profile";
import {shuffleArray, Timer} from "@/components/helper";
import {required} from "@/components/tasks/helper";

export default defineComponent({
  name: "ProfileRelationSearch",
  emits: ['selected'],
  watch: {
    selectProfileType: {
      handler() {
        this.selectedProfiles = []
        this.searchProfiles()
      }
    }
  },
  props: {
    num: {
      type: String,
      required: true,
    },
    pId: {
      type: String,
      required: true,
    },
    p1SubType: {
      type: String as PropType<ProfileSubType>,
      required: true,
    },
    p1Type: {
      type: String as PropType<ProfileType>,
      required: true,
    },
    p2SubType: {
      type: String as PropType<ProfileSubType>,
      required: true,
    },
    p2Type: {
      type: String as PropType<ProfileType>,
      required: true,
    },
    disabled: {
      type: Boolean,
      required: false
    },
  },
  data() {
    return {
      selected: null,
      profilesLoading: false,
      suggestedProfile: "",
      suggestedProfiles: [] as UnrelatedProfile[],
      timer: new Timer(),
      hideNoData: true,
    }
  },
  computed: {
    ProfileType() {
      return ProfileType
    },
    selectedProfiles: {
      set(a: UnrelatedProfile) {

        const tmp: ProfileRelation = {
          id: this.pId,
          type: this.p1Type,
          subType: this.p1SubType,
          num: this.num,
          rel: {
            id: a.id,
            type: a.type,
            subType: a.subType,
            num: a.num,
          }
        }

        this.$emit('selected', tmp)
      },
      get(): Profile {
        return this.selected
      }
    }
  },
  methods: {
    async click() {
      const res = await profileService.profileServiceSearchProfileRelation({
        body: {
          p1SubType: this.p2SubType,
          p1Type: this.p2Type,
          p2SubType: this.p1SubType,
          p2Type: this.p1Type,
        }
      })
      this.suggestedProfiles = res.items
    },
    required,
    async searchProfiles(v: string) {
      this.timer.add(100)
      this.timer.cb(async () => {
        this.hideNoData = false
        if (v === "" || v === null || v === undefined) {
          v = ''
        }
        try {
          this.profilesLoading = true
          const res = await profileService.profileServiceSearchProfileRelation({
            body: {
              p1SubType: this.p2SubType,
              p1Type: this.p2Type,
              p2SubType: this.p1SubType,
              p2Type: this.p1Type,
            }
          })
          this.suggestedProfiles = []
          this.suggestedProfiles.push(...res.items)
        } finally {
          this.profilesLoading = false
        }
      })
    },
  },
})


</script>


<style scoped>
</style>

