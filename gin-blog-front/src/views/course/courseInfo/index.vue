<script setup>
import { computed, onMounted, ref, watch } from 'vue'
import { storeToRefs } from 'pinia'
import { useRoute } from 'vue-router'
import { useCourseStore } from '@/store/course'
import { useTrainerStore } from '@/store/trainer'
import { useAppStore, useUserStore } from '@/store'
import AppFooter from '@/components/layout/AppFooter.vue'
import CourseBookingModal from '@/components/modal/CourseBookingModal.vue'

const route = useRoute()
const courseStore = useCourseStore()
const { loading, error, courses, courseCount } = storeToRefs(courseStore)

const trainerStore = useTrainerStore()
const userStore = useUserStore()
const appStore = useAppStore()

// 控制购买课程模态框
const showBookingModal = ref(false)
const selectedCourse = ref(null)

// 搜索条件
const searchForm = ref({
  title: '',
  coach_id: null,
  is_single: null,
  tag_ids: [],
  page: 1,
  size: 8,
})

// 单人/团体课程选项
const singleOptions = [
  { label: '全部', value: null },
  { label: '单人课程', value: true },
  { label: '团体课程', value: false },
]

// 分页相关
const currentPage = computed({
  get: () => searchForm.value.page,
  set: (val) => {
    searchForm.value.page = val
    fetchCourses()
  },
})
const totalPages = computed(() => Math.ceil(courseCount.value / searchForm.value.size))

// 格式化日期时间
function formatDateTime(dateTimeStr) {
  const date = new Date(dateTimeStr)
  return `${date.toLocaleDateString()} ${date.toLocaleTimeString(undefined, { hour: '2-digit', minute: '2-digit' })}`
}

// 获取课程列表
async function fetchCourses() {
  await courseStore.getCourseList(searchForm.value)
}

// 重置搜索条件
function resetSearch() {
  searchForm.value = {
    title: '',
    coach_id: null,
    is_single: null,
    tag_ids: [],
    page: 1,
    size: 8,
  }
  fetchCourses()
}

// 监听搜索条件变化，重置页码
watch(
  () => [searchForm.value.title, searchForm.value.coach_id, searchForm.value.is_single, searchForm.value.tag_ids],
  () => {
    searchForm.value.page = 1
    fetchCourses()
  },
  { deep: true },
)

// 打开购买课程模态框
function openBookingModal(course) {
  // 检查是否登录
  if (!userStore.userId) {
    // 未登录，显示登录框
    appStore.setLoginFlag(true)
    return
  }

  selectedCourse.value = course
  showBookingModal.value = true
}

// 处理购买成功
function handleBookingSuccess() {
  // 刷新课程列表
  fetchCourses()
}

// 初始化
onMounted(async () => {
  // 获取教练列表和标签列表
  await Promise.all([
    courseStore.getTagList(),
    trainerStore.getTrainerList(),
  ])

  // 检查URL参数中是否包含coach_id
  const coachId = route.query.coach_id
  if (coachId) {
    searchForm.value.coach_id = Number(coachId)
  }

  // 获取课程列表
  await fetchCourses()
})
</script>

<template>
  <main class="min-h-[100vh] flex flex-col px-4 py-6">
    <div class="mx-auto container">
      <!-- 页面标题 -->
      <h1 class="mb-6 text-3xl text-gray-800 font-bold">
        课程信息
      </h1>

      <!-- 搜索条件 -->
      <div class="mb-6 rounded-lg bg-white p-4 shadow-md">
        <div class="grid grid-cols-1 gap-4 lg:grid-cols-4 md:grid-cols-2">
          <!-- 课程名称搜索 -->
          <div class="flex flex-col">
            <label class="mb-1 text-sm text-gray-600">课程名称</label>
            <input
              v-model="searchForm.title"
              type="text"
              placeholder="请输入课程名称"
              class="border border-gray-300 rounded-md px-3 py-2 focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
            >
          </div>

          <!-- 教练选择 -->
          <div class="flex flex-col">
            <label class="mb-1 text-sm text-gray-600">教练</label>
            <select
              v-model="searchForm.coach_id"
              class="border border-gray-300 rounded-md px-3 py-2 focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
            >
              <option :value="null">
                全部教练
              </option>
              <option v-for="coach in trainerStore.trainers" :key="coach.id" :value="coach.id">
                {{ coach.name }}
              </option>
            </select>
          </div>

          <!-- 课程类型选择 -->
          <div class="flex flex-col">
            <label class="mb-1 text-sm text-gray-600">课程类型</label>
            <select
              v-model="searchForm.is_single"
              class="border border-gray-300 rounded-md px-3 py-2 focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
            >
              <option v-for="option in singleOptions" :key="option.value" :value="option.value">
                {{ option.label }}
              </option>
            </select>
          </div>

          <!-- 标签选择 -->
          <div class="flex flex-col">
            <label class="mb-1 text-sm text-gray-600">标签</label>
            <select
              v-model="searchForm.tag_ids"
              multiple
              class="border border-gray-300 rounded-md px-3 py-2 focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
            >
              <option v-for="tag in courseStore.tags" :key="tag.id" :value="tag.id">
                {{ tag.Name }}
              </option>
            </select>
          </div>
        </div>

        <!-- 搜索按钮 -->
        <div class="mt-4 flex justify-end space-x-2">
          <button
            class="border border-gray-300 rounded-md px-4 py-2 text-gray-700 transition duration-300 hover:bg-gray-100"
            @click="resetSearch"
          >
            重置
          </button>
          <button
            class="rounded-md bg-blue-600 px-4 py-2 text-white transition duration-300 hover:bg-blue-700"
            @click="fetchCourses"
          >
            搜索
          </button>
        </div>
      </div>

      <!-- 加载状态和错误提示 -->
      <div v-if="loading" class="flex items-center justify-center py-10">
        <div class="h-12 w-12 animate-spin border-b-2 border-t-2 border-blue-500 rounded-full" />
      </div>

      <div v-else-if="error" class="relative mb-4 border border-red-400 rounded bg-red-100 px-4 py-3 text-red-700">
        <span class="block sm:inline">{{ error }}</span>
      </div>

      <!-- 课程列表 -->
      <div v-else-if="courses.length > 0" class="grid grid-cols-1 gap-6 lg:grid-cols-4 md:grid-cols-2">
        <div
          v-for="course in courses"
          :key="course.id"
          class="overflow-hidden rounded-lg bg-white shadow-md transition duration-300 hover:shadow-lg"
        >
          <div class="p-4">
            <!-- 课程标题 -->
            <h3 class="mb-2 text-xl text-gray-800 font-semibold">
              {{ course.title }}
            </h3>

            <!-- 标签 -->
            <div class="mb-3 flex flex-wrap gap-1">
              <span
                v-for="(tag, index) in course.tags"
                :key="index"
                class="rounded-full bg-blue-100 px-2 py-1 text-xs text-blue-800"
              >
                {{ tag }}
              </span>
            </div>

            <!-- 课程描述 -->
            <p class="line-clamp-2 mb-3 text-sm text-gray-600">
              {{ course.description }}
            </p>

            <!-- 课程信息 -->
            <div class="text-sm text-gray-500 space-y-1">
              <p class="flex justify-between">
                <span>教练:</span>
                <span class="text-gray-700 font-medium">{{ course.coach ? course.coach.name : '未分配' }}</span>
              </p>
              <p class="flex justify-between">
                <span>课程类型:</span>
                <span class="text-gray-700 font-medium">{{ course.is_single ? '单人课程' : '团体课程' }}</span>
              </p>
              <p v-if="course.is_single === 0" class="flex justify-between">
                <span>人数上限:</span>
                <span class="text-gray-700 font-medium">{{ course.max_capacity }}人</span>
              </p>
              <p class="flex justify-between">
                <span>开始时间:</span>
                <span class="text-gray-700 font-medium">{{ formatDateTime(course.start_time) }}</span>
              </p>
              <p class="flex justify-between">
                <span>结束时间:</span>
                <span class="text-gray-700 font-medium">{{ formatDateTime(course.end_time) }}</span>
              </p>
            </div>

            <!-- 购买按钮 -->
            <button
              class="mt-4 w-full rounded-lg bg-blue-600 py-2 text-sm text-white transition duration-300 hover:bg-blue-700"
              @click="openBookingModal(course)"
            >
              购买课程
            </button>
          </div>
        </div>
      </div>

      <!-- 无数据提示 -->
      <div v-else class="flex flex-col items-center justify-center py-10">
        <svg class="mb-4 h-16 w-16 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <p class="text-lg text-gray-500">
          暂无课程数据
        </p>
      </div>

      <!-- 分页控件 -->
      <div v-if="courses.length > 0" class="mt-6 flex justify-center">
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
              <!-- 始终显示第一页 -->
              <button
                class="border border-gray-300 rounded-md px-3 py-1 transition duration-300"
                :class="[currentPage === 1 ? 'bg-blue-500 text-white' : 'bg-white text-gray-700 hover:bg-gray-100']"
                @click="currentPage = 1"
              >
                1
              </button>

              <!-- 显示省略号 -->
              <span v-if="currentPage > 3" class="border border-gray-300 rounded-md bg-white px-3 py-1 text-gray-700">
                ...
              </span>

              <!-- 显示当前页附近的页码 -->
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

              <!-- 显示省略号 -->
              <span v-if="currentPage < totalPages - 2" class="border border-gray-300 rounded-md bg-white px-3 py-1 text-gray-700">
                ...
              </span>

              <!-- 始终显示最后一页 -->
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
  </main>

  <!-- 课程购买模态框 -->
  <CourseBookingModal
    v-model="showBookingModal"
    :course="selectedCourse"
    @booking-success="handleBookingSuccess"
  />

  <AppFooter />
</template>
