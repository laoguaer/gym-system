<script setup>
import { onMounted, reactive, ref } from 'vue'

import TrainerInfo from './components/tainnerInfo.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
import api from '@/api'
</script>

<template>
  <div class="min-h-[100vh] flex flex-col px-4 py-4">
    <TrainerInfo />
  </div>

  <AppFooter />
</template>
