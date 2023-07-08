<script setup lang="ts">
import { ref } from 'vue'

import apiClient from '@/apis'
import { traPId } from '@/apis/generated'
import TrapIdListInputForm from '@/components/TrapIdListInputForm.vue'

const contestants = ref<traPId[] | undefined>()

apiClient.user.getContestants().then((res) => (contestants.value = res))

const submit = (ids: traPId[]) => {
  apiClient.adminOnly.putContestants(ids)
}
</script>

<template>
  <v-card>
    <v-card-title> 競技者 </v-card-title>

    <v-card-text v-if="contestants">
      <TrapIdListInputForm :initial-list="contestants" :submit-func="submit" />
    </v-card-text>
  </v-card>
</template>
