<template>
  <v-dialog
    v-model="dialog"
    width="800"
    :close-on-content-click="false"
    :close-on-back="false"
    :persistent="true"
  >
    <template v-slot:activator="{ props }">
      <v-btn density="comfortable" @click="dialog=true">Обменять промокод на 10$</v-btn>
    </template>

    <v-card>
      <v-card-title>Промокод</v-card-title>
      <v-card-text>
        <v-form ref="form" class="d-flex justify-space-between my-4">
          <v-text-field
            class="mr-3"
            type="text"
            v-model="promo"
            density="compact"
            variant="outlined"
            placeholder="введите промокод"
            :rules="[required]"
          />
          <v-btn rounded="false" @click="usePromo">Next</v-btn>
        </v-form>

      </v-card-text>
      <v-card-actions class="d-flex justify-end">
        <v-btn @click="dialog = false">Cancel</v-btn>
      </v-card-actions>
    </v-card>

  </v-dialog>
</template>


<script lang="ts">

import {defineComponent} from 'vue';
import {helperService} from "@/generated/services";
import {onlyInteger, required} from "@/components/tasks/helper";
import {mapStores} from "pinia";
import {useSysStore, useUserStore} from "@/plugins/pinia";

export default defineComponent({
  name: "Promo",
  created() {
  },
  computed: {
    ...mapStores(useSysStore),
    ...mapStores(useUserStore),
  },
  data() {
    return {
      promo: "",
      dialog: false,
      loading: false,
    }
  },
  methods: {
    required,
    async usePromo(): Promise<void> {
      if (!await this.validateForm()) {
        return
      }

      try {
        this.loading = true
        const res = await helperService.helperServiceUsePromo({body: {promo: this.promo}})
        if (!res.valid) {
          this.sysStore.showSnackBar("неверный промокод", true)
        } else {
          await this.userStore.syncUser()
          this.sysStore.showSnackBar(`вам начислено ${res.bonus} USD`, false)
          this.dialog = false
        }
      } catch (e) {
        this.sysStore.showSnackBar("неверный промокод", true)
      } finally {
        this.loading = false
      }
    },
    async validateForm(): Promise<boolean> {
      // @ts-ignore попизди мне еще что руки из жопы у меня ага
      // спасибо китайцам скажи лучше
      const {valid} = await this["$refs"]['form'].validate()

      return valid
    },
  }

})


</script>


<style>
</style>

