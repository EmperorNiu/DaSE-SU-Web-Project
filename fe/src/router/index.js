import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../components/Login.vue'
import Register from '../components/Register.vue'
import Frame from '../components/Frame.vue'
import BlogIntro from '../components/Blog/BlogIntro.vue'
import WriteBlog from '../components/Blog/WriteBlog.vue'
// import Home from '../components/Home.vue'
import MyBlogs from '../components/Blog/MyBlogs.vue'
import WatchLater from '../components/Blog/WatchLater.vue'
import User from '../components/User.vue'
import Activity from '../components/Activity.vue'
import Projlist from '../components/Achievement/Projlist.vue'
import Myproject from '../components/Achievement/Myproject.vue'
import Projinfo from '../components/Achievement/Projinfo.vue'
// import Meminfo from '../components/Achievement/Meminfo.vue'
import comment from '../components/subcomponents/comment.vue'
import show1 from '../components/subcomponents/show1.vue'
import WriteProj from '../components/Achievement/WriteProj.vue'
import GraduatesGo from '../components/Alumni/GraduatesGo.vue'
import GraduatesJob from '../components/Alumni/GraduatesJob.vue'
import MyCareer from '../components/Alumni/MyCareer.vue'
import Calendar from '../components/Calendar/Calendar2.vue'
import ViewBlog from '../components/Blog/ViewBlog.vue'
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
      { path: '/testViewBlog',component: ViewBlog}, 
      { path: '/viewblog',component: ViewBlog}, 
      
      // { path: '/home', component: Home },
      { path: '/home', component: Calendar },
      { path: '/blog', component: BlogIntro },
      { path: '/writeblog', component: WriteBlog },
      { path: '/playlist', component: WatchLater },
      { path: '/myblogs', component: MyBlogs },
      { path: '/user', component: User },
      { path: '/activity', component: Activity },
      { path: '/projlist', component: Projlist },
      { path: '/myproject', component: Myproject },
      { path: '/projinfo/:id', component: Projinfo, show1 },
      // { path: '/meminfo/:id', component: Meminfo },
      { path: '/comment/:id', component: comment },
      { path: '/writeproj', component: WriteProj },
      { path: '/graduatesgo', component: GraduatesGo },
      { path: '/graduatesjob', component: GraduatesJob },
      { path: '/mycareer', component: MyCareer }
    ]
  }
]

const router = new VueRouter({
  routes
})

// 路由导航守卫
router.beforeEach((to, from, next) => {
  if (to.path === '/login') return next()
  const token = window.sessionStorage.getItem('token')
  if (!token) return next('/login')
  next()
})

export default router
