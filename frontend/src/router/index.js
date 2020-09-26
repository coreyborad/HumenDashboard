import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../pages/Login.vue'
import MainLayout from '../pages/MainLayout.vue'

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
    meta: { title: 'Main' },
    component: MainLayout,
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

export default router
