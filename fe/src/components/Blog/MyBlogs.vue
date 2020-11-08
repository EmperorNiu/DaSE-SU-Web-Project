<template>
  <div>
    <div
      class="blog-intro-container"
      v-for="blogIntro in blogs"
      v-bind:key="blogIntro.blogId"
    >
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
      </el-card>
    </div>
<!-- 分页 -->
    <!-- <div style="float:left;margin:15px">
    <el-switch v-model="value" style="margin-bottom:0px">
    </el-switch>
    <el-pagination
      :hide-on-single-page="value"
      :total="blogs.length"
      layout="prev, pager, next">
    </el-pagination>
    </div>
     -->
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
      value: false,
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
        .get('blog/getMyBlogList')
        .then(result => {
          this.blogs = result.data.blogs
          // console.log(result)
          // eslint-disable-next-line handle-callback-err
        })
        .catch(err => {
          console.log(err)
        })
    }

  }
}
</script>

<style lang="less" scoped>
.blog-intro-container {
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
    width: 95%;
    margin: 15px;

    .later-item{
      float: right; position:stastic; margin-bottom: 10px
    }
  }
  .recommend {
    color: #7e7e7e;
  }
}

</style>
