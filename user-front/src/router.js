import { createRouter, createWebHistory } from 'vue-router'

import NProgress from 'nprogress'
import './styles/nprogress.css'

const basicRoutes = [
  {
    name: 'Home',
    path: '/',
    component: () => import('@/views/home/index.vue'),
    meta: {
      title: '首页',
    },
  },
  {
    name: 'Article',
    path: '/article/:id',
    component: () => import('@/views/teaching/article/detail/index.vue'),
  },
  {
    name: 'Archive',
    path: '/archives',
    component: () => import('@/views/discover/archive/index.vue'),
    meta: {
      title: '归档',
    },
  },
  {
    name: 'Category',
    path: '/categories',
    component: () => import('@/views/discover/category/index.vue'),
    meta: {
      title: '分类',
    },
  },
  {
    name: 'CategoryArticles',
    path: '/categories/:categoryId',
    component: () => import('@/views/teaching/article/list/index.vue'),
    meta: {
      title: '分类',
    },
  },
  {
    name: 'Tag',
    path: '/tags',
    component: () => import('@/views/discover/tag/index.vue'),
    meta: {
      title: '标签',
    },
  },
  {
    name: 'TagArticles',
    path: '/tags/:tagId',
    component: () => import('@/views/teaching/article/list/index.vue'),
    meta: {
      title: '标签',
    },
  },
  {
    name: 'About',
    path: '/about',
    component: () => import('@/views/about/index.vue'),
    meta: {
      title: '关于我',
    },
  },
  {
    name: 'MessageBoard',
    path: '/message',
    component: () => import('@/views/message/index.vue'),
    meta: {
      title: '留言',
    },
  },
  {
    name: 'User',
    path: '/user',
    component: () => import('@/views/user/index.vue'),
    meta: {
      title: '个人中心',
    },
  },
  {
    name: 'TrainerInfo',
    path: '/trainerInfo',
    component: () => import('@/views/trainer/trainerInfo/index.vue'),
  },
  {
    name: 'MyTrainer',
    path: '/myTrainer',
    component: () => import('@/views/trainer/myTrainer/index.vue'),
  },
  {
    name: 'TeachingText',
    path: '/teachingText',
    component: () => import('@/views/teaching/teachingText/index.vue'),
  },
  {
    name: 'TeachingVideo',
    path: '/teachingVideo',
    component: () => import('@/views/teaching/teachingVideo/index.vue'),
  },
  {
    name: 'Video',
    path: '/video/:id',
    component: () => import('@/views/teaching/teachingVideo/detail/index.vue'),
  },
  {
    name: 'CourseInfo',
    path: '/courseInfo',
    component: () => import('@/views/course/courseInfo/index.vue'),
  },
  {
    name: 'MyCourse',
    path: '/myCourses',
    component: () => import('@/views/course/myCourse/index.vue'),
  },
  {
    name: '404',
    path: '/404',
    component: () => import('@/views/error-page/404.vue'),
  },
  // 无匹配路由跳转 404
  {
    name: 'NotFound',
    path: '/:pathMatch(.*)*',
    redirect: '/404',
    isHidden: true,
  },
]

export const router = createRouter({
  history: createWebHistory('/'),
  routes: basicRoutes,
  scrollBehavior: () => ({ left: 0, top: 0 }),
})

router.afterEach((to) => {
  document.title = `${to.meta?.title ?? import.meta.env.VITE_APP_TITLE}`
})

NProgress.configure({ showSpinner: false })

router.beforeEach((to, from, next) => {
  NProgress.start()
  for (let i = 0; i < 5; i++) NProgress.inc()
  setTimeout(() => NProgress.done(), 300)
  next()
})
