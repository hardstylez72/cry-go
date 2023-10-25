<template>
  <div class="my-2">
    <v-autocomplete
      chips="true"
      closable-chips
      return-object
      v-model="selectedProfiles"
      :loading="profilesLoading"
      @update:search="searchProfiles"
      :items="suggestedProfiles"
      :item-title="profileTitle"
      item-value="id"
      :multiple="multiple"
      :rules="[required]"
      ref="label-input"
      label="Нужно выбрать профили, Например 1-10"
      density="comfortable"
      variant="outlined"
      clearable="true"
      :disabled="disabled"
      no-data-text="Петушки не нашлись"
      :hide-no-data="hideNoData"
    >
      <template v-slot:chip="{ props, item }">
        <v-chip
          v-bind="props"
          :text="item.value.num"
        ></v-chip>
      </template>
    </v-autocomplete>
    <div class="d-inline-flex align-center">
      <v-radio-group v-model="selectProfileType" inline="" hide-details density="compact" :disabled="disabled">
        Выберите тип профилей:
        <v-radio v-if="profileTypeAllowed(ProfileType.EVM)" class="mx-8" :value="ProfileType.EVM">EVM</v-radio>
        <v-radio v-if="profileTypeAllowed(ProfileType.StarkNet)" class="mx-8" :value="ProfileType.StarkNet">StarkNet
        </v-radio>
      </v-radio-group>
      <v-switch class="mx-2" v-if="shufflable && selectedProfiles.length" v-model="shuffle" label="перетасовать"
                density="compact" hide-details/>
    </div>

  </div>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {profileService} from "@/generated/services"
import {Profile, ProfileType} from "@/generated/profile";
import {profileTitle, shuffleArray, Timer} from "@/components/helper";
import {required} from "@/components/tasks/menu/helper";

export default defineComponent({
  name: "ProfileSearch",
  emits: ['update:modelValue'],
  watch: {
    selectProfileType: {
      handler() {
        this.selectedProfiles = []
        this.searchProfiles()
      }
    },
    shuffle: {
      handler() {
        if (this.shuffle) {
          shuffleArray(this.selectedProfiles)
        } else {
          this.selectedProfiles.sort((a, b) => a.num - b.num)
        }
      }
    }
  },
  props: {
    shufflable: {
      type: Boolean,
      required: false,
      default: true
    },
    profileTypeFilter: {
      type: Array as PropType<ProfileType[]>,
      required: false,
      default: [ProfileType.EVM, ProfileType.StarkNet]
    },
    multiple: {
      type: Boolean,
      required: false,
      default: true
    },
    disabled: {
      type: Boolean,
      required: false
    },
    modelValue: {
      type: Array as PropType<Profile[]>,
      required: true
    }
  },
  data() {
    return {
      profilesLoading: false,
      suggestedProfile: "",
      suggestedProfiles: [] as Profile[],
      suggestedProfilesOrig: [] as Profile[],
      timer: new Timer(),
      hideNoData: false,
      selectProfileType: ProfileType.EVM,
      shuffle: false,
    }
  },
  computed: {
    ProfileType() {
      return ProfileType
    },
    selectedProfiles: {
      set(a: Profile[]) {
        this.$emit('update:modelValue', a)
      },
      get(): Profile[] {
        return this.modelValue
      }
    }
  },
  methods: {
    profileTypeAllowed(t: ProfileType): boolean {
      if (!this.profileTypeFilter) {
        return true
      }

      return this.profileTypeFilter.some(p => p === t)
    },
    profileTitle,
    required,
    async searchProfiles(v: string) {
      this.timer.add(100)
      this.timer.cb(async () => {
        this.hideNoData = false
        if (v === "" || v === null || v === undefined) {
          v = ''
        }
        if (v.split("-").length > 1) {
          try {
            this.hideNoData = true
            this.profilesLoading = true
            const res = await profileService.profileServiceSearchProfile({
              body: {
                pattern: v,
                type: this.selectProfileType
              }
            })

            this.suggestedProfiles = []
            this.suggestedProfiles.push(...res.profiles)

            while (this.selectedProfiles.pop()) {

            }
            this.selectedProfiles.push(...this.suggestedProfiles)

            this.suggestedProfilesOrig.push(...this.selectedProfiles)

          } finally {
            this.profilesLoading = false
          }
        } else {
          try {
            this.profilesLoading = true
            const res = await profileService.profileServiceSearchProfile({
              body: {
                pattern: v,
                type: this.selectProfileType
              }
            })
            this.suggestedProfiles = []
            this.suggestedProfiles.push(...res.profiles)
            this.suggestedProfilesOrig.push(...this.selectedProfiles)
          } finally {
            this.profilesLoading = false
          }
        }
      })
    },
  },
  async mounted() {
    const res = await profileService.profileServiceSearchProfile({body: {pattern: "", type: this.selectProfileType}})
    this.suggestedProfiles = res.profiles
  }
})


</script>


<style scoped>
</style>

