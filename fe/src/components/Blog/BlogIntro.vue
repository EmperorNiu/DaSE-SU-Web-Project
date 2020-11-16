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
          <el-button style="float: right; padding: 3px 0" type="text" v-on:click="toDetail(blogIntro.blog_id)">阅读</el-button>

        </div>
        <div class="text item">
          <div class="author">作者：{{ blogIntro.author_name }}</div>
          <div class="recommend">{{ blogIntro.labels }}</div>
          <div class="statistics">
            <div class="comment-num">收藏数：{{ blogIntro.star_times }}</div>
            <div class="view-num">观看人数： {{ blogIntro.read_times }}</div>
            <div class="view-num">赞： {{ blogIntro.thumbs_times }}</div>
            <!-- <div class="view-time">预计观看时间：{{viewTime}}</div> -->
          </div>
        </div>
        <el-tooltip class="item" effect="dark" content="稍后浏览" placement="left-start">
            <!-- TODO  -->
            <!-- 添加到稍后浏览中, 应该加入的该博客的id -->
          <el-button class="later-item" v-on:click="add(blogIntro.blog_id)" type="primary" icon="el-icon-edit" circle></el-button>
        </el-tooltip>
      </el-card>
    </div>
<!-- 分页 -->
<!-- 每页显示5篇 -->
    <div style="float:left;margin:15px">

    <el-switch v-model="value" style="margin-bottom:0px">
    </el-switch>
    <el-pagination
      :hide-on-single-page="value"
      :total="blogNums"
      layout="prev, pager, next">
    </el-pagination>
    </div>
  </div>
</template>

<script>
export default {
  created() {
    this.initBlogList()
    this.getBlogNums()
    // this.blogs = this.blogs.concat(this.blogIntro)
  },
  data() {
    return {
      value: false,
      blogs: [
      ],
      blogIntro: {
        blogId: 0,
        title: '深度学习优化器',
        author_name: 'Dylan Niu',
        labels:
          '随机梯度下降 SGD、动量 Momentum、Momentum+SGD算法、牛顿动量 Nesterov、Nesterov+SGD算法、AdaGrad',
        star_times: 0,
        read_times: 0,
        thumbs_times: 0
      },
      blogNums: 1
    }
  },
  methods: {
    toDetail(blogId){
      
    this.$router.push({path:'/viewblog',query:{id:blogId}})
    },
    initBlogList() {
      this.$http
        .get('blog/getBlogList', {
          headers: {
            token: window.sessionStorage.getItem('token')
          }
        })
        .then(result => {
          this.blogs = result.data.blogs
          // console.log(result)
          // eslint-disable-next-line handle-callback-err
        })
    },
    getBlogNums() {
      this.$http
        .get('blog/getBlogNums')
        .then(
          result => {
            this.blogNums = Math.ceil(result.data.blogNums / 5) * 10
          }
        )
    },
    saveMd(value, render) {
      //   console.log('this is render' + render)
      window.sessionStorage.setItem('blog', render)
    },
    // 添加到稍后浏览
    // 将blogID回传, 后端应该存入到该用户的某个数据库下
    // 在稍后看中从数据库中获取之前所存的博客
    add(value) {
      this.$http
        .post('blog/watchLater', { blog_id: value , userId: sessionStorage.getItem('user_id') })
        .then(result => {
          console.log(result)
        })
    }
  }
}
</script>

<style lang="less" scoped>
.blog-intro-container {
  .text {
    font-size: 14px;
    color: #161515;
  }
  .title{
    margin: 60%;
    color: #181616;
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
