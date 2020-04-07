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
Vue.use(mavonEditor)
Vue.use(contentmenu)
Vue.config.productionTip = false

// 配置请求根路径
axios.defaults.baseURL = 'http://127.0.0.1:8888/api/v1/'
axios.interceptors.request.use(config => {
  // console.log(config)
  config.headers.Authorization = window.sessionStorage.getItem('token')
  // 在最后必须 return config
  return config
})
Vue.prototype.$http = axios

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
