<script setup lang="ts">
import apiClient from '../apis'
import UserIcon from './UserIcon.vue'
import { ref } from 'vue'

import { UserInfo } from '@/apis/generated'

const me = ref<UserInfo>()
const drawer = ref(false)
const pcDrawer = ref(true)

apiClient.user.getMe().then((res) => (me.value = res))
</script>

<template>
  <v-app-bar color="primary">
    <!-- スマホ画面でのみメニューアイコンを表示 -->
    <v-app-bar-nav-icon @click="drawer = !drawer" class="d-flex d-lg-none"></v-app-bar-nav-icon>

    <v-app-bar-title>speQ</v-app-bar-title>

    <template v-slot:append>
      <div v-if="me !== undefined">
        <UserIcon :trap-id="me.traPId" />
        {{ me.traPId }}
      </div>
    </template>
  </v-app-bar>

  <!-- PC画面用 -->
  <v-navigation-drawer class="d-none d-sm-flex" v-model="pcDrawer">
    <v-list>
      <v-list-item> test </v-list-item>
    </v-list>
  </v-navigation-drawer>

  <!-- スマホ画面用 -->
  <v-navigation-drawer class="d-flex d-lg-none" v-model="drawer" temporary>
    <v-list>
      <v-list-item> test </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>
