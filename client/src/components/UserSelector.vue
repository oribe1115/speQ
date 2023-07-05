<script setup lang="ts">
import UserIcon from './UserIcon.vue'
import { ref } from 'vue'
import { watch } from 'vue'

import { traPId } from '@/apis/generated'

const props = defineProps<{
  items: traPId[]
}>()

const emit = defineEmits<{
  (e: 'selected', value: traPId): traPId
}>()

const selected = ref<traPId>('')

watch(selected, () => {
  emit('selected', selected.value)
})
</script>

<template>
  <v-select v-model="selected" :items="props.items" clearable style="min-width: 300px">
    <!-- 選択肢の表示形式をカスタムすると選択できなくなる -->
    <!-- <template #item="{ item }"><UserIcon :trap-id="item.value" />@{{ item.value }}</template> -->

    <template #selection="{ item }">
      <UserIcon v-if="item.value !== ''" :trap-id="item.value" class="mr-2" />@{{ item.value }}
    </template>
  </v-select>
</template>
