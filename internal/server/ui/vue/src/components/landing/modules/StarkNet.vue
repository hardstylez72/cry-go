<template>
  <div class="mx-8 my-8">
    <div class="text-h3 text-center my-5">StarkNet</div>
    <div>
      <v-table>
        <thead>
        <tr>
          <th class="text-left">Product</th>
          <th class="text-left">Status</th>
          <th class="text-left">Actions</th>
          <th class="text-left">Quest</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="item in integratedZksync" :key="item.name">
          <td><a :href="item.link" class="d-inline-flex">
            <v-img height="20px" :src="item.img" class="mr-2"/>
            {{ item.name }}
          </a></td>
          <td>{{ item.status }}</td>
          <td>{{ item.action }}</td>
          <td>{{ item.quest }}</td>
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
import {taskProps} from "@/components/tasks/tasks";

enum Status {
  'integrated' = 'integrated',
  'soon' = 'soon',
  'not planned' = 'not planned',
  'no decision' = 'no decision'
}

export default defineComponent({
  name: "StarkNet",
  created() {
  },
  data() {
    return {
      integratedZksync: [
        {
          name: taskProps.DeployStarkNetAccount.service.name,
          link: taskProps.DeployStarkNetAccount.service.link,
          img: taskProps.DeployStarkNetAccount.service.img,
          status: Status.integrated,
          action: 'Deploys UrgentX account',
          quest: ''
        },
        {
          name: taskProps.Swap10k.service.name,
          link: taskProps.Swap10k.service.link,
          img: taskProps.Swap10k.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> USDC)',
          quest: ''
        },
        {
          name: taskProps.SithSwap.service.name,
          link: taskProps.SithSwap.service.link,
          img: taskProps.SithSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> USDC)',
          quest: ''
        },
        {
          name: taskProps.JediSwap.service.name,
          link: taskProps.JediSwap.service.link,
          img: taskProps.JediSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> USDC)',
          quest: ''
        }
      ],
    }
  }
})


</script>


<style>
</style>

