<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { ref } from 'vue'
import { computed } from 'vue'

import apiClient from '@/apis'
import { TripleVote, traPId } from '@/apis/generated'
import UserSelector from '@/components/UserSelector.vue'
import { useContestantsStore } from '@/store/contestants'

const { fetchContestants } = useContestantsStore()
const { contestants } = storeToRefs(useContestantsStore())

const targetContestant = ref<traPId>()
const tripleVote = ref<TripleVote>()

const showSuccessToast = ref<boolean>(false)
const timeout = ref(3000)

const submitTargetContestant = () => {
  if (targetContestant.value === '') {
    return
  }

  apiClient.vote.postVote(targetContestant.value).then((res) => {
    showSuccessToast.value = true
    targetContestant.value = res
  })
}

const submitVoteDisabled = computed(() => targetContestant.value === '')

const submitTripleVote = () => {
  apiClient.vote.postVoteTriple(tripleVote.value).then(() => {
    showSuccessToast.value = true
  })
}
const tripleVoteSubmitDisable = computed(
  () =>
    tripleVote.value?.first === '' ||
    tripleVote.value?.second === '' ||
    tripleVote.value?.third === ''
)

apiClient.vote.getVote().then((res) => (targetContestant.value = res))
apiClient.vote.getVoteTriple().then((res) => (tripleVote.value = res))
fetchContestants()
</script>

<template>
  <v-card>
    <v-card-title>1位予想</v-card-title>

    <v-card-text class="d-flex flex-column mt-4">
      <v-lazy :model-value="targetContestant !== undefined" v-if="targetContestant !== undefined">
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
      <v-lazy :model-value="tripleVote !== undefined" v-if="tripleVote">
        <div>
          <div>1位</div>
          <UserSelector :items="contestants" v-model:value="tripleVote.first" />
        </div>
      </v-lazy>
      <v-lazy :model-value="tripleVote !== undefined" v-if="tripleVote">
        <div>
          <div>2位</div>
          <UserSelector :items="contestants" v-model:value="tripleVote.second" />
        </div>
      </v-lazy>
      <v-lazy :model-value="tripleVote !== undefined" v-if="tripleVote">
        <div>
          <div>3位</div>
          <UserSelector :items="contestants" v-model:value="tripleVote.third" />
        </div>
      </v-lazy>
    </v-card-text>

    <v-card-actions class="d-flex flex-column mb-4">
      <v-btn :onclick="submitTripleVote" :disabled="tripleVoteSubmitDisable">Submit</v-btn>
    </v-card-actions>
  </v-card>

  <v-snackbar v-model="showSuccessToast" :timeout="timeout" color="green"> Success! </v-snackbar>
</template>
