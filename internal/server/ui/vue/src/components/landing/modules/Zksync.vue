<template>
  <div class="mx-8 my-8">
    <div class="text-h3 text-center my-5">ZkSync</div>
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
  name: "Zksync",
  created() {
  },
  data() {
    return {
      integratedZksync: [
        {
          name: 'zk bridge',
          link: taskProps.ZkSyncOfficialBridgeFromEthereum.service.link,
          img: taskProps.ZkSyncOfficialBridgeFromEthereum.service.img,
          status: Status.integrated,
          action: 'Bridge ZkSync (ETH) <-> Ethereum (ETH)',
          quest: 'https://layer3.xyz/quests/welcome-to-a-new-era'
        },
        {
          name: 'Block.vue',
          link: taskProps.SyncSwap.service.link,
          img: taskProps.SyncSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> USDC, LUSD, LSD, MUTE). LP deposit/withdraw (ETH/USDC)',
          quest: 'https://layer3.xyz/quests/syncswap-on-zksync'
        },
        {
          name: 'MuteIO',
          link: taskProps.MuteioSwap.service.link,
          img: taskProps.MuteioSwap.service.img,
          status: Status.integrated,
          action: 'Swap ZkSync (ETH <-> USDC)',
          quest: ''
        },
        {
          name: 'Maverick',
          link: taskProps.MaverickSwap.service.link,
          img: taskProps.MaverickSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> MAV, LUSD, USDC)',
          quest: ''
        },
        {
          name: 'SpaceFi',
          link: taskProps.SpaceFISwap.service.link,
          img: taskProps.SpaceFISwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> SPACE, USDC)',
          quest: ''
        },
        {
          name: taskProps.VelocoreSwap.service.name,
          link: taskProps.VelocoreSwap.service.link,
          img: taskProps.VelocoreSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> VC, USDC)',
          quest: ''
        },
        {
          name: taskProps.IzumiSwap.service.name,
          link: taskProps.IzumiSwap.service.link,
          img: taskProps.IzumiSwap.service.img,
          status: Status.integrated,
          action: 'Swap (WETH/ETH <-> USDC/IZI)',
          quest: ''
        },
        {
          name: taskProps.VeSyncSwap.service.name,
          link: taskProps.VeSyncSwap.service.link,
          img: taskProps.VeSyncSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> USDC)',
          quest: ''
        },
        {
          name: taskProps.EzkaliburSwap.service.name,
          link: taskProps.EzkaliburSwap.service.link,
          img: taskProps.EzkaliburSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> USDC)',
          quest: ''
        },
        {
          name: taskProps.ZkSwap.service.name,
          link: taskProps.ZkSwap.service.link,
          img: taskProps.ZkSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> USDC)',
          quest: ''
        },
        {
          name: taskProps.MerklyMintAndBridgeNFT.service.name,
          link: taskProps.MerklyMintAndBridgeNFT.service.link,
          img: taskProps.MerklyMintAndBridgeNFT.service.img,
          status: Status.integrated,
          action: 'Mint and Bridge NFT',
          quest: ''
        },
        {
          name: taskProps.PancakeSwap.service.name,
          link: taskProps.PancakeSwap.service.link,
          img: taskProps.PancakeSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> USDC)',
          quest: ''
        },
        {
          name: taskProps.OdosSwap.service.name,
          link: taskProps.OdosSwap.service.link,
          img: taskProps.OdosSwap.service.img,
          status: Status.integrated,
          action: 'Swap (ETH <-> USDC, USDC <-> USDT)',
          quest: ''
        },
        {
          name: taskProps.AcrossBridge.service.name,
          link: taskProps.AcrossBridge.service.link,
          img: taskProps.AcrossBridge.service.img,
          status: Status.integrated,
          action: 'EVM мост (ввод/вывод ETH/USDC)',
          quest: ''
        }
      ],
      links: [
        {
          name: 'explorer',
          link: 'https://explorer.zksync.io/',
          desc: 'default explorer'
        },
        {
          name: 'zksync activity stats',
          link: 'https://byfishh.github.io/zk-flow/',
          desc: 'shows platform activity stats'
        },

      ]
    }
  }
})


</script>


<style>
</style>

