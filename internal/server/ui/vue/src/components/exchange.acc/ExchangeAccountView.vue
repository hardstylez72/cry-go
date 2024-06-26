<template>
  <v-container>
    <NavBar :title="`Биржевой аккаунт ${withdrawer.label}`"/>
    <Loader v-if="loading"/>
    <v-card v-else>
      <v-card-text>

        <v-form ref="formm">
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
          <v-row v-if="withdrawer.exchangeType === ExchangeType.Okex">
            <v-expansion-panels>
              <v-expansion-panel title="`Депозитные адреса">
                <v-expansion-panel-text>
                  <OkexWithdrawOptionSubAcc :id="withdrawer.id" :hide-base="true"/>
                </v-expansion-panel-text>
              </v-expansion-panel>
            </v-expansion-panels>
          </v-row>

          <v-row class="d-flex justify-center my-1">
            <v-col>
              <div>
                <v-btn width="100%" v-if="WithdrawerChanged" :loading="updating" @click="update">Update
                </v-btn>
              </div>
            </v-col>

          </v-row>
          <v-row class="d-flex justify-center my-2">
            <v-col>
              <div style="color: red" v-if="updatingError">{{ updatingError }}</div>
            </v-col>
          </v-row>
        </v-form>

        <v-expansion-panels v-if="withdrawer.exchangeType === ExchangeType.Okex" variant="accordion" multiple>
          <v-expansion-panel
            title="Добавить суб-аккаунт"
          >
            <v-expansion-panel-text>
              <CreateWithdrawerSubAcc :id="withdrawerId" @withdrawer-created="List"/>
            </v-expansion-panel-text>
          </v-expansion-panel>
          <v-expansion-panel v-for="sub in subaccs" :title="`Суб-аккаунт: ${sub.label}`">
            <v-expansion-panel-text>
              <b class="d-flex justify-center">Адреса депозитных суб-аккаунтов</b>
              <OkexWithdrawOptionSubAcc :id="sub.id"/>
            </v-expansion-panel-text>
          </v-expansion-panel>
        </v-expansion-panels>

      </v-card-text>
    </v-card>
  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {withdrawerService} from "@/generated/services"
import {
  CreateOkexWithdrawerRequest,
  ExchangeType, ExchangeWithdrawNetwork,
  ExchangeWithdrawOptions,
  Withdrawer
} from "@/generated/withdraw";
import ProxyInput from "@/components/ProxyInput.vue";
import OkexWithdrawOptionSubAcc from "@/components/exchange.acc/OkexWithdrawOptionSubAcc.vue";
import CreateWithdrawerSubAcc from "@/components/exchange.acc/CreateWithdrawerSubAcc.vue";
import {required} from "@/components/tasks/helper";
import ProfileSearch from "@/components/profile/ProfileSearch.vue";
import {Profile} from "@/generated/profile";
import {Timer} from "@/components/helper";
import NavBar from "@/components/NavBar.vue";
import Loader from "@/components/Loader.vue";

interface Network {
  network: string
  title: string
}

interface Token {
  token: string
  title: string
}

export default defineComponent({
  name: "ExchangeAccountView",
  computed: {
    ExchangeType() {
      return ExchangeType
    },
    withdrawerId: {
      set(id: string) {
        this.$router.push({query: {wId: id, ...this.$route.query}})
      },
      get(): string {
        const s = this.$route.query.wId
        if (s !== null && !Array.isArray(s)) {
          return s
        }
        return ''
      }
    },
    WithdrawerChanged() {
      if (this.withdrawer.proxy !== this.originalWithdrawer.proxy) {
        return true
      }
      if (this.withdrawer.label !== this.originalWithdrawer.label) {
        return true
      }
      return false
    },
    subaccs(): Withdrawer[] {
      return this.withdrawers
    }
  },
  components: {
    Loader,
    NavBar,
    ProfileSearch,
    CreateWithdrawerSubAcc,
    OkexWithdrawOptionSubAcc,
    ProxyInput
  },
  data() {
    return {
      loading: false,
      updatingError: "" as string | undefined,
      error: null as any,
      show: false,
      item: {} as CreateOkexWithdrawerRequest,
      saveLoading: false,
      listLoading: false,
      withdrawers: [] as Withdrawer[],
      withdrawerStatus: 'OK',
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
    required,
    async add() {
      this.show = true
      await this.List()
    },
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

    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['formm'].validate()

      return valid
    },

    async validateWithdrawForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
    async List() {
      try {
        this.listLoading = true
        const res = await withdrawerService.withdrawerServiceListSubWithdrawer({body: {withdrawerId: this.withdrawerId}})
        this.withdrawers = res.withdrawers
        this.$emit("change", false)
      } finally {
        this.listLoading = false
      }
    },
    Close() {
      this.show = false
    },
    async getWithdrawer() {
      const res = await withdrawerService.withdrawerServiceGetWithdrawer({body: {withdrawerId: this.withdrawerId}})
      this.withdrawer = res.withdrawer
      this.withdrawerStatus = res.error ? res.error : 'OK'
      this.originalWithdrawer = Object.assign({}, res.withdrawer)
    },
    async update() {

      try {
        this.updating = true
        this.updatingError = ''

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
        if (!this.updatingError) {
          await this.getWithdrawer()
        }

      } finally {
        this.updating = false
      }
    },
    GetTokenTitle(item: ExchangeWithdrawOptions): string {
      if (item) {
        return item.token + " - " + item.amount
      }
      return ""
    },
    GetNetworkTitle(item: ExchangeWithdrawNetwork): string {
      if (item) {
        return `${item.network} min:${item.minAmount} max:${item.maxAmount} fee:${item.fee}`
      }
      return ""
    },
  },
  async created() {
    try {
      this.loading = true
      await this.getWithdrawer()
      await this.List()
    } finally {
      this.loading = false
    }

  }
})


</script>


<style>


</style>

