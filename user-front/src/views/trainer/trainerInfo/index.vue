<script setup>
import { onMounted, reactive, ref } from 'vue'

import ChatBot from '../../../components/layout/ChatBot.vue'
import TrainerInfo from './components/tainnerInfo.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
import { useUserStore } from '@/store'

// const userStore = useUserStore()
</script>

<template>
  <div class="min-h-[100vh] flex flex-col px-4 py-4">
    <TrainerInfo />
    <!-- <template v-if="userStore.userId">
      <ChatBot />
    </template> -->
  </div>

  <AppFooter />
</template>
