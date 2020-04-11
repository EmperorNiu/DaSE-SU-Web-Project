import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/element.js'
import './assets/css/global.css'
import './assets/fonts/iconfont.css'
import 'katex/dist/katex.min.css'
import axios from 'axios'
import './lib/mui/css/mui.min.css'
import './lib/mui/css/icons-extra.css'

// md编辑器插件
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
// 右键自定义菜单插件
import contentmenu from 'v-contextmenu/dist/index.js'
import 'v-contextmenu/dist/index.css'
// 背景特效
import VueParticles from 'vue-particles'

Vue.use(mavonEditor)
Vue.use(contentmenu)
Vue.use(VueParticles)
Vue.config.productionTip = false

// 配置请求根路径
axios.defaults.baseURL = 'http://localhost:8001/api/'
axios.interceptors.request.use(config => {
  // console.log(config)
  config.headers.Authorization = window.sessionStorage.getItem('token')
  // 在最后必须 return config
  return config
})
Vue.prototype.$http = axios
axios.defaults.headers.post['Content-Type'] = 'application/json'
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
