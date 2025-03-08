import { defineStore } from 'pinia'
import { convertImgUrl } from '@/utils'

import api from '@/api'

const trainers = [
  {
    id: 1,
    avatar: 'https://randomuser.me/api/portraits/men/32.jpg',
    name: '李教练',
    specialty: '力量训练、有氧运动',
    rating: 4.8,
    experience: '5年',
    intro: '国家健身教练认证，专注于帮助客户实现健康减脂和肌肉塑造目标。',
    students: 128,
    courses: 15,
  },
  {
    id: 2,
    avatar: 'https://randomuser.me/api/portraits/men/33.jpg',
    name: '王教练',
    specialty: '瑜伽、普拉提',
    rating: 4.9,
    experience: '7年',
    intro: '资深瑜伽教练，擅长帮助学员提高身体柔韧性和核心力量。',
    students: 156,
    courses: 12,
  },
  {
    id: 3,
    avatar: 'https://randomuser.me/api/portraits/men/34.jpg',
    name: '张教练',
    specialty: '功能性训练、康复训练',
    rating: 4.7,
    experience: '4年',
    intro: '专注于功能性训练和运动康复，帮助客户恢复运动机能和预防伤害。',
    students: 98,
    courses: 8,
  },
  {
    id: 4,
    avatar: 'https://randomuser.me/api/portraits/men/35.jpg',
    name: '刘教练',
    specialty: '搏击、拳击',
    rating: 4.6,
    experience: '6年',
    intro: '前职业拳击运动员，现专注于搏击和拳击训练，提供高强度的有氧训练课程。',
    students: 112,
    courses: 10,
  },
  {
    id: 5,
    avatar: 'https://randomuser.me/api/portraits/men/36.jpg',
    name: '陈教练',
    specialty: '游泳、水中有氧',
    rating: 4.9,
    experience: '8年',
    intro: '前国家游泳队队员，擅长游泳技术指导和水中有氧训练。',
    students: 145,
    courses: 14,
  },
  {
    id: 6,
    avatar: 'https://randomuser.me/api/portraits/men/37.jpg',
    name: '赵教练',
    specialty: '健美、增肌训练',
    rating: 4.8,
    experience: '5年',
    intro: '健美比赛冠军，专注于肌肉增长和体型塑造训练。',
    students: 132,
    courses: 11,
  },
  {
    id: 7,
    avatar: 'https://randomuser.me/api/portraits/men/37.jpg',
    name: 'test',
    specialty: '健美、增肌训练',
    rating: 4.8,
    experience: '5年',
    intro: '健美比赛冠军，专注于肌肉增长和体型塑造训练。',
    students: 132,
    courses: 11,
  },
]

export const useTrainerStore = defineStore('trainer', {
  state: () => ({
    trainerList: [],
    currentTrainer: null,
    loading: false,
    error: null,
  }),
  getters: {
    trainers: state => state.trainerList || [],
    trainerCount: state => state.trainerList.length,
    // 获取特定教练信息的getter
    getTrainerById: state => id => state.trainerList.find(trainer => trainer.id === id),
    // 按评分排序的教练列表
    trainersByRating: state => [...state.trainerList].sort((a, b) => b.rating - a.rating),
    // 按经验排序的教练列表
    trainersByExperience: state => [...state.trainerList].sort((a, b) => b.experience - a.experience),
    // 按学生数量排序的教练列表
    trainersByStudents: state => [...state.trainerList].sort((a, b) => b.students - a.students),
  },
  actions: {
    // 获取所有教练信息
    async getTrainerList() {
      this.loading = true
      this.error = null
      try {
        // 注意：这里假设API中有getTrainerList方法，如果实际API名称不同，需要修改
        // const resp = await api.getTrainerInfo()
        // if (resp.code === 0) {
        //   // 处理教练数据，确保图片URL正确
        //   this.trainerList = resp.data.map(trainer => ({
        //     ...trainer,
        //     avatar: trainer.avatar ? convertImgUrl(trainer.avatar) : 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png',
        //   }))
        //   return Promise.resolve(resp.data)
        // }
        // else {
        //   this.error = resp.msg || '获取教练信息失败'
        //   return Promise.reject(resp)
        // }
        this.trainerList = trainers
        return Promise.resolve(trainers)
      }
      catch (error) {
        this.error = error.message || '获取教练信息失败'
        return Promise.reject(error)
      }
      finally {
        this.loading = false
      }
    },
    // 重置状态
    resetState() {
      this.$reset()
    },
  },
})
