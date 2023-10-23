<template>
  <div class="mx-8 my-8">
    <div class="text-h3 text-center my-5">Base</div>
    <div>
      <v-table>
        <thead>
        <tr>
          <th class="text-left">Product</th>
          <th class="text-left">Actions</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="item in integratedZksync" :key="item.name">
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
import {taskProps} from "@/components/tasks/tasks";


export default defineComponent({
  name: "Base",
  created() {
  },
  data() {
    return {
      integratedZksync: [
        {
          name: taskProps.PancakeSwap.service.name,
          link: taskProps.PancakeSwap.service.link,
          img: taskProps.PancakeSwap.service.img,
          action: 'Swap ETH,USDC,USDC.b',
        },
        {
          name: taskProps.AcrossBridge.service.name,
          link: taskProps.AcrossBridge.service.link,
          img: taskProps.AcrossBridge.service.img,
          action: 'EVM мост (ввод/вывод ETH/USDC)',
        },
        {
          name: taskProps.OdosSwap.service.name,
          link: taskProps.OdosSwap.service.link,
          img: taskProps.OdosSwap.service.img,
          action: 'Swap ETH,USDC,USDC.b',
        }
      ],
    }
  }
})


</script>


<style>
</style>

