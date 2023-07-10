<script setup lang="ts">
import ProblemInputForm from './ProblemInputForm.vue'
import { ref } from 'vue'

import { ProblemInfo } from '@/apis/generated'

const props = defineProps<{
  initialList: ProblemInfo[]
  submitFunc: (problems: ProblemInfo[]) => void
}>()

const problems = ref<ProblemInfo[]>([...props.initialList])

const removeProblem = (index: number) => {
  problems.value.splice(index, 1)
  correctIds()
}
const appendEmptyItem = () => {
  problems.value.push({
    id: 0,
    title: ''
  })

  correctIds()
}

const submittable = () => {
  return problems.value.every((problem) => problem.title !== '')
}

const clickOnSave = () => {
  correctIds()
  props.submitFunc(problems.value)
}

const correctIds = () => {
  problems.value.forEach((problem, index) => (problem.id = index + 1))
}
</script>

<template>
  <div class="d-flex flex-column">
    <v-list :key="problems.length">
      <v-list-item v-for="(_text, i) in problems" :key="i">
        <ProblemInputForm
          v-model:title="problems[i].title"
          v-model:description="problems[i].description"
          :prepend-text="(i + 1).toString()"
          @delete="removeProblem"
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
