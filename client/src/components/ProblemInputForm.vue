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
  <div class="d-flex align-center" style="column-gap: 16px">
    {{ prependText }}
    <v-form>
      <v-text-field
        v-model="currentTitle"
        @update:model-value="$emit('update:title', currentTitle)"
      >
      </v-text-field>
      <v-textarea
        v-model="currentDescription"
        @update:model-value="$emit('update:description', currentDescription)"
      ></v-textarea>
    </v-form>

    <v-btn icon="mdi-close" color="red" class="mb-4" @click="$emit('delete')"></v-btn>
  </div>
</template>
