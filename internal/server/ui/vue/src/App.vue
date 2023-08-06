<template>
  <v-app class="font">

    <v-navigation-drawer
      v-if="userLoggedIn"
      location="left"
      v-model="drawer"
      :permanent="!isMobile"
      expand-on-hover
    >
      <v-list-item
        prepend-avatar="/favicon-32x32.png"
        :title="userEmail"
        nav
        @click="goGeneral"
      >
        <div style="height: 20px" v-if="ass">Balance: {{ ass }}</div>
        <template v-slot:append>
          <v-btn
            variant="text"
            icon="mdi-logout"
            @click="logout"
          ></v-btn>
        </template>
      </v-list-item>

      <v-divider/>

      <v-list density="compact" nav height="vh 90">
        <v-list-item prepend-icon="mdi-account-multiple-outline" color="green" title="Профили" value="home"
                     @click="$router.push({name: 'Profiles'})"></v-list-item>
        <v-list-item prepend-icon="mdi-account-cash-outline" color="green" title="Биржи" value="Withdraw"
                     @click="$router.push({name: 'Withdraw'})"></v-list-item>
        <v-list-item prepend-icon="mdi-account-hard-hat-outline" color="green" title="Конструктор" value="Constructor"
                     @click="$router.push({name: 'Constructor'})"></v-list-item>
        <v-list-item prepend-icon="mdi-blender-outline" color="green" title="Процесы" value="Processes"
                     @click="$router.push({name: 'Processes'})"></v-list-item>
        <v-list-item prepend-icon="mdi-cog" title="Настройки" color="green" value="Настройки"
                     @click="$router.push({name: 'Settings'})"></v-list-item>
        <v-list-item prepend-icon="mdi-cash-multiple" title="Биллинг" color="green" value="Мани"
                     @click="$router.push({name: 'Billing'})"></v-list-item>


        <div class="h-auto py-4 pl-3">
          <div class="flex-column  justify-start">

            <a href="https://t.me/drop_hunter_alert_bot" target="_blank">
              <div style="height: 20px" class="d-inline-flex">
                <v-img src="/icons/telegram.png"/>
                <span class="ml-2">Бот</span>
              </div>
            </a>
            <br/>
            <a href="https://t.me/+WX8iCLaJBelhNTVi" target="_blank">
              <div style="height: 20px" class="d-inline-flex">
                <v-img src="/icons/telegram.png"/>
                <span class="ml-2">Сообщество</span>
              </div>
            </a>
          </div>


          <v-list-item class="d-flex pt-8" style="font-size: 14px">
            <div v-if="impact.top">
              <i>{{ `24 Hour statistics` }}
                <br/>
                <div>{{ ` total - ${impact.total} tx` }}</div>
                <div>{{ ` top - ${impact.top} tx` }}</div>
                <div>{{ ` me - ${impact.my} tx` }}</div>
              </i>
            </div>
          </v-list-item>
        </div>

      </v-list>

    </v-navigation-drawer>

    <v-main class="mt-7">
      <notifications/>
      <router-view/>
      <Snackbar/>
    </v-main>

  </v-app>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import SideBar from "@/components/SideBar.vue";
import Login from "@/components/Login.vue";
import {helperService} from "@/generated/services";
import {mapActions, mapStores} from "pinia";
import {useUserStore, useSysStore} from "@/plugins/pinia";
import Snackbar from "@/components/Snackbar.vue";
import NavBar from "@/components/NavBar.vue";

export default defineComponent({
  name: 'App',
  components: {
    NavBar,
    Snackbar,
    SideBar,
    Login,
  },
  data() {

  },
  computed: {
    drawer: {
      get() {
        return this.userStore.drawer
      },
      set(v: boolean) {
        this.userStore.$patch((s) => {
          s.drawer = v
        })
      },
    },
    ...mapStores(useUserStore),
    impact() {
      return this.userStore.impact
    },
    userLoggedIn(): boolean {
      return this.userStore.login
    },
    userEmail(): string {
      return this.userStore.email
    },
    isMobile(): boolean {
      return this.userStore.isMobile
    },
    ass(): string {
      return this.userStore.ass
    },
    ...mapActions(useUserStore, ['syncUser', "syncDailyImpact"])
  },
  methods: {
    goGeneral() {
      this.$router.push({name: 'LandingPage'})
    },
    logout() {
      document.cookie.split(";").forEach(function (c) {
        document.cookie = c.replace(/^ +/, "").replace(/=.*/, "=;expires=" + new Date().toUTCString() + ";path=/");
      });
      localStorage.clear();
      window.location.reload()
    }
  },
  async mounted() {
    const period = import.meta.env.DEV ? 60_000 : 5000

    await this.userStore.syncUser()
    await this.userStore.syncDailyImpact()

    setInterval(async () => {
      await this.userStore.syncUser()
      await this.userStore.syncDailyImpact()
    }, period)
  }
})


</script>

<style>

@font-face {
  font-family: "Space_Mono";
  src: local("Space_Mono"),
  url(~@/assets/fonts/Space_Mono/SpaceMono-Bold.ttf) format("truetype");
}

.font {
  /*font-family: "Space_Mono", monospace;*/
  /*font-size: 16px;*/
}

</style>
