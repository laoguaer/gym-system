const Layout = () => import('@/layout/index.vue')

export default {
  name: 'Reservation',
  path: '/',
  component: Layout,
  redirect: '/reservation',
  meta: {
    order: 8,
  },
  isCatalogue: true,
  children: [
    {
      name: 'Reservation',
      path: '/reservation',
      component: () => import('./index.vue'),
      meta: {
        title: '预约管理',
        icon: 'mdi:account',
        order: 0,
      },
    },
  ],
}
