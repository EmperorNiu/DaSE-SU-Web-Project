<template>
  <div>
    <!-- <projinfo></projinfo>
    <projinfo></projinfo>
    <projinfo></projinfo> -->
    <ul class="mui-table-view">
      <li class="mui-table-view-cell mui-media" v-for="item in projlist" :key="item.project_id">
        <router-link :to="'/home/projinfo/' + item.project_id">
          <div class="mui-media-body">
            <h1>{{ item.project_leader }}</h1>
            <!-- <p class='mui-ellipsis'>
              <span>发表时间：{{ item.add_time}}</span>
            </p> -->
          </div>
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script>
import { Toast } from 'mint-ui'

export default {
  data() {
    return {
      projlist: [] // 项目列表
    }
  },
  created() {
    this.getProjList()
  },
  methods: {
    getProjList() {
      // 获取项目列表
      this.$http.get('/getProjList').then(result => {
        if (result.body.message === 'success') {
          // 如果没有失败，应该把数据保存到 data 上
          this.projlist = result.body.project
        } else {
          Toast('获取项目列表失败！')
        }
      })
    }
  }
}
</script>

<style lang="less" scoped>
.projinfo-container {
  width: 85%;
  height: 150px;
  // margin-left: 5%;
  background-color: rgb(247, 247, 247);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
  padding: 15px;
}
.title {
  font-size: 20px;
  font-weight: 600;
}
.recommend {
  color: #7e7e7e;
}
</style>
