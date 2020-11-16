/* eslint-disable */
<template>
  <div>
    <div class="blog_container">

    <div class="blog_title">
      标题：
      <el-input placeholder="请输入文章标题" v-model="title" clearable></el-input>
      标签：
      <el-input placeholder="请输入标签,以空格分隔 " v-model="labels" clearable></el-input>
    </div>
    <mavon-editor v-model="value" :ishljs="true" @save="saveMd" />
    <!-- <div class="button">
    上传一个md文件
    <el-button round @click="upload">上传</el-button>
    <el-button round @click="submit">提交</el-button>
    </div> -->
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      md: '',
      title: '',
      labels:''
    }
  },
  methods: {
    saveMd(value, render) {
      //   console.log('this is render' + render)
      var newBlog = {
        title: this.title,
        author_id: parseInt(window.sessionStorage.getItem('user_id')),
        author_name: window.sessionStorage.getItem('user_name'),
        content_html: this.render,
        content_md: this.value,
        labels: this.labels
      }
      console.log(newBlog)

      this.$http
        .post('blog/publishBlog', newBlog, {
          headers: {
            token: window.sessionStorage.getItem('token')
          }
        })
        .then(result => {
          this.$message({
            message: '发布成功',
            type: 'success'
          })
        })
        // eslint-disable-next-line handle-callback-err
        .catch(err => {
          this.$message('发布失败')
        })
      window.sessionStorage.setItem('blog', render)
    }
  }
}
</script>

<style lang="less" scoped>
.blog_container{
width: 95%;
height: 95%;
.button{
  display: flex;
  justify-content: flex-end;
  padding: 20px 0px;
}
}
</style>
