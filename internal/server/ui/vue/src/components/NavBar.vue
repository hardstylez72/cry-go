<template>
  <v-app-bar :elevation="2" absolute="true">
    <template v-slot:prepend>
      <v-icon class="mx-2 my-2 " size="30px" style="z-index: 9999"
              v-if="isMobile && userLoggedIn && !drawer"
              @click="clickMenuOnMobile"
              icon="mdi-menu"></v-icon>
    </template>

    <template v-slot:append>
      <slot name="default"/>
    </template>
    <template v-slot:title>
      {{ title }}
    </template>


  </v-app-bar>

</template>

<script lang="ts">
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";

export default {
  name: "NavBar",
  props: {
    title: {
      type: String,
      required: true
    }
  },
  computed: {
    userLoggedIn(): boolean {
      return this.userStore.login
    },
    ...mapStores(useUserStore),
    drawer: {
      get() {
        return this.userStore.drawer
      },
      set(v: boolean) {
        this.userStore.$patch((s) => {
          s.drawer = v
        })
      }
    },
    isMobile() {
      if (/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {
        return true
      } else {
        return false
      }
    },
  },
  data() {
    return {}
  },
  methods: {
    clickMenuOnMobile() {
      this.drawer = true
    },
  },
  created() {
    if (this.isMobile) {
      this.drawer = false
    } else {
      this.drawer = true
    }
  }
}
</script>

<style scoped>

</style>
