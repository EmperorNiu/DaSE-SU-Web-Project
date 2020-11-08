<template>
  <!-- <div> -->
    <!-- <projinfo></projinfo>
    <projinfo></projinfo>
    <projinfo></projinfo> -->
    <!-- <ul class="mui-table-view">
      <li class="mui-table-view-cell mui-media" v-for="item in projlist" :key="item.project_id">
        <router-link :to="'/projinfo/' + item.project_id">
          <div class="mui-media-body">
            <h1>{{ item.project_leader }}</h1>
          </div>
        </router-link>
      </li>
    </ul> -->
    <!-- <ul class="mui-table-view">
      <li class="mui-table-view-cell mui-media">
        <router-link :to="'/projinfo/' + 1">
          <div class="mui-media-body">
            <h1>项目负责人</h1>
          </div>
          <br /><br />
        </router-link>
        <router-link :to="'/projinfo/' + 2">
          <div class="mui-media-body">
            <h1>项目负责人</h1>
          </div>
        </router-link>
        <router-link :to="'/projinfo/' + 2">
          <div class="mui-media-body">
            <h1>项目负责人</h1>
          </div>
        </router-link>
      </li>
    </ul> -->
    <!-- <el-table :data="tableData" style="width: 100%">
      <el-table-column prop="name" label="项目名称" width="180"> </el-table-column>
      <el-table-column prop="address" label="项目负责人"> </el-table-column>
    </el-table> -->
    <div>
    <div class="proj-intro-container" v-for="table_intro_Data in the_table" v-bind:key="table_intro_Data.proj_id">
      <div class="intro-info">
        <div class="title">{{table_intro_Data.proj_name}}</div>
        <div class="author">负责人：{{table_intro_Data.leader}}</div>
        <div class="others">其他成员：{{table_intro_Data.members}}</div>
      </div>
    </div>
  </div>
  <!-- </div> -->
</template>

<script>
// import { Toast } from 'mint-ui'

// export default {
//   data() {
//     return {
//       projlist: [] // 项目列表
//     }
//   },
//   created() {
//     this.getProjList()
//   },
//   methods: {
//     getProjList() {
//       // 获取项目列表
//       this.$http.get('/getProjList').then(result => {
//         if (result.body.message === 'success') {
//           // 如果没有失败，应该把数据保存到 data 上
//           this.projlist = result.body.project
//         } else {
//           Toast('获取项目列表失败！')
//         }
//       })
//     }
//   }
// }
export default {
  data() {
    return {
      the_table: [],
      table_intro_Data: [{
        proj_id: '',
        leader: '',
        members: [],
        proj_name: ''
      }]
    }
  },
  methods: {
    GetStats() {
      var url = 'Achievement/get_proj_List'
      this.$http.get(url).then((result) => {
        // console.log(result)
        this.the_table = result.data.proj
        // this.circleUrl = this.profile.avatar_url
      })
        .catch(err => {
          console.log(err)
        })
    }
  }
}
</script>

<style lang="less" scoped>
.proj-intro-container {
  width: 85%;
  height: 150px;
  // margin-left: 5%;
  background-color: rgb(247, 247, 247);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
  padding: 15px;
  margin-bottom: 15px;
}
.title {
  font-size: 20px;
  font-weight: 600;
  color: rgb(0, 0, 2);
}
// .recommend {
//   color: #7e7e7e;
// }
</style>
