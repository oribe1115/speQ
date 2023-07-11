<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { ref } from 'vue'
import { computed } from 'vue'

import apiClient from '@/apis'
import { ScoreInfo, traPId } from '@/apis/generated'
import UserSelector from '@/components/UserSelector.vue'
import { useContestantsStore } from '@/store/contestants'

const { fetchContestants } = useContestantsStore()
const { contestants } = storeToRefs(useContestantsStore())

const targetContestant = ref<traPId>('')
const score = ref<string>()

const disabled = computed(() => targetContestant.value === '' || score.value === undefined)

const clickOnSave = () => {
  if (targetContestant.value === '' || score.value === undefined) {
    return
  }
  const req: ScoreInfo = {
    contestantId: targetContestant.value,
    score: Number(score.value)
  }
  apiClient.adminOnly.postScore(req).then()
}

fetchContestants()
</script>

<template>
  <v-card class="d-flex flex-column">
    <v-card-title>スコア登録</v-card-title>
    <v-card-subtitle>現在のスコアはダッシュボードから確認してください。</v-card-subtitle>

    <v-card-text class="mt-8">
      <UserSelector :items="contestants" v-model:value="targetContestant" />
      <v-text-field type="number" placeholder="50.134" v-model="score"></v-text-field>
    </v-card-text>

    <v-card-actions class="align-self-center mb-4">
      <v-btn @click="clickOnSave" :disabled="disabled" class="align-self-center">Update</v-btn>
    </v-card-actions>
  </v-card>
</template>
