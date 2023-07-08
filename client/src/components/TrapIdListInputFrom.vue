<script setup lang="ts">
import { ref } from 'vue'

import { traPId } from '@/apis/generated'
import TrapIdInputForm from '@/components/TrapIdInputForm.vue'

const props = defineProps<{
  initialList: traPId[]
}>()

const trapIds = ref<traPId[]>([...props.initialList])
</script>

<template>
  <v-list>
    <v-list-item v-for="(_id, i) in trapIds" :key="i">
      <TrapIdInputForm v-model:trap-id="trapIds[i]" />
    </v-list-item>
  </v-list>
</template>
