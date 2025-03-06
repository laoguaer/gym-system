<script setup>
import { onMounted, reactive, ref } from 'vue'

// 无限轮播图
import InfiniteLoading from 'v3-infinite-loading'

// Markdown => Html
import { marked } from 'marked'

import ArticleCard from './components/ArticleCard.vue'
import AuthorInfo from './components/AuthorInfo.vue'
import UserInfo from './components/UserInfo.vue'
import WebsiteInfo from './components/WebsiteInfo.vue'
import HomeBanner from './components/HomeBanner.vue'
import Announcement from './components/Announcement.vue'
import TalkingCarousel from './components/TalkingCarousel.vue'
import ChatBot from './components/ChatBot.vue'
import AppFooter from '@/components/layout/AppFooter.vue'

import api from '@/api'

import { useUserStore } from '@/store'

const articleList = ref([])
const loading = ref(false)
const userStore = useUserStore() // 获取用户状态

// 无限加载文章
const params = reactive({ page_size: 5, page_num: 1 }) // 列表加载参数

onMounted(async () => {
  loading.value = true
  // 首次加载
  const resp = await api.getArticles(params)
  articleList.value = resp.data
  // 过滤 Markdown 符号
  articleList.value.forEach(e => e.content = filterMdSymbol(e.content))
  params.page_num++
  loading.value = false
})

// 过滤 Markdown 符号: 先转 Html 再去除 Html 标签
function filterMdSymbol(md) {
  return marked(md) // 转 HTML
    .replace(/<\/?[^>]*>/g, '') // 正则去除 Html 标签
    .replace(/[|]*\n/, '')
    .replace(/&npsp;/gi, '')
}
</script>

<template>
  <!-- 首页封面图 -->
  <div class="relative">
    <HomeBanner />
    <!-- 用户信息卡片 -->
    <div v-if="userStore.userId" class="absolute right-5 top-15 z-10">
      <UserInfo />
    </div>
    <!-- 聊天机器人 -->
    <div v-if="userStore.userId" class="absolute left-5 top-60 z-10 w-80 -translate-y-1/2">
      <ChatBot />
    </div>
  </div>
  <!-- 内容 -->
  <div class="mx-auto max-w-[1230px] flex flex-col justify-center px-3" style="margin-top: calc(100vh)">
    <div class="grid grid-cols-12 gap-4">
      <!-- 左半部分 -->
      <div class="col-span-12 lg:col-span-9 space-y-5" />
      <!-- 右半部分 -->
      <div class="col-span-0 lg:col-span-3">
        <!-- sticky 实现悬浮固定效果 -->
        <div class="sticky top-5 space-y-5" />
      </div>
    </div>
  </div>
  <AppFooter />
</template>
