import { createRouter, createWebHistory } from 'vue-router'
 
const routes = [
  {
    path: '/home',
    name: '/home',
    component: () => import('../components/home.vue')
  },
  {
    path: '/msg',
    name: '/msg',
    component: () => import('../components/msg.vue')
  },
  {
    path: '/login',
    name: '/login',
    component: () => import('../components/index.vue')
  },
 
] 
 
const router = createRouter({
  history: createWebHistory(),
  routes
})
 
export default router