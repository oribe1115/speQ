<script setup lang="ts">
import { useRouter } from 'vue-router'

import { RouteName } from '@/router'

const router = useRouter()
</script>

<template>
  <v-list>
    <v-list-item nav @click="router.push(RouteName.Vote)"> 投票する </v-list-item>
  </v-list>
</template>
