<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { useRouter } from 'vue-router'

import { RouteName } from '@/router'
import { useMeStore } from '@/store/me'

const router = useRouter()

const { isAdminOrRoot } = storeToRefs(useMeStore())
</script>

<template>
  <v-list>
    <v-list-item nav @click="router.push(RouteName.Vote)"> 投票する </v-list-item>

    <v-list-group value="Admin" v-if="isAdminOrRoot">
      <template v-slot:activator="{ props }">
        <v-list-item v-bind="props" title="Admin"></v-list-item>
      </template>
      <v-list-item>test</v-list-item>
    </v-list-group>
  </v-list>
</template>
