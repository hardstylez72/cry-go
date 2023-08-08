<template>
  <v-container>
    <NavBar title="Настройки">
    </NavBar>
    <div>
      <v-expansion-panels class="mt-4" multiple="multiple">
        <SettingsNetwork v-for="n in networks" :network="n"/>
      </v-expansion-panels>
    </div>
  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {Network} from "@/generated/flow";
import SettingsNetwork from "@/components/settings/SettingsNetwork.vue";
import WEIInputField from "@/components/WEIInputField.vue";
import Loader from "@/components/Loader.vue";
import NavBar from "@/components/NavBar.vue";

export default defineComponent({
  name: "Settings",
  components: {NavBar, Loader, WEIInputField, SettingsNetwork},
  data() {
    return {
      networks: [
        Network.ZKSYNCERA,
        Network.OPTIMISM,
        Network.ARBITRUM,
        Network.AVALANCHE,
        Network.BinanaceBNB,
        Network.POLIGON,
        Network.Etherium,
        Network.ZKSYNCLITE
      ] as Network[],
    }
  },
})


</script>


<style scoped>

</style>

