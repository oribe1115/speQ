<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  prependText: string
  text: string
}>()

const currentInput = ref<string>(props.text)

defineEmits(['update:text', 'delete'])
</script>

<template>
  <div class="d-flex align-center" style="column-gap: 16px">
    <v-text-field v-model="currentInput" @update:model-value="$emit('update:text', currentInput)">
      <template v-slot:prepend v-if="props.prependText !== ''">
        {{ props.prependText }}
      </template>
    </v-text-field>
    <v-btn icon="mdi-close" color="red" class="mb-4" @click="$emit('delete')"></v-btn>
  </div>
</template>
