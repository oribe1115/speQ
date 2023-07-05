<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import UserIcon from '@/components/UserIcon.vue'
import { RouteName } from '@/router'
import { useMeStore } from '@/store/me'

const router = useRouter()

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
    <v-list>
      <v-list-item nav @click="router.push(RouteName.Vote)"> 投票する </v-list-item>
    </v-list>
  </v-navigation-drawer>

  <!-- スマホ画面用 -->
  <v-navigation-drawer class="d-flex d-lg-none" v-model="drawer" temporary>
    <v-list>
      <v-list-item nav @click="router.push(RouteName.Vote)"> 投票する </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>
