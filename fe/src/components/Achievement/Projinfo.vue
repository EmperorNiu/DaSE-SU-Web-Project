<template>
  <div>
    <el-container>
      <el-header>{{ projinfo.project_name }}</el-header>
      <el-header class="th">
        项目负责人：{{ projinfo.project_leader }}
      </el-header>
      <el-container>
        <el-main>
          {{ projinfo.ProjectDescription}}
          <el-button
            type="primary"
            @click="Download(projinfo.url)"
          >
            下载
          </el-button>
          <!-- <el-link href={{ projinfo.url }}>下载链接</el-link> -->
          <el-button
            type="primary"
            @click="Apply(projinfo.project_id)"
          >
            申请授权
          </el-button>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
export default {
  data() {
    return {
      id: this.$route.params.id, // 将 URL 地址中传递过来的 Id值，挂载到 data上，方便以后调用
      projinfo: {
        // project_id: 0,
        // project_leader: '',
        // project_members: [],
        // project_name: '',
        // projectDescription: ''
      } // 项目对象
    }
  },
  created() {
    this.getProjInfo()
  },
  methods: {
    getProjInfo() {
      // 获取项目详情
      this.$http.get(`project/getProject?userid=${this.id}`).then(result => {
        console.log(result)
        if (result.data.message === 'success') {
          this.projinfo = result.data.project
          // console.log(this.projinfo.project_name)
        } else {
          // Toast('获取项目失败！')
        }
      })
    },
    // Apply(userid) {
    //   this.$http.post(`http://example.api?userid=${userid}?`).then(function(result) {
    //     if (result.data.message !== 'success') {
    //     }
    //   })
    // },
    Download(url) {
      console.log(url)
      location.href = url
    }
  // components: {
    // 用来注册子组件的节点
    // 'comment-box': comment,
    // eslint-disable-next-line vue/no-unused-components
    // show1
  // }
  }
}
</script>

<style lang="less" scoped>
.proj-info-container{
  width: 85%;
  // height: 150px;
  height: 15000 px;
  // margin-left: 5%;
  background-color: rgb(247, 247, 247);
  box-shadow: 0 2 px 4 px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
  padding: 15 px;
  margin-bottom: 15 px;
}
.title {
  font-size: 50px;
  font-weight: 600;
  color: black;
}
.author {
  font-size: 20px;
  font-weight: 600;
  color: black;
  margin: 0 auto;
  text-align: center;
}
.recommend {
  width: 85%;
  // height: 150px;
  height: 10000 px;
  // margin-left: 5%;
  background-color: rgb(247, 247, 247);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
  padding: 15px;
  margin-bottom: 15px;
  color: #7e7e7e;
  margin: 0 auto;
  text-align: center;
}
.el-header{
    background-color: #B3C0D1;
    color: #333;
    text-align: center;
    line-height: 60px;
    // height:200px;
    font-size: 50px;
    font-weight: 600;
  }
.th{
    background-color: #B3C0D1;
    color: #333;
    text-align: center;
    line-height: 60px;
    // height:200px;
    font-size: 30px;
    font-weight: 600;
  }
.el-main {
    background-color: #E9EEF3;
    color: #333;
    text-align: center;
    line-height: 160px;
  }
  .el-aside {
    background-color: #D3DCE6;
    color: #333;
    text-align: center;
    line-height: 1000px;
    font-size: 50px;
    font-weight: 600;
  }
</style>
