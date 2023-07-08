<script setup lang="ts">
import { ref } from 'vue'

import ReactiveTextForm from '@/components/ReactiveTextForm.vue'

const props = defineProps<{
  initialList: string[]
  submitFunc: (values: string[]) => void
}>()

const values = ref<string[]>([...props.initialList])

const removeItem = (index: number) => {
  values.value.splice(index, 1)
}
const appendEmptyItem = () => {
  values.value.push('')
}

const submittable = () => {
  return values.value.every((text) => text !== '')
}

const clickOnSave = () => {
  props.submitFunc(values.value)
}
</script>

<template>
  <div class="d-flex flex-column">
    <v-list>
      <v-list-item v-for="(_text, i) in values" :key="i">
        <ReactiveTextForm
          v-model:text="values[i]"
          :prepend-text="(i + 1).toString()"
          @delete="removeItem"
        />
      </v-list-item>
    </v-list>
    <v-btn icon="mdi-plus" class="align-self-center" :onclick="appendEmptyItem"></v-btn>

    <v-btn
      @click="clickOnSave"
      :disabled="!submittable()"
      :width="40"
      class="align-self-center mt-10"
      >Save</v-btn
    >
  </div>
</template>
