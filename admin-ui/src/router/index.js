import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Service from '../views/service/List.vue'
import NamingService from '../views/naming/List.vue'
import Instance from '../views/naming/Instance.vue'
import KvGroup from '../views/kvgroup/List.vue'
import KvGroupConfig from '../views/kvgroup/Config.vue'

Vue.use(VueRouter)
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch((err) => err)
}

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/service',
    name: 'Service',
    component: Service
  },
  {
    path: '/naming-service',
    name: 'NamingService',
    component: NamingService
  },
  {
    path: '/naming-service/instance/:ns/:name',
    name: 'Instance',
    component: Instance
  },
  {
    path: '/kvgroup',
    name: 'KvGroup',
    component: KvGroup
  },
  {
    path: '/kvgroup/:ns/:name',
    name: 'KvGroupConfig',
    component: KvGroupConfig
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
