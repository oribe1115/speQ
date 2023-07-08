<script setup lang="ts">
import { ref } from 'vue'

import { ContestInfo } from '@/apis/generated'

const contestInfo = ref<ContestInfo>({})
</script>

<template>
  <v-card>
    <v-card-item>
      <v-card-title>イベント名</v-card-title>

      <v-text-field v-model="contestInfo.title"></v-text-field>
    </v-card-item>
  </v-card>
</template>
