<script setup lang="ts">
import AppHeader from '@/components/AppHeader.vue'
import { useMeStore } from '@/store/me'

const { fetchMe } = useMeStore()

fetchMe()
</script>

<template>
  <v-app>
    <AppHeader />

    <v-main>
      <RouterView />
    </v-main>
  </v-app>
</template>
