import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../components/Login.vue'
import Register from '../components/Register.vue'
import Frame from '../components/Frame.vue'
import Blog from '../components/Blog.vue'
import Home from '../components/Home.vue'
import User from '../components/User.vue'
import Activity from '../components/Activity.vue'
import Achievement from '../components/Achievement.vue'
Vue.use(VueRouter)

const routes = [
  { path: '/', redirect: '/login' },
  { path: '/login', component: Login },
  { path: '/register', component: Register },
  {
    path: '/frame',
    component: Frame,
    redirect: '/home',
    children: [
      { path: '/home', component: Home },
      { path: '/blog', component: Blog },
      { path: '/user', component: User },
      { path: '/activity', component: Activity },
      { path: '/achievement', component: Achievement }
    ]
  }
]

const router = new VueRouter({
  routes
})

// 路由导航守卫
// router.beforeEach((to, from, next) => {
//   if (to.path === '/login') return next()
//   const token = window.sessionStorage.getItem('token')
//   if (!token) return next('/login')
//   next()
// })

export default router
