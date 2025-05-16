<script setup>
import { onMounted, ref } from 'vue'
import { storeToRefs } from 'pinia'
import { RouterLink } from 'vue-router'

// 假设将来会有一个教练相关的store
// import { useTrainerStore } from '@/store'

// 模拟一句励志名言
const quote = ref('"You don\'t get what you wish for, you get what you work for."')
const trainer = ref('Elvis Lam')

// 模拟教练数据数组
const trainers = ref([
  {
    id: 1,
    avatar: 'https://randomuser.me/api/portraits/men/32.jpg',
    name: '李教练',
    specialty: '力量训练、有氧运动',
    rating: 4.8,
    experience: '5年',
    intro: '国家健身教练认证，专注于帮助客户实现健康减脂和肌肉塑造目标。',
    students: 128,
    courses: 15,
  },
  {
    id: 2,
    avatar: 'https://randomuser.me/api/portraits/men/33.jpg',
    name: '王教练',
    specialty: '瑜伽、普拉提',
    rating: 4.9,
    experience: '7年',
    intro: '资深瑜伽教练，擅长帮助学员提高身体柔韧性和核心力量。',
    students: 156,
    courses: 12,
  },
  {
    id: 3,
    avatar: 'https://randomuser.me/api/portraits/men/34.jpg',
    name: '张教练',
    specialty: '功能性训练、康复训练',
    rating: 4.7,
    experience: '4年',
    intro: '专注于功能性训练和运动康复，帮助客户恢复运动机能和预防伤害。',
    students: 98,
    courses: 8,
  },
  {
    id: 4,
    avatar: 'https://randomuser.me/api/portraits/men/35.jpg',
    name: '刘教练',
    specialty: '搏击、拳击',
    rating: 4.6,
    experience: '6年',
    intro: '前职业拳击运动员，现专注于搏击和拳击训练，提供高强度的有氧训练课程。',
    students: 112,
    courses: 10,
  },
  {
    id: 5,
    avatar: 'https://randomuser.me/api/portraits/men/36.jpg',
    name: '陈教练',
    specialty: '游泳、水中有氧',
    rating: 4.9,
    experience: '8年',
    intro: '前国家游泳队队员，擅长游泳技术指导和水中有氧训练。',
    students: 145,
    courses: 14,
  },
  {
    id: 6,
    avatar: 'https://randomuser.me/api/portraits/men/37.jpg',
    name: '赵教练',
    specialty: '健美、增肌训练',
    rating: 4.8,
    experience: '5年',
    intro: '健美比赛冠军，专注于肌肉增长和体型塑造训练。',
    students: 132,
    courses: 11,
  },
])

// 未来可以从API获取数据
onMounted(async () => {
//   const resp = await api.getTrainerInfo()
//   trainers.value = resp.data
})
</script>

<template>
  <div class="card-view w-full animate-zoom-in animate-duration-600 p-6 space-y-8">
    <h2 class="mb-6 text-center text-2xl font-bold">
      我们的专业教练团队
    </h2>
    <!-- 左右布局：左侧文字，右侧教练头像 -->
    <div class="flex flex-col items-center justify-between gap-8 md:flex-row">
      <!-- 左侧文字区域 -->
      <div class="flex-1 text-center space-y-6 md:text-left">
        <div>
          <p class="text-2xl font-medium">
            {{ quote }}
          </p>
          <p class="mt-3 text-gray-600">
            Trainer
          </p>
          <p class="text-xl font-semibold">
            {{ trainer }}
          </p>
        </div>
        <div>
          <RouterLink to="/trainerInfo" class="inline-block border border-gray-300 rounded-lg px-6 py-3 transition duration-300 hover:scale-105 hover:bg-gray-100 hover:shadow-md">
            Know More →
          </RouterLink>
        </div>
      </div>

      <!-- 右侧教练头像区域 -->
      <div class="grid grid-cols-2 max-w-65% gap-4 rounded-lg bg-gray-50 p-4 md:grid-cols-3 md:w-auto">
        <div v-for="trainer in trainers" :key="trainer.id" class="flex flex-col items-center">
          <div class="overflow-hidden rounded-lg bg-white shadow-md transition duration-300 hover:scale-105 hover:shadow-xl">
            <img :src="trainer.avatar" class="h-64 w-full object-cover" alt="trainer avatar">
            <div class="p-2 text-center" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
