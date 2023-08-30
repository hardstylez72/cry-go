<template>
  <div class="mx-8 my-8">
    <div class="text-h3 text-center my-5">Other</div>
    <div>
      <v-table>
        <thead>
        <tr>
          <th class="text-left">Product</th>
          <th class="text-left">Status</th>
          <th class="text-left">Actions</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="item in integrated" :key="item.name">
          <td><a :href="item.link" class="d-inline-flex">
            <v-img height="20px" :src="item.img" class="mr-2"/>
            {{ item.name }}
          </a></td>
          <td>{{ item.status }}</td>
          <td>{{ item.action }}</td>
        </tr>
        </tbody>
      </v-table>
      <div>
        <v-expansion-panels>
          <v-expansion-panel
            title="Мнения"
          >
            <v-expansion-panel-text>
              <div>
                Возможет дроп, токен не вышел
                <a href="https://galxe.com/OrbiterFinance/campaign/GCbnmUNe9g" target="_blank">nft и критерии</a>

                <div>
                  Почему Orbiter Finance может оказаться «темной лошадкой-гемом»

                  Как мы знаем, Orbiter finance привлекли инвестиции от различных фондов, и сумма инвестиций не
                  разглашается (очень странно почему?)

                  Одним из инвесторов стал фонд Tiger Global. Ну и что нам с него? Давайте попробуем разобраться.

                  Этот же фонд возглавил раунд в 298 млн$ в 2021 году для TradingView

                  На сегодняшний день TradingView самая популярная в мире платформа-соц.сеть для трейдеров с 30
                  миллионами посещаемости в месяц, и с оценкой в 3 млрд$ (данные 2021г.)

                  Денис Глоба (Дада, вы не ослышались, CEO самой популярной в мире компании для трейдеров является
                  простым парнем-ITшником из Ростова-на-Дону) генеральный директор и соучредитель TradingView, написал:
                  "Мы рады сотрудничать с Tiger Global. Их глобальный финтех-опыт и идеи внесут значительный вклад в
                  развитие видения TradingView информированной финансовой торговли и инвестирования для всего мира»

                  Так же, Tiger global является инвестором таких всемирно известных гигантов, как: Revolut (раунд на 800
                  млн$ от 2021г.), Coinbase (раунд на 300 млн$ от 2018г.)

                  А еще, вы не поверите, Tiger Global старейший акционер таких компания, как Яндекс и Qiwi

                  Из этого понятно, что Tiger Global не инвестирует в мелкие проекты, и тот момент, что данные по
                  инвестициям Orbiter Finance являются закрытыми - только наводит на мысль, что перед потенциальным
                  аирдропом эти данные раскроют, ну а мы просто продолжаем минтить NFT, что бы потом сказать, что
                  «каждый мог»
                </div>
              </div>
            </v-expansion-panel-text>
          </v-expansion-panel>
        </v-expansion-panels>


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

