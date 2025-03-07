<script setup>
import { onMounted, ref } from 'vue'
import { storeToRefs } from 'pinia'

// 假设将来会有一个教练相关的store
// import { useTrainerStore } from '@/store'

// 模拟教练数据
const trainerInfo = ref({
  avatar: 'https://randomuser.me/api/portraits/men/32.jpg',
  name: '李教练',
  specialty: '力量训练、有氧运动',
  rating: 4.8,
  experience: '5年',
  intro: '国家健身教练认证，专注于帮助客户实现健康减脂和肌肉塑造目标。',
  students: 128,
  courses: 15,
})

// 未来可以从API获取数据
onMounted(async () => {
//   const resp = await api.getTrainerInfo(id)
//   trainerInfo.value = resp.data
})
</script>

<template>
  <div class="card-view animate-zoom-in animate-duration-600 space-y-3">
    <p class="flex items-center text-lg">
      <span class="i-icon-park:fitness mr-1.5" />
      <span>教练信息</span>
    </p>

    <!-- 教练头像和基本信息 -->
    <div class="flex flex-col items-center space-y-2">
      <img :src="trainerInfo.avatar" class="border-primary h-20 w-20 border-2 rounded-full duration-300 hover:scale-110" alt="trainer avatar">
      <h3 class="text-lg font-bold">
        {{ trainerInfo.name }}
      </h3>
      <div class="text-sm text-gray-600">
        {{ trainerInfo.specialty }}
      </div>

      <!-- 评分 -->
      <div class="flex items-center space-x-1">
        <span class="i-icon-park:star-fill text-yellow-500" />
        <span>{{ trainerInfo.rating }}</span>
      </div>
    </div>

    <!-- 教练详细信息 -->
    <div class="mt-3 space-y-2">
      <p>
        <span>教龄：</span>
        <span class="float-right">{{ trainerInfo.experience }}</span>
      </p>
      <p>
        <span>学员数：</span>
        <span class="float-right">{{ trainerInfo.students }}</span>
      </p>
      <p>
        <span>课程数：</span>
        <span class="float-right">{{ trainerInfo.courses }}</span>
      </p>
    </div>

    <!-- 教练简介 -->
    <div class="mt-3">
      <p class="text-sm text-gray-700">
        {{ trainerInfo.intro }}
      </p>
    </div>

    <!-- 预约按钮 -->
    <div class="mt-4">
      <button class="bg-primary hover:bg-primary-dark w-full rounded-md py-2 text-white transition-300">
        预约课程
      </button>
    </div>
  </div>
</template>
