import { defineStore } from 'pinia'
import { ref } from 'vue'

import apiClient from '@/apis'
import { traPId } from '@/apis/generated'

export const useContestantsStore = defineStore('contestants', () => {
  const contestants = ref<traPId[]>([])

  const fetchContestants = () => {
    apiClient.user.getContestants().then((res) => (contestants.value = res))
  }

  return {
    contestants,
    fetchContestants
  }
})
