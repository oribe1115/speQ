<script setup lang="ts">
const props = defineProps({
  trapId: String
})

const iconBaseUrl = 'https://q.trap.jp/api/v3/public/icon/'
const iconUrl = iconBaseUrl + props.trapId
</script>

<template>
  <v-avatar>
    <v-img :src="iconUrl"> </v-img>
  </v-avatar>
</template>
