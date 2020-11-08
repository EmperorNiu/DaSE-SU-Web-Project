<template>
  <div>
    <div class="proj_title">
      <el-input placeholder="项目名称" v-model="title" clearable></el-input>
    </div>
    <div class="proj_author">
      <el-input placeholder="作者" v-model="author" clearable ></el-input>
    </div>
    <p></p>
    <mavon-editor v-model="value" :ishljs="true" @save="saveMd" />
    <p></p>
    <label>请选择上传源代码:</label>
    <p></p>
    <el-upload
  class="upload-demo"
  ref="upload"
  action=""
  :on-preview="handlePreview"
  :on-remove="handleRemove"
  :file-list="fileList"
  :auto-upload="false">
  <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
  <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload">上传到服务器</el-button>
</el-upload>
  </div>
</template>

<script>
export default {
  data() {
    return {
      md: '',
      title: '',
      author: '',
      fileList: [
      ]
    }
  },
  methods: {
    submitUpload() {
      this.$refs.upload.submit()
    },
    handleRemove(file, fileList) {
      console.log(file, fileList)
    },
    handlePreview(file) {
      console.log(file)
    },
    saveMd(value, render) {
      //   console.log('this is render' + render)
      var newProj = {
        title: this.title,
        author_id: parseInt(window.sessionStorage.getItem('user_id')),
        author_name: window.sessionStorage.getItem('user_name'),
        content_html: render,
        content_md: value
      }
      console.log(newProj)
      this.$http
        .post('projlist/publishProj', newProj, {
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
      window.sessionStorage.setItem('project', render)
    }
  }
}
</script>
