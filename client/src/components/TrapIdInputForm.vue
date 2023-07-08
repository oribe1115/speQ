<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  trapId: string
}>()

const currentInput = ref<string>(props.trapId)

defineEmits(['update:trapId'])
</script>

<template>
  <v-text-field
    prefix="@"
    v-model="currentInput"
    @update:model-value="$emit('update:trapId', currentInput)"
  ></v-text-field>
</template>
