<script setup lang="ts">
import UserIcon from './UserIcon.vue'

const props = defineProps<{
  trapIds: string[]
}>()
</script>

<template>
  <div class="overlapped-list">
    <UserIcon
      v-for="(traPId, i) in props.trapIds"
      :key="i"
      :trap-id="traPId"
      class="overlapped-item"
    />
  </div>
</template>

<style>
.overlapped-item {
  margin-left: -15px;
}

.overlapped-list:first-child {
  margin-left: 0px;
}
</style>
