<script setup>
import { computed, ref } from 'vue'
import UModal from '@/components/ui/UModal.vue'
import { useAppStore, useUserStore } from '@/store'
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

// 购买次数
const bookingCount = ref(1)

// 格式化日期时间
function formatDateTime(dateTimeStr) {
  const date = new Date(dateTimeStr)
  return `${date.toLocaleDateString()} ${date.toLocaleTimeString(undefined, { hour: '2-digit', minute: '2-digit' })}`
}

// 确认购买
async function confirmBooking() {
  // 检查是否登录
  if (!userStore.token) {
    // 未登录，显示登录框
    appStore.setLoginFlag(true)
    return
  }

  try {
    // 调用课程预约/购买API
    const response = await api.buyCourse({
      course_id: props.course.id,
      count: bookingCount.value,
      user_id: userStore.userId,
    })
    if (response.code !== 0) {
      throw new Error('购买失败')
    }
    // 购买成功提示
    window.$notify?.success('课程购买成功!')

    // 重置购买次数
    bookingCount.value = 1

    // 关闭模态框
    showModal.value = false

    // 触发购买成功事件
    emit('bookingSuccess')
  }
  catch (error) {
    window.$notify?.error(`购买失败: ${error.message || '未知错误'}`)
  }
}

// 取消购买
function cancelBooking() {
  bookingCount.value = 1
  showModal.value = false
}
</script>

<template>
  <UModal v-model="showModal" :width="480">
    <div class="mx-2 my-1">
      <div class="mb-4 text-xl font-bold">
        购买课程
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
          <p class="flex justify-between">
            <span>开始时间:</span>
            <span class="font-medium">{{ formatDateTime(course.start_time) }}</span>
          </p>
          <p class="flex justify-between">
            <span>结束时间:</span>
            <span class="font-medium">{{ formatDateTime(course.end_time) }}</span>
          </p>
        </div>
      </div>

      <!-- 购买次数输入 -->
      <div class="mb-4">
        <label class="mb-1 block text-sm text-gray-700 font-medium">购买次数</label>
        <input
          v-model.number="bookingCount"
          type="number"
          min="1"
          class="w-full border border-gray-300 rounded-md px-3 py-2 focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
        >
      </div>

      <!-- 按钮 -->
      <div class="mt-6 flex justify-end space-x-3">
        <button
          class="border border-gray-300 rounded-md bg-white px-4 py-2 text-sm text-gray-700 font-medium hover:bg-gray-50"
          @click="cancelBooking"
        >
          返回
        </button>
        <button
          class="rounded-md bg-blue-600 px-4 py-2 text-sm text-white font-medium hover:bg-blue-700"
          @click="confirmBooking"
        >
          确认购买
        </button>
      </div>
    </div>
  </UModal>
</template>
