<script setup lang="ts">
import { computed } from 'vue'

import useIconUrl from '@/composables/useIconUrl'
import { useExistTrapIds } from '@/store/existTrapIds'

const props = defineProps({
  trapId: String
})

const { genIconUrl } = useIconUrl()
const { isExist } = useExistTrapIds()

const showIcon = computed(() => {
  return isExist(props.trapId ?? '')
})
const iconUrl = computed(() => genIconUrl(props.trapId ?? ''))
</script>

<template>
  <v-avatar style="background-color: white">
    <v-img v-if="showIcon" :src="iconUrl"> </v-img>
  </v-avatar>
</template>
