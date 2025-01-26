<script setup>
import { nextTick, ref, watch } from 'vue'
import api from '@/api'

const messages = ref([])
const userInput = ref('')
const isSending = ref(false) // 用于跟踪是否正在发送消息
const inputRef = ref(null) // 用于引用输入框

async function sendMessage() {
  if (userInput.value.trim() === '' || isSending.value)
    return

  isSending.value = true // 开始发送消息

  // 添加用户消息到聊天
  messages.value.push({ type: 'user', text: userInput.value })

  try {
    // 调用后端API
    const response = await api.chatWithBot({ message: userInput.value })
    // 添加机器人回复到聊天
    messages.value.push({ type: 'bot', text: response.data.reply })
  }
  catch (error) {
    console.error('与机器人通信时出错:', error)
    // 添加错误提示到聊天
    messages.value.push({ type: 'bot', text: '抱歉，我遇到了一些问题，请稍后再试。' })
  }

  // 清空输入框
  userInput.value = ''
  isSending.value = false // 消息发送完成

  // 等待 DOM 更新后滚动到底部
  await nextTick(() => {
    scrollToBottom()
    inputRef.value?.focus() // 重新聚焦到输入框
  })
}

// 滚动到消息列表底部
function scrollToBottom() {
  const messagesContainer = document.querySelector('.messages')
  if (messagesContainer) {
    messagesContainer.scrollTop = messagesContainer.scrollHeight
  }
}
</script>

<template>
  <div class="chat-bot card-view animate-zoom-in animate-duration-600 space-y-2">
    <h2 class="mb-2 text-center text-lg font-bold">
      AI管家
    </h2>
    <div class="messages">
      <div v-for="(message, index) in messages" :key="index" class="message" :class="[message.type]">
        {{ message.text }}
      </div>
    </div>
    <input ref="inputRef" v-model="userInput" class="input-field" placeholder="输入消息..." @keyup.enter="sendMessage">
  </div>
</template>

<style scoped>
.chat-bot {
  border: 1px solid #ccc; /* 添加边框 */
  padding: 10px;
  max-width: 600px; /* 增加宽度 */
  margin: 0 auto;
  overflow: hidden; /* 隐藏溢出内容 */
  background-color: #f9f9f9; /* 背景颜色 */
  border-radius: 8px; /* 圆角 */
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* 阴影 */
}

.messages {
  max-height: 300px; /* 增加消息区域高度 */
  overflow-y: auto; /* 添加垂直滚动条 */
  padding: 5px; /* 内边距 */
}

.message {
  margin-bottom: 5px; /* 消息间距 */
  padding: 5px; /* 消息内边距 */
  border-radius: 4px; /* 消息圆角 */
}

.user {
  text-align: right;
  color: blue;
  background-color: #e0f7fa; /* 用户消息背景色 */
}

.bot {
  text-align: left;
  color: green;
  background-color: #e8f5e9; /* 机器人消息背景色 */
}

.input-field {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}
</style>
