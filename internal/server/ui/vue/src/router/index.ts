// Composables
import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'

const routes = [
  {
    name: 'Collector',
    path: '/scripts/collector',
    component: () => import('@/components/scripts/Collector.vue'),
  },
  {
    name: 'ZkSyncStat',
    path: '/stats/zksync',
    component: () => import('@/components/stats/checker/ZkSync.vue'),
  },
  {
    name: 'Issue',
    path: '/issue/:id',
    component: () => import('@/components/issue/Issue.vue'),
    props: {
      id: String,
    }
  },
  {
    name: 'Issues',
    path: '/issues',
    component: () => import('@/components/issue/Issues.vue'),
  },
  {
    name: 'Stats',
    path: '/stats',
    component: () => import('@/components/stats/Stats.vue'),
  },
  {
    name: 'StatSwap',
    path: '/stats/swap',
    component: () => import('@/components/stats/Swaps.vue'),
  },
  {
    name: 'ProfileRelations',
    path: '/profile/relations',
    component: () => import('@/components/profile/Relations.vue'),
  },
  {
    name: 'Modules',
    path: '/modules',
    component: () => import('@/components/landing/modules/General.vue'),
  },
  {
    name: 'About',
    path: '/about',
    component: () => import('@/components/About.vue'),
  },
  {
    name: 'LandingPage',
    path: '/',
    component: () => import('@/components/landing/LandingPage.vue'),
  },
  {
    name: 'GetStarted',
    path: '/get_started',
    component: () => import('@/components/guides/GetStarted.vue'),
  },
  {
    name: 'Profiles',
    path: '/profiles',
    component: () => import('@/components/profile/Profiles.vue'),
  },
  {
    name: 'Withdraw',
    path: '/exchange_accounts',
    component: () => import('@/components/exchange.acc/ExchangeAccounts.vue'),
  },
  {
    name: 'Constructor',
    path: '/constructor',
    component: () => import('@/components/flow/PageList.vue'),
  },
  {
    name: 'SharedFlowList',
    path: '/constructor/shared',
    component: () => import('@/components/flow/SharedFlowList.vue'),
  },
  {
    name: 'Processes',
    path: '/processes',
    component: () => import('@/components/process/Processes.vue'),
  },
  {
    name: 'SharedFlow',
    path: '/shared-flow/:id',
    component: () => import('@/components/flow/SharedFlow.vue'),
    props: {
      id: String,
    }
  },
  {
    name: 'Flow',
    path: '/flow/:id',
    component: () => import('@/components/flow/OldView.vue'),
    props: {
      id: String,
    }
  },
  {
    name: 'FlowViewV2',
    path: '/flow/random/:id',
    component: () => import('@/components/flow/View.vue'),
    props: {
      id: String,
    }
  },

  {
    name: 'ViewProcess',
    path: '/process/:id',
    component: () => import('@/components/process/ViewProcess.vue'),
    props: {
      id: String,
    }
  },
  {
    name: 'Settings',
    path: '/settings',
    component: () => import('@/components/settings/Settings.vue'),
  },
  {
    name: 'CreateFlow',
    path: '/flow/create',
    component: () => import('@/components/flow/OldCreateFlow.vue'),
  },
  {
    name: 'CreateFlowV2',
    path: '/flow/random/create',
    component: () => import('@/components/flow/CreateFlow.vue'),
  },
  {
    name: 'ExchangeAccountView',
    path: '/exchange_account',
    component: () => import('@/components/exchange.acc/ExchangeAccountView.vue'),
  },
  {
    name: 'ExchangeAccountBinanceDocs',
    path: '/docs/binance_account_instruction',
    component: () => import('@/components/exchange.acc/ExchangeAccountBinanceDocs.vue'),
  },
  {
    name: 'ExchangeAccountOkexDocs',
    path: '/docs/okex_account_instruction',
    component: () => import('@/components/exchange.acc/ExchangeAccountOkexDocs.vue'),
  },
  {
    name: 'CreateProfileBatch',
    path: '/profiles/batch/create',
    component: () => import('@/components/profile/CreateProfile.vue'),
  },
  {
    name: 'Billing',
    path: '/billing',
    component: () => import('@/components/billing/Billing.vue'),
  },


] as RouteRecordRaw[]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
