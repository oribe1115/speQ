<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  trapId: string
}>()

const currentInput = ref<string>(props.trapId)

defineEmits(['update:trapId', 'delete'])
</script>

<template>
  <div class="d-flex">
    <v-text-field
      prefix="@"
      v-model="currentInput"
      @update:model-value="$emit('update:trapId', currentInput)"
    ></v-text-field>
    <v-btn icon="mdi-close" color="red" class="ml-4" @click="$emit('delete')"></v-btn>
  </div>
</template>
