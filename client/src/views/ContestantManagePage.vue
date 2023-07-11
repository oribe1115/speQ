<script setup lang="ts">
import { ref } from 'vue'

import apiClient from '@/apis'
import { traPId } from '@/apis/generated'
import TrapIdListInputForm from '@/components/TrapIdListInputForm.vue'

const contestants = ref<traPId[] | undefined>()
const showSuccessToast = ref<boolean>(false)
const timeout = ref(3000)

apiClient.user.getContestants().then((res) => {
  if (res === null) {
    res = []
  }
  contestants.value = res
})

const submit = (ids: traPId[]) => {
  apiClient.adminOnly.putContestants(ids).then(() => (showSuccessToast.value = true))
}
</script>

<template>
  <v-card>
    <v-card-title> 競技者 </v-card-title>

    <v-card-text>
      <v-lazy :model-value="contestants !== undefined" v-if="contestants">
        <TrapIdListInputForm :initial-list="contestants" :submit-func="submit" />
      </v-lazy>
    </v-card-text>
  </v-card>

  <v-snackbar v-model="showSuccessToast" :timeout="timeout" color="green"> Success! </v-snackbar>
</template>
