<script setup lang="ts">
import { ref } from 'vue'

import apiClient from '@/apis'
import { traPId } from '@/apis/generated'

const rootUsers = ref<traPId[]>([])

apiClient.user.getRootUsers().then((res) => (rootUsers.value = [res]))
</script>

<template>
  <v-card>
    <v-card-title> Root </v-card-title>

    <v-card-subtitle>rootユーザーはアプリケーション実行時の環境変数で設定されます</v-card-subtitle>
    <v-card-text> </v-card-text>
  </v-card>

  <v-card class="mt-8">
    <v-card-title> Admin </v-card-title>

    <v-card-subtitle>イベントに関する設定や情報更新ができます</v-card-subtitle>
    <v-card-text> </v-card-text>
  </v-card>
</template>
