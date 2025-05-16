const Layout = () => import('@/layout/index.vue')

export default {
  name: 'Student',
  path: '/student',
  component: Layout,
  redirect: '/student/list',
  meta: {
    title: '学员管理',
    icon: 'mdi:account-school',
    order: 7,
    // role: ['admin'],
    // requireAuth: true,
  },
  children: [
    {
      name: 'StudentList',
      path: 'list',
      component: () => import('./index.vue'),
      meta: {
        title: '学员列表',
        icon: 'mdi:account-group',
        keepAlive: true,
      },
    },
  ],
}
