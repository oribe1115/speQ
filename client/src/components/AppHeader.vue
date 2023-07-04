<script setup lang="ts">
import apiClient from '../apis'
import UserIcon from './UserIcon.vue'
import { ref } from 'vue'

import { UserInfo } from '@/apis/generated'

const me = ref<UserInfo>()
apiClient.user.getMe().then((res) => (me.value = res))
</script>

<template>
  <v-app-bar>
    <v-app-bar-title>speQ</v-app-bar-title>
    <template v-slot:append>
      <div v-if="me !== undefined">
        <UserIcon :trap-id="me.traPId" />
        {{ me.traPId }}
      </div>
    </template>
  </v-app-bar>
</template>
