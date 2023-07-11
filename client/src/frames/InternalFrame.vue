<script setup lang="ts">
import { storeToRefs } from 'pinia'

import BaseFrame from '@/frames/BaseFrame.vue'
import { useMeStore } from '@/store/me'

const { isLogin } = storeToRefs(useMeStore())
</script>

<template>
  <base-frame>
    <div v-if="isLogin">
      <router-view></router-view>
    </div>
    <div v-else>
      <p>Login Required</p>
    </div>
  </base-frame>
</template>
