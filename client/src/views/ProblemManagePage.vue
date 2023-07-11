<script setup lang="ts">
import { ref } from 'vue'

import apiClient from '@/apis'
import { ProblemInfo } from '@/apis/generated'
import ProblemListInputForm from '@/components/ProblemListInputForm.vue'

const problems = ref<ProblemInfo[] | undefined>()

// TODO: 現在のproblemsの取得
problems.value = []

const submit = (problems: ProblemInfo[]) => {
  apiClient.adminOnly.putProblems(problems)
}
</script>

<!-- TODO: descriptionも扱えるようにする -->
<template>
  <v-card>
    <v-card-title>競技中に解かれる問題</v-card-title>
    <v-card-subtitle>競技中は追加のみ行うことができます</v-card-subtitle>

    <v-card-text>
      <v-lazy :model-value="problems !== undefined" v-if="problems">
        <ProblemListInputForm :initial-list="problems" :submit-func="submit" />
      </v-lazy>
    </v-card-text>
  </v-card>
</template>
