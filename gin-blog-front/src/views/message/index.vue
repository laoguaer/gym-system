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

const dmRef = ref(null) // 弹幕 ref 对象
const isHide = ref(false) // 隐藏弹幕
const isLoop = ref(true) // 循环播放

// 弹幕列表
const danmus = ref([{
  avatar: 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png',
  content: '大家好，我是作者，欢迎给我点一颗 Star!',
  nickname: '赵宇锋',
}])

// 弹幕速度
const danmakuSpeed = ref(200)
// 弹幕颜色
const danmakuColors = ['#fff', '#0ff', '#0f0', '#ff0', '#f0f', '#f00']

onMounted(async () => {
  const resp = await api.getMessages()
  await nextTick()
  // 去重处理，避免相同内容的弹幕同时出现
  const uniqueMessages = resp.data.filter((item) => {
    return !danmus.value.some(existing =>
      existing.content === item.content && existing.nickname === item.nickname,
    )
  })
  danmus.value = [...danmus.value, ...uniqueMessages]
})

// 防抖标志
const isSending = ref(false)

async function send() {
  content.value = content.value.trim()
  if (!content.value) {
    window?.$message?.info('消息不能为空!')
    return
  }

  // 防抖处理
  if (isSending.value)
    return
  isSending.value = true

  try {
    const data = {
      avatar: userStore.avatar,
      nickname: userStore.nickname,
      content: content.value,
      color: danmakuColors[Math.floor(Math.random() * danmakuColors.length)],
    }

    await api.saveMessage(data)
    dmRef.value.push(data)
    content.value = ''
    window?.$message?.success('发送成功!')
  }
  catch (error) {
    console.error('发送弹幕失败:', error)
    window?.$message?.error('发送失败，请稍后重试!')
  }
  finally {
    setTimeout(() => {
      isSending.value = false
    }, 1000)
  }
}

watch(isHide, val => val ? dmRef.value.hide() : dmRef.value.show())

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
    <div class="absolute inset-x-1 bottom-[8%] z-5 mx-auto w-1/2 animate-zoom-in border-1 rounded-3xl px-1 py-5 text-center text-light shadow-2xl">
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
      <div class="mt-4 flex flex-wrap justify-center gap-2">
        <button
          class="border-1 rounded-lg px-2 py-1 text-sm transition-colors hover:bg-white/10"
          @click="dmRef.play"
        >
          <span class="i-carbon:play mr-1" />播放
        </button>
        <button
          class="border-1 rounded-lg px-2 py-1 text-sm transition-colors hover:bg-white/10"
          @click="dmRef.pause"
        >
          <span class="i-carbon:pause mr-1" />暂停
        </button>
        <button
          class="border-1 rounded-lg px-2 py-1 text-sm transition-colors hover:bg-white/10"
          @click="dmRef.stop"
        >
          <span class="i-carbon:stop mr-1" />停止
        </button>
        <button
          class="border-1 rounded-lg px-2 py-1 text-sm transition-colors hover:bg-white/10"
          :class="{ 'bg-white/20': isHide }"
          @click="isHide = !isHide"
        >
          <span :class="isHide ? 'i-carbon:view-off' : 'i-carbon:view'" class="mr-1" />
          {{ isHide ? '显示弹幕' : '隐藏弹幕' }}
        </button>
        <button
          class="border-1 rounded-lg px-2 py-1 text-sm transition-colors hover:bg-white/10"
          :class="{ 'bg-white/20': isLoop }"
          @click="isLoop = !isLoop"
        >
          <span :class="isLoop ? 'i-carbon:repeat' : 'i-carbon:no-repeat'" class="mr-1" />
          {{ isLoop ? '停止循环' : '循环播放' }}
        </button>
        <div class="flex items-center gap-1 border-1 rounded-lg px-2 py-1 text-sm">
          <span class="i-carbon:speed-fast mr-1" />
          <span>速度:</span>
          <button
            class="px-1 hover:text-#f00"
            @click="danmakuSpeed = Math.max(100, danmakuSpeed - 50)"
          >
            -
          </button>
          <span>{{ danmakuSpeed }}</span>
          <button
            class="px-1 hover:text-#0f0"
            @click="danmakuSpeed = Math.min(400, danmakuSpeed + 50)"
          >
            +
          </button>
        </div>
      </div>
    </div>
    <!-- 弹幕列表 -->
    <div class="absolute inset-0 bottom-[35%] top-[60px]">
      <vue-danmaku
        ref="dmRef"
        v-model:danmus="danmus"
        class="h-full w-full"
        use-slot
        :loop="isLoop"
        :speeds="danmakuSpeed"
        :channels="8"
        :top="5"
        :bottom="5"
        :debounce="80"
        :is-suspend="true"
      >
        <template #dm="{ danmu }">
          <div
            class="flex items-center rounded-3xl bg-#00000080 px-2 py-1 backdrop-blur-sm transition-all duration-300 hover:scale-105 hover:bg-#00000095 lg:px-4 lg:py-2"
            :style="{ color: danmu.color || '#fff' }"
          >
            <img class="h-[28px] w-[28px] rounded-full object-cover" :src="convertImgUrl(danmu.avatar)" alt="avatar">
            <span class="ml-2 text-sm font-medium"> {{ `${danmu.nickname}: ${danmu.content}` }}</span>
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
