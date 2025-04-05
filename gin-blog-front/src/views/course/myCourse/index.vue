<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import myCourse from './components/myCourse.vue'
import { useAppStore, useUserStore } from '@/store'

const userStore = useUserStore()
// const router = useRouter()
const isLoggedIn = ref(false)

onMounted(async () => {
  await userStore.getUserInfo()
  isLoggedIn.value = !!userStore.userId
})

function goToLogin() {
  // 设置登录弹窗标志
  const appStore = useAppStore()
  appStore.setLoginFlag(true)

  // 添加事件监听，当登录成功后刷新页面状态
  const unsubscribe = appStore.$subscribe((mutation, state) => {
    if (mutation.type === 'direct' && mutation.events.key === 'loginFlag' && !state.loginFlag) {
      // 登录弹窗关闭，可能是登录成功
      userStore.getUserInfo().then(() => {
        isLoggedIn.value = !!userStore.userId
        unsubscribe() // 取消订阅，避免内存泄漏
      })
    }
  })
}
</script>

<template>
  <div class="mx-auto px-4 py-6 container">
    <template v-if="userStore.userId">
      <myCourse />
    </template>
    <template v-else>
      <div class="flex flex-col items-center justify-center py-12">
        <div class="text-center">
          <h2 class="mb-4 text-2xl text-gray-800 font-bold">
            需要登录
          </h2>
          <p class="mb-6 text-gray-600">
            请先登录后查看您的课程信息
          </p>
          <button
            class="rounded-md bg-blue-600 px-6 py-2 text-white transition duration-300 hover:bg-blue-700"
            @click="goToLogin"
          >
            立即登录
          </button>
        </div>
      </div>
    </template>
  </div>
</template>
