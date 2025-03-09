<script setup>
import { onMounted, reactive, ref } from 'vue'

// 无限轮播图
import InfiniteLoading from 'v3-infinite-loading'

import VideoCard from './components/VideoCard.vue'
import VideoCarousel from './components/VideoCarousel.vue'
import AppFooter from '@/components/layout/AppFooter.vue'

import api from '@/api'

const videoList = ref([])
const loading = ref(false)

// 无限加载视频
const params = reactive({ page_size: 5, page_num: 1 }) // 列表加载参数
async function getVideosInfinite($state) {
  if (!loading.value) {
    try {
      console.log('infnit 加载视频')
      const resp = await api.getVideos(params)
      // 加载完成
      if (!resp.data.page_data.length) {
        $state.complete()
        return
      }
      // 非首次加载, 都是往列表中添加数据
      videoList.value.push(...resp.data)
      params.page_num++
      $state.loaded()
    }
    catch (error) {
      console.error('加载视频失败:', error)
      $state.error(error)
    }
  }
}

onMounted(async () => {
  loading.value = true
  // 首次加载
  try {
    const resp = await api.getVideos(params)
    videoList.value = resp.data
    console.log('首次加载视频', resp)
    params.page_num++
  }
  catch (error) {
    console.error('加载视频失败:', error)
  }
  finally {
    loading.value = false
  }
})

function backTop() {
  window.scrollTo({ behavior: 'smooth', top: 0 })
}
</script>

<template>
  <!-- 内容 -->
  <div class="mx-auto mb-8 mt-20 max-w-[1230px] flex flex-col justify-center px-3">
    <div class="grid grid-cols-12 gap-4">
      <div />
      <!-- 左半部分 -->
      <div class="col-span-10 space-y-5">
        <!-- 视频轮播 -->
        <VideoCarousel />
        <!-- 视频列表 -->
        <div class="space-y-5">
          <VideoCard v-for="(item, idx) in videoList" :key="item.id" :video="item" :idx="idx" />
        </div>
        <!-- 无限加载 -->
        <div class="f-c-c">
          <InfiniteLoading class="mt-2 lg:mt-5" @infinite="getVideosInfinite">
            <!-- TODO: 优化界面 -->
            <template #spinner>
              <span class="animate-pulse text-xl">
                loading...
              </span>
            </template>
            <template #complete>
              <span class="flex gap-2 text-gray">
                没有更多视频啦!
                <button class="flex items-center" @click="backTop">
                  点击回到顶部 <span class="i-mdi:arrow-up-bold-box ml-1 inline-block text-xl" />
                </button>
              </span>
            </template>
          </InfiniteLoading>
        </div>
      </div>
    </div>
  </div>
  <!-- 底部 -->
  <AppFooter />
</template>
