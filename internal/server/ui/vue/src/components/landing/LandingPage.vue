<template>

  <v-container>
    <NavBar title="Главная">
      <template v-slot:default>
        <Login v-if="!userLoggedIn"/>
      </template>
    </NavBar>

    <h2 class="text-center text-h4 font-weight-bold text-blue-darken-2 my-5">
      Экономьте свое время с помощью автоматизированного исполнения airdrop-активностей
    </h2>

    <div>
      <b>Проблема</b>
      <p> В случае если вы занимаетесь дроп-хантингом, используя анти-детект браузер и команду людей, которые
        руками совершают активности
        - вы сталкиваетесь с проблемой человеческого фактора. Такой подход сложно масштабировать и контролировать.
        Помимо этого,
        активности в большей мере
        довольно монотонны и не каждый способен продолжительное время повторять однообразные действия.
      </p>
      <b>Решение</b>
      <p>
        Сервис поможет вам масштабировать число ваших аккаунтов и сэкономить время на рутинных задачах
        с помощью удобного рабочего места и автоматизации активностей.
      </p>
      <b>Принцип работы:</b>
      <ol class="ml-8">
        <li><a href="/profiles" target="_blank">Добавьте или сгенерируйте evm аккаунты</a></li>
        <li><a href="/constructor" target="_blank">Соберите сценарий исполнения активностей</a></li>
        <li><a href="https://t.me/drop_hunter_alert_bot" target="_blank">Зарегистрируйтесь в телеграм-боте</a></li>
        <li><a href="/constructor" target="_blank">Запустите сценарий для выбранных аккаунтов</a></li>
      </ol>
    </div>

    <v-expansion-panels class="mt-4" v-if="standalone">
      <v-expansion-panel
        key="var"
        title="Возможности"
      >
        <v-expansion-panel-text>
          <div>
            <ul>
              <li>Удобная генерация новых evm аккаунтов (1, 10, 100, 1000 штук и более)</li>
              <li>Возможность импортировать существующие emv аккаунты пачкой в 2 клика</li>
              <li>Возможность закрепить за каждым аккаунтов socks-5 proxy, который позволит скрыть ip, и выполнять все
                активности
                через proxy
              </li>
              <li>Интерактивный просмотр баланса в 1 клик</li>
              <li>Удобное пополнение баланса с биржи в 4 клика</li>
              <li>Интеграция с binance и okex</li>
              <li>Создание своего сценария активностей с помощью конструктора</li>
              <li>Запуск сценария для выбранных аккаунтов</li>
              <li>Возможность регулировать настройки газа. Ограничение на максимальный газ и урезание газа с помощью
                множителя
              </li>
              <li>Удобная работа с ошибками, возможность интерактивной работы с процессами.
                Можно останавливать, запускать, пропускать, и повторно выполнять шаги в которых возникла ошибка
              </li>
              <li>Можно получать уведомления с помощью бота в телеграме о ходе процессов</li>
              <li>Оценка стоимости транзакции в 1 клик как в метамаске</li>
            </ul>
          </div>
        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel
        key="airdrop"
        title="Интеграции"
      >
        <v-expansion-panel-text>
          <General/>
        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel title="О доверии">
        <v-expansion-panel-text>
          <iframe
            width="560"
            height="315"
            src="https://www.youtube.com/watch?v=UZFfwq_0tw0"
            srcdoc="<style>*{padding:0;margin:0;overflow:hidden}
    html,body{height:100%}
    img,span{position:absolute;width:100%;top:0;bottom:0;margin:auto}
    span{height:1.5em;text-align:center;font:48px/1.5 sans-serif;color:white;text-shadow:0 0 0.5em black}
    </style>
    <a href=https://www.youtube.com/embed/UZFfwq_0tw0?autoplay=1>
    <img src=https://img.youtube.com/vi/UZFfwq_0tw0/hqdefault.jpg alt='Demo video'>
    <span>▶</span>
    </a>"
            frameborder="0"
            allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture"
            allowfullscreen
            title="Demo video">
          </iframe>
        </v-expansion-panel-text>
      </v-expansion-panel>
    </v-expansion-panels>
    <v-expansion-panels v-else class="mt-4">
      <v-expansion-panel
        key="sec"
        title="Безопасность"
      >
        <v-expansion-panel-text>
          <div>
            <p>
              В крипте много скама, думаю каждый сталкивался с ним хоть раз. Выполнение on-chain активностей на сетях
              блокчейна невозможно без
              приватного ключа. Поэтому вопрос доверия встает играет ключевую роль.
            </p>
            <br/>

            Расскажу подробнее о drop-hunter:
            Проект разрабатывается, поддерживается и амнистируется одним человеком.
            <br/>
            <v-img src="/me.jpg" max-height="600px" class="mt-3"></v-img>
            <div class="text-center">Вот этот человек, это я, меня зовут Сергей, приятно познакомиться:)</div>
            <br/>
            Я являюсь профессиональным разработчиком WEB-приложений на протяжении последних 5 лет.
            <br/>
            Ссылки на мои соцсети:
            <a href="https://www.linkedin.com/in/sergey-k-317a80164/" target="_blank" class="d-inline-flex mx-1">
              <v-img height="20px" src="/icons/linkedin.png" @click=""></v-img>
              Linkedin
            </a>
            <a href="https://t.me/shinybozdo" target="_blank" class="d-inline-flex mx-1">
              <v-img height="20px" src="/icons/telegram.png" @click=""></v-img>
              Telegram
            </a>
            <a href="https://github.com/hardstylez72?tab=repositories" target="_blank" class="d-inline-flex mx-1">
              <v-img height="20px" src="/icons/github.png" @click=""></v-img>
              Github
            </a>

            <br/>
            Основной сервер приложения имеет открытый исходный код, который вы можете посмотреть по <a
            href="https://github.com/hardstylez72/cry" target="_blank" class="my-2">ссылке</a>

            <ul class="ml-4 mt-3">
              Какие меры предосторожности приняты:
              <li>чувствительные данные зашифрованы (слив бд не позволит получить ваши приватные ключи)</li>
              <li>чувствительные данные не логируются</li>
              <li>чувствительные данные доступны для пользователя только в момент создания/генерации</li>
            </ul>


          </div>
        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel
        key="var"
        title="Возможности"
      >
        <v-expansion-panel-text>
          <div>
            <ul>
              <li>Удобная генерация новых evm аккаунтов (1, 10, 100, 1000 штук и более)</li>
              <li>Возможность импортировать существующие emv аккаунты пачкой в 2 клика</li>
              <li>Возможность закрепить за каждым аккаунтов socks-5 proxy, который позволит скрыть ip, и выполнять все
                активности
                через proxy
              </li>
              <li>Интерактивный просмотр баланса в 1 клик</li>
              <li>Удобное пополнение баланса с биржи в 4 клика</li>
              <li>Интеграция с binance и okex</li>
              <li>Создание своего сценария активностей с помощью конструктора</li>
              <li>Запуск сценария для выбранных аккаунтов</li>
              <li>Возможность регулировать настройки газа. Ограничение на максимальный газ и урезание газа с помощью
                множителя
              </li>
              <li>Удобная работа с ошибками, возможность интерактивной работы с процессами.
                Можно останавливать, запускать, пропускать, и повторно выполнять шаги в которых возникла ошибка
              </li>
              <li>Можно получать уведомления с помощью бота в телеграме о ходе процессов</li>
              <li>Оценка стоимости транзакции в 1 клик как в метамаске</li>
            </ul>
          </div>
        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel
        key="money"
        title="Монетизация"
      >
        <v-expansion-panel-text>
          Каждому новому пользователю при регистрации будет создан внутренний аккаунт на котором будет доступно 10$.
          Вы не можете вывести эти деньги, их можно потратить на исполнение активностей.
          <br/>
          Есть бесплатные (вывод/депозит на биржу и задержка) активности и платные (все остальные).
          Платная активность имеет фиксированную цену в 0.2$.
          Списывание с баланса происходит при успешном выполнении каждой платной активности.
          Начального баланса хватит на прогон 50 активностей, что достаточно для ознакомления c работой сервиса.
          Баланс можно пополнить и также смотреть детализацию платы за активности.
        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel
        key="proc"
        title="Преимущества относительно конкурентов"
      >
        <v-expansion-panel-text>
          <ul>
            <li>Какие активности делать - выбор каждого. Сервис в этом смысле предоставляет полную свободу и
              прозрачность.
            </li>
            <li>Не нужно платить за возможность опробовать продукт</li>
            <li>Полный контроль над процессами (повторный запуск в случае ошибки, остановка, возобновление)</li>
            <li>Возможно использование на мобильном устройстве</li>
          </ul>

        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel
        key="airdrop"
        title="Интеграции"
      >
        <v-expansion-panel-text>
          <General/>
        </v-expansion-panel-text>
      </v-expansion-panel>
      <v-expansion-panel
        v-for="i in items"
        :key="i.title"
        :title="i.title"
        :text="i.details"
      />
    </v-expansion-panels>


  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import NavBar from "@/components/NavBar.vue";
import Login from "@/components/Login.vue";
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";
import General from "@/components/landing/modules/General.vue";

export default defineComponent({
  name: "LandingPage",
  components: {General, Login, NavBar},
  methods: {},
  computed: {
    standalone(): boolean {
      return import.meta.env.STANDALONE
    },
    userLoggedIn(): boolean {
      return this.userStore.login
    },
    ...mapStores(useUserStore),
  },
  data() {
    return {
      onboarding: 0,
      items: [
        {
          title: 'Поддерживаемые сети',
          details: 'Ethereum, Binance BNB, Optimism, Arbitrum, Polygon, ZkSync ERA, ZkSync lite, Avalanche, StarkNet'
        },
      ]
    }
  }
})


</script>


<style>
</style>

