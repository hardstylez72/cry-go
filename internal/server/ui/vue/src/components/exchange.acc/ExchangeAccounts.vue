<template>
  <div>
    <NavBar title="Биржи">
      <template v-slot:default class="d-flex justify-end ma-3">
        <ExchangeAccountDocs class="mr-3"/>
        <v-btn v-if="selectedSome" color="black" variant="flat" @click=DeleteWithdrawers class="mr-3">Удалить
        </v-btn>
        <v-btn v-else @click=CreateWithdrawer variant="flat" :class="noAccounts ? 'onboarding mx-2 my-2' : `mx-2 my-2`">
          Добавить
        </v-btn>
      </template>
    </NavBar>
    <Loader v-if="loading"/>
    <div v-else>


      <div v-if="noAccounts" class="my-8 mx-8 text-center">
        Для удобного пополнения/депозита "кошельков" привязанных к профилям можно добавить аккаунты бирж okex/binance.
        Это займет больше времени чем создание профилей, потому что сначала необходимо создать API ключи на биржах и
        настроить их.
        <a href="/docs/binance_account_instruction" target="_blank">инструкция для binance</a>,
        <a href="/docs/okex_account_instruction" target="_blank">инструкция для okex</a>

      </div>

      <v-list v-else max-width="96vw" class="px-5" nav="true">
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
          style="border: 0px solid "
        >

          <v-row align="center">
            <v-col cols="1">
              <CheckBox style="height: 30px" density="compact" :id="item.id" focused
                        @CheckboxChanged="CheckboxChanged"/>
            </v-col>

            <v-col>
              <b v-if="item.label">{{ item.label }}</b>
            </v-col>
            <v-col>
              <b>{{ item.exchangeType }}</b>
            </v-col>
            <v-col>
              <BtnCheckProxy v-if="item.proxy" :proxy="item.proxy"/>
            </v-col>
            <v-col>
              <v-btn density="compact" @click="viewAccount(item.id)">View</v-btn>
            </v-col>

          </v-row>
        </v-list-item>
      </v-list>

      <CreateWithdraw :showProp="showCreateWithdrawerDialog" @change="CreateWithdrawerChange"
                      @profileCreated="UpdateList"/>
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {withdrawerService} from "@/generated/services"
import CreateWithdraw from "@/components/exchange.acc/CreateWithdraw.vue";
import CheckBox from "@/components/CheckBox.vue";
import {ExchangeType, Withdrawer} from "@/generated/withdraw";
import OkexWithdrawOption from "@/components/exchange.acc/ExchangeAccountView.vue";
import ExchangeAccountDocs from "@/components/exchange.acc/ExchangeAccountDocs.vue";
import BtnCheckProxy from "@/components/BtnCheckProxy.vue";
import ExportExchangeAccounts from "@/components/exchange.acc/ExportExchangeAccounts.vue";
import Loader from "@/components/Loader.vue";
import NavBar from "@/components/NavBar.vue";

export default defineComponent({
  name: "Withdraw",
  components: {
    NavBar,
    Loader,
    ExportExchangeAccounts,
    BtnCheckProxy,
    ExchangeAccountDocs,
    OkexWithdrawOption,
    CheckBox,
    CreateWithdraw,
  },
  data() {
    return {
      list: [] as Withdrawer[],
      showCreateWithdrawerDialog: false,
      selected: new Set<string>(),
      loading: true,
      loadingError: false,
    }
  },
  computed: {
    noAccounts() {
      return this.getList.length === 0
    },
    ExchangeType() {
      return ExchangeType
    },
    getList(): Withdrawer[] {
      return this.list
    },
    selectedSome(): boolean {
      return this.selected.size > 0
    }
  },
  methods: {
    viewAccount(id: string) {
      this.$router.push({name: 'ExchangeAccountView', query: {wId: id}})
    },
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
        const res = await withdrawerService.withdrawerServiceListWithdrawer()
        this.list = res.withdrawers
        this.selected = new Set<string>()
      } catch (e) {
        this.loadingError = true
      } finally {
        this.loading = false
      }

    },
    CreateWithdrawer() {
      this.showCreateWithdrawerDialog = true
    },
    CreateWithdrawerChange(b: boolean) {
      this.showCreateWithdrawerDialog = b
    },
    async DeleteWithdrawers() {
      for (let id of this.selected.keys()) {
        await withdrawerService.withdrawerServiceDeleteWithdrawer({body: {id: id}})
      }
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

