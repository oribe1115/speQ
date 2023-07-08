<script setup lang="ts">
import { ref } from 'vue'

import apiClient from '@/apis'
import { ContestInfo } from '@/apis/generated'

const contestInfo = ref<ContestInfo>({})
const sampleDateText = '2023-04-01T14:15:22+09:00'

const clickOnSave = () => {
  apiClient.adminOnly.putContestInfo(contestInfo.value)
}

apiClient.contest.getContestInfo().then((res) => (contestInfo.value = res))
</script>

<template>
  <v-card class="d-flex flex-column">
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

    <v-card-item>
      <v-card-title>開始予定時刻</v-card-title>
      <v-card-subtitle>競技開始前にトップページで表示</v-card-subtitle>

      <v-card-text>
        この時間にはシステム上ではイベントは開始されません。

        <v-text-field
          v-model="contestInfo.scheduled_start_time"
          :placeholder="sampleDateText"
        ></v-text-field>
      </v-card-text>
    </v-card-item>

    <v-card-item>
      <v-card-title>競技開始時刻</v-card-title>
      <v-card-subtitle>システム上で競技が開始される時刻</v-card-subtitle>

      <v-card-text>
        競技中は様々な情報の変更が行えなくなります。
        競技の開始準備が整ってから入力することを推奨します。

        <v-text-field v-model="contestInfo.start_time" :placeholder="sampleDateText"></v-text-field>
      </v-card-text>
    </v-card-item>

    <v-card-item>
      <v-card-title>競技終了時刻</v-card-title>
      <v-card-subtitle>システム上で競技が終了される時刻</v-card-subtitle>

      <v-card-text>
        競技の開始準備が整ってから入力することを推奨します。 競技中も編集が可能です。

        <v-text-field v-model="contestInfo.end_time" :placeholder="sampleDateText"></v-text-field>
      </v-card-text>
    </v-card-item>

    <v-card-item>
      <v-card-title>投票フリーズ時刻</v-card-title>
      <v-card-subtitle>投票が行えなくなる時刻</v-card-subtitle>

      <v-card-text>
        競技中も編集が可能です。

        <v-text-field
          v-model="contestInfo.voting_freeze_time"
          :placeholder="sampleDateText"
        ></v-text-field>
      </v-card-text>
    </v-card-item>

    <v-card-actions class="align-self-center mb-4">
      <v-btn @click="clickOnSave" class="align-self-center">Save</v-btn>
    </v-card-actions>
  </v-card>
</template>
