<script setup>
import { computed, nextTick, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'

import BannerInfo from './components/BannerInfo.vue'
import LastNext from './components/LastNext.vue'
import Recommend from './components/Recommend.vue'

import AppFooter from '@/components/layout/AppFooter.vue'
import Comment from '@/components/comment/Comment.vue'

import { convertImgUrl } from '@/utils'
import api from '@/api'

const route = useRoute()

const data = ref({
  id: 0,
  title: '',
  content: '',
  description: '',
  created_at: '',
  updated_at: '',
  like_count: 0,
  view_count: 0,
  comment_count: 0,
  img: '',
  cover: '',
  video_url: '',
  category: {},
  next_video: {},
  last_video: {},
  recommend_videos: [],
})

// 视频内容
const videoRef = ref(null)
const loading = ref(true)

onMounted(async () => {
  try {
    // 这里需要后端提供获取视频详情的API，暂时使用文章详情API模拟
    const resp = await api.getArticleDetail(route.params.id)
    data.value = resp.data
    // 这里假设视频URL存储在content字段中，实际应根据API返回调整
    data.value.video_url = data.value.content
  }
  catch (err) {
    console.error(err)
  }
  finally {
    loading.value = false
  }
})

const styleVal = computed(() =>
  data.value.cover || data.value.img
    ? `background: url('${convertImgUrl(data.value.cover || data.value.img)}') center center / cover no-repeat;`
    : 'background: rgba(0,0,0,0.1) center center / cover no-repeat;',
)
</script>

<template>
  <!-- 头部 -->
  <div :style="styleVal" class="banner-fade-down absolute inset-x-0 top-0 h-[360px] f-c-c lg:h-[400px]">
    <BannerInfo v-if="!loading" :video="data" />
  </div>
  <!-- 主体内容 -->
  <main class="flex-1">
    <div class="card-fade-up grid grid-cols-12 mx-auto mb-3 mt-[380px] gap-4 px-1 lg:mt-[440px] lg:max-w-[1200px]">
      <!-- 视频主体 -->
      <div class="card-view col-span-12 mx-2 pt-7 lg:col-span-9 lg:mx-0">
        <!-- 视频播放器 -->
        <div ref="videoRef" class="aspect-video w-full lg:mx-10">
          <video 
            v-if="data.video_url" 
            class="h-full w-full" 
            controls 
            :src="data.video_url"
          ></video>
          <div v-else class="h-full w-full bg-gray-200 f-c-c">
            <p class="text-gray-500">视频加载中或不存在</p>
          </div>
        </div>
        
        <!-- 视频描述 -->
        <div class="my-5 lg:mx-10">
          <h2 class="text-xl font-bold">视频描述</h2>
          <p class="mt-2 text-gray-700">{{ data.description || '暂无描述' }}</p>
        </div>
        
        <!-- 评论区 -->
        <Comment 
          class="my-5 lg:mx-10" 
          :article-id="data.id" 
          :comment-count="data.comment_count" 
        />
        
        <!-- 上一个/下一个视频 -->
        <LastNext 
          :last-video="data.last_video" 
          :next-video="data.next_video" 
        />
      </div>
      
      <!-- 右侧边栏 -->
      <div class="col-span-12 space-y-5 lg:col-span-3">
        <!-- 推荐视频 -->
        <Recommend :recommend-list="data.recommend_videos" />
      </div>
    </div>
  </main>
  <!-- 底部 -->
  <AppFooter />
</template>