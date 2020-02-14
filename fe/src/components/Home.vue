<template>
  <div>
    <div v-html="ss" v-contextmenu:contextmenu id="ss"></div>
    <el-drawer
      title="我是标题"
      :visible.sync="drawer"
      :direction="direction"
      :before-close="handleClose"
    >
      <span>我来啦!</span>
      <el-input type="textarea" :rows="2" placeholder="请输入内容" v-model="textarea"></el-input>
    </el-drawer>

    <v-contextmenu ref="contextmenu">
      <v-contextmenu-item @click="handleClick">菜单1</v-contextmenu-item>
      <v-contextmenu-item @click="handleClick">菜单2</v-contextmenu-item>
      <v-contextmenu-item @click="handleClick">菜单3</v-contextmenu-item>
    </v-contextmenu>
  </div>
</template>

<script>
import Highlighter from 'web-highlighter'
export default {
  mounted() {
    // this.hl = window.sessionStorage.getItem('hl')
    var highlighter = new Highlighter({
      $root: document.getElementById('ss')
    })
    // if (this.hl != null) {
    //   console.log(this.hl)
    //   for (let index = 0; index < this.hl.length; index++) {
    //     const s = this.hl[index]
    //     highlighter.fromStore(s.startMeta, s.endMeta, s.id, s.text)
    //   }
    // }
    // highlighter.on(Highlighter.event.CREATE, ({ sources }) => {
    //   // var sources = { data: data, hs: inst, e: e }
    //   this.save(sources)
    // })
    // highlighter.on(Highlighter.event.CLICK, (data, inst, e) => {
    //   this.drawer = true
    //   highlighter.getDoms(data.id)
    //   var temp = document.getElementById(data.id)
    //   console.log(data.id)
    //   console.log(highlighter.getDoms(data.id))
    //   inst.remove(data.id)
    // })
    highlighter.run()
    // var blog = window.sessionStorage.getItem('blog')
    // this.ss = blog
  },
  destroyed() {
    // highlighter.stop()
  },
  data() {
    return {
      // highlighter: {}
      hl: {},
      drawer: false,
      direction: 'rtl',
      textarea: '',
      ss: '<h2>这是一个h2标签</h2>',
      highId: '',
      comment: [
        { hlId: 0, writer: 'test', content: '这是一个评论或者订正呀！！' }
      ]
    }
  },
  methods: {
    handleClose(done) {
      this.drawer = false
    },
    storeToJson() {
      const store = window.sessionStorage.getItem('hl')
      let sources
      try {
        sources = JSON.parse(store) || []
      } catch (e) {
        sources = []
      }
      return sources
    },
    save(data) {
      console.log(data)
      const stores = this.storeToJson()
      const map = {}
      stores.forEach((store, idx) => (map[store.hs.id] = idx))

      if (!Array.isArray(data)) {
        data = [data]
      }

      data.forEach(store => {
        // update
        if (map[store.id] !== undefined) {
          stores[map[store.id]] = store
        } else {
          stores.push(store)
        }
      })
      console.log(stores)
      window.sessionStorage.setItem('hl', stores)
    },
    handleClick(vm, event) {
      alert(`「${vm.$slots.default[0].text}」被点击啦！`)
    }
  }
}
</script>
