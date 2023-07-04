<script setup lang="ts">
import { ref } from 'vue'

import apiClient from '@/apis'
import { ContestInfo } from '@/apis/generated'

const contestInfo = ref<ContestInfo>()

apiClient.contest.getContestInfo().then((res) => (contestInfo.value = res))
</script>

<template>
  <h1>{{ contestInfo?.title }}</h1>
  <p>{{ contestInfo?.description }}</p>
</template>
