<template>
  <div>
    <v-chip v-if="deposit.profileId" closable close-icon="mdi-close"
            @click:close="remove(deposit.profileId, deposit.addr)">
      {{ `${deposit.profileLabel} (${deposit.subType})` }}
    </v-chip>
    <span v-else @click="searchProfiles">
    <v-autocomplete
      return-object
      v-model="selectedProfile"
      :loading="profilesLoading"
      :items="suggestedProfiles"
      :item-title="profileTitle"
      item-value="id"
      label="Выберите EMV профиль"
      density="compact"
      no-data-text="нет свободных EMV профилей"
    />
  </span>
  </div>
</template>

<script lang="ts">

import {defineComponent, PropType} from 'vue';
import {profileService, withdrawerService} from "@/generated/services"
import {DepositAddresses,} from "@/generated/withdraw";
import {Profile, ProfileSubType} from "@/generated/profile";
import {profileTitle, Timer} from "@/components/helper";
import {mapStores} from "pinia";
import {useSysStore} from "@/plugins/pinia";

export default defineComponent({
  name: "OkexDepositProfilesEMV",
  emits: ['update:modelValue', 'updated'],
  computed: {
    ...mapStores(useSysStore),
  },
  props: {
    modelValue: {
      required: true,
      type: Object as PropType<DepositAddresses>
    },
    withdrawerId: {
      required: true,
      type: String
    }
  },
  watch: {
    deposit: {
      handler() {
        this.$emit('update:modelValue', this.deposit)
      },
      deep: true
    },
    selectedProfile: {
      handler(p: Profile) {
        this.attach(p)
      }
    }
  },
  data() {
    return {
      deposit: {} as DepositAddresses,
      selectedProfile: null as Profile | null,
      suggestedProfile: "",
      suggestedProfiles: [] as Profile[],
      profilesLoading: false,
      timer: new Timer()
    }
  },
  methods: {
    profileTitle,
    async remove(profileId: string | undefined, addr: string) {
      try {
        await withdrawerService.withdrawerServiceOkexDepositAddrDetach({
          body: {
            profileId: profileId ? profileId : '',
            okexDepositAddr: addr,
            withdrawerId: this.withdrawerId,
            subType: ProfileSubType.Metamask
          }
        })
        this.deposit.profileId = undefined
        this.deposit.profileLabel = undefined
        this.sysStore.showSnackBar("профиль отвязан от адреса депозита", false)
      } catch (e) {
        this.sysStore.showSnackBar("ошибка при отвязке профиля от адреса депозита", true)
      }
    },
    async attach(p: Profile) {
      if (p && p.id) {
        try {
          await withdrawerService.withdrawerServiceOkexDepositAddrAttach({
            body: {
              profileId: p.id,
              okexDepositAddr: this.deposit.addr,
              withdrawerId: this.withdrawerId,
              subType: ProfileSubType.Metamask,
            }
          })

          this.deposit.profileId = p.id
          this.deposit.subType = ProfileSubType.Metamask
          this.deposit.profileLabel = `${p.num}`

          this.$emit('updated')
          this.sysStore.showSnackBar("профиль привязан к адресу депозита", false)
        } catch (e) {
          this.sysStore.showSnackBar("ошибка при привязке профиля к адресу депозита", true)
        }
      }
    },
    async searchProfiles() {
      this.timer.add(0)
      this.timer.cb(async () => {
        try {
          this.profilesLoading = true
          const res = await profileService.profileServiceSearchProfilesNotConnectedToOkexDeposit({body: {subType: ProfileSubType.Metamask}})

          this.suggestedProfiles = res.profiles
        } catch (e) {
          this.sysStore.showSnackBar("ошибка при поиске профилей", true)
        } finally {
          this.profilesLoading = false
        }
      })
    },
  },
  async created() {
    Object.assign(this.deposit, this.modelValue)
  }
})


</script>


<style>


</style>

