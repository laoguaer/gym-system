const Layout = () => import('@/layout/index.vue')

export default {
  name: 'Reservation',
  path: '/reservation',
  component: Layout,
  redirect: '/reservation/list',
  meta: {
    title: '预约管理',
    icon: 'mdi:calendar-clock',
    order: 9,
    // role: ['admin'],
    // requireAuth: true,
  },
  children: [
    {
      name: 'ReservationList',
      path: 'list',
      component: () => import('./index.vue'),
      meta: {
        title: '预约列表',
        icon: 'mdi:format-list-checks',
        keepAlive: true,
      },
    },
  ],
}
