<script setup>
import { onMounted, ref } from 'vue'

const sentence = ref('视频是最直观的学习方式，让知识更加生动。')

onMounted(() => {
  // 视频相关名言
  const videoSentences = [
    '视频是最直观的学习方式，让知识更加生动。',
    '一个好的教学视频胜过千言万语。',
    '学无止境，视频学习让你事半功倍。',
    '观看视频学习，提高效率，拓展视野。',
    '视频教学，让复杂知识变得简单易懂。'
  ]
  
  // 随机选择一条名言
  sentence.value = videoSentences[Math.floor(Math.random() * videoSentences.length)]
})
</script>

<template>
  <div class="card-view animate-zoom-in animate-duration-600">
    <div class="flex text-center">
      <button class="i-mdi-video-outline text-xl" />
      <div class="flex-1">
        {{ sentence }}
      </div>
      <button class="animate-arrow i-mdi-chevron-double-right text-2xl" />
    </div>
  </div>
</template>

<style scoped>
.animate-arrow {
  animation: 1s passing infinite;
}

/* 左 -> 右 闪的特效 */
@keyframes passing {
  0% { transform: translateX(-50%); opacity: 0; }
  50% { transform: translateX(0); opacity: 1; }
  100% { transform: translateX(50%); opacity: 0; }
}
</style>