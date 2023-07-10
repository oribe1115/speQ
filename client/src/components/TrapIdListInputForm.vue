<script setup lang="ts">
import { ref } from 'vue'

import { traPId } from '@/apis/generated'
import TrapIdInputForm from '@/components/TrapIdInputForm.vue'
import { useExistTrapIds } from '@/store/existTrapIds'

const props = defineProps<{
  initialList: traPId[]
  submitFunc: (ids: traPId[]) => void
}>()

const trapIds = ref<traPId[]>([...props.initialList])

const { isExist } = useExistTrapIds()

const removeTrapId = (index: number) => {
  trapIds.value.splice(index, 1)
}
const appendEmptyTrapId = () => {
  trapIds.value.push('')
}

const submittable = () => {
  return trapIds.value.every((id) => isExist(id))
}

const clickOnSave = () => {
  props.submitFunc(trapIds.value)
}
</script>

<template>
  <div class="d-flex flex-column">
    <v-list :key="trapIds.length">
      <v-list-item v-for="(_id, i) in trapIds" :key="i">
        <TrapIdInputForm v-model:trap-id="trapIds[i]" @delete="removeTrapId" />
      </v-list-item>
    </v-list>
    <v-btn icon="mdi-plus" class="align-self-center" :onclick="appendEmptyTrapId"></v-btn>

    <v-btn
      @click="clickOnSave"
      :disabled="!submittable()"
      :width="40"
      class="align-self-center mt-10"
      >Save</v-btn
    >
  </div>
</template>
