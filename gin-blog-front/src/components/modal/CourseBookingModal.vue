<script setup>
import { computed, onMounted, ref, watch } from 'vue'
import UModal from '@/components/ui/UModal.vue'
import { useAppStore, useUserStore } from '@/store'
import { formatDate } from '@/utils'
import api from '@/api'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false,
  },
  course: {
    type: Object,
    default: () => ({}),
  },
})

const emit = defineEmits(['update:modelValue', 'bookingSuccess'])

const userStore = useUserStore()
const appStore = useAppStore()

// 控制模态框显示
const showModal = computed({
  get: () => props.modelValue,
  set: val => emit('update:modelValue', val),
})

// 预约相关状态
const selectedDate = ref(new Date()) // 选择的日期
const availableTimeSlots = ref([]) // 可用时间段
const selectedTimeSlot = ref(null) // 选择的时间段
const loading = ref(false) // 加载状态

// 生成未来15天的日期选项
const dateOptions = computed(() => {
  const options = []
  const today = new Date()
  today.setHours(0, 0, 0, 0)

  for (let i = 0; i < 15; i++) {
    const date = new Date(today)
    date.setDate(today.getDate() + i)
    options.push({
      date,
      label: `${formatDate(date, 'MM月DD日')} (${getWeekDay(date)})`,
    })
  }

  return options
})

// 获取星期几
function getWeekDay(date) {
  const weekDays = ['周日', '周一', '周二', '周三', '周四', '周五', '周六']
  return weekDays[date.getDay()]
}

// 选择日期后获取可用时间段
async function fetchAvailableTimeSlots() {
  if (!selectedDate.value)
    return

  loading.value = true
  try {
    // 调用接口获取可用时间段
    const response = await api.getAvailableTimeSlots({
      date: formatDate(selectedDate.value, 'YYYY-MM-DD'),
      course_id: props.course.id,
    })

    if (response.code !== 0) {
      throw new Error('获取可用时间段失败')
    }

    // 生成10点到20点的时间段
    const timeSlots = []
    for (let hour = 10; hour < 20; hour++) {
      const startTime = new Date(selectedDate.value)
      startTime.setHours(hour, 0, 0)

      const endTime = new Date(selectedDate.value)
      endTime.setHours(hour + 1, 0, 0)

      // 检查时间段是否可用
      const isAvailable = response.data.some((timeStr) => {
        const slotStart = new Date(timeStr)
        return slotStart.getHours() === hour
      })

      timeSlots.push({
        startTime,
        endTime,
        label: `${hour}:00 - ${hour + 1}:00`,
        isAvailable,
      })
    }

    availableTimeSlots.value = timeSlots
    // 重置选择的时间段
    selectedTimeSlot.value = null
  }
  catch (error) {
    console.error('获取可用时间段失败:', error)
    window.$notify?.error(`获取可用时间段失败: ${error.message || '未知错误'}`)
  }
  finally {
    loading.value = false
  }
}

// 确认预约
async function confirmBooking() {
  // 检查是否登录
  if (!userStore.token) {
    // 未登录，显示登录框
    appStore.setLoginFlag(true)
    return
  }

  // 检查是否选择了时间段
  if (!selectedTimeSlot.value) {
    window.$notify?.error('请选择预约时间段')
    return
  }

  loading.value = true
  try {
    // 调用预约API
    const response = await api.bookCourse({
      user_id: userStore.userId,
      course_id: props.course.id,
      start_time: formatDate(selectedTimeSlot.value.startTime, 'YYYY-MM-DD HH:mm:ss'),
      end_time: formatDate(selectedTimeSlot.value.endTime, 'YYYY-MM-DD HH:mm:ss'),
    })

    if (response.code !== 0) {
      throw new Error('预约失败')
    }

    // 预约成功提示
    window.$notify?.success('课程预约成功!')

    // 关闭模态框
    showModal.value = false

    // 触发预约成功事件
    emit('bookingSuccess')
  }
  catch (error) {
    window.$notify?.error(`预约失败: ${error.message || '未知错误'}`)
  }
  finally {
    loading.value = false
  }
}

// 取消预约
function cancelBooking() {
  selectedDate.value = new Date()
  selectedTimeSlot.value = null
  showModal.value = false
}

// 监听日期变化，获取可用时间段
watch(selectedDate, () => {
  fetchAvailableTimeSlots()
})

// 初始化
onMounted(() => {
  // 默认选择今天，并获取可用时间段
  selectedDate.value = new Date()
  fetchAvailableTimeSlots()
})
</script>

<template>
  <UModal v-model="showModal" :width="480">
    <div class="mx-2 my-1">
      <div class="mb-4 text-xl font-bold">
        预约课程
      </div>

      <!-- 课程信息 -->
      <div v-if="course" class="mb-4 rounded-lg bg-gray-50 p-4">
        <h3 class="mb-2 text-lg text-gray-800 font-semibold">
          {{ course.title }}
        </h3>

        <div class="mb-2 flex flex-wrap gap-1">
          <span
            v-for="(tag, index) in course.tags"
            :key="index"
            class="rounded-full bg-blue-100 px-2 py-1 text-xs text-blue-800"
          >
            {{ tag }}
          </span>
        </div>

        <div class="text-sm text-gray-600 space-y-1">
          <p class="flex justify-between">
            <span>教练:</span>
            <span class="font-medium">{{ course.coach ? course.coach.name : '未分配' }}</span>
          </p>
          <p class="flex justify-between">
            <span>课程类型:</span>
            <span class="font-medium">{{ course.is_single ? '单人课程' : '团体课程' }}</span>
          </p>
          <p v-if="course.is_single === 0" class="flex justify-between">
            <span>人数上限:</span>
            <span class="font-medium">{{ course.max_capacity }}人</span>
          </p>
        </div>
      </div>

      <!-- 日期选择 -->
      <div class="mb-4">
        <label class="mb-1 block text-sm text-gray-700 font-medium">选择日期</label>
        <select
          v-model="selectedDate"
          class="w-full border border-gray-300 rounded-md px-3 py-2 focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
        >
          <option v-for="option in dateOptions" :key="option.label" :value="option.date">
            {{ option.label }}
          </option>
        </select>
      </div>

      <!-- 时间段选择 -->
      <div class="mb-4">
        <label class="mb-1 block text-sm text-gray-700 font-medium">选择时间段</label>
        <div v-if="loading" class="flex justify-center py-4">
          <div class="h-6 w-6 animate-spin border-b-2 border-blue-500 rounded-full" />
        </div>
        <div v-else class="grid grid-cols-3 gap-2">
          <button
            v-for="slot in availableTimeSlots"
            :key="slot.label"
            class="border rounded-md px-3 py-2 text-sm font-medium transition-colors"
            :class="{
              'border-blue-500 bg-blue-50 text-blue-700': selectedTimeSlot === slot,
              'border-gray-300 bg-white text-gray-700 hover:bg-gray-50': selectedTimeSlot !== slot && slot.isAvailable,
              'border-gray-200 bg-gray-100 text-gray-400 cursor-not-allowed': !slot.isAvailable,
            }"
            :disabled="!slot.isAvailable"
            @click="selectedTimeSlot = slot"
          >
            {{ slot.label }}
          </button>
        </div>
      </div>

      <!-- 按钮 -->
      <div class="mt-6 flex justify-end space-x-3">
        <button
          class="border border-gray-300 rounded-md bg-white px-4 py-2 text-sm text-gray-700 font-medium hover:bg-gray-50"
          @click="cancelBooking"
        >
          取消
        </button>
        <button
          class="rounded-md bg-blue-600 px-4 py-2 text-sm text-white font-medium hover:bg-blue-700"
          :disabled="loading || !selectedTimeSlot"
          @click="confirmBooking"
        >
          确认预约
        </button>
      </div>
    </div>
  </UModal>
</template>
