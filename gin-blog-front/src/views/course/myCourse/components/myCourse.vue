<script setup>
import { computed, onMounted, ref } from 'vue'
import CourseCalendar from './CourseCalendar.vue'
import { useAppStore, useUserStore } from '@/store'
import { formatDate } from '@/utils'
import CourseBookingModal from '@/components/modal/CourseBookingModal.vue'

const userStore = useUserStore()
const appStore = useAppStore()
const courseList = ref([])
const bookingList = ref([])
const loading = ref(true)

// 控制预约模态框
const showBookingModal = ref(false)
const selectedCourse = ref(null)

// 按教练分组的课程列表
const coursesByCoach = computed(() => {
  const groupedCourses = {}

  courseList.value.forEach((course) => {
    const coachId = course.coach_id
    const coachName = course.coach?.name || '未分配教练'

    if (!groupedCourses[coachId]) {
      groupedCourses[coachId] = {
        coach: course.coach || { id: coachId, name: coachName },
        courses: [],
      }
    }

    groupedCourses[coachId].courses.push(course)
  })

  return Object.values(groupedCourses)
})

// 格式化时间
function formatCourseTime(startTime, endTime) {
  return `${formatDate(new Date(startTime), 'YYYY-MM-DD HH:mm')} - ${formatDate(new Date(endTime), 'HH:mm')}`
}

// 获取课程类型
function getCourseType(isSingle) {
  return isSingle === 1 ? '私教课' : '团体课'
}

// 打开预约模态框
function openBookingModal(course) {
  // 检查是否登录
  if (!userStore.token) {
    appStore.setLoginFlag(true)
    return
  }
  selectedCourse.value = course
  showBookingModal.value = true
}

// 预约成功后刷新数据
async function refreshData() {
  try {
    // 获取课程列表
    const courseResult = await userStore.getMyCourseList({ user_id: userStore.userId })
    courseList.value = courseResult || []

    // 获取预约记录
    const bookingResult = await userStore.getMyBookings({ user_id: userStore.userId })
    bookingList.value = bookingResult || []
  }
  catch (error) {
    console.error('获取数据失败:', error)
  }
}

onMounted(async () => {
  try {
    // 获取课程列表
    const courseResult = await userStore.getMyCourseList({ user_id: userStore.userId })
    courseList.value = courseResult || []

    // 获取预约记录
    const bookingResult = await userStore.getMyBookings({ user_id: userStore.userId })
    bookingList.value = bookingResult || []
  }
  catch (error) {
    console.error('获取数据失败:', error)
  }
  finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="my-course-container">
    <!-- 预约模态框 -->
    <CourseBookingModal
      v-model="showBookingModal"
      :course="selectedCourse"
      @booking-success="refreshData"
    />
    <h1 class="mb-6 text-2xl text-gray-800 font-bold">
      我的课程
    </h1>

    <div v-if="loading" class="flex justify-center py-10">
      <div class="h-10 w-10 animate-spin border-b-2 border-blue-500 rounded-full" />
    </div>

    <div v-else-if="courseList.length === 0" class="py-10 text-center">
      <div class="text-lg text-gray-500">
        您还没有预约任何课程
      </div>
    </div>

    <div v-else>
      <!-- 课程表组件 -->
      <CourseCalendar v-model:booking-list="bookingList" @refresh="refreshData" />

      <!-- 课程列表 -->
      <div class="my-courses-list space-y-8">
        <h2 class="mb-4 text-xl text-gray-800 font-semibold">
          已购买的课程列表
        </h2>

        <!-- 按教练分组显示课程 -->
        <div v-for="group in coursesByCoach" :key="group.coach.id" class="coach-group mb-6">
          <div class="coach-info mb-4 flex items-center">
            <div class="mr-3 h-10 w-10 flex items-center justify-center rounded-full bg-blue-100">
              <span class="text-blue-600 font-bold">{{ group.coach.name?.substring(0, 1) || '教' }}</span>
            </div>
            <h3 class="text-lg text-gray-700 font-semibold">
              {{ group.coach.name || '未分配教练' }}
            </h3>
          </div>

          <div class="course-list flex flex-nowrap gap-4 overflow-x-auto pb-2">
            <div
              v-for="course in group.courses"
              :key="course.id"
              class="course-card w-72 flex-shrink-0 overflow-hidden rounded-lg bg-white shadow-md transition-shadow duration-300 hover:shadow-lg"
            >
              <div class="p-5">
                <div class="mb-3 flex items-start justify-between">
                  <h4 class="flex-1 truncate text-lg text-gray-800 font-semibold">
                    {{ course.title }}
                  </h4>
                  <span
                    class="rounded-full px-2 py-1 text-xs"
                    :class="course.is_single === 1 ? 'bg-purple-100 text-purple-700' : 'bg-green-100 text-green-700'"
                  >
                    {{ getCourseType(course.is_single) }}
                  </span>
                </div>

                <p class="line-clamp-2 mb-3 h-10 text-sm text-gray-600">
                  {{ course.description }}
                </p>

                <div v-if="course.is_single === 0" class="mb-2 text-sm text-gray-500">
                  <i class="fas fa-clock mr-1" />
                  {{ formatCourseTime(course.start_time, course.end_time) }}
                </div>

                <div class="flex items-center justify-between">
                  <div class="text-sm text-blue-600 font-medium">
                    购买次数: {{ course.buy_cnt }} 次
                  </div>
                  <div class="text-sm text-blue-600 font-medium">
                    已使用: {{ course.use_cnt }} 次
                  </div>
                </div>

                <!-- 预约按钮 -->
                <div class="mt-3 border-t border-gray-100 pt-3">
                  <button
                    class="w-full rounded-md bg-blue-600 px-3 py-2 text-sm text-white font-medium transition-colors hover:bg-blue-700"
                    @click="openBookingModal(course)"
                  >
                    预约课程
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.my-course-container {
  min-height: 60vh;
}

.course-card {
  border: 1px solid rgba(0, 0, 0, 0.05);
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.course-list {
  scrollbar-width: thin;
  scrollbar-color: rgba(0, 0, 0, 0.2) transparent;
}

.course-list::-webkit-scrollbar {
  height: 6px;
}

.course-list::-webkit-scrollbar-track {
  background: transparent;
}

.course-list::-webkit-scrollbar-thumb {
  background-color: rgba(0, 0, 0, 0.2);
  border-radius: 3px;
}
</style>
