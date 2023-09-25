<template>
  <v-container>
    <NavBar title="Статистика">
    </NavBar>

    <v-list>
      <v-list-item
        variant="flat"
        density="comfortable"
        to="/stats/swap"
        rounded
        title="Свапалки и мосты (Swap/Bridge)"
        elevation="1"
      />

      <v-list-item
        variant="flat"
        density="comfortable"
        to="/stats/zksync"
        rounded
        title="ZkSync"
        elevation="1"
      />

    </v-list>
  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import Loader from "@/components/Loader.vue";
import NavBar from "@/components/NavBar.vue";
import {networkProps} from "@/components/helper";

export default defineComponent({
  name: "Stats",
  computed: {
    networkProps() {
      return networkProps
    }
  },
  components: {NavBar, Loader},
})


</script>


<style scoped>

</style>

