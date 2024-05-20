<template>
  <Loader v-if="listLoading"/>
  <div v-else>
    <v-form ref="formm" v-if="!hideBase">
      <v-row>
        <v-col><b>Status</b>: {{ withdrawerStatus }}</v-col>
      </v-row>
      <v-row>
        <v-col cols="12">
          <v-text-field
            ref="label-input"
            v-model="withdrawer.label"
            label="label"
            density="comfortable"
            variant="outlined"
            :rules="labelRules()"
            @input="labelChanged"
          ></v-text-field>
        </v-col>
        <v-col cols="12">
          <ProxyInput v-model="withdrawer.proxy" :required="true"/>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-btn v-if="WithdrawerChanged" class="my-3" :loading="updating" @click="update">Update</v-btn>
          <div v-if="updatingError">{{ updatingError }}</div>
        </v-col>
      </v-row>
    </v-form>


    <b>EMV </b> <span v-if="emv.length"><b>Networks: </b>{{ emv[0].networks }}</span>
    <v-list max-width="96vw" class="px-5" :nav="true">
      <v-list-item
        density="compact"
        variant="plain"
        class="my-0"
        v-for="item in emv"
        :key="item.addr"
        rounded
        height="auto"
        width="auto"
        elevation="1"
        style="border: 0px solid "
      >

        <v-row align="center">

          <v-col>
            {{ `${item.addr}` }}
          </v-col>
          <v-col>
            <OkexWithdrawProfiles :model-value="item" :withdrawer-id="id"
                                  @updated="OkexWithdrawProfileUpdated"/>
          </v-col>

        </v-row>
      </v-list-item>
    </v-list>

    <b>StarkNet </b><span v-if="urgent.length"><b>Networks: </b>{{ urgent[0].networks }}</span>
    <v-list max-width="96vw" class="px-5" :nav="true">
      <v-list-item
        density="compact"
        variant="plain"
        class="my-0"
        v-for="(item, i) in urgent"
        :key="item.addr"
        rounded
        height="auto"
        width="auto"
        elevation="1"
        style="border: 0px solid "
      >

        <v-row align="center">

          <v-col>
            {{ `${item.addr}` }}
          </v-col>
          <v-col>
            <OkexDepositProfilesStarkNetUrgentX :model-value="urgent[i]" :withdrawer-id="id"
                                                @updated="OkexWithdrawProfileUpdated"/>
            <OkexDepositProfilesStarkNetBravos :model-value="braavos[i]" :withdrawer-id="id"
                                               @updated="OkexWithdrawProfileUpdated"/>
          </v-col>

        </v-row>
      </v-list-item>
    </v-list>

  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {helperService, profileService, withdrawerService} from "@/generated/services"
import {
  CreateOkexWithdrawerRequest,
  CreateWithdrawerRequest,
  DepositAddresses,
  ExchangeType,
  Withdrawer
} from "@/generated/withdraw";
import ProxyInput from "@/components/ProxyInput.vue";
import {Profile} from "@/generated/profile";
import {Timer} from "@/components/helper";
import OkexWithdrawProfiles from "@/components/exchange.acc/OkexDepositProfilesEMV.vue";
import Loader from "@/components/Loader.vue";
import OkexDepositProfilesStarkNetUrgentX from "@/components/exchange.acc/OkexDepositProfilesStarkNetUrgentX.vue";
import OkexDepositProfilesStarkNetBravos from "@/components/exchange.acc/OkexDepositProfilesStarkNetBravos.vue";
import {mapStores} from "pinia";
import {useSysStore} from "@/plugins/pinia";

export default defineComponent({
  name: "OkexWithdrawOptionSubAcc",
  computed: {
    WithdrawerChanged() {
      if (this.withdrawer.proxy !== this.originalWithdrawer.proxy) {
        return true
      }
      if (this.withdrawer.label !== this.originalWithdrawer.label) {
        return true
      }

      return false
    },
    ...mapStores(useSysStore),
  },
  components: {
    OkexDepositProfilesStarkNetBravos,
    OkexDepositProfilesStarkNetUrgentX,
    Loader,
    OkexWithdrawProfiles,
    ProxyInput
  },
  props: {
    id: {
      required: true,
      type: String
    },
    hideBase: {
      required: false,
      type: Boolean,
      default: false,
    }
  },
  data() {
    return {
      updatingError: "" as string | undefined,
      listLoading: false as boolean,
      emv: [] as DepositAddresses[],
      urgent: [] as DepositAddresses[],
      braavos: [] as DepositAddresses[],
      withdrawerStatus: "",
      withdrawer: {
        label: '',
        proxy: '',
        id: ''
      } as Withdrawer,
      originalWithdrawer: {
        label: '',
        proxy: '',
        id: ''
      } as Withdrawer,
      updating: false,
    }
  },
  methods: {
    labelRules() {
      return [
        (v: string) => this.withdrawer.label ? true : 'label is required'
      ]
    },
    async labelChanged(e: InputEvent) {
      await this.validate()
    },
    async validate(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['label-input'].validate()

      return valid
    },
    async getWithdrawer() {
      try {
        const res = await withdrawerService.withdrawerServiceGetWithdrawer({body: {withdrawerId: this.id}})
        this.withdrawer = res.withdrawer
        this.withdrawerStatus = res.error ? res.error : 'OK'
        this.originalWithdrawer = Object.assign({}, res.withdrawer)
      } catch (e) {
        this.sysStore.showSnackBar("ошибка при получении данных", true)
      }

    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['formm'].validate()

      return valid
    },
    async update() {

      try {
        this.updating = true
        this.updatingError = ""

        if (!await this.validateForm()) {
          this.updating = false
          return
        }

        const res = await withdrawerService.withdrawerServiceUpdateWithdrawer({
          body: {
            proxy: this.withdrawer.proxy,
            label: this.withdrawer.label,
            withdrawerId: this.withdrawer.id,
          }
        })

        this.updatingError = res.error

        await this.getWithdrawer()
      } finally {
        this.updating = false
      }
    },

    async OkexWithdrawProfileUpdated() {
      // await this.List()
    },
    async List() {
      try {
        this.listLoading = true
        const res = await withdrawerService.withdrawerServiceListDepositAddresses({body: {withdrawerId: this.id}})
        this.emv = res.emv.sort((a, b) => {
          const nameA = a.addr.toUpperCase(); // ignore upper and lowercase
          const nameB = b.addr.toUpperCase(); // ignore upper and lowercase
          if (nameA < nameB) {
            return -1;
          }
          if (nameA > nameB) {
            return 1;
          }

          // names must be equal
          return 0;
        })
        this.braavos = res.braavos.sort((a, b) => {
          const nameA = a.addr.toUpperCase(); // ignore upper and lowercase
          const nameB = b.addr.toUpperCase(); // ignore upper and lowercase
          if (nameA < nameB) {
            return -1;
          }
          if (nameA > nameB) {
            return 1;
          }

          // names must be equal
          return 0;
        })
        this.urgent = res.urgentx.sort((a, b) => {
          const nameA = a.addr.toUpperCase(); // ignore upper and lowercase
          const nameB = b.addr.toUpperCase(); // ignore upper and lowercase
          if (nameA < nameB) {
            return -1;
          }
          if (nameA > nameB) {
            return 1;
          }

          // names must be equal
          return 0;
        })
      } catch (e) {
        this.sysStore.showSnackBar("ошибка при получении списка адресов с okx", true)

      } finally {
        this.listLoading = false
      }
    },
  },
  async created() {
    this.List()
    this.getWithdrawer()
  }
})


</script>


<style>


</style>

