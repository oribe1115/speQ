<script setup lang="ts">
import NavigationList from './NavigationList.vue'
import { storeToRefs } from 'pinia'
import { ref } from 'vue'

import UserIcon from '@/components/UserIcon.vue'
import { useMeStore } from '@/store/me'

const { isLogin, trapId } = storeToRefs(useMeStore())
const drawer = ref(false)
const pcDrawer = ref(true)
</script>

<template>
  <v-app-bar color="primary">
    <!-- スマホ画面でのみメニューアイコンを表示 -->
    <v-app-bar-nav-icon @click="drawer = !drawer" class="d-flex d-lg-none"></v-app-bar-nav-icon>

    <v-app-bar-title>speQ</v-app-bar-title>

    <template v-slot:append>
      <div v-if="isLogin">
        <UserIcon :trap-id="trapId" />
        {{ trapId }}
      </div>
    </template>
  </v-app-bar>

  <!-- PC画面用 -->
  <v-navigation-drawer class="d-none d-sm-flex" v-model="pcDrawer">
    <NavigationList />
  </v-navigation-drawer>

  <!-- スマホ画面用 -->
  <v-navigation-drawer class="d-flex d-lg-none" v-model="drawer" temporary>
    <NavigationList />
  </v-navigation-drawer>
</template>
