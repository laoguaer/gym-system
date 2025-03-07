<script setup>
import { ref } from 'vue'

// 模拟多个教练数据
const trainers = ref([
  {
    id: 1,
    avatar: 'https://static.wixstatic.com/media/c50fb9_d761471e57d246769bbf6bfacade5a92~mv2.jpg/v1/fill/w_248,h_296,al_c,q_80,usm_0.66_1.00_0.01,enc_auto/c50fb9_d761471e57d246769bbf6bfacade5a92~mv2.jpg',
    name: '李教练',
    specialty: '力量训练',
  },
  {
    id: 2,
    avatar: 'https://static.wixstatic.com/media/c50fb9_06359b65b9ae47e481b7d5ee288bb69e~mv2.jpg/v1/fill/w_248,h_296,al_c,q_80,usm_0.66_1.00_0.01,enc_auto/c50fb9_06359b65b9ae47e481b7d5ee288bb69e~mv2.jpg',
    name: '王教练',
    specialty: '瑜伽',
  },
  {
    id: 3,
    avatar: 'https://static.wixstatic.com/media/c50fb9_9fc6e96f250d4c93ba9647bb739d6c1b~mv2.jpg/v1/fill/w_248,h_296,al_c,q_80,usm_0.66_1.00_0.01,enc_auto/c50fb9_9fc6e96f250d4c93ba9647bb739d6c1b~mv2.jpg',
    name: '张教练',
    specialty: '有氧运动',
  },
  {
    id: 4,
    avatar: 'https://static.wixstatic.com/media/c50fb9_e5de07b0d45f4dbcb7c62d9db51ecc62~mv2.jpg/v1/fill/w_248,h_296,al_c,q_80,usm_0.66_1.00_0.01,enc_auto/c50fb9_e5de07b0d45f4dbcb7c62d9db51ecc62~mv2.jpg',
    name: '刘教练',
    specialty: '普拉提',
  },
])
</script>

<template>
  <div class="card-view animate-zoom-in animate-duration-600">
    <p class="mb-3 flex items-center text-lg">
      <span class="i-icon-park:fitness mr-1.5" />
      <span>热门教练</span>
    </p>

    <!-- 教练头像网格 -->
    <div class="grid grid-cols-2 gap-3">
      <div
        v-for="trainer in trainers"
        :key="trainer.id"
        class="group flex flex-col cursor-pointer items-center"
      >
        <div class="relative overflow-hidden rounded-full">
          <img
            :src="trainer.avatar"
            class="border-primary h-16 w-16 border-2 rounded-full transition-all duration-300 group-hover:rotate-6 group-hover:scale-110"
            :alt="trainer.name"
          >
          <div class="bg-primary/30 absolute inset-0 flex items-center justify-center rounded-full opacity-0 transition-opacity duration-300 group-hover:opacity-100">
            <span class="text-xs text-white font-bold">查看</span>
          </div>
        </div>
        <p class="mt-1 text-center text-sm">
          {{ trainer.name }}
        </p>
      </div>
    </div>
  </div>
</template>
