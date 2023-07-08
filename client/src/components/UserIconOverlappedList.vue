<script setup lang="ts">
import UserIcon from './UserIcon.vue'

const props = defineProps<{
  trapIds: string[]
}>()
</script>

<template>
  <div>
    <UserIcon
      v-for="(traPId, i) in props.trapIds"
      :key="i"
      :trap-id="traPId"
      v-bind:class="i !== 0 ? 'overlapped-item' : ''"
    />
  </div>
</template>

<style>
.overlapped-item {
  margin-left: -15px;
}
</style>
