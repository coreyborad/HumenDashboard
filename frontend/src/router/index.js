import Vue from 'vue'
import VueRouter from 'vue-router'
import { getToken } from '@/utils/auth'
Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/pages/Login.vue'),
  },
  {
    path: '/',
    name: 'MainLayout',
    meta: { title: 'Main', requiresAuth: true },
    component: () => import('@/pages/MainLayout.vue'),
    redirect: 'summary',
    children: [
      {
        path: 'summary',
        name: 'Summary',
        meta: { title: 'Summary', icon: 'mdi-view-dashboard' },
        component: () => import('@/pages/Summary')
      },
      {
        path: 'stocks',
        name: 'Stocks',
        meta: { title: 'Stocks', icon: 'mdi-image' },
        component: () => import('@/pages/Stocks')
      }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

router.beforeEach((to, from, next) => {
  if (
    to.matched.some(record => {
      return record.meta.requiresAuth
    }
    )) {
    const token = getToken()
    if (token !== undefined) {
      next()
    } else {
      next('/login')
    }
  } else {
    next()
  }
})

export default router
