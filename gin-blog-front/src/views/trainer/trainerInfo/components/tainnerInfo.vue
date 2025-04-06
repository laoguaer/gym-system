<script setup>
import { computed, onMounted, ref } from 'vue'
import { storeToRefs } from 'pinia'

// 导入教练相关的store
import { useTrainerStore } from '@/store/trainer'

const trainerStore = useTrainerStore()
const { loading, error } = storeToRefs(trainerStore)

// 分页相关 - 减小每页显示数量以适应一个屏幕
const currentPage = ref(1)
const pageSize = ref(6)
const totalPages = computed(() => Math.ceil(trainerStore.trainerCount / pageSize.value))

// 当前页面显示的教练数据
const displayedTrainers = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return trainerStore.trainers.slice(start, end)
})

// 排序方式
const sortOptions = [
  { label: '默认排序', value: 'default' },
  { label: '按评分排序', value: 'rating' },
  { label: '按经验排序', value: 'experience' },
  { label: '按学员数排序', value: 'students' },
]
const currentSort = ref('default')

// 根据排序方式获取教练列表
const trainers = computed(() => {
  switch (currentSort.value) {
    case 'rating':
      return [...displayedTrainers.value].sort((a, b) => b.rating - a.rating)
    case 'experience':
      // 由于experience是字符串（如"5年"），需要提取数字部分进行比较
      return [...displayedTrainers.value].sort((a, b) => {
        const aYears = Number.parseInt(a.experience.replace(/[^0-9]/g, ''))
        const bYears = Number.parseInt(b.experience.replace(/[^0-9]/g, ''))
        return bYears - aYears
      })
    case 'students':
      return [...displayedTrainers.value].sort((a, b) => b.students - a.students)
    default:
      return displayedTrainers.value
  }
})

// 从store获取教练数据
onMounted(async () => {
  await trainerStore.getTrainerList()
})
</script>

<template>
  <div class="flex flex-col">
    <!-- 加载状态和错误提示 -->
    <div v-if="loading" class="flex items-center justify-center py-10">
      <div class="h-12 w-12 animate-spin border-b-2 border-t-2 border-blue-500 rounded-full" />
    </div>

    <div v-else-if="error" class="relative mb-4 border border-red-400 rounded bg-red-100 px-4 py-3 text-red-700">
      <span class="block sm:inline">{{ error }}</span>
    </div>

    <div v-else>
      <!-- 排序选项 -->
      <div class="flex-wra mt-4 h-35 flex items-center justify-between gap-2">
        <h2 class="text-2xl text-gray-800 font-bold">
          教练团队
        </h2>
        <div class="flex items-center space-x-2">
          <span class="text-gray-600">排序方式:</span>
          <select v-model="currentSort" class="border border-gray-300 rounded-md px-3 py-1 focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500">
            <option v-for="option in sortOptions" :key="option.value" :value="option.value">
              {{ option.label }}
            </option>
          </select>
        </div>
      </div>

      <!-- 教练卡片网格 - 减小间距和内边距 -->
      <div class="grid grid-cols-1 flex-1 gap-3 overflow-hidden lg:grid-cols-3 md:grid-cols-2 -mt-8">
        <div v-for="trainer in trainers" :key="trainer.id" class="rounded-lg bg-white p-3 shadow-md transition duration-300 hover:shadow-lg">
          <div class="mb-2 flex items-center">
            <img :src="trainer.avatar" class="mr-4 h-16 w-16 rounded-full" alt="trainer avatar">
            <div>
              <h3 class="text-lg font-semibold">
                {{ trainer.name }}
              </h3>
              <p class="text-sm text-gray-600">
                {{ trainer.specialty }}
              </p>
            </div>
          </div>
          <div class="space-y-1">
            <p class="flex justify-between">
              <span class="text-gray-600">评分:</span>
              <span class="font-medium">{{ trainer.rating }}/5.0</span>
            </p>
            <p class="flex justify-between">
              <span class="text-gray-600">教龄:</span>
              <span class="font-medium">{{ trainer.experience }}</span>
            </p>
            <p class="flex justify-between">
              <span class="text-gray-600">课程数:</span>
              <span class="font-medium">{{ trainer.courses }}门</span>
            </p>
          </div>

          <p class="line-clamp-2 mt-2 text-xs text-gray-700">
            {{ trainer.intro }}
          </p>

          <button 
            class="mt-2 w-full rounded-lg bg-blue-600 py-1 text-sm text-white transition duration-300 hover:bg-blue-700"
            @click="$router.push(`/courseInfo?coach_id=${trainer.id}`)"
          >
            预约课程
          </button>
        </div>
      </div>
      <!-- // 分页控件 - 减小上边距 -->
      <div class="mt-3 flex justify-center">
        <div class="flex space-x-2">
          <button
            :disabled="currentPage === 1"
            class="border border-gray-300 rounded-md px-3 py-1 transition duration-300 disabled:cursor-not-allowed hover:bg-gray-100 disabled:opacity-50"
            @click="currentPage = Math.max(1, currentPage - 1)"
          >
            上一页
          </button>

          <div class="flex space-x-1">
            <!-- 页数较少时显示所有页码 -->
            <template v-if="totalPages <= 7">
              <button
                v-for="page in totalPages"
                :key="page"
                class="border border-gray-300 rounded-md px-3 py-1 transition duration-300"
                :class="[currentPage === page ? 'bg-blue-500 text-white' : 'bg-white text-gray-700 hover:bg-gray-100']"
                @click="currentPage = page"
              >
                {{ page }}
              </button>
            </template>

            <!-- 页数较多时显示部分页码 -->
            <template v-else>
              <!-- 第一页 -->
              <button
                class="border border-gray-300 rounded-md px-3 py-1 transition duration-300"
                :class="[currentPage === 1 ? 'bg-blue-500 text-white' : 'bg-white text-gray-700 hover:bg-gray-100']"
                @click="currentPage = 1"
              >
                1
              </button>

              <!-- 左省略号 -->
              <span v-if="currentPage > 3" class="border border-gray-300 rounded-md px-3 py-1">...</span>

              <!-- 当前页附近的页码 -->
              <template v-for="page in totalPages" :key="page">
                <button
                  v-if="page !== 1 && page !== totalPages && Math.abs(page - currentPage) <= 1"
                  class="border border-gray-300 rounded-md px-3 py-1 transition duration-300"
                  :class="[currentPage === page ? 'bg-blue-500 text-white' : 'bg-white text-gray-700 hover:bg-gray-100']"
                  @click="currentPage = page"
                >
                  {{ page }}
                </button>
              </template>

              <!-- 右省略号 -->
              <span v-if="currentPage < totalPages - 2" class="border border-gray-300 rounded-md px-3 py-1">...</span>

              <!-- 最后一页 -->
              <button
                class="border border-gray-300 rounded-md px-3 py-1 transition duration-300"
                :class="[currentPage === totalPages ? 'bg-blue-500 text-white' : 'bg-white text-gray-700 hover:bg-gray-100']"
                @click="currentPage = totalPages"
              >
                {{ totalPages }}
              </button>
            </template>
          </div>

          <button
            :disabled="currentPage === totalPages"
            class="border border-gray-300 rounded-md px-3 py-1 transition duration-300 disabled:cursor-not-allowed hover:bg-gray-100 disabled:opacity-50"
            @click="currentPage = Math.min(totalPages, currentPage + 1)"
          >
            下一页
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
