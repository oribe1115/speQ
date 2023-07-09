<script setup lang="ts">
import apiClient from '@/apis'
import { ProblemInfo } from '@/apis/generated'
import ProblemListInputForm from '@/components/ProblemListInputForm.vue'

// TODO: 現在のproblemsの取得

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
      <ProblemListInputForm :initial-list="[]" :submit-func="submit" />
    </v-card-text>
  </v-card>
</template>
