<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { ref } from 'vue'

import apiClient from '@/apis'
import { traPId } from '@/apis/generated'
import UserIcon from '@/components/UserIcon.vue'
import UserSelector from '@/components/UserSelector.vue'
import { useContestantsStore } from '@/store/contestants'

const { fetchContestants } = useContestantsStore()
const { contestants } = storeToRefs(useContestantsStore())

const currentTargetContestant = ref<traPId>('')
const newTargetContestant = ref<traPId>('')

const selected = (value: traPId) => {
  newTargetContestant.value = value
}
const submitTargetContestant = () => {
  if (newTargetContestant.value === '') {
    return
  }

  apiClient.vote
    .postVote(newTargetContestant.value)
    .then((res) => (currentTargetContestant.value = res))
}

apiClient.vote.getVote().then((res) => (currentTargetContestant.value = res))
fetchContestants()
</script>

<template>
  <v-card width="100%">
    <template v-slot:title> 現在の一位予想 </template>

    <v-card-text v-if="currentTargetContestant !== ''">
      <UserIcon :trap-id="currentTargetContestant" />
      {{ currentTargetContestant }}
    </v-card-text>

    <v-card-text v-else> 未設定 </v-card-text>
  </v-card>

  <UserSelector :items="contestants" @selected="selected" />
  <v-btn :onclick="submitTargetContestant" :disabled="newTargetContestant === ''">Submit</v-btn>
</template>
