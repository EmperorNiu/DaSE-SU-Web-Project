<template>
  <div>
    <el-card style="width:94%; margin-left:3%;">
    <el-alert
      title="在这里上传的目的，是为了您的项目能够更好地发展的同时，给后届的学弟学妹一个学习的机会。感谢您所做出的贡献！"
      type="success"
      :closable="false">
    </el-alert>
    <p></p>
    <el-form :inline="true" :model="formInline" class="demo-form-inline">
      <el-form-item label="项目名称" style = "margin-left:210px;">
        <el-input v-model="title" placeholder="项目名称"></el-input>
      </el-form-item>
      <el-form-item label="作者" style = "margin-left:280px">
        <el-input v-model="author" placeholder="作者"></el-input>
      </el-form-item>
    </el-form>
    <p></p>
    <el-form ref="form" :model="form" label-width="80px">
      <el-form-item label="项目概要简介" style = "margin-left:200px;margin-right:200px;">
        <el-input type="textarea" v-model="desc" rows = "10"></el-input>
      </el-form-item>
      <el-form-item label="授权方式" style = "margin-left:210px;">
        <el-select v-model="region" placeholder="请选择授权方式">
          <el-option label="完全开源" value="1"></el-option>
          <el-option label="私下联络授权" value="2"></el-option>
          <el-option label="付费授权" value="3"></el-option>
          <el-option label="按时限授权" value="4"></el-option>
        </el-select>
      </el-form-item>
    </el-form>
    <p></p>
    <label style = "margin-left:220px;">请选择上传源代码:</label>
    <p></p>
    <el-upload
  class="upload-demo"
  ref="upload"
  action=""
  :on-preview="handlePreview"
  :on-remove="handleRemove"
  :file-list="fileList"
  :auto-upload="false">
  <el-button slot="trigger" size="small" type="primary" style = "margin-left:220px;">选取文件</el-button>
  <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload">上传到服务器</el-button>
</el-upload>
<el-button  style = "margin-left:580px;" @click = "SubmitProj">确定提交</el-button>
</el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: '',
      author: '',
      fileList: [
      ],
      desc: '',
      region: ''
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
    SubmitProj() {
      //   console.log('this is render' + render)
      var newProj = {
        author_id: parseInt(window.sessionStorage.getItem('user_id')),
        author_name: window.sessionStorage.getItem('user_name'),
        title: this.title,
        // 标题
        author: this.author,
        // 作者
        fileList: this.fileList,
        // 上传的文件
        desc: this.desc,
        // 内容
        region: this.region
        // 授权方式
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
      window.sessionStorage.setItem('project')
    }
  }
}
</script>
