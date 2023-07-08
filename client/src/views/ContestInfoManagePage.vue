<script setup lang="ts">
import { ref } from 'vue'

import apiClient from '@/apis'
import { ContestInfo } from '@/apis/generated'

const contestInfo = ref<ContestInfo>({})

apiClient.contest.getContestInfo().then((res) => (contestInfo.value = res))
</script>

<template>
  <v-card>
    <v-card-item>
      <v-card-title>イベント名</v-card-title>
      <v-card-text>
        <v-text-field v-model="contestInfo.title"></v-text-field>
      </v-card-text>
    </v-card-item>

    <v-card-item>
      <v-card-title>イベント概要</v-card-title>
      <v-card-text>
        <v-textarea v-model="contestInfo.description"></v-textarea>
      </v-card-text>
    </v-card-item>
  </v-card>
</template>
