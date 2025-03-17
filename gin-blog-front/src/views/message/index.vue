<script setup>
import { computed, nextTick, onMounted, ref, watch } from 'vue'
import { storeToRefs } from 'pinia'
import vueDanmaku from 'vue3-danmaku'

import api from '@/api'
import { convertImgUrl } from '@/utils'
import { useAppStore, useUserStore } from '@/store'

const userStore = useUserStore()
const { pageList } = storeToRefs(useAppStore())

const content = ref('')
const showBtn = ref(false)

const isLoop = ref(true) // 循环播放

// 用于跟踪当前显示的弹幕内容，避免重复显示
const displayingContents = ref(new Set())

// 弹幕列表
const danmus = ref([{
  avatar: 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png',
  content: '大家好，我是作者，欢迎给我点一颗 Star!',
  nickname: '赵宇锋',
  id: 'welcome-message', // 添加唯一标识
}])

onMounted(async () => {
  const resp = await api.getMessages()
  await nextTick()

  danmus.value = [...danmus.value, ...resp.data]

  // 初始化弹幕组件后，添加自定义事件处理
  if (dmRef.value) {
    // 监听弹幕显示事件
    dmRef.value.$el.addEventListener('danmaku-show', handleDanmakuShow)
    // 监听弹幕结束事件
    dmRef.value.$el.addEventListener('danmaku-end', handleDanmakuEnd)

    // 将初始弹幕内容添加到displayingContents集合中
    danmus.value.forEach((danmu) => {
      if (danmu && danmu.content) {
        displayingContents.value.add(danmu.content)
      }
    })
  }
})

async function send() {
  content.value = content.value.trim()
  if (!content.value) {
    window?.$message?.info('消息不能为空!')
    return
  }

  // 检查是否已经在显示相同内容的弹幕
  if (displayingContents.value.has(content.value)) {
    window?.$message?.info('相同内容的弹幕已在显示中!')
    return
  }

  const data = {
    avatar: userStore.avatar,
    nickname: userStore.nickname,
    content: content.value,
    id: `msg-${Date.now()}-${Math.random().toString(36).substr(2, 9)}`, // 添加唯一标识
  }
  await api.saveMessage(data)
  content.value = ''
}

// 根据后端配置动态获取封面
const coverStyle = computed(() => {
  const page = pageList.value.find(e => e.label === 'message')
  return page
    ? `background: url('${page?.cover}') center center / cover no-repeat;`
    : 'background: url("https://static.talkxj.com/config/83be0017d7f1a29441e33083e7706936.jpg") center center / cover no-repeat;'
})
</script>

<template>
  <div :style="coverStyle" class="banner-fade-down absolute inset-x-0 h-screen overflow-hidden">
    <!-- 弹幕输入框 -->
    <div class="absolute inset-x-1 bottom-[15%] z-5 mx-auto w-1/2 animate-zoom-in border-1 rounded-3xl px-1 py-5 text-center text-light shadow-2xl">
      <h1 class="text-2xl font-bold">
        留言板
      </h1>
      <div class="mt-6 h-9 flex justify-center lg:mt-6">
        <input
          v-model="content"
          class="w-3/4 border-1 rounded-2xl bg-transparent px-4 text-sm text-#eee outline-none"
          placeholder="说点什么吧？"
          @click.stop="showBtn = true"
          @keyup.enter="send"
        >
        <button
          v-if="showBtn"
          class="ml-3 animate-back-in-right border-1 rounded-2xl px-4"
          @click="send"
        >
          发送
        </button>
      </div>
    </div>
    <!-- 弹幕列表 -->
    <div class="absolute inset-0 bottom-[35%] top-[60px]">
      <vue-danmaku
        v-model:danmus="danmus"
        class="h-full w-full"
        use-slot
        :loop="isLoop"
        :speeds="100"
        :channels="5"
        :top="10"
        :is-suspend="true"
        debounce="3000"
        random-channel
      >
        <template #dm="{ danmu }">
          <div class="flex items-center rounded-3xl bg-#00000060 px-2 py-1 text-white lg:px-4 lg:py-2">
            <img class="h-[28px] rounded-full" :src="convertImgUrl(danmu.avatar)" alt="avatar">
            <span class="ml-2 text-sm"> {{ `${danmu.nickname} : ${danmu.content}` }}</span>
          </div>
        </template>
      </vue-danmaku>
    </div>
  </div>
</template>

<style scoped>
input::-webkit-input-placeholder {
  color: #eee;
}
</style>
