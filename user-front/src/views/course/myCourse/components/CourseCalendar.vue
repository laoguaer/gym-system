<script setup>
import { computed, ref, watch } from 'vue'
import { formatDate } from '@/utils'
import { useUserStore } from '@/store'
import api from '@/api'

const props = defineProps({
  bookingList: {
    type: Array,
    default: () => [],
  },
})

const emit = defineEmits(['update:bookingList', 'refresh', 'update:selectedDate'])

const userStore = useUserStore()

// 日期预约数据加载状态
const dayLoading = ref(false)

// 当前选中的日期
const selectedDate = ref(new Date())

// 当前显示的周
const currentWeek = ref([])

// 当前周的开始日期
const weekStartDate = ref(new Date())

// 初始化当前周
function initCurrentWeek() {
  const today = new Date()
  // 设置为当天的0点0分0秒
  today.setHours(0, 0, 0, 0)

  // 获取当前是星期几 (0-6, 0代表星期日)
  const currentDay = today.getDay()

  // 计算本周的星期一日期 (如果今天是星期日，则往前推6天，否则往前推 currentDay-1 天)
  const mondayOffset = currentDay === 0 ? -6 : -(currentDay - 1)
  const monday = new Date(today)
  monday.setDate(today.getDate() + mondayOffset)

  weekStartDate.value = new Date(monday)
  generateWeekDays()

  // 默认选中今天
  selectedDate.value = new Date(today)
}

// 生成一周的日期
function generateWeekDays() {
  const week = []
  const startDate = new Date(weekStartDate.value)

  for (let i = 0; i < 7; i++) {
    const date = new Date(startDate)
    date.setDate(startDate.getDate() + i)
    week.push({
      date,
      dayOfWeek: i,
      dayName: getDayName(i),
      dayOfMonth: date.getDate(),
      isToday: isSameDay(date, new Date()),
      isSelected: isSameDay(date, selectedDate.value),
    })
  }

  currentWeek.value = week
}

// 获取星期几的名称
function getDayName(dayIndex) {
  const dayNames = ['一', '二', '三', '四', '五', '六', '日']
  return dayNames[dayIndex]
}

// 判断两个日期是否是同一天
function isSameDay(date1, date2) {
  return (
    date1.getFullYear() === date2.getFullYear()
    && date1.getMonth() === date2.getMonth()
    && date1.getDate() === date2.getDate()
  )
}

// 切换到上一周
function goToPrevWeek() {
  const newStartDate = new Date(weekStartDate.value)
  newStartDate.setDate(newStartDate.getDate() - 7)
  weekStartDate.value = newStartDate
  generateWeekDays()
}

// 切换到下一周
function goToNextWeek() {
  const newStartDate = new Date(weekStartDate.value)
  newStartDate.setDate(newStartDate.getDate() + 7)
  weekStartDate.value = newStartDate
  generateWeekDays()
}

// 切换到当前周
function goToCurrentWeek() {
  initCurrentWeek()
}

// 选择日期
async function selectDate(day) {
  selectedDate.value = new Date(day.date)
  generateWeekDays() // 更新选中状态

  // 发出选中日期更新事件
  emit('update:selectedDate', selectedDate.value)

  // 获取选中日期的预约记录
  await fetchDayBookings(day.date)
}

// 获取指定日期的预约记录
async function fetchDayBookings(date) {
  dayLoading.value = true
  try {
    const formattedDate = formatDate(date, 'YYYY-MM-DD')
    const params = {
      UserID: userStore.userId,
      Day: formattedDate,
    }

    const result = await userStore.getUserBookingWithDay(params)
    if (result && Array.isArray(result)) {
      // 更新父组件的预约列表
      emit('update:bookingList', result)
    }
  }
  catch (error) {
    console.error('获取日期预约记录失败:', error)
  }
  finally {
    dayLoading.value = false
  }
}

// 获取当前选中日期的预约
const selectedDateCourses = computed(() => {
  if (!props.bookingList || props.bookingList.length === 0)
    return []

  return props.bookingList.filter((booking) => {
    const bookingDate = new Date(booking.StartTime)
    return isSameDay(bookingDate, selectedDate.value)
  })
})

// 格式化时间
function formatCourseTime(startTime, endTime) {
  return `${formatDate(new Date(startTime), 'HH:mm')} - ${formatDate(new Date(endTime), 'HH:mm')}`
}

// 获取预约状态
function getBookingStatus(status) {
  const statusMap = {
    0: '已预约',
    1: '已完成',
    2: '已取消',
  }
  return statusMap[status] || '未知状态'
}

// 获取状态对应的样式
function getStatusClass(status) {
  const classMap = {
    0: 'bg-yellow-100 text-yellow-700',
    1: 'bg-blue-100 text-blue-700',
    2: 'bg-gray-100 text-gray-700',
    3: 'bg-green-100 text-green-700',
  }
  return classMap[status] || 'bg-gray-100 text-gray-700'
}

// 取消预约
async function cancelBooking(booking) {
  try {
    console.log(booking)
    // 调用取消预约API
    const response = await api.cancelBooking({
      user_id: userStore.userId,
      booking_id: booking.Id,
    })

    if (response.code !== 0) {
      throw new Error('取消预约失败')
    }

    // 提示成功
    window.$notify?.success('预约已取消!')

    // 重新获取当天预约
    await fetchDayBookings(selectedDate.value)

    // 通知父组件刷新数据
    emit('refresh')
  }
  catch (error) {
    window.$notify?.error(`取消失败: ${error.message || '未知错误'}`)
  }
}

// 监听selectedDate变化，更新日历选中状态
watch(selectedDate, () => {
  generateWeekDays()
})

// 初始化
initCurrentWeek()
// 初始化完成后加载当天数据
fetchDayBookings(selectedDate.value)
</script>

<template>
  <div class="course-calendar mb-8">
    <div class="calendar-header mb-4 flex items-center justify-between">
      <h2 class="text-xl text-gray-800 font-semibold">
        课程表
      </h2>
      <div class="flex space-x-2">
        <button
          class="rounded bg-gray-100 px-3 py-1 text-sm transition-colors hover:bg-gray-200"
          @click="goToPrevWeek"
        >
          上一周
        </button>
        <button
          class="rounded bg-blue-500 px-3 py-1 text-sm text-white transition-colors hover:bg-blue-600"
          @click="goToCurrentWeek"
        >
          本周
        </button>
        <button
          class="rounded bg-gray-100 px-3 py-1 text-sm transition-colors hover:bg-gray-200"
          @click="goToNextWeek"
        >
          下一周
        </button>
      </div>
    </div>

    <div class="calendar-days grid grid-cols-7 mb-4 gap-1">
      <div
        v-for="day in currentWeek"
        :key="day.dayOfWeek"
        class="day-item flex flex-col cursor-pointer items-center justify-center rounded p-2 transition-colors"
        :class="{
          'bg-blue-500 text-white': day.isSelected,
          'bg-blue-100': day.isToday && !day.isSelected,
          'hover:bg-gray-100': !day.isSelected && !day.isToday,
        }"
        @click="selectDate(day)"
      >
        <span class="text-sm font-medium">周{{ day.dayName }}</span>
        <span class="mt-1 text-lg" :class="{ 'font-bold': day.isToday || day.isSelected }">
          {{ day.dayOfMonth }}
        </span>
      </div>
    </div>

    <div class="selected-date-courses">
      <div class="date-header mb-3 flex items-center text-gray-600">
        <span>{{ formatDate(selectedDate, 'YYYY年MM月DD日') }} 课程安排</span>
        <div v-if="dayLoading" class="ml-2 h-4 w-4 animate-spin border-b-2 border-blue-500 rounded-full" />
      </div>

      <div v-if="selectedDateCourses.length === 0" class="empty-courses py-6 text-center text-gray-500">
        当天没有预约记录
      </div>

      <div v-else class="course-list space-y-3">
        <div
          v-for="booking in selectedDateCourses"
          :key="booking.Id"
          class="course-item border border-gray-100 rounded-lg bg-white p-3 shadow-sm"
        >
          <div class="flex items-start justify-between">
            <h3 class="text-lg text-gray-800 font-medium">
              课程名: {{ booking.CourseTitle }}
            </h3>
            <span
              class="rounded-full px-2 py-1 text-xs"
              :class="getStatusClass(booking.Status)"
            >
              {{ getBookingStatus(booking.Status) }}
            </span>
          </div>

          <div class="mt-2 text-sm text-gray-600">
            <div class="flex items-center">
              <i class="fas fa-clock mr-1" />
              {{ formatCourseTime(booking.StartTime, booking.EndTime) }}
            </div>

            <div class="mt-1 flex items-center">
              <i class="fas fa-calendar-check mr-1" />
              预约时间: {{ formatDate(booking.CreatedAt, 'YYYY-MM-DD HH:mm') }}
            </div>

            <!-- 取消预约按钮 -->
            <div v-if="booking.Status === 0" class="mt-3 flex justify-end">
              <button
                class="rounded-md bg-red-500 px-3 py-1 text-xs text-white font-medium transition-colors hover:bg-red-600"
                @click="cancelBooking(booking)"
              >
                取消预约
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.day-item {
  height: 70px;
}

.course-item {
  transition: transform 0.2s;
}

.course-item:hover {
  transform: translateY(-2px);
}
</style>
