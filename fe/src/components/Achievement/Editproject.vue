<template>
  <div>
    <el-card style="width:94%; margin-left:3%;">
      <el-form
        :inline="true"
        :model="form"
        class="demo-form-inline"
      >
        <el-form-item
          label="项目名称"
          style="margin-left:210px;"
        >
          <el-input
            v-model="form.project_name"
            placeholder="项目名称"
          />
        </el-form-item>
        <el-form-item
          label="作者"
          style="margin-left:280px"
        >
          <el-input
            v-model="form.project_leader"
            placeholder="作者"
          />
        </el-form-item>
      </el-form>
      <p />
      <el-form
        ref="form"
        :model="form"
        label-width="80px"
      >
        <el-form-item
          label="项目概要简介"
          style="margin-left:200px;margin-right:200px;"
        >
          <el-input
            v-model="form.ProjectDescription"
            type="textarea"
            rows="10"
          />
        </el-form-item>
        <el-form-item
          label="项目链接"
          style="margin-left:210px;"
        >
          <el-input
            v-model="form.url"
            placeholder="项目链接"
          />
        </el-form-item>
      </el-form>
      <!-- <p></p> -->
      <!-- <label style = "margin-left:220px;">请选择上传源代码:</label> -->
      <!-- <p></p> -->
      <!-- <el-upload
  class="upload-demo"
  ref="upload"
  action=""
  :on-preview="handlePreview"
  :on-remove="handleRemove"
  :file-list="fileList"
  :auto-upload="false">
  <el-button slot="trigger" size="small" type="primary" style = "margin-left:220px;">选取文件</el-button>
  <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload">上传到服务器</el-button>
</el-upload> -->
      <el-button
        style="margin-left:580px;"
        @click="SubmitProj()"
      >
        确定提交
      </el-button>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      id: this.$route.params.id,
      form: {
        project_name: '',
        project_leader: '',
        url: '',
        ProjectDescription: ''
      },
      projinfo: {
        project_name: '',
        project_leader: '',
        url: '',
        ProjectDescription: ''
      }
    }
  },
  created() {
    this.Getproj()
  },
  methods: {
    Getproj() {
      this.$http.get(`project/getProject?userid=${this.id}`).then(result => {
        console.log(result)
        if (result.data.message === 'success') {
          this.form = result.data.project
          console.log(this.form)
        } else {
          // Toast('获取项目失败！')
        }
      })
    },
    SubmitProj() {
      this.projinfo = this.form
      console.log(this.projinfo)
      var url = 'project/createProject'
      this.$http.put(url, this.projinfo, {
        headers: {
          token: window.sessionStorage.getItem('token')
        }
      }).then((result) => {
        this.$message({
          message: '上传成功',
          type: 'success'
        })
      })
    }
  }
}
</script>
