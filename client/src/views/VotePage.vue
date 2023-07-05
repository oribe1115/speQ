<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { ref } from 'vue'

import apiClient from '@/apis'
import { traPId } from '@/apis/generated'
import UserSelector from '@/components/UserSelector.vue'
import { useContestantsStore } from '@/store/contestants'

const { fetchContestants } = useContestantsStore()
const { contestants } = storeToRefs(useContestantsStore())

const targetContestant = ref<traPId>('')

const selected = (value: traPId) => {
  targetContestant.value = value
}
const submitTargetContestant = () => {
  if (targetContestant.value === '') {
    return
  }

  apiClient.vote.postVote(targetContestant.value).then()
}

fetchContestants()
</script>

<template>
  <p>vote page</p>
  <UserSelector :items="contestants" @selected="selected" />

  <v-btn :onclick="submitTargetContestant" :disabled="targetContestant === ''">Submit</v-btn>
</template>
