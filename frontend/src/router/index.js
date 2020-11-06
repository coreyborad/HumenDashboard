import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../pages/Login.vue'
import MainLayout from '../pages/MainLayout.vue'
import { getToken } from '@/utils/auth'
Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/',
    name: 'MainLayout',
    meta: { title: 'Main', requiresAuth: true },
    component: MainLayout,
    redirect: 'summary',
    children: [
      {
        path: 'summary',
        name: 'Summary',
        meta: { title: 'Summary', icon: 'summary' }
      },
      {
        path: 'my_stock',
        name: 'My Stock',
        meta: { title: 'My Stock', icon: 'stock' }
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
