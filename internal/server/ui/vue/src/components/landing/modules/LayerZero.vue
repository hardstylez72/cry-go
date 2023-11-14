<template>
  <div class="mx-8 my-8">
    <div class="text-h3 text-center my-5">LayerZero</div>
    <div>
      <v-table>
        <thead>
        <tr>
          <th class="text-left">Product</th>
          <th class="text-left">Actions</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="item in integrated" :key="item.name">
          <td><a :href="item.link" class="d-inline-flex">
            <v-img height="20px" :src="item.img" class="mr-2"/>
            {{ item.name }}
          </a></td>
          <td>{{ item.action }}</td>
        </tr>
        </tbody>
      </v-table>
      <div>


      </div>

    </div>

  </div>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import {taskProps} from "@/components/tasks/tasks";

enum Status {
  'integrated' = 'integrated',
  'soon' = 'soon',
  'not planned' = 'not planned',
  'no decision' = 'no decision'
}

export default defineComponent({
  name: "LayerZero",
  created() {
  },
  data() {
    return {
      integrated: [
        {
          name: 'Stargate.finance',
          link: taskProps.StargateBridge.service.link,
          img: taskProps.StargateBridge.service.img,
          status: Status.integrated,
          action: 'Swaps (ETH, STG, USDT, USDC)',
          quest: ''
        },
        {
          name: 'TestNetBridge',
          link: taskProps.TestNetBridgeSwap.service.link,
          img: taskProps.TestNetBridgeSwap.service.img,
          status: Status.integrated,
          action: 'Arbitrum, Optimism, Ethereum (ETH) -> Goerli (ETH)',
          quest: ''
        },
      ],
    }
  }
})


</script>


<style>
</style>

