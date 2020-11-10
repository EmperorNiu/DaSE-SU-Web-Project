<template>
  <div>
    <el-container>
      <el-header>项目标题</el-header>
      <el-header class="th">项目负责人</el-header>
      <el-container>
        <el-main>
          项目内容
          <el-button type="primary" disabled>下载</el-button>
          <el-button type="primary" disabled>申请授权</el-button>
        </el-main>
      </el-container>
    </el-container>
    <!-- <div class="proj-info-container">
      <div class="info">
        <div class="title">
          {{ projinfo.project_name }}
        </div>
        <div class="author">
          {{ projinfo.project_leader }}
        </div>
        <div class="title">
          项目名称:{{ projinfo.proj_name }}
        </div>
        <br><br>
        <div class="author">
          项目负责人:{{ projinfo.leader }}
        </div>
        <br><br>
        <div class="members">
          项目其他成员:{{ projinfo.members }}
        </div>
        <br><br>
      </div>
      <div class="recommend">
        {{ projinfo.ProjectDescription }}
      </div>
      <div class="recommend">
        我们已经知道，在zookeeper中，写请求转发给leader，读请求follower自己处理，没有sync的情况下达不到linearizability。Mongodb呢?官方文档是这么说的，“MongoDB
        is consistent by default: reads and writes are issued to the primary
        member of a replica set.
        ”这句话误导了我很多年，让我一直以为默认情况下，读写操作都发给primary，这总应该linearizable了吧?结果得知真相之后让我眼泪掉下来。Write
        Concern，Read Concern和Read Preference
      </div>
      <button onclick="Apply(userid, projid)">
        申请
      </button>
      <button onclick="Download(userid, projid)">
        下载
      </button>
    </div> -->
  </div>
</template>

<script>
// import comment from '../subcomponents/comment.vue'
// import show1 from '../subcomponents/show1.vue'
// import { Toast } from 'mint-ui'
export default {
  data() {
    return {
      id: this.$route.params.id, // 将 URL 地址中传递过来的 Id值，挂载到 data上，方便以后调用
      projinfo: {
        project_id: 0,
        project_leader: '',
        project_members: [],
        project_name: '',
        projectDescription: ''
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
        if (result.body.message === 'success') {
          this.projinfo = result.body.project
        } else {
          // Toast('获取项目失败！')
        }
      })
    },
    Apply(userid, projid) {
      // projid.appliers.add(userid, 0)
      this.$http.post(`http://example.api?userid=${userid}?`).then(function(result) {
        if (result.data.message !== 'success') {
          // Toast('申请失败！')
        }
      })
    },
    Download(userid, projid) {
      this.$http.get(`/123456?userid=${userid}?projid=${projid}`).then(result => {
        if (result.body.message === 'success') {
          if (result.body.yes_or_no === 1) {
            this.$http.get(`/123456?userid=${userid}?projid=${projid}`).then(result => {
              // 下载
            })
          } else {
            // Toast('无权限！')
          }
        } else {
          // Toast('网络连接失败！')
        }
      })
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
