<script setup>
import { onMounted, ref } from 'vue'
import { RouterLink } from 'vue-router'
import dayjs from 'dayjs'

const { video } = defineProps({
  video: {},
})

function formatDate(date) {
  return dayjs(date).format('YYYY-MM-DD')
}
</script>

<template>
  <!-- PC 端显示 -->
  <div class="mx-4 mt-12 hidden text-center text-light lg:block">
    <h1 class="text-3xl">
      {{ video.title }}
    </h1>
    <p class="f-c-c py-2">
      <span class="i-mdi:calendar mr-1 text-lg" /> 发布于 {{ formatDate(video.created_at) }}
      <span class="px-2">|</span>
      <span class="i-mdi:update mr-1 text-lg" /> 更新于 {{ formatDate(video.updated_at) }}
      <span class="px-2">|</span>
      <RouterLink :to="`/categories/${video.category?.id}?name=${video.category?.name}`" class="f-c-c">
        <span class="i-material-symbols:menu mr-1 text-lg" /> {{ video.category?.name }}
      </RouterLink>
    </p>
    <div class="f-c-c">
      <span class="i-mdi:play-circle-outline mr-1 text-lg" /> 播放量：{{ video.view_count || 0 }}
      <span class="px-2">|</span>
      <span class="i-ic:outline-insert-comment mr-1 text-lg" /> 评论数：{{ video.comment_count || 0 }}
    </div>
  </div>
  <!-- 移动端显示 -->
  <div class="mx-4 mt-12 block text-left text-light lg:hidden space-y-1.5">
    <h1 class="text-2xl">
      {{ video.title }}
    </h1>
    <div class="mb-1 mt-2 flex flex-wrap items-center lg:justify-center">
      <span class="i-mdi:calendar mr-1" /> 发布于 {{ formatDate(video.created_at) }}
      <span class="px-2">|</span>
      <span class="i-mdi:update mr-1" /> 更新于 {{ formatDate(video.updated_at) }}
    </div>
    <div class="flex gap-2">
      <div class="f-c-c">
        <RouterLink :to="`/categories/${video.category?.id}?name=${video.category?.name}`">
          <span class="i-material-symbols:menu mr-1" /> {{ video.category?.name }}
        </RouterLink>
      </div>
    </div>
    <div>
      <span class="i-mdi:play-circle-outline mr-1" /> 播放量：{{ video.view_count || 0 }}
      <span class="px-2">|</span>
      <span class="i-ic:outline-insert-comment mr-1" /> 评论数：{{ video.comment_count || 0 }}
    </div>
  </div>
</template>