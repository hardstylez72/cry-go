<template>
  <v-container>
    <NavBar title="Статистика по свапалкам">
    </NavBar>
    <Loader v-if="loading"/>
    <div v-else-if="!data.updated">
      <i>Статистика обновляется. Примерное время ожидания - 10 минут. Ваша позиция в очереди:
        {{ Math.floor(Math.random() * 1000) }}</i>
    </div>
    <div v-else>
      <div class="mx-2 my-2"><i>Обновлено: {{ formatTime(data.updated) }}</i></div>
      <div class="mx-2 my-5"><b>Отклонение курса в пулах Dapp от курса binance в процентах</b> <i>(чем выше число - тем
        больше
        теряется на
        операциях swap/bridge)</i>
      </div>

      <br/>

      <div class="d-flex align-center" style="width: 150px">
        <v-img height="22px" :src="networkProps.ZKSYNCERA.img" class="mx-4"/>
        <b>{{ networkProps.ZKSYNCERA.name }}</b>
      </div>
      <v-list>
        <v-list-item>
          <span class="d-flex justify-space-between">
           <b>Dapp</b>
          <b>ETH -> USDC</b>
          <b>USDC -> ETH</b>
          </span>

        </v-list-item>
        <v-list-item v-for="item in zksync" elevation="1" rounded>
          <span class="d-flex justify-space-between">
          <a class="d-flex align-center" style="width: 150px" :href="item.link" target="_blank">
                  <b>{{ item.name }}</b>
            <v-img height="22px" :src="item.img" class="mx-4"/>
          </a>
          <span>{{ item.eth_usdc }}%</span>
          <span>{{ item.usdc_eth }}%</span>
               </span>
        </v-list-item>
      </v-list>
      <br/>

      <div class="d-flex align-center" style="width: 150px">
        <v-img height="22px" :src="networkProps.StarkNet.img" class="mx-4"/>
        <b>{{ networkProps.StarkNet.name }}</b>
      </div>
      <v-list>
        <v-list-item>
          <span class="d-flex justify-space-between">
           <b>Dapp</b>
          <b>ETH -> USDC</b>
          <b>USDC -> ETH</b>
          </span>

        </v-list-item>
        <v-list-item v-for="item in starknet" elevation="1" rounded>
          <span class="d-flex justify-space-between">
          <a class="d-flex align-center" style="width: 150px" :href="item.link" target="_blank">
                  <b>{{ item.name }}</b>
            <v-img height="22px" :src="item.img" class="mx-4"/>
          </a>
          <span>{{ item.eth_usdc }}%</span>
          <span>{{ item.usdc_eth }}%</span>
               </span>
        </v-list-item>
      </v-list>
      <br/>
    </div>

    <div class="mx-2 my-5"><b>Как пользоваться:</b> <i>у вас есть желание конвертировать 1000 USDC в ETH и обратно.
      В если значение в таблице к примеру 0.5% в обе стороны - то на пуле при конвертации туда-обратно у вас без учета
      комиссии останется -
      1000 USDC * 2 * (0.5%/100%) = 990 USDC. получается пул забрал себе 10 USDC. Для объемов выбирайте наименьший
      процент расхождения</i></div>
    <div class="mx-2 my-5"><b>Как формируется статистика:</b> <i>
      <br/>
      1) симулируется транзакция с определенным проскальзыванием при которой в dapp транзакция проходит
      <br/>
      2) измеряется расхождение курса пула с курсом binance
      <br/>
      3) значения проскальзываня и разницей с бинансом суммируются пула по формуле:
      <b>Result = slippage{binance-pool} + slippage{pool} </b></i>
    </div>
  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import Loader from "@/components/Loader.vue";
import NavBar from "@/components/NavBar.vue";
import {formatTime, networkProps} from "@/components/helper";
import {publicService} from "@/generated/services";
import {SwapStatItem, SwapStatRes} from "@/generated/public";
import {TaskType} from "@/generated/flow";
import {taskProps} from "@/components/tasks/tasks";

type Dapp = {
  img: string
  name: string
  link: string
  eth_usdc: string
  usdc_eth: string
}

export default defineComponent({
  name: "Swaps",
  data() {
    return {
      data: {} as SwapStatRes,
      loading: false,
      zksync: [] as Dapp[],
      starknet: [] as Dapp[]
    }
  },
  computed: {
    networkProps() {
      return networkProps
    },
  },
  components: {NavBar, Loader},
  methods: {
    formatTime,
    async loadData() {
      try {
        this.loading = true
        this.data = await publicService.publicServiceSwapStat()
        this.buildZkSync()
        this.buildStarkNet()

      } catch (e) {

      } finally {
        this.loading = false
      }
    },
    buildZkSync() {
      this.zksync = []
      const zksync = new Map<TaskType, Dapp>()

      this.data.zkSyncETHUSDC.forEach(item => {
        zksync.set(item.type, {
          img: taskProps[item.type].service.img,
          name: taskProps[item.type].service.name,
          link: taskProps[item.type].service.link,
          eth_usdc: Number(item.sum).toPrecision(2).toString(),
          usdc_eth: ''
        })
      })

      this.data.zkSyncUSDCETH.forEach(item => {
        const v = zksync.get(item.type)
        if (v) {
          v.usdc_eth = Number(item.sum).toPrecision(2).toString()
          v.img = taskProps[item.type].service.img
          v.name = taskProps[item.type].service.name
          v.link = taskProps[item.type].service.link
          zksync.set(item.type, v)
        }
      })

      zksync.forEach(v => {
        this.zksync.push(v)
      })
    },
    buildStarkNet() {
      this.starknet = []
      const zksync = new Map<TaskType, Dapp>()

      this.data.starknetETHUSDC.forEach(item => {
        zksync.set(item.type, {
          img: taskProps[item.type].service.img,
          name: taskProps[item.type].service.name,
          link: taskProps[item.type].service.link,
          eth_usdc: Number(item.sum).toPrecision(2).toString(),
          usdc_eth: ''
        })
      })

      this.data.starknetUSDCETH.forEach(item => {
        const v = zksync.get(item.type)
        if (v) {
          v.usdc_eth = Number(item.sum).toPrecision(2).toString()
          v.img = taskProps[item.type].service.img
          v.name = taskProps[item.type].service.name
          v.link = taskProps[item.type].service.link
          zksync.set(item.type, v)
        }
      })

      zksync.forEach(v => {
        this.starknet.push(v)
      })
    }
  },
  created() {
    this.loadData()
  }
})


</script>


<style scoped>

</style>

