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
  <v-card>
    <v-card-title>1位予想</v-card-title>

    <v-card-text>
      <div>
        <div>現在の1位予想</div>
        <div v-if="currentTargetContestant !== ''">
          <UserIcon :trap-id="currentTargetContestant" />
          {{ currentTargetContestant }}
        </div>
        <div v-else>未設定</div>
      </div>
    </v-card-text>

    <v-card-actions class="d-flex flex-column">
      <UserSelector :items="contestants" @selected="selected" class="w-75" />
      <v-btn :onclick="submitTargetContestant" :disabled="newTargetContestant === ''">Submit</v-btn>
    </v-card-actions>
  </v-card>
</template>
