<script setup>
import { useRoute, useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'

import { Icon } from '@iconify/vue'
import UDrawer from '@/components/ui/UDrawer.vue'

import { useAppStore, useUserStore } from '@/store'
import api from '@/api'

const { collapsed, blogConfig } = storeToRefs(useAppStore())

const [route, router] = [useRoute(), useRouter()]
const [userStore, appStore] = [useUserStore(), useAppStore()]

const menuOptions = [
  { text: '首页', icon: 'mdi:home', path: '/' },
  { text: '关于', icon: 'mdi:information-outline', path: '/about' },
  { text: '客户见证', icon: 'mdi:forum', path: '/message' },
  { text: '教练团队', icon: 'mdi:account-group', path: '/trainerInfo' },
  { text: '我的教练', icon: 'mdi:menu', path: '/myTrainer' },
  { text: '课程信息预览', icon: 'mdi:view-gallery', path: '/courseInfo' },
  { text: '我的课程', icon: 'mdi:menu', path: '/myCourses' },
  { text: '健身文章', icon: 'mdi:view-list', path: '/teachingText' },
  { text: '健身教学短片', icon: 'mdi:menu', path: '/teachingVideo' },
]

async function logout() {
  await userStore.logout()
  if (route.name === 'User') { // 如果在个人信息页登出则回到首页
    router.push('/')
  }
  window.$notify?.success('退出登录成功！')
}
</script>

<template>
  <UDrawer v-model="collapsed" placement="right" :width="250">
    <div class="mx-5">
      <div class="pt-4 text-center space-y-3">
        <div class="flex justify-center">
          <img :src="blogConfig.website_avatar" class="h-20 rounded-full" alt="作者头像">
        </div>
      </div>
      <!-- 分隔线 -->
      <hr class="my-4 border-2 border-color-#d2ebfd border-dashed">
      <!-- 菜单 -->
      <div v-for="item of menuOptions" :key="item.text" class="m-2 p-1">
        <RouterLink :to="item.path" class="flex items-center" @click="appStore.setCollapsed(false)">
          <Icon :icon="item.icon" class="text-lg" />
          <span class="ml-5"> {{ item.text }} </span>
        </RouterLink>
      </div>
      <!-- 登录 -->
      <div>
        <template v-if="!userStore.userId">
          <div class="m-2 flex items-center p-1" @click="appStore.setLoginFlag(true)">
            <Icon icon="ph:user-bold" class="text-lg" />
            <span class="ml-5"> 登录 </span>
          </div>
        </template>
        <template v-else>
          <RouterLink to="/user">
            <div class="m-2 flex items-center p-1" @click="appStore.setCollapsed(false)">
              <Icon icon="mdi:account-circle" class="text-lg" />
              <span class="ml-5"> 个人中心 </span>
            </div>
          </RouterLink>
          <div class="m-2 flex items-center p-1" @click="logout">
            <Icon icon="mdi:logout" class="text-lg" />
            <span class="ml-5"> 退出登录 </span>
          </div>
        </template>
      </div>
    </div>
  </UDrawer>
</template>
