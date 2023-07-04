import { RouteRecordRaw, createRouter, createWebHistory } from 'vue-router'

import PublicFrame from '@/frames/PublicFrame.vue'
import HomePage from '@/views/HomePage.vue'

export enum RouteName {
  Home = 'home'
}

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: PublicFrame,
    children: [
      {
        path: '/',
        name: RouteName.Home,
        component: HomePage
      }
    ]
  }
]

const routerHistory = createWebHistory(import.meta.env.BASE_URL)

const router = createRouter({
  history: routerHistory,
  routes
})

export default router
