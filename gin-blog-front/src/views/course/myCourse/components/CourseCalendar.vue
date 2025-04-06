<script setup>
import { computed, ref, watch } from 'vue'
import { formatDate } from '@/utils'

const props = defineProps({
  courseList: {
    type: Array,
    default: () => [],
  },
})

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
    date1.getFullYear() === date2.getFullYear() &&
    date1.getMonth() === date2.getMonth() &&
    date1.getDate() === date2.getDate()
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
function selectDate(day) {
  selectedDate.value = new Date(day.date)
  generateWeekDays() // 更新选中状态
}

// 获取当前选中日期的课程
const selectedDateCourses = computed(() => {
  if (!props.courseList || props.courseList.length === 0) return []
  
  return props.courseList.filter(course => {
    const courseDate = new Date(course.start_time)
    return isSameDay(courseDate, selectedDate.value)
  })
})

// 格式化时间
function formatCourseTime(startTime, endTime) {
  return `${formatDate(new Date(startTime), 'HH:mm')} - ${formatDate(new Date(endTime), 'HH:mm')}`
}

// 获取课程类型
function getCourseType(isSingle) {
  return isSingle === 1 ? '私教课' : '团体课'
}

// 监听selectedDate变化，更新日历选中状态
watch(selectedDate, () => {
  generateWeekDays()
})

// 初始化
initCurrentWeek()
</script>

<template>
  <div class="course-calendar mb-8">
    <div class="calendar-header flex items-center justify-between mb-4">
      <h2 class="text-xl text-gray-800 font-semibold">
        课程表
      </h2>
      <div class="flex space-x-2">
        <button
          class="px-3 py-1 text-sm bg-gray-100 hover:bg-gray-200 rounded transition-colors"
          @click="goToPrevWeek"
        >
          上一周
        </button>
        <button
          class="px-3 py-1 text-sm bg-blue-500 text-white hover:bg-blue-600 rounded transition-colors"
          @click="goToCurrentWeek"
        >
          本周
        </button>
        <button
          class="px-3 py-1 text-sm bg-gray-100 hover:bg-gray-200 rounded transition-colors"
          @click="goToNextWeek"
        >
          下一周
        </button>
      </div>
    </div>
    
    <div class="calendar-days grid grid-cols-7 gap-1 mb-4">
      <div
        v-for="day in currentWeek"
        :key="day.dayOfWeek"
        class="day-item flex flex-col items-center justify-center p-2 rounded cursor-pointer transition-colors"
        :class="{
          'bg-blue-500 text-white': day.isSelected,
          'bg-blue-100': day.isToday && !day.isSelected,
          'hover:bg-gray-100': !day.isSelected && !day.isToday
        }"
        @click="selectDate(day)"
      >
        <span class="text-sm font-medium">周{{ day.dayName }}</span>
        <span class="text-lg mt-1" :class="{ 'font-bold': day.isToday || day.isSelected }">
          {{ day.dayOfMonth }}
        </span>
      </div>
    </div>
    
    <div class="selected-date-courses">
      <div class="date-header mb-3 text-gray-600">
        {{ formatDate(selectedDate, 'YYYY年MM月DD日') }} 课程安排
      </div>
      
      <div v-if="selectedDateCourses.length === 0" class="empty-courses py-6 text-center text-gray-500">
        当天没有预约课程
      </div>
      
      <div v-else class="course-list space-y-3">
        <div
          v-for="course in selectedDateCourses"
          :key="course.id"
          class="course-item p-3 bg-white rounded-lg border border-gray-100 shadow-sm"
        >
          <div class="flex justify-between items-start">
            <h3 class="text-lg font-medium text-gray-800">
              {{ course.title }}
            </h3>
            <span
              class="rounded-full px-2 py-1 text-xs"
              :class="course.is_single === 1 ? 'bg-purple-100 text-purple-700' : 'bg-green-100 text-green-700'"
            >
              {{ getCourseType(course.is_single) }}
            </span>
          </div>
          
          <div class="mt-2 text-sm text-gray-600">
            <div class="flex items-center">
              <i class="fas fa-clock mr-1" />
              {{ formatCourseTime(course.start_time, course.end_time) }}
            </div>
            
            <div v-if="course.coach" class="flex items-center mt-1">
              <i class="fas fa-user mr-1" />
              教练: {{ course.coach.name }}
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