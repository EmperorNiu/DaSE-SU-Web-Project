<template>
  <div class="projinfo-container">
    <!-- 大标题 -->
    <h3 class="title">{{ projinfo.title }}</h3>
    <!-- 子标题 -->
    <p class="subtitle">
      <span>发表时间：{{ projinfo.add_time | dateFormat }}</span>
    </p>
    <hr>
    <!-- 内容区域 -->
    <div class="content" v-html="projinfo.content"></div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      id: this.$route.params.id, // 将 URL 地址中传递过来的 Id值，挂载到 data上，方便以后调用
      projinfo: {} // 项目对象
    };
  },
  created() {
    this.getProjInfo();
  },
  methods: {
    getProjInfo() {
      // 获取项目详情
      this.$http.get("api/getproj/" + this.id).then(result => {
        if (result.body.status === 0) {
          this.projinfo = result.body.message[0];
        } else {
          Toast("获取项目失败！");
        }
      });
    }
  }
};
</script>

<style lang="scss">
.projinfo-container {
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
  }
  .content {
    img {
      width: 100%;
    }
  }
}
</style>
