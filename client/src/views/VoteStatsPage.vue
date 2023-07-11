<script setup lang="ts">
import { ref } from 'vue'

import apiClient from '@/apis'
import { VoteStatsItem } from '@/apis/generated'
import UserIcon from '@/components/UserIcon.vue'
import UserIconOverlappedList from '@/components/UserIconOverlappedList.vue'

const voteStats = ref<VoteStatsItem[]>()

apiClient.vote.getVoteStats().then((res) => (voteStats.value = res))
</script>

<template>
  <v-card>
    <v-card-title>みんなの1位予想</v-card-title>

    <v-card-text class="mt-4">
      <v-lazy :model-value="voteStats !== undefined" v-if="voteStats">
        <v-table>
          <tbody>
            <tr v-for="item in voteStats" :key="item.contestant" style="height: 70px">
              <td class="w-25">
                <UserIcon :trap-id="item.contestant" class="mr-4" />{{ item.contestant }}
              </td>
              <td><UserIconOverlappedList :trap-ids="item.voters" /></td>
            </tr>
          </tbody>
        </v-table>
      </v-lazy>
    </v-card-text>
  </v-card>
</template>
