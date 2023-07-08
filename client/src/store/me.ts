import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

import apiClient from '@/apis'
import { UserInfo } from '@/apis/generated'

export const useMeStore = defineStore('me', () => {
  const myInfo = ref<UserInfo>()

  const isLogin = computed(() => myInfo.value !== undefined)
  const trapId = computed(() => myInfo.value?.traPId ?? '')

  const fetchMe = () => {
    apiClient.user.getMe().then((res) => (myInfo.value = res))
  }

  return {
    isLogin,
    trapId,
    fetchMe
  }
})
