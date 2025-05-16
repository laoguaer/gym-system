<script setup>
import { computed } from 'vue'
import dayjs from 'dayjs'

import { convertImgUrl } from '@/utils'

const props = defineProps({
  idx: Number,
  video: {},
})

// 判断图片放置位置 (左 or 右)
const isRightClass = computed(() => props.idx % 2 === 0
  ? 'rounded-t-xl md:order-0 md:rounded-l-xl md:rounded-tr-0'
  : 'rounded-t-xl md:order-1 md:rounded-r-xl md:rounded-tl-0')
</script>

<template>
  <div class="group h-[430px] w-full flex flex-col animate-zoom-in animate-duration-700 items-center rounded-xl bg-white shadow-md transition-600 md:h-[280px] md:flex-row hover:shadow-2xl">
    <!-- 封面图 -->
    <div :class="isRightClass" class="relative h-[230px] w-full overflow-hidden md:h-full md:w-45/100">
      <RouterLink :to="`/video/${video.id}`">
        <img class="h-full w-full transition-600 hover:scale-110" :src="video.cover">
        <!-- 播放按钮覆盖在封面上 -->
        <div class="absolute inset-0 flex items-center justify-center opacity-0 transition-300 group-hover:opacity-100">
          <span class="i-mdi-play-circle rounded-full bg-black bg-opacity-50 text-5xl text-white" />
        </div>
      </RouterLink>
    </div>
    <!-- 视频信息 -->
    <div class="my-4 w-9/10 md:w-55/100 space-y-4 md:px-10">
      <RouterLink :to="`/video/${video.id}`">
        <span class="text-2xl font-bold transition-300 group-hover:text-violet">
          {{ video.title }}
        </span>
      </RouterLink>
      <div class="flex flex-wrap text-sm color-[#858585]">
        <!-- 置顶 -->
        <span v-if="video.is_top === 1" class="flex items-center color-[#ff7242]">
          <span class="i-carbon:align-vertical-top mr-1" /> 置顶
        </span>
        <span v-if="video.is_top === 1" class="mx-1.5">|</span>
        <!-- 日期 -->
        <span class="flex items-center">
          <span class="i-mdi-calendar-month-outline mr-1" /> {{ dayjs(video.created_at).format('YYYY-MM-DD') }}
        </span>
        <span class="mx-1.5">|</span>
        <!-- 分类 -->
        <RouterLink :to="`/categories/${video.category_id}?name=${video.category?.name}`" class="flex items-center">
          <span class="i-mdi-inbox-full mr-1" /> {{ video.category?.name }}
        </RouterLink>
        <span class="mx-1.5">|</span>
        <!-- 播放量 -->
        <span class="flex items-center">
          <span class="i-mdi-play-circle-outline mr-1" /> {{ video.view_count || 0 }} 播放
        </span>
      </div>
      <div class="ell-4 text-sm leading-6">
        {{ video.description || video.content }}
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 省略文字最多 N 行 */
.ell-4 {
  display: -webkit-box;
  overflow: hidden;
  text-overflow: ellipsis;
  -webkit-line-clamp: 4;
  -webkit-box-orient: vertical;
}
</style>
