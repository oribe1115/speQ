<script setup lang="ts">
import { ref } from 'vue'

import UserIcon from '@/components/UserIcon.vue'
import { useExistTrapIds } from '@/store/existTrapIds'

const props = defineProps<{
  trapId: string
}>()

const currentInput = ref<string>(props.trapId)

defineEmits(['update:trapId', 'delete'])

const { isExist } = useExistTrapIds()
const rules = {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  required: (value: any) => {
    if (value === '') {
      return 'Field is required'
    } else if (!isExist(value)) {
      return '存在しないtraP IDです'
    }
    return true
  }
}
</script>

<template>
  <div class="d-flex align-center" style="column-gap: 16px">
    <UserIcon :trap-id="currentInput" class="mb-4" />
    <v-text-field
      prefix="@"
      v-model="currentInput"
      @update:model-value="$emit('update:trapId', currentInput)"
      :rules="[rules.required]"
    ></v-text-field>
    <v-btn icon="mdi-close" color="red" class="mb-4" @click="$emit('delete')"></v-btn>
  </div>
</template>
