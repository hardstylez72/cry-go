<template>
  <div class="mx-8 my-8">
    <div class="text-h3 text-center my-5">Other</div>
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
import {link} from "@/components/tasks/links";
import {taskProps, taskTypes} from "@/components/tasks/tasks";

enum Status {
  'integrated' = 'integrated',
  'soon' = 'soon',
  'not planned' = 'not planned',
  'no decision' = 'no decision'
}

export default defineComponent({
  name: "Other",
  created() {
  },
  data() {
    return {
      integrated: [
        {
          name: taskProps.OrbiterBridge.service.name,
          link: taskProps.OrbiterBridge.service.link,
          status: Status.integrated,
          action: 'Bridge (ARBI, BNB, MATIC, AVAX, ETH, ZkSync, ZkSyncLite) ETH',
          img: taskProps.OrbiterBridge.service.img
        },
        {
          name: taskProps.TraderJoeSwap.service.name,
          link: taskProps.TraderJoeSwap.service.link,
          status: Status.integrated,
          action: 'DEX swap arbitrum (STG, USDT, USDC -> ETH)',
          img: taskProps.TraderJoeSwap.service.img,
        },
        {
          name: taskProps.SnapshotVote.service.name,
          link: taskProps.SnapshotVote.service.link,
          status: Status.integrated,
          action: 'follow and vote automatic on any guild',
          img: taskProps.SnapshotVote.service.img,
        },
        {
          name: taskProps.Dmail.service.name,
          link: taskProps.Dmail.service.link,
          status: Status.integrated,
          action: 'send emails',
          img: taskProps.Dmail.service.img,
        },
      ],
    }
  }
})


</script>


<style>
</style>

