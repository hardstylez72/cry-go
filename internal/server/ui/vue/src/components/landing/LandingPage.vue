<template>

  <v-container>
    <NavBar title="Главная">
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
    </div>

    <div class="my-5">
      <h2 class="text-center text-h5 font-weight-bold text-blue-darken-2 my-5">Интеграции</h2>
      <TaskSelector @taskSelected="addStep"/>

      <v-card v-if="tasks">
        <v-card-title class="d-flex justify-space-between px-0 align-center">
          <div class="d-inline-flex align-center">
            <a><h4 :id="tasks.weight" class="mx-2">{{ `${tasks.taskType}` }}</h4></a>
            <a target="_blank" :href="taskSpec(tasks).service.link" class="mx-1">
              <v-img height="22px" :src="taskSpec(tasks).service.img"/>
            </a>
          </div>
        </v-card-title>
        <component

          style="box-shadow: rgba(0, 0, 0, 0.16) 0 3px 6px, rgba(0, 0, 0, 0.23) 0 3px 6px;"
          :is="tasks.component"
          :weight="tasks.weight"
          :task="tasks.task"
          :spec="taskSpec(tasks)"
        />
      </v-card>
    </div>

    <div class="my-5">
      <h2 class="text-center text-h5 font-weight-bold text-blue-darken-2 my-5">Безопасность</h2>
      <p>
        В крипте много скама, думаю каждый сталкивался с ним хоть раз. Выполнение on-chain активностей на сетях
        блокчейна невозможно без
        приватного ключа. Поэтому вопрос доверия встает играет ключевую роль.
      </p>
      <br/>

      Расскажу подробнее о drop-hunter:
      Проект разрабатывается, поддерживается и амнистируется одним человеком.
      <br/>
      Я являюсь профессиональным разработчиком WEB-приложений на протяжении последних 5 лет.
      <br/>

      <ul class="ml-4 mt-3">
        Какие меры предосторожности приняты:
        <li>чувствительные данные зашифрованы (слив бд не позволит получить ваши приватные ключи)</li>
        <li>чувствительные данные не логируются</li>
        <li>чувствительные данные доступны для пользователя только в момент создания/генерации</li>
      </ul>
    </div>


    <div class="my-5">
      <h2 class="text-center text-h5 font-weight-bold text-blue-darken-2 my-5">Возможности</h2>
      <ul>
        <li>Удобная генерация новых evm/StarkNet аккаунтов (1, 10, 100, 1000 штук и более)</li>
        <li>Возможность импортировать существующие аккаунты пачкой в 2 клика</li>
        <li>Возможность закрепить за каждым аккаунтов socks-5 proxy, который позволит скрыть ip, и выполнять все
          активности
          через proxy, или не указывать прокси - тогда будет использован внутренний rotating-proxy (каждые 5 минут
          меняет ip)
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
        <li>Возможность смотреть статистику по обьемам/транзакциям</li>
        <li>Ограничение проскальзывания, ограничения проскальзывания по курсу бинанса</li>
      </ul>
    </div>

  </v-container>
</template>

<script lang="ts">

import {defineComponent} from 'vue';
import NavBar from "@/components/NavBar.vue";
import Login from "@/components/Login.vue";
import {mapStores} from "pinia";
import {useUserStore} from "@/plugins/pinia";
import General from "@/components/landing/modules/General.vue";
import TaskSelector from "@/components/flow/TaskSelector.vue";
import {TaskType} from "@/generated/flow";
import {TaskArg, taskComponentMap, taskProps} from "@/components/tasks/tasks";
import {TaskSpec} from "@/components/tasks/utils";

export default defineComponent({
  name: "LandingPage",
  components: {TaskSelector, General, Login, NavBar},
  methods: {
    async addStep(task: TaskType) {
      this.tasks = {
        component: taskComponentMap.get(task),
        weight: 1,
        taskType: task
      }
    },
    taskSpec(item: TaskArg): TaskSpec {
      return taskProps[item.taskType]
    },
  },
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
      tasks: null,
    }
  }
})


</script>


<style>
</style>

