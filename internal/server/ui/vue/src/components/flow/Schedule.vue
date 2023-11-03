<template>
  <v-dialog
    v-model="dialog"
    width="1000"
    :close-on-content-click="false"
    :close-on-back="false"
    persistent="true"
  >
    <template v-slot:activator="{ props }">
      <v-btn width="100%" :disabled="disabled" @click="dialog=true">Запланировать запуск
      </v-btn>
    </template>

    <v-card>
      <v-card-title class="d-flex justify-space-between">
        <span>Расписание</span>
        <div>
          <v-btn @click="callBoba" :loading="loading" variant="outlined">Запланировать</v-btn>
          <v-icon icon="mdi-close" @click="dialog = false"/>
        </div>

      </v-card-title>
      <v-card-text>

        <div class="d-flex flex-wrap align-center">
          Запускать
          <div style="width: 100px">
            <v-text-field class="mx-1" hide-details type="number" density="compact"
                          v-model="DayCount" variant="outlined"/>
          </div>
          раз с интервалом в
          <div style="width: 100px">
            <v-text-field class="mx-1" hide-details type="number" density="compact"
                          v-model="interval" variant="outlined"/>
          </div>
          дней

          с отклонением менее
          <div style="width: 100px">
            <v-text-field class="mx-1" hide-details type="number" density="compact"
                          v-model="deltaInHours" variant="outlined"/>
          </div>
          часов
        </div>

        <v-list variant="plain" rounded>
          <v-list-item elevation="1" width="100%" v-for="(item, i) in runAfterSchedule" height="auto">
            <div class="d-flex  flex-wrap align-center">
              <div style="max-width: 200px" class="mx-1 my-2">
                <v-text-field :rules="[required]" type="date" v-model="item.date" label="Дата запуска" density="compact"
                              variant="outlined" hide-details/>
              </div>
              <div style="max-width: 220px; min-width: 140px" class="mx-1 my-2">
                <v-text-field :rules="[required]" type="time" v-model="item.time" label="Время запуска"
                              density="compact"
                              variant="outlined" hide-details/>
              </div>
              <v-icon v-if="runAfterSchedule.length -1 === i" icon="mdi-plus" color="green" @click="addRunAfterTime"/>
              <v-icon v-if="runAfterSchedule.length -1 === i && i !== 0" icon="mdi-close" color="red"
                      @click="removeRunAfterTime"/>
              <i class="mx-1">{{ formatTime(runAfterList[i]) }}</i>

            </div>


          </v-list-item>
        </v-list>
      </v-card-text>

    </v-card>

  </v-dialog>
</template>

<script lang="ts">
import {formatTime, getDate, getTime, isMobile, ts} from "@/components/helper";
import {required} from "@/components/tasks/menu/helper";


interface Item {
  date: any
  time: any
}

export default {
  name: "Schedule",
  emits: ['bobaSatisfied'],
  props: {
    disabled: {
      type: Boolean,
      required: false,
    }
  },
  watch: {
    interval: {
      handler() {
        this.generate()
      }
    },
    DayCount: {
      handler() {
        this.generate()
      }
    },
    deltaInHours: {
      handler() {
        this.generate()
      }
    },
    dialog: {
      handler() {
        if (this.dialog) {
          this.addRunAfterTime()
        }
      }
    },
    runAfterSchedule: {
      handler() {
        this.runAfterList = []
        this.runAfterSchedule.forEach(e => {
          this.runAfterList.push(ts(e.date, e.time))
        })
      },
      deep: true,
    }
  },
  data() {
    return {
      interval: 0,
      DayCount: 1,
      deltaInHours: 0,

      dialog: false,
      runAfterSchedule: [] as Item[],
      runAfterList: [] as Date[],
      runAfter: true,

      loading: false,
    }
  },
  methods: {
    callBoba() {
      this.loading = true
      this.$emit('bobaSatisfied', this.runAfterList)
    },
    generate() {
      this.runAfterSchedule = []


      for (let i = 0; i < this.DayCount; i++) {

        let delta = 0
        if (this.deltaInHours) {
          while (true) {
            const rand = Math.floor(Math.random() * 100)
            if (rand < this.deltaInHours * 60) {
              delta = rand
              break
            }
          }
        }

        this.runAfterSchedule.push(this.genDate(this.interval * i, delta))
      }
    },
    isMobile,
    required,
    formatTime,
    removeRunAfterTime() {
      this.runAfterSchedule.pop()
    },
    genDate(addDays: number, addMinutes: number): Item {
      const date = getDate(addDays)
      const time = getTime(addMinutes)
      return {date, time}
    },
    addRunAfterTime() {
      const date = getDate()
      const time = getTime()
      this.runAfterSchedule.push({date: date, time: time})
    },
  }
}
</script>

<style scoped>

</style>
