<script setup>
import { onMounted, reactive, ref } from 'vue'

// 无限轮播图
import InfiniteLoading from 'v3-infinite-loading'

// Markdown => Html
import { marked } from 'marked'

import ChatBot from '../../components/layout/ChatBot.vue'
import AuthorInfo from './components/AuthorInfo.vue'
import UserInfo from './components/UserInfo.vue'
import WebsiteInfo from './components/WebsiteInfo.vue'
import HomeBanner from './components/HomeBanner.vue'
import Announcement from './components/Announcement.vue'
import TrainerOverview from './components/TrainerOverview.vue'
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
    <div v-if="userStore.userId" class="absolute right-5 top-40 z-10">
      <UserInfo />
    </div>
    <!-- 聊天机器人 - 现在是全局浮动组件
    <ChatBot v-if="userStore.userId" /> -->
  </div>
  <!-- 内容 -->
  <div class="min-h-[100vh] flex flex-col px-4 py-4" style="margin-top: calc(100vh)">
    <TrainerOverview />
  </div>
  <AppFooter />
</template>
