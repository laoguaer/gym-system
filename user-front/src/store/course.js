import { defineStore } from 'pinia'
import api from '@/api'

export const useCourseStore = defineStore('course', {
  state: () => ({
    courses: [],
    courseCount: 0,
    tags: [],
    loading: false,
    error: null,
  }),

  actions: {
    /**
     * 获取课程列表
     * @param {object} params 查询参数
     */
    async getCourseList(params = {}) {
      this.loading = true
      this.error = null

      try {
        const response = await api.getCourseList(params)
        this.courses = response.data.page_data || []
        this.courseCount = response.data.total || 0
        return response.data
      }
      catch (err) {
        this.error = err.message || '获取课程列表失败'
        console.error('获取课程列表失败:', err)
        return null
      }
      finally {
        this.loading = false
      }
    },

    /**
     * 获取标签列表
     */
    async getTagList() {
      try {
        const response = await api.getCourseTags()
        this.tags = response.data || []
        console.log('获取标签列表成功:', this.tags)
        return this.tags
      }
      catch (err) {
        console.error('获取标签列表失败:', err)
        return []
      }
    },

    /**
     * 重置状态
     */
    resetState() {
      this.courses = []
      this.courseCount = 0
      this.loading = false
      this.error = null
    },
  },

  persist: {
    key: 'gym-course',
    storage: localStorage,
    paths: ['coaches', 'tags'],
  },
})
