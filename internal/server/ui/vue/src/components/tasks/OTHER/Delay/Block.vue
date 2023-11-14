<template>
  <v-card-actions>
    <br/>
    <v-container>
      <v-row>
        <v-checkbox
          :disabled="disabled"
          density="compact"
          focused
          name="randomize"
          v-model="random"
          label="randomize"
          @input="CheckboxChanged">
        </v-checkbox>
      </v-row>
      <v-row v-if="!random">


        <v-slider
          v-model="duration"
          density="compact"
          label="Фиксированная задержка"
          :min="1"
          step="1"
          :max="1800"
          persistent-hint
          center-affix
          thumb-label="always"
          hide-details
          :disabled="disabled"

        >
          <template v-slot:thumb-label="{ modelValue }">
            <div style="width: 120px" class="text-center">{{ humanDuration(modelValue) }}</div>
          </template>

        </v-slider>
      </v-row>

      <div v-if="random">
        <div>Рандомная задержка</div>
        <br/>
        <v-range-slider
          v-model="randomRange"
          density="compact"
          :min="1"
          step="10"
          :max="1800"
          persistent-hint
          label=""
          center-affix
          hide-details
          :disabled="disabled"
        >
          <template v-slot:prepend>
            <div style="width: 60px">{{ humanDuration(randomRange[0]) }}</div>
          </template>
          <template v-slot:append>
            <div style="width: 60px">{{ humanDuration(randomRange[1]) }}</div>
          </template>
        </v-range-slider>
      </div>
    </v-container>
  </v-card-actions>
</template>

<script lang="ts">
import {DelayTask, Task, TaskType} from "@/generated/flow";
import {defineComponent, PropType} from "vue";
import assert from 'assert'
import {humanDuration} from "@/components/helper";
import {required} from "@/components/tasks/helper";

export default defineComponent({
  name: "TaskDelay",
  emits: ['taskChanged'],
  props: {
    weight: {
      type: Number,
      required: true,
    },
    disabled: {
      type: Boolean,
      required: true,
    },
    task: {
      type: Object as PropType<Task>,
      required: false,
      deep: true,
    }
  },
  watch: {
    duration: {
      handler() {
        this.$emit('taskChanged', this.getTask())
      },
    },
    random: {
      handler() {
        this.$emit('taskChanged', this.getTask())
      },
    },
    randomRange: {
      handler() {
        this.$emit('taskChanged', this.getTask())
      },
      deep: true
    },
    task: {
      handler(b, a) {
        if (JSON.stringify(a) !== JSON.stringify(b)) {
          this.syncTask()
        }
      },
      deep: true
    }
  },

  data() {
    return {
      duration: "0",
      random: true,
      randomRange: [180, 300] as Array<Number>
    }
  },
  methods: {
    required,
    humanDuration,
    getTask(): Task {
      return {
        weight: this.weight.toString(),
        delayTask: {
          duration: this.duration ? this.duration : '0',
          random: this.random,
          minRandom: this.randomRange[0],
          maxRandom: this.randomRange[1],
        },
        taskType: TaskType.Delay,
        description: ""
      }
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['delay-form'].validate()

      return valid
    },
    CheckboxChanged() {

    },
    inputChanged() {
      this.duration = this.duration.toString().replace(/^0+/, '');

      return this.validateForm()
    },
    syncTask() {
      if (this.task) {
        if (this.task.delayTask) {
          const t = this.task.delayTask
          this.duration = t.duration
          this.random = t.random
          this.randomRange[0] = t.minRandom ? Number(t.minRandom) : Number(this.minRandom)
          this.randomRange[1] = t.maxRandom ? Number(t.maxRandom) : Number(this.maxRandom)
          this.$emit('taskChanged', this.getTask())
        }
      }
    }
  },
  created() {
    this.$emit('taskChanged', this.getTask())
    this.syncTask()
  }
})
</script>

<style scoped>

</style>
