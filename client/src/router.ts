import { RouteRecordRaw, createRouter, createWebHistory } from 'vue-router'

import AdminFrame from '@/frames/AdminFrame.vue'
import InternalFrame from '@/frames/InternalFrame.vue'
import PublicFrame from '@/frames/PublicFrame.vue'
import AdminHomePage from '@/views/AdminHomePage.vue'
import AdminManagePage from '@/views/AdminManagePage.vue'
import ContestInfoManagePage from '@/views/ContestInfoManagePage.vue'
import ContestantManagePage from '@/views/ContestantManagePage.vue'
import HomePage from '@/views/HomePage.vue'
import ProblemManagePage from '@/views/ProblemManagePage.vue'
import ScoreManagePage from '@/views/ScoreManagePage.vue'
import VotePage from '@/views/VotePage.vue'

export enum RouteName {
  Home = 'home',
  Vote = 'vote',
  AdminHome = 'admin',
  AdminManage = 'admin-manage',
  ContestantManage = 'contestant-manage',
  ContestInfoManage = 'contest-info-manage',
  ProblemMange = 'problems-manage',
  ScoreManage = 'score-manage'
}

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: PublicFrame,
    children: [
      {
        path: '',
        name: RouteName.Home,
        component: HomePage
      }
    ]
  },
  {
    path: '/vote',
    component: InternalFrame,
    children: [
      {
        path: '',
        name: RouteName.Vote,
        component: VotePage
      }
    ]
  },
  {
    path: '/admin',
    component: AdminFrame,
    children: [
      {
        path: '',
        name: RouteName.AdminHome,
        component: AdminHomePage
      },
      {
        path: 'users/admin',
        name: RouteName.AdminManage,
        component: AdminManagePage
      },
      {
        path: 'users/contestant',
        name: RouteName.ContestantManage,
        component: ContestantManagePage
      },
      {
        path: 'contest',
        name: RouteName.ContestInfoManage,
        component: ContestInfoManagePage
      },
      {
        path: 'problems',
        name: RouteName.ProblemMange,
        component: ProblemManagePage
      },
      {
        path: 'score',
        name: RouteName.ScoreManage,
        component: ScoreManagePage
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
