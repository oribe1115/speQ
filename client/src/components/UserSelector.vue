<script setup lang="ts">
import UserIcon from './UserIcon.vue'
import { ref } from 'vue'

import { traPId } from '@/apis/generated'

const props = defineProps<{
  items: traPId[]
  value: traPId | undefined
}>()

defineEmits(['update:value'])

const selected = ref<traPId>(props.value ?? '')
</script>

<template>
  <v-select
    v-model="selected"
    :items="props.items"
    @update:model-value="$emit('update:value', selected)"
    clearable
    style="min-width: 300px"
  >
    <!-- 選択肢の表示形式をカスタムすると選択できなくなる -->
    <!-- <template #item="{ item }"><UserIcon :trap-id="item.value" />@{{ item.value }}</template> -->

    <template #selection="{ item }">
      <UserIcon v-if="item.value !== ''" :trap-id="item.value" class="mr-2" />@{{ item.value }}
    </template>
  </v-select>
</template>
