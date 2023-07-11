<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { ref } from 'vue'
import { computed } from 'vue'

import apiClient from '@/apis'
import { traPId } from '@/apis/generated'
import UserSelector from '@/components/UserSelector.vue'
import { useContestantsStore } from '@/store/contestants'

const { fetchContestants } = useContestantsStore()
const { contestants } = storeToRefs(useContestantsStore())

const targetContestant = ref<traPId>()

const tripleVoteFirst = ref<traPId>()
const tripleVoteSecond = ref<traPId>()
const tripleVoteThird = ref<traPId>()

const submitTargetContestant = () => {
  if (targetContestant.value === '') {
    return
  }

  apiClient.vote.postVote(targetContestant.value).then((res) => (targetContestant.value = res))
}

const submitVoteDisabled = computed(() => targetContestant.value === '')

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

apiClient.vote.getVote().then((res) => (targetContestant.value = res))
fetchContestants()
</script>

<template>
  <v-card>
    <v-card-title>1位予想</v-card-title>

    <v-card-text class="d-flex flex-column mt-4">
      <v-lazy :model-value="targetContestant !== undefined" v-if="targetContestant">
        <UserSelector :items="contestants" v-model:value="targetContestant" />
      </v-lazy>
    </v-card-text>

    <v-card-actions class="d-flex flex-column mb-4">
      <v-btn :onclick="submitTargetContestant" :disabled="submitVoteDisabled">Submit</v-btn>
    </v-card-actions>
  </v-card>

  <v-card class="mt-8">
    <v-card-title>三連単</v-card-title>
    <v-card-subtitle>この情報はダッシュボードには反映されません。</v-card-subtitle>

    <v-card-text>
      <div>
        <div>1位</div>
        <UserSelector :items="contestants" v-model:value="tripleVoteFirst" />
      </div>
      <div>
        <div>2位</div>
        <UserSelector :items="contestants" v-model:value="tripleVoteSecond" />
      </div>
      <div>
        <div>3位</div>
        <UserSelector :items="contestants" v-model:value="tripleVoteThird" />
      </div>
    </v-card-text>

    <v-card-actions class="d-flex flex-column mb-4">
      <v-btn :onclick="submitTripleVote" :disabled="tripleVoteSubmitDisable">Submit</v-btn>
    </v-card-actions>
  </v-card>
</template>
