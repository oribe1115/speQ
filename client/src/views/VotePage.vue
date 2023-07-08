<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { ref } from 'vue'

import apiClient from '@/apis'
import { traPId } from '@/apis/generated'
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
  <p>vote page</p>
  <p>現在の一位予想</p>
  <p>{{ currentTargetContestant }}</p>
  <UserSelector :items="contestants" @selected="selected" />

  <v-btn :onclick="submitTargetContestant" :disabled="newTargetContestant === ''">Submit</v-btn>
</template>
