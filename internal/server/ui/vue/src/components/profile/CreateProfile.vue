<template>
  <div class="mx-8 my-8" style="min-height: 500px">
    <v-container>
      <NavBar title="Добавление профилей"/>
      <div class="pb-5 font-italic text-pre-wrap text-break">
        <p>
          При создании профиля рекомендуется указывать socks5 proxy, если прокси указан -
          вся активность по профилю будет совершаться через прокси.
          Это повышает анонимность и уменьшает риск стать сибилом.
        </p>
        <br/>
        <div>Форма для ввода ниже позволяет импортировать профили в формате CSV.</div>
        <br/>
        <div> каждая строка должна быть вида:
          "приватный ключ (обязательно)", "socks5 proxy (опционально)", "название (опционально)"
        </div>
      </div>

      <v-btn :loading="loading" @click="validate" class="mb-4 mx-4">
        <v-tooltip
          activator="parent"
          location="bottom"
          width="300px"
          text="Валидирует форму, полезна для проверки работоспособности прокси и корректности ввода приватных ключей"
        />
        Валидация
      </v-btn>
      <v-btn v-if="!demo" :loading="loading" @click="save" :class="needSave ? 'onboarding mb-4 mx-4': 'mb-4 mx-4'">
        <v-tooltip
          activator="parent"
          location="bottom"
          width="300px"
          text="Последовательное добавление профилей из формы ввода на сервис"
        />
        Добавить
      </v-btn>

      <GenerationProfiles @generated="generated" :profile-sub-type="selectProfileSubType" :demo="demo"
                          :loading="loading" :profile-type="selectProfileType"/>
      <div class="mr-9 d-flex justify-start">
        <v-radio-group
          class="my-2"
          v-model="selectProfileType"
          :inline="true"
          hide-details
          density="compact"
          :disabled="needSave">
          Выберите тип профилей:
          <v-radio class="mx-8" :value="ProfileType.EVM">EVM</v-radio>
          <v-radio class="mx-8" :value="ProfileType.StarkNet">StarkNet</v-radio>
        </v-radio-group>
      </div>
      <v-radio-group
        class="my-2"
        v-model="selectProfileSubType"
        :inline="true"
        hide-details
        density="compact"
        :disabled="needSave">
        Выберите тип аккаунта:
        <v-radio v-for="item in profileSubTypes" style="margin-left: 50px; margin-right: 50px" :value="item">{{
            item
          }}
        </v-radio>
      </v-radio-group>

      <div>
        <v-list>
          <v-list-item
            elevation="1"
            v-for="item in items"
          >
            <div>
              <span class="mx-1" :style="`color: ${getStatusColor(item)}`">{{ `[ status: ${item.status} ]` }}</span>
              <span class="mx-1">PK: {{ item.pk }}</span>
              <span class="mx-1">Proxy: {{ item.proxy || 'Не задан' }}</span>
              <span class="mx-1">Label: {{ item.label || 'Не задан' }}</span>

            </div>


          </v-list-item>
        </v-list>
      </div>

      <v-textarea
        auto-grow
        @input="preparseText"
        density="compact"
        :clearable="true"
        variant="outlined"
        placeholder="<metamask-pk> required, <proxy> optional, <custom label> optional] ..."
        v-model="text"
        :error-messages="errorMessage"
        :disabled="loading"
      />
      <div v-if="valid">Number of profiles: <b>{{ items.length }}</b></div>
      <br/>

    </v-container>
  </div>

</template>

<script lang="ts">

import {defineComponent} from 'vue';
//@ts-ignore
import Papa from 'papaparse';
import {helperService, profileService} from "@/generated/services";
import GenerationProfiles from "@/components/profile/GenerationProfiles.vue";
import NavBar from "@/components/NavBar.vue";
import {ProfileSubType, ProfileType} from "@/generated/profile";


interface Item {
  pk: string
  proxy?: string
  label?: string

  status?: string
}

export default defineComponent({
  watch: {
    selectProfileType: {
      handler() {
        if (this.selectProfileType === ProfileType.EVM) {
          this.profileSubTypes = [ProfileSubType.Metamask]
        }
        if (this.selectProfileType === ProfileType.StarkNet) {
          this.profileSubTypes = [ProfileSubType.UrgentX, ProfileSubType.Braavos]
        }

        this.selectProfileSubType = this.profileSubTypes[0]
      }
    }
  },
  computed: {
    ProfileSubType() {
      return ProfileSubType
    },
    ProfileType() {
      return ProfileType
    }
  },
  props: {
    demo: {
      required: false,
      type: Boolean,
      default: false
    }
  },
  components: {NavBar, GenerationProfiles},
  name: "CreateProfile",
  data() {
    return {
      needSave: false,
      loading: false,
      mode: '',
      dialog: false,
      items: [] as Item[],
      validError: '',
      valid: false,
      generateCount: 10,
      errorMessage: '',
      text: '',
      selectProfileType: ProfileType.EVM,
      selectProfileSubType: ProfileSubType.Metamask,
      profileSubTypes: [ProfileSubType.Metamask] as ProfileSubType[]
    }
  },
  methods: {
    generated(text: string) {
      this.needSave = true
      this.text = text
      this.preparseText()
    },
    getStatusColor(item: Item): string {
      if (item.status === 'ok') {
        return 'green'
      }
      if (item.status === 'validating') {
        return 'blue'
      }
      if (item.status === 'saving') {
        return 'blue'
      }
      if (item.status === 'ready') {
        return 'grey'
      }
      return 'red'
    },
    print(data: string[][]) {
      this.text = this.format(data)
    },

    preparseText() {
      try {
        this.errorMessage = ''
        this.items = []
        const data = this.preparse(this.text)
        data.forEach((line) => {
          this.items.push({
            pk: line[0],
            proxy: line[1] ? line[1] : '',
            label: line[2] ? line[2] : '',
            status: 'ready'
          })
        })
        this.print(data)
      } catch (e) {
        this.errorMessage = e.message
      }

    },
    async save() {
      this.needSave = false
      this.mode = 'validate and save'

      this.loading = true
      for (let i = 0; i < this.items.length; i++) {
        let item = this.items[i]
        try {
          item.status = 'saving'
          const res = await profileService.profileServiceCreateProfile({
            body: {
              proxy: item.proxy,
              mmskPk: item.pk,
              label: item.label ? item.label : '',
              type: this.selectProfileType,
              subType: this.selectProfileSubType
            }
          })
          item.status = 'ok'
        } catch (e) {
          item.status = 'error saving: ' + e
        }
      }
      this.loading = false
    },
    async validate() {
      this.loading = true
      this.mode = 'validate only'

      for (let i = 0; i < this.items.length; i++) {
        let item = this.items[i]
        this.items[i] = await this.validateItem(item)
      }

      this.loading = false
    },
    async validateItem(item: Item): Promise<Item> {
      item.status = 'validating'

      if (item.proxy) {
        try {
          const stat = await helperService.helperServiceValidateProxy({body: {proxy: item.proxy}})
          if (!stat.valid) {
            item.status = stat.errorMessage
            return item
          } else {
          }
        } catch (e) {
          item.status = 'proxy error'
          return item
        }
      }

      try {
        const pkStat = await helperService.helperServiceValidatePk({
          body: {
            pk: item.pk,
            type: this.selectProfileType,
            subType: this.selectProfileSubType,
          }
        })
        if (!pkStat.valid) {
          item.status = 'pk invalid'
          return item
        }
      } catch (e) {
        item.status = 'pk invalid'
        return item
      }

      item.status = 'ok'

      return item
    },
    preparse(text: string): string[][] {
      text = text.replaceAll(" ", "")
      this.errorMessage = ''
      const result = Papa.parse<string[]>(text, {skipEmptyLines: true, delimitersToGuess: [';', ',']})
      result.data.forEach((line: string[], i) => {
        try {
          if (line.length >= 1) {
            this.simpleValidatePK(line[0])
          }

          if (line.length >= 2) {
            this.simpleValidateProxy(line[1])
          }
          if (line.length !== 3) {
            throw new Error(`Строка должна иметь вид: [pk,proxy,label]. В качестве разделителя можно использовать: ',' или ';'. \n
          Пример(metamask): 4374d77b231f6ed16f2fc23a6e28706615f852bdaf223685475f7337487ae51d,,`)
          }
        } catch (e) {
          throw new Error(`Ошибка в строке ${i + 1}. ` + e.message)
        }

      })
      return result.data
    },
    simpleValidateProxy(proxy: string) {
      proxy = proxy.replaceAll(" ", "")
      if (proxy === '') {
        return
      }

      const parts = proxy.split(':')
      if (parts.length != 4) {
        throw new Error(`Proxy должен иметь вид [ip:port:login:password]. Пример: 123.456.67.89:5436:user1:password1`)
      }
    },
    simpleValidatePK(pk: string) {
      if (this.selectProfileType === ProfileType.EVM) {
        const res = pk.search(/(^[a-f0-9]{64}$)/g)
        if (res === -1) {
          throw new Error(`Приватный ключ должен состоять из 64 символов. Пример: 69efc3b091b176542c7d7eee1bbf7321244a5c83d026d693a7288e072e663b81`)
        }
      }

    },
    format(data: string[][]): string {
      let formatted = ""
      data.forEach((line: string[]) => {
        formatted += line.join(";")
        formatted += '\r'
      })
      return formatted
    }
  }
})


</script>


<style>


</style>

