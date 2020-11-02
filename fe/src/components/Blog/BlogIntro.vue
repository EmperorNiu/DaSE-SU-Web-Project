<template>
  <div>
    <div class="blog-intro-container">
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span>{{ blogIntro.title }}</span>
          <!-- <a href=></a> -->
          <el-button style="float: right; padding: 3px 0" type="text"
            >阅读</el-button
          >
        </div>
        <div class="text item">
          <div class="author">作者：{{ blogIntro.author_name }}</div>
          <div class="recommend">{{ blogIntro.introduction }}</div>
          <div class="statistics">
            <div class="comment-num">收藏数：{{ blogIntro.star_times }}</div>
            <div class="view-num">观看人数： {{ blogIntro.read_times }}</div>
            <div class="view-num">赞： {{ blogIntro.thumbs_times }}</div>
            <!-- <div class="view-time">预计观看时间：{{viewTime}}</div> -->
          </div>
        </div>
        <el-button type="primary" icon="el-icon-edit" circle></el-button>
      </el-card>
      <!-- <div class="blog-intro-container" v-for="blogIntro in blogs" v-bind:key="blogIntro.BlogIds"> -->
      <!-- <div class="info">
        <div class="title">{{ blogIntro.title }}</div>
      </div> -->
    </div>
  </div>
</template>

<script>
export default {
  created() {
    this.initBlogList()
    // this.blogs = this.blogs.concat(this.blogIntro)
  },
  data() {
    return {
      blogs: [],
      blogIntro: {
        blogId: 0,
        title: '深度学习优化器',
        author_name: 'Dylan Niu',
        introduction:
          '随机梯度下降 SGD、动量 Momentum、Momentum+SGD算法、牛顿动量 Nesterov、Nesterov+SGD算法、AdaGrad',
        star_times: 0,
        read_times: 0,
        thumbs_times: 0
      }
    }
  },
  methods: {
    initBlogList() {
      this.$http
        .get('blog/getBlogList', {
          headers: {
            token: window.sessionStorage.getItem('token')
          }
        })
        .then(result => {
          this.blogs = result.data.blogs
        })
    }
  }
}
</script>

<style lang="less" scoped>
.blog-intro-container {
  // width: 85%;
  // height: 150px;
  // // margin-left: 5%;
  // // background-color: rgb(247, 247, 247);
  // box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
  // padding: 15px;
  // margin-bottom: 15px;
  .text {
    font-size: 14px;
  }

  .item {
    margin-bottom: 18px;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
    content: '';
  }
  .clearfix:after {
    clear: both;
  }

  .box-card {
    width: 480px;
  }
  .recommend {
    color: #7e7e7e;
  }
}
// .title {
//   font-size: 20px;
//   font-weight: 600;
//   color: rgb(0, 0, 2);
// }
</style>
