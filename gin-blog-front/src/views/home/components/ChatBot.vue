<script setup>
import { ref } from 'vue'
import api from '@/api'

const messages = ref([])
const userInput = ref('')

async function sendMessage() {
  if (userInput.value.trim() === '')
    return

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
  }

  // 清空输入框
  userInput.value = ''
}
</script>

<template>
  <div class="chat-bot">
    <div class="messages">
      <div v-for="(message, index) in messages" :key="index" :class="message.type">
        {{ message.text }}
      </div>
    </div>
    <input v-model="userInput" placeholder="输入消息..." @keyup.enter="sendMessage">
  </div>
</template>

<style scoped>
.chat-bot {
  /* 添加样式 */
}
.user {
  /* 用户消息样式 */
}
.bot {
  /* 机器人消息样式 */
}
</style>
