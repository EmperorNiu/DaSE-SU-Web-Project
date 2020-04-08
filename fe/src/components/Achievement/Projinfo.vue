<template>
  <div>
    <div class="projinfo-container">
    <div class="info">
      <div class="title">{{projinfo.project_name}}</div>
      <div class="author">{{projinfo.project_leader}}</div>
    </div>
    <div class="recommend">{{projinfo.ProjectDescription}}</div>
    <show1></show1>
    <comment-box :id="this.id"></comment-box>
    </div>
  </div>
</template>

<script>
import comment from '../subcomponents/comment.vue'
import show1 from '../subcomponents/show1.vue'
import { Toast } from 'mint-ui'
export default {
  data() {
    return {
      id: this.$route.params.id, // 将 URL 地址中传递过来的 Id值，挂载到 data上，方便以后调用
      projinfo: {} // 项目对象
    }
  },
  created() {
    this.getProjInfo()
  },
  methods: {
    getProjInfo() {
      // 获取项目详情
      this.$http.get('/getprojlist' + this.id).then(result => {
        if (result.body.message === 'success') {
          this.projinfo = result.body.project[0]
        } else {
          Toast('获取新闻失败！')
        }
      })
    }
  },
  components: {
    // 用来注册子组件的节点
    'comment-box': comment,
    // eslint-disable-next-line vue/no-unused-components
    show1
  }
}
</script>

<style lang="less" scoped>
.projinfo-container {
  width: 85%;
  height: 150px;
  // margin-left: 5%;
  background-color: rgb(247, 247, 247);
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);
  padding: 15px;
  margin-bottom: 15px;
}
.title {
  font-size: 20px;
  font-weight: 600;
}
.recommend {
  color: #7e7e7e
}
</style>
