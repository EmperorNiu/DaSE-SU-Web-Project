<template>
  <div>
    <div class="blog_title">
      标题：
      <el-input
        v-model="title"
        placeholder="请输入文章标题"
        clearable
      />
    </div>
    <mavon-editor
      v-model="value"
      :ishljs="true"
      @save="saveMd"
    />
  </div>
</template>

<script>
export default {
  data() {
    return {
      md: '',
      title: ''
    }
  },
  methods: {
    saveMd(value, render) {
      //   console.log('this is render' + render)
      var newBlog = {
        title: this.title,
        author_id: parseInt(window.sessionStorage.getItem('user_id')),
        author_name: window.sessionStorage.getItem('user_name'),
        content_html: render,
        content_md: value
      }
      // console.log(newBlog)
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
