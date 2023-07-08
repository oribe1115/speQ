<script setup lang="ts">
import { ref } from 'vue'

import { traPId } from '@/apis/generated'
import TrapIdInputForm from '@/components/TrapIdInputForm.vue'

const props = defineProps<{
  initialList: traPId[]
}>()

const trapIds = ref<traPId[]>([...props.initialList])

const removeTrapId = (index: number) => {
  trapIds.value.splice(index, 1)
}
const appendEmptyTrapId = () => {
  trapIds.value.push('')
}
</script>

<template>
  <div class="d-flex flex-column">
    <v-list>
      <v-list-item v-for="(_id, i) in trapIds" :key="i">
        <TrapIdInputForm v-model:trap-id="trapIds[i]" @delete="removeTrapId" />
      </v-list-item>
    </v-list>
    <v-btn icon="mdi-plus" class="align-self-center" :onclick="appendEmptyTrapId"></v-btn>
  </div>
</template>
