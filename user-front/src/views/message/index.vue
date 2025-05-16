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
})

async function send() {
  content.value = content.value.trim()
  if (!content.value) {
    window?.$message?.info('消息不能为空!')
    return
  }

  const data = {
    avatar: userStore.avatar,
    nickname: userStore.nickname,
    content: content.value,
  }
  await api.saveMessage(data)
  danmus.value.push({
    avatar: convertImgUrl(userStore.avatar),
    nickname: userStore.nickname,
    content: content.value,
  })
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
    <div class="absolute inset-x-1 bottom-[15%] z-5 mx-auto w-1/2 animate-zoom-in border-1 rounded-3xl px-1 py-5 text-center text-light color-red shadow-2xl">
      <h1 class="text-2xl font-bold">
        留言板
      </h1>
      <div class="mt-6 h-9 flex justify-center lg:mt-6">
        <input
          v-model="content"
          class="w-3/4 border-0 rounded-2xl bg-transparent px-4 py-2 text-sm text-white shadow-inner outline-none ring-1 ring-gray-300/30 ring-inset transition-all duration-300 placeholder:text-gray-300 hover:shadow-md focus:ring-2 focus:ring-emerald-400/70"
          placeholder="说点什么吧？"
          @click.stop="showBtn = true"
          @keyup.enter="send"
        >
        <button
          v-if="userStore.userId && showBtn"
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
        :debounce="3000"
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
  color: rgba(238, 238, 238, 0.8);
  transition: color 0.3s ease;
}

input:hover::-webkit-input-placeholder {
  color: rgba(238, 238, 238, 1);
}

input:focus {
  box-shadow: 0 0 10px rgba(255, 255, 255, 0.2);
}
</style>
