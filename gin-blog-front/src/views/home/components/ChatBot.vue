<script setup>
import { nextTick, onMounted, ref, watch } from 'vue'
import api from '@/api'
import { useAppStore, useUserStore } from '@/store'

const messages = ref([])
const userInput = ref('')
const isSending = ref(false) // 用于跟踪是否正在发送消息
const inputRef = ref(null) // 用于引用输入框

const messageMap = {
  login: '<button class="login-btn" onclick="appStore.setLoginFlag(true)">登录</button>',
  welcome: '欢迎来到AI管家，有什么可以帮助您的吗？',
  error: '抱歉，发生了错误，请稍后再试。',
  // 添加更多消息类型
}

// 初始欢迎消息
onMounted(() => {
  messages.value.push({ type: 'bot', text: getSystemMessage('welcome') })
})

function getSystemMessage(key) {
  return messageMap[key] || '未知消息类型' // 默认返回未知消息
}

const userStore = useUserStore() // 获取用户状态
async function sendMessage() {
  try {
    if (userInput.value.trim() === '' || isSending.value)
      return
    if (!userStore.userId) {
    // 如果用户未登录, 发送一个系统的消息
      messages.value.push({ type: 'bot', text: '使用客服功能请先点击下面按钮进行登录哦' })
      messages.value.push({ type: 'system', text: getSystemMessage('login') })
      return
    }
    const input = userInput.value
    userInput.value = ''
    isSending.value = true // 开始发送消息

    // 添加用户消息到聊天
    messages.value.push({ type: 'user', text: input })

    // 调用后端API
    const response = await api.chatWithBot({ chat_string: input })
    console.log('API 响应:', response)
    const resp = JSON.parse(response) // 确保这里的response是有效的JSON
    const output = resp.data // 直接使用data字段
    // 添加机器人回复到聊天
    messages.value.push({ type: 'bot', text: output }) // 直接使用output
  }
  catch (error) {
    console.error('与机器人通信时出错:', error)
    messages.value.push({ type: 'system', text: getSystemMessage('error') })
  }
  finally {
    // 等待 DOM 更新后滚动到底部
    isSending.value = false // 消息发送完成
    await nextTick(() => {
      scrollToBottom()
      inputRef.value?.focus() // 重新聚焦到输入框
    })
  }
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
  <div class="chat-bot card-view animate-zoom-in animate-duration-600 lg:block space-y-2">
    <h2 class="chat-title">
      <i class="chat-icon">💬</i> AI管家
    </h2>
    <div class="messages">
      <div v-for="(message, index) in messages" :key="index" class="message" :class="[message.type]">
        <div v-if="message.type === 'system'" class="message-content" v-html="message.text" />
        <div v-else class="message-content">
          {{ message.text }}
        </div>
      </div>
      <div v-if="isSending" class="typing-indicator">
        <span />
        <span />
        <span />
      </div>
    </div>
    <div class="input-container">
      <input
        ref="inputRef"
        v-model="userInput"
        class="input-field"
        placeholder="输入消息..."
        @keyup.enter="sendMessage"
      >
      <button class="send-button" :disabled="isSending || !userInput.trim()" @click="sendMessage">
        <i class="send-icon">➤</i>
      </button>
    </div>
  </div>
</template>

<style scoped>
.chat-bot {
  border: none;
  padding: 18px;
  max-width: 600px;
  margin: 0 auto;
  overflow: hidden;
  background-color: #ffffff;
  border-radius: 16px;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.12);
  transition: all 0.3s ease;
}

.chat-bot:hover {
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.15);
  transform: translateY(-2px);
}

.chat-title {
  margin-bottom: 18px;
  text-align: center;
  font-size: 1.3rem;
  font-weight: 600;
  color: #333;
  padding-bottom: 12px;
  border-bottom: 1px solid #f0f0f0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.chat-icon {
  margin-right: 10px;
  font-size: 1.4rem;
}

.messages {
  max-height: 320px;
  overflow-y: auto;
  padding: 15px 10px;
  margin-bottom: 15px;
  scroll-behavior: smooth;
  background-color: #f8f9fa;
  border-radius: 12px;
  box-shadow: inset 0 0 10px rgba(0, 0, 0, 0.03);
}

.message {
  margin-bottom: 14px;
  max-width: 85%;
  animation: fadeIn 0.4s ease;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(12px); }
  to { opacity: 1; transform: translateY(0); }
}

.message-content {
  padding: 12px 16px;
  border-radius: 18px;
  word-break: break-word;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.08);
  line-height: 1.5;
  font-size: 0.95rem;
}

.system {
  margin: 18px auto;
  max-width: 90%;
  text-align: center;
}

.system .message-content {
  display: inline-block;
  color: #ffffff;
  background: linear-gradient(135deg, #6366f1, #8b5cf6);
  border-radius: 20px;
  padding: 10px 18px;
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.2);
}

.login-btn {
  background: linear-gradient(135deg, #4ade80, #22c55e);
  color: white;
  border: none;
  padding: 6px 16px;
  border-radius: 6px;
  cursor: pointer;
  font-weight: bold;
  transition: all 0.2s ease;
  box-shadow: 0 2px 8px rgba(34, 197, 94, 0.2);
}

.login-btn:hover {
  background: linear-gradient(135deg, #22c55e, #16a34a);
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(34, 197, 94, 0.3);
}

.user {
  margin-left: auto;
  text-align: right;
}

.user .message-content {
  background: linear-gradient(135deg, #3b82f6, #2563eb);
  color: white;
  border-top-right-radius: 4px;
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.2);
}

.bot {
  margin-right: auto;
  text-align: left;
}

.bot .message-content {
  background-color: white;
  color: #333;
  border-top-left-radius: 4px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  border-left: 3px solid #e5e7eb;
}

.input-container {
  display: flex;
  position: relative;
  border-radius: 30px;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;
  border: 1px solid #e5e7eb;
}

.input-container:focus-within {
  box-shadow: 0 4px 15px rgba(59, 130, 246, 0.25);
  border-color: #3b82f6;
}

.input-field {
  flex: 1;
  padding: 14px 18px;
  border: none;
  outline: none;
  font-size: 0.95rem;
  background-color: #ffffff;
  transition: all 0.3s ease;
}

.input-field::placeholder {
  color: #9ca3af;
}

.send-button {
  background: linear-gradient(135deg, #3b82f6, #2563eb);
  color: white;
  border: none;
  width: 45px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s ease;
}

.send-button:hover {
  background: linear-gradient(135deg, #2563eb, #1d4ed8);
}

.send-button:disabled {
  background: linear-gradient(135deg, #bfdbfe, #93c5fd);
  cursor: not-allowed;
}

.send-icon {
  font-size: 1rem;
}

/* 打字指示器动画 */
.typing-indicator {
  display: flex;
  padding: 10px 14px;
  width: 50px;
  border-radius: 20px;
  background-color: white;
  margin-bottom: 10px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  border-left: 3px solid #e5e7eb;
}

.typing-indicator span {
  height: 8px;
  width: 8px;
  margin: 0 2px;
  background-color: #6b7280;
  border-radius: 50%;
  display: inline-block;
  animation: typing 1.4s infinite ease-in-out both;
}

.typing-indicator span:nth-child(1) {
  animation-delay: 0s;
}

.typing-indicator span:nth-child(2) {
  animation-delay: 0.2s;
}

.typing-indicator span:nth-child(3) {
  animation-delay: 0.4s;
}

@keyframes typing {
  0% { transform: scale(1); opacity: 0.7; }
  50% { transform: scale(1.5); opacity: 1; }
  100% { transform: scale(1); opacity: 0.7; }
}

/* 滚动条样式 */
.messages::-webkit-scrollbar {
  width: 6px;
}

.messages::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 10px;
}

.messages::-webkit-scrollbar-thumb {
  background: #c5c5c5;
  border-radius: 10px;
}

.messages::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

/* 响应式调整 */
@media (max-width: 768px) {
  .chat-bot {
    max-width: 100%;
    border-radius: 12px;
    margin: 0 10px;
  }

  .message-content {
    font-size: 0.9rem;
    padding: 10px 14px;
  }
}
</style>
