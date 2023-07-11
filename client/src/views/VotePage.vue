<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { ref } from 'vue'
import { computed } from 'vue'

import apiClient from '@/apis'
import { traPId } from '@/apis/generated'
import UserIcon from '@/components/UserIcon.vue'
import UserSelector from '@/components/UserSelector.vue'
import { useContestantsStore } from '@/store/contestants'

const { fetchContestants } = useContestantsStore()
const { contestants } = storeToRefs(useContestantsStore())

const currentTargetContestant = ref<traPId>('')
const newTargetContestant = ref<traPId>('')

const tripleVoteFirst = ref<traPId>()
const tripleVoteSecond = ref<traPId>()
const tripleVoteThird = ref<traPId>()

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

const selectedFirst = (value: traPId) => {
  tripleVoteFirst.value = value
}
const selectedSecond = (value: traPId) => {
  tripleVoteSecond.value = value
}
const selectedThird = (value: traPId) => {
  tripleVoteThird.value = value
}
const submitTripleVote = () => {
  if (
    tripleVoteFirst.value === undefined ||
    tripleVoteSecond.value === undefined ||
    tripleVoteThird.value === undefined
  ) {
    return
  }
  apiClient.vote
    .postVoteTriple({
      first: tripleVoteFirst.value,
      second: tripleVoteSecond.value,
      third: tripleVoteThird.value
    })
    .then()
}
const tripleVoteSubmitDisable = computed(
  () =>
    tripleVoteFirst.value === undefined ||
    tripleVoteSecond.value === undefined ||
    tripleVoteThird.value === undefined
)

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

    <v-card-actions class="d-flex flex-column mb-4">
      <UserSelector :items="contestants" @selected="selected" class="w-75" />
      <v-btn :onclick="submitTargetContestant" :disabled="newTargetContestant === ''">Submit</v-btn>
    </v-card-actions>
  </v-card>

  <v-card class="mt-8">
    <v-card-title>三連単</v-card-title>
    <v-card-subtitle>この情報はダッシュボードには反映されません。</v-card-subtitle>

    <v-card-text>
      <div>
        <div>1位</div>
        <UserSelector :items="contestants" @selected="selectedFirst" />
      </div>
      <div>
        <div>2位</div>
        <UserSelector :items="contestants" @selected="selectedSecond" />
      </div>
      <div>
        <div>3位</div>
        <UserSelector :items="contestants" @selected="selectedThird" />
      </div>
    </v-card-text>

    <v-card-actions class="d-flex flex-column mb-4">
      <v-btn :onclick="submitTripleVote" :disabled="tripleVoteSubmitDisable">Submit</v-btn>
    </v-card-actions>
  </v-card>
</template>
