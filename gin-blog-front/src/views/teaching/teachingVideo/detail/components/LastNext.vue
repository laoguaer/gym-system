<script setup>
import { convertImgUrl } from '@/utils'

const { lastVideo, nextVideo } = defineProps({
  lastVideo: {},
  nextVideo: {},
})
</script>

<template>
  <div class="flex flex-wrap bg-#000 text-white lg:flex-nowrap">
    <div v-if="lastVideo.id" class="group relative h-[150px] w-full overflow-hidden">
      <RouterLink :to="`/video/${lastVideo.id}`">
        <img :src="lastVideo.cover || convertImgUrl(lastVideo.img)" class="h-full w-full object-cover opacity-40 transition-600 group-hover:scale-110 group-hover:opacity-80">
        <!-- top:50%; translateY: -50%; 实现绝对定位中的垂直居中 -->
        <div class="absolute top-1/2 w-full px-10 text-left leading-6 -translate-y-1/2">
          <p> 上一个视频 </p>
          <p> {{ lastVideo.title }} </p>
        </div>
      </RouterLink>
    </div>
    <div v-if="nextVideo.id" class="group relative h-[150px] w-full overflow-hidden">
      <RouterLink :to="`/video/${nextVideo.id}`">
        <img :src="nextVideo.cover || convertImgUrl(nextVideo.img)" class="h-full w-full object-cover opacity-40 transition-600 group-hover:scale-110 group-hover:opacity-80">
        <div class="absolute top-1/2 w-full px-10 text-right leading-6 -translate-y-1/2">
          <p> 下一个视频 </p>
          <p> {{ nextVideo.title }} </p>
        </div>
      </RouterLink>
    </div>
  </div>
</template>