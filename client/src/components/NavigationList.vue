<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useRouter } from 'vue-router'

import { RouteName } from '@/router'
import { useMeStore } from '@/store/me'

const router = useRouter()

const { isAdminOrRoot } = storeToRefs(useMeStore())

type ListItem = {
  text: string
  linkTo: RouteName
}

const internalLinks: ListItem[] = [
  {
    text: '投票する',
    linkTo: RouteName.Vote
  }
]

const adminLinks: ListItem[] = [
  {
    text: 'top',
    linkTo: RouteName.AdminHome
  },
  {
    text: 'adminユーザー管理',
    linkTo: RouteName.AdminManage
  },
  {
    text: '競技者管理',
    linkTo: RouteName.ContestantManage
  },
  {
    text: 'イベント情報管理',
    linkTo: RouteName.ContestInfoManage
  },
  {
    text: '問題情報管理',
    linkTo: RouteName.ProblemMange
  }
]
</script>

<template>
  <v-list>
    <v-list-item
      v-for="(item, i) in internalLinks"
      :key="i"
      @click="router.push({ name: item.linkTo })"
    >
      {{ item.text }}
    </v-list-item>
  </v-list>

  <v-list v-if="isAdminOrRoot">
    <v-divider></v-divider>
    <v-list-subheader>Admin</v-list-subheader>

    <v-list-item
      v-for="(item, i) in adminLinks"
      :key="i"
      @click="router.push({ name: item.linkTo })"
    >
      {{ item.text }}
    </v-list-item>
  </v-list>
</template>
