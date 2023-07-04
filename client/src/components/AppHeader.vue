<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import apiClient from '@/apis'
import { UserInfo } from '@/apis/generated'
import UserIcon from '@/components/UserIcon.vue'
import { RouteName } from '@/router'

const router = useRouter()

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
