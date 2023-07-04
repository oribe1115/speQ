<script setup lang="ts">
import useMe from '../composables/users/useMe'

const { isLogin } = useMe()
</script>

<template>
  <div class="d-flex justify-center">
    <div v-if="isLogin()">
      <router-view></router-view>
    </div>
    <div v-else>
      <p>Login Required</p>
    </div>
  </div>
</template>
