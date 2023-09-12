<template>
  <v-app class="font">
    <v-navigation-drawer
      location="left"
      v-model="drawer"
      :permanent="!isMobile"
    >
      <v-list-item v-if="userLoggedIn"
                   prepend-avatar="/favicon-32x32.png"
                   :title="userEmail"
                   nav
                   @click="goGeneral"
      >
        <div style="height: 20px; font-size: 14px" v-if="ass">Баланс: {{ ass }}</div>
        <template v-slot:append>
          <v-btn
            variant="text"
            icon="mdi-logout"
            @click="logout"
          ></v-btn>
        </template>
      </v-list-item>
      <v-list-item v-else class="mx-2 my-2" prepend-avatar="/favicon-32x32.png"
                   :title="userEmail"
                   nav
                   @click="goGeneral">
        <Login/>
      </v-list-item>

      <v-list-item>
        <v-switch @click="toggleTheme" :label="getTheme()">toggle theme</v-switch>
      </v-list-item>

      <v-divider/>

      <v-list density="compact" height="vh 90" nav="true">
        <v-list-item v-if="userLoggedIn" prepend-icon="mdi-account-multiple-outline" color="green" title="Профили"
                     value="home"
                     @click="$router.push({name: 'Profiles'})"></v-list-item>
        <v-list-item v-if="userLoggedIn" prepend-icon="mdi-account-cash-outline" color="green" title="Биржи"
                     value="Withdraw"
                     @click="$router.push({name: 'Withdraw'})"></v-list-item>
        <v-list-item v-if="userLoggedIn" prepend-icon="mdi-account-hard-hat-outline" color="green" title="Конструктор"
                     value="Constructor"
                     @click="$router.push({name: 'Constructor'})"></v-list-item>
        <v-list-item v-if="userLoggedIn" prepend-icon="mdi-blender-outline" color="green" title="Процессы"
                     value="Processes"
                     @click="$router.push({name: 'Processes'})"></v-list-item>
        <v-list-item v-if="userLoggedIn" prepend-icon="mdi-cog" title="Настройки" color="green" value="Настройки"
                     @click="$router.push({name: 'Settings'})"></v-list-item>
        <v-list-item v-if="userLoggedIn" prepend-icon="mdi-cash-multiple" title="Биллинг" color="green" value="Мани"
                     @click="$router.push({name: 'Billing'})"></v-list-item>

        <v-list-item prepend-icon="mdi-chart-bar" title="Статистика" color="green"
                     value="Статистика"
                     @click="$router.push({name: 'Stats'})"></v-list-item>

        <v-list-item v-if="userLoggedIn" prepend-icon="mdi-lifebuoy" title="Поддержка" color="green"
                     value="поддержка"
                     @click="$router.push({name: 'Issues'})"></v-list-item>

        <div class="h-auto py-4 pl-3">
          <div class="flex-column  justify-start">

            <a v-if="userLoggedIn" href="https://t.me/drop_hunter_alert_bot" target="_blank">
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


          <v-list-item class="d-flex pt-8" style="font-size: 14px" v-if="userLoggedIn">
            <div v-if="impact.top">
              <i>{{ `Статистика за 24 часа` }}
                <br/>
                <div>{{ ` Всего - ${impact.total} транзакций` }}</div>
                <div>{{ ` Лидер - ${impact.top} транзакций` }}</div>
                <div>{{ ` Я - ${impact.my} транзакций` }}</div>
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
import Support from "@/components/issue/Support.vue";
import {ThemeInstance} from 'vuetify'
import {myCustomLightTheme} from "@/plugins/vuetify";

export default defineComponent({
  name: 'App',
  components: {
    Support,
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
    toggleTheme() {
      const theme = this.getTheme() === 'dark' ? 'light' : 'dark'
      this.setTheme(theme)
    },
    goGeneral() {
      this.$router.push({name: 'LandingPage'})
    },
    logout() {
      document.cookie.split(";").forEach(function (c) {
        document.cookie = c.replace(/^ +/, "").replace(/=.*/, "=;expires=" + new Date().toUTCString() + ";path=/");
      });
      localStorage.clear();
      window.location.reload()
    },
    getTheme() {
      const t = window.localStorage.getItem('cry_theme')
      if (!t) {
        return this.$vuetify.theme.global.name
      }

      return t
    },
    setTheme(theme: string) {
      window.localStorage.setItem('cry_theme', theme)
      this.$vuetify.theme.global.name = theme
    }
  },

  async mounted() {
    this.$vuetify.theme.global.name = this.getTheme()
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


.onboarding {
  border: none; /* Убираем рамку */
  animation: flip 2s infinite;
}

@keyframes flip {
  0%, 100% {
    opacity: 1;
    transform: scale(1.1);
  }
  50% {
    opacity: 0.7;
    transform: scale(1);
  }
}
</style>
