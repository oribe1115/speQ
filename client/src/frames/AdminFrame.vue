<script setup lang="ts">
import { storeToRefs } from 'pinia'

import BaseFrame from '@/frames/BaseFrame.vue'
import { useMeStore } from '@/store/me'

const { isLogin, isAdminOrRoot } = storeToRefs(useMeStore())
</script>

<template>
  <base-frame>
    <div v-if="!isLogin">
      <p>Login Required</p>
    </div>
    <div v-else-if="!isAdminOrRoot">
      <p>This is Admin page</p>
    </div>
    <div v-else>
      <router-view></router-view>
    </div>
  </base-frame>
</template>
