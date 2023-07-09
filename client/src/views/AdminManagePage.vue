<script setup lang="ts">
import { ref } from 'vue'

import apiClient from '@/apis'
import { traPId } from '@/apis/generated'
import TrapIdListInputForm from '@/components/TrapIdListInputForm.vue'
import UserIcon from '@/components/UserIcon.vue'

const rootUsers = ref<traPId[]>([])
const adminUsers = ref<traPId[]>([])

apiClient.user.getRootUsers().then((res) => (rootUsers.value = [res]))
apiClient.user.getAdmins().then((res) => (adminUsers.value = res))

const submit = (ids: traPId[]) => {
  apiClient.adminOnly.putAdminUsers(ids)
}
</script>

<template>
  <v-card>
    <v-card-title> Root </v-card-title>

    <v-card-subtitle>rootユーザーはアプリケーション実行時の環境変数で設定されます</v-card-subtitle>
    <v-card-text>
      <div v-for="(id, i) in rootUsers" :key="i"><UserIcon :trap-id="id" /> {{ id }}</div>
    </v-card-text>
  </v-card>

  <v-card class="mt-8">
    <v-card-title> Admin </v-card-title>

    <v-card-subtitle>イベントに関する設定や情報更新ができます</v-card-subtitle>
    <v-card-text>
      <TrapIdListInputForm :initial-list="adminUsers" :submit-func="submit" />
    </v-card-text>
  </v-card>
</template>
