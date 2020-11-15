<template>
  <div>
    <el-table
      :data="the_table"
      style="width: 100%"
      height="300"
    >
      <el-table-column
        label="项目名称"
        prop="proj_name"
        width="180"
      />
      <el-table-column
        label="项目负责人"
        prop="leader"
        width="180"
      />
      <el-table-column
        label="项目内容"
        prop="proj_id"
        width="150"
      >
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="cimsInputClick (scope.row)"
          >
            查看项目内容
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- <div>
    <div class="proj-intro-container" v-for="table_intro_Data in the_table" v-bind:key="table_intro_Data.proj_id">
      <div class="intro-info">
        <div class="title">项目名称：{{table_intro_Data.proj_name}}</div>
        <div class="author">负责人：{{table_intro_Data.leader}}</div>
      </div>
    </div>
  </div> -->
  </div>
</template>

<script>
export default {
  data() {
    return {
      the_table: []
    }
  },
  mounted() {
    this.initProjList()
  },
  methods: {
    initProjList() {
      var url = 'project/getProjectList'
      this.$http
        .get(url)
        .then((result) => {
          // console.log(result)
          this.the_table = result.data.proj
          // console.log(this.the_table)
        })
        .catch((err) => {
          console.log(err)
        })
    },
    cimsInputClick (row) {
      this.$router.push('/projinfo/' + row.proj_id)
    }
  }
}
// export default {
//   data() {
//     return {
//       tableData: [{
//         // date: '2016-05-02',
//         name: '王小虎',
//         leader: '上海市普陀区金沙江路 1518 弄'
//       }, {
//         // date: '2016-05-04',
//         name: '王小虎',
//         leader: '上海市普陀区金沙江路 1517 弄'
//       }, {
//         // date: '2016-05-01',
//         name: '王小虎',
//         leader: '上海市普陀区金沙江路 1519 弄'
//       }, {
//         // date: '2016-05-03',
//         name: '王小虎',
//         leader: '上海市普陀区金沙江路 1516 弄'
//       }]
//     }
//   },
//   methods: {
//     cimsInputClick () {
//       // window.loction.href = '/projinfo/1'
//       this.$router.push('/projinfo/1')
//     }
//   }
// }
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
