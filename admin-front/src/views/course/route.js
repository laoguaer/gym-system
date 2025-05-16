const Layout = () => import('@/layout/index.vue')

export default {
  name: 'Course',
  path: '/',
  component: Layout,
  redirect: '/course',
  meta: {
    order: 8,
  },
  isCatalogue: true,
  children: [
    {
      name: 'Course',
      path: '/course',
      component: () => import('./index.vue'),
      meta: {
        title: '课程管理',
        icon: 'mdi:account',
        order: 0,
      },
    },
  ],
}
