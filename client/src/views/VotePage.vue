<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { ref } from 'vue'

import { traPId } from '@/apis/generated'
import UserSelector from '@/components/UserSelector.vue'
import { useContestantsStore } from '@/store/contestants'

const { fetchContestants } = useContestantsStore()
const { contestants } = storeToRefs(useContestantsStore())

const targetContestant = ref<traPId>('')

const selected = (value: traPId) => {
  targetContestant.value = value
}

fetchContestants()
</script>

<template>
  <p>vote page</p>
  <UserSelector :items="contestants" @selected="selected" />
</template>
