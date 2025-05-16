<script setup>
import { computed, onMounted, ref } from 'vue'
import { storeToRefs } from 'pinia'

// 导入教练相关的store
import { useTrainerStore } from '@/store/trainer'

const trainerStore = useTrainerStore()
const { loading, error } = storeToRefs(trainerStore)

// 假设这是用户已选择的教练ID列表，实际应用中可能需要从用户数据中获取
const myTrainerIds = ref([1, 3, 5]) // 示例数据，实际应从用户数据中获取

// 获取我的教练列表
const myTrainers = computed(() => {
  return trainerStore.trainers.filter(trainer => myTrainerIds.value.includes(trainer.id))
})

// 课程安排数据（示例数据）
const schedules = ref([
  {
    trainerId: 1,
    schedules: [
      { day: '周一', time: '10:00-11:30', course: '力量训练' },
      { day: '周三', time: '14:00-15:30', course: '有氧运动' },
    ],
  },
  {
    trainerId: 3,
    schedules: [
      { day: '周二', time: '09:00-10:30', course: '功能性训练' },
      { day: '周五', time: '16:00-17:30', course: '康复训练' },
    ],
  },
  {
    trainerId: 5,
    schedules: [
      { day: '周四', time: '08:00-09:30', course: '游泳训练' },
      { day: '周六', time: '13:00-14:30', course: '水中有氧' },
    ],
  },
])

// 训练记录数据（示例数据）
const trainingRecords = ref([
  {
    trainerId: 1,
    records: [
      { date: '2023-10-15', duration: '90分钟', content: '完成力量训练，重点锻炼上肢肌群' },
      { date: '2023-10-12', duration: '60分钟', content: '完成有氧训练，心率控制良好' },
    ],
  },
  {
    trainerId: 3,
    records: [
      { date: '2023-10-14', duration: '60分钟', content: '功能性训练，改善身体姿态' },
      { date: '2023-10-10', duration: '45分钟', content: '康复训练，缓解肩部疼痛' },
    ],
  },
  {
    trainerId: 5,
    records: [
      { date: '2023-10-13', duration: '60分钟', content: '游泳训练，提高游泳技巧' },
      { date: '2023-10-09', duration: '45分钟', content: '水中有氧，全身协调性训练' },
    ],
  },
])

// 获取特定教练的课程安排
function getTrainerSchedules(trainerId) {
  const trainerSchedule = schedules.value.find(s => s.trainerId === trainerId)
  return trainerSchedule ? trainerSchedule.schedules : []
}

// 获取特定教练的训练记录
function getTrainerRecords(trainerId) {
  const trainerRecord = trainingRecords.value.find(r => r.trainerId === trainerId)
  return trainerRecord ? trainerRecord.records : []
}

// 当前展开的教练ID
const expandedTrainerId = ref(null)

// 切换展开状态
function toggleExpand(trainerId) {
  expandedTrainerId.value = expandedTrainerId.value === trainerId ? null : trainerId
}

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
      <h2 class="mb-4 text-2xl text-gray-800 font-bold">
        我的教练
      </h2>

      <!-- 没有教练时的提示 -->
      <div v-if="myTrainers.length === 0" class="py-10 text-center">
        <p class="text-gray-500">
          您还没有选择任何教练
        </p>
        <button class="mt-4 rounded-lg bg-blue-600 px-4 py-2 text-white transition duration-300 hover:bg-blue-700">
          浏览教练
        </button>
      </div>

      <!-- 我的教练列表 -->
      <div v-else class="space-y-4">
        <div v-for="trainer in myTrainers" :key="trainer.id" class="overflow-hidden border border-gray-200 rounded-lg">
          <!-- 教练基本信息 -->
          <div class="bg-white p-4">
            <div class="flex items-start justify-between">
              <div class="flex items-center">
                <img :src="trainer.avatar" class="mr-4 h-16 w-16 rounded-full" alt="trainer avatar">
                <div>
                  <h3 class="text-lg font-semibold">
                    {{ trainer.name }}
                  </h3>
                  <p class="text-sm text-gray-600">
                    {{ trainer.specialty }}
                  </p>
                  <div class="mt-1 flex space-x-4">
                    <span class="text-sm text-gray-600">评分: {{ trainer.rating }}/5.0</span>
                    <span class="text-sm text-gray-600">教龄: {{ trainer.experience }}</span>
                  </div>
                </div>
              </div>
              <button
                class="rounded-md bg-blue-100 px-3 py-1 text-blue-700 transition duration-300 hover:bg-blue-200"
                @click="toggleExpand(trainer.id)"
              >
                {{ expandedTrainerId === trainer.id ? '收起' : '详情' }}
              </button>
            </div>
          </div>

          <!-- 展开的详细信息 -->
          <div v-if="expandedTrainerId === trainer.id" class="border-t border-gray-200 bg-gray-50 p-4">
            <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
              <!-- 课程安排 -->
              <div class="rounded-lg bg-white p-3 shadow-sm">
                <h4 class="mb-2 text-gray-800 font-medium">
                  课程安排
                </h4>
                <div v-if="getTrainerSchedules(trainer.id).length > 0" class="space-y-2">
                  <div v-for="(schedule, index) in getTrainerSchedules(trainer.id)" :key="index" class="flex justify-between border-b border-gray-100 pb-1 text-sm last:border-0">
                    <span class="text-gray-600">{{ schedule.day }} {{ schedule.time }}</span>
                    <span class="font-medium">{{ schedule.course }}</span>
                  </div>
                </div>
                <p v-else class="text-sm text-gray-500">
                  暂无课程安排
                </p>
              </div>

              <!-- 训练记录 -->
              <div class="rounded-lg bg-white p-3 shadow-sm">
                <h4 class="mb-2 text-gray-800 font-medium">
                  训练记录
                </h4>
                <div v-if="getTrainerRecords(trainer.id).length > 0" class="space-y-2">
                  <div v-for="(record, index) in getTrainerRecords(trainer.id)" :key="index" class="border-b border-gray-100 pb-2 text-sm last:border-0">
                    <div class="flex justify-between">
                      <span class="text-gray-600">{{ record.date }}</span>
                      <span class="text-gray-600">{{ record.duration }}</span>
                    </div>
                    <p class="mt-1 text-gray-700">
                      {{ record.content }}
                    </p>
                  </div>
                </div>
                <p v-else class="text-sm text-gray-500">
                  暂无训练记录
                </p>
              </div>
            </div>

            <!-- 操作按钮 -->
            <div class="mt-3 flex justify-end space-x-2">
              <button
                class="rounded-md bg-green-600 px-3 py-1 text-sm text-white transition duration-300 hover:bg-green-700"
                @click="$router.push(`/courseInfo?coach_id=${trainer.id}`)"
              >
                预约课程
              </button>
              <button class="rounded-md bg-red-100 px-3 py-1 text-sm text-red-700 transition duration-300 hover:bg-red-200">
                取消关注
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
