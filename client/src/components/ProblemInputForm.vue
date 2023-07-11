<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  prependText: string
  title: string
  description: string | undefined
}>()

const currentTitle = ref<string>(props.title)
const currentDescription = ref<string>(props.description ?? '')

defineEmits(['update:title', 'update:description', 'delete'])
</script>

<template>
  <div class="d-flex flex-column">
    <div class="d-flex justify-space-between">
      <div class="d-flex align-content-center">
        <div style="width: 2rem; font-size: large">{{ prependText }}</div>
        <v-text-field
          v-model="currentTitle"
          @update:model-value="$emit('update:title', currentTitle)"
        >
        </v-text-field>
      </div>
      <v-btn icon="mdi-close" color="red" @click="$emit('delete')"></v-btn>
    </div>

    <v-textarea
      v-model="currentDescription"
      @update:model-value="$emit('update:description', currentDescription)"
      class="ml-8"
    ></v-textarea>
  </div>
</template>
