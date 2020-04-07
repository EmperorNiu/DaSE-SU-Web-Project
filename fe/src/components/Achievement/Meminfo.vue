<template>
  <div class="meminfo-container">
    <!-- 大标题 -->
    <h3 class="title">{{ meminfo.author }}</h3>
    <!-- 子标题 -->
    <!-- <p class="subtitle">
      <span>年级：{{ meminfo.add_time | dateFormat }}</span>
    </p> -->
    <hr>
    <!-- 内容区域 -->
    <div class="content" v-html="meminfo.article_html"></div>
  </div>
</template>

<script>
import { Toast } from 'mint-ui'

export default {
  data() {
    return {
      id: this.$route.params.memid, // 将 URL 地址中传递过来的 Id值，挂载到 data上，方便以后调用
      meminfo: {}
    }
  },
  created() {
    this.getMemInfo()
  },
  methods: {
    getMemInfo() {
      // 获取成员信息详情
      this.$http.get('/getmemlist/' + this.id).then(result => {
        if (result.body.status === 0) {
          this.meminfo = result.body.message[0]
        } else {
          Toast('获取成员信息失败！')
        }
      })
    }
  }
}
</script>

<style lang="scss">
.newsinfo-container {
  padding: 0 4px;
  .title {
    font-size: 16px;
    text-align: center;
    margin: 15px 0;
    color: red;
  }
  .subtitle {
    font-size: 13px;
    color: #226aff;
    display: flex;
    justify-content: space-between;
  }
  .content {
    img {
      width: 100%;
    }
  }
}
</style>
