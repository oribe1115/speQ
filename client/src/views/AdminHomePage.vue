<script setup lang="ts">
import { ref } from 'vue'
import { computed } from 'vue'

import apiClient from '@/apis'
import { ContestInfo, traPId } from '@/apis/generated'
import UserIconOverlappedList from '@/components/UserIconOverlappedList.vue'

type ContestInfoItem = {
  key: string
  value: string | undefined
}

const rootUsers = ref<traPId[]>([])
const adminUsers = ref<traPId[]>([])
const contestInfo = ref<ContestInfo>()
const contestants = ref<traPId[]>([])

const contestInfoAsList = computed<ContestInfoItem[]>(() => {
  return [
    { key: 'イベント名', value: contestInfo.value?.title },
    { key: 'イベントの概要', value: contestInfo.value?.description },
    { key: '開始予定時刻', value: contestInfo.value?.scheduled_start_time },
    { key: '開始時刻', value: contestInfo.value?.start_time },
    { key: '終了時刻', value: contestInfo.value?.end_time },
    { key: '投票フリーズ時刻', value: contestInfo.value?.voting_freeze_time }
  ]
})

apiClient.user.getRootUsers().then((res) => (rootUsers.value = [res]))
apiClient.user.getAdmins().then((res) => (adminUsers.value = res))
apiClient.contest.getContestInfo().then((res) => (contestInfo.value = res))
apiClient.user.getContestants().then((res) => (contestants.value = res))
</script>

<template>
  <v-card>
    <v-card-title> 管理者 </v-card-title>

    <v-card-text>
      Root:
      <UserIconOverlappedList :trap-ids="rootUsers" />

      Admin:
      <UserIconOverlappedList :trap-ids="adminUsers" />
    </v-card-text>
  </v-card>

  <v-card class="mt-8">
    <v-card-title>コンテスト情報</v-card-title>

    <v-card-text class="d-flex flex-column" style="row-gap: 1rem">
      <div v-for="(item, i) in contestInfoAsList" :key="i">
        <div>{{ item.key }}</div>
        <div class="ml-4">{{ item.value ?? '未設定' }}</div>
      </div>
    </v-card-text>
  </v-card>

  <v-card class="mt-8">
    <v-card-title>競技者</v-card-title>

    <v-card-text>
      <UserIconOverlappedList :trap-ids="contestants" />
    </v-card-text>
  </v-card>
</template>
