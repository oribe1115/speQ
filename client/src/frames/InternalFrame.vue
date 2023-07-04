<script setup lang="ts">
import useMe from '@/composables/users/useMe'
import BaseFrame from '@/frames/BaseFrame.vue'

const { isLogin } = useMe()
</script>

<template>
  <base-frame>
    <div v-if="isLogin()">
      <router-view></router-view>
    </div>
    <div v-else>
      <p>Login Required</p>
    </div>
  </base-frame>
</template>
