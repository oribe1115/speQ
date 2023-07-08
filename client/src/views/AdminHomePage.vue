<script setup lang="ts">
import { ref } from 'vue'

import apiClient from '@/apis'
import { traPId } from '@/apis/generated'
import UserIconOverlappedList from '@/components/UserIconOverlappedList.vue'

const rootUsers = ref<traPId[]>([])
const adminUsers = ref<traPId[]>([])

apiClient.user.getRootUsers().then((res) => (rootUsers.value = [res]))
apiClient.user.getAdmins().then((res) => (adminUsers.value = res))
</script>

<template>
  <v-card>
    <v-card-title> 管理者 </v-card-title>

    <v-cart-text>
      Root:
      <UserIconOverlappedList :trap-ids="rootUsers" />
    </v-cart-text>

    <v-cart-text>
      Admin:
      <UserIconOverlappedList :trap-ids="adminUsers" />
    </v-cart-text>
  </v-card>
</template>
