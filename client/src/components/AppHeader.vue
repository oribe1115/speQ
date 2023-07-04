<script setup lang="ts">
import useMe from '../composables/users/useMe'
import UserIcon from './UserIcon.vue'
import { ref } from 'vue'

const { isLogin, getId } = useMe()

const traPId = ref<string>('')
traPId.value = getId()
</script>

<template>
  <v-app-bar>
    <v-app-bar-title>speQ</v-app-bar-title>
    <template v-slot:append>
      <div v-if="isLogin()">
        <UserIcon :trap-id="traPId" />
      </div>
    </template>
  </v-app-bar>
</template>
