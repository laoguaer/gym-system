const Layout = () => import('@/layout/index.vue')

export default {
  name: 'Coach',
  path: '/coach',
  component: Layout,
  redirect: '/coach/list',
  meta: {
    title: '教练管理',
    icon: 'mdi:human-male-board',
    order: 6,
    // role: ['admin'],
    // requireAuth: true,
  },
  children: [
    {
      name: 'CoachList',
      path: 'list',
      component: () => import('./index.vue'),
      meta: {
        title: '教练列表',
        icon: 'mdi:account-group',
        keepAlive: true,
      },
    },
  ],
}
