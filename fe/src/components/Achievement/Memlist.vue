<template>
<div>
<ul class="mui-table-view">
      <li class="mui-table-view-cell mui-media" v-for="item in memlist" :key="item.id">
        <router-link :to="'/home/meminfo/' + item.id">
          <img class="mui-media-object mui-pull-left" :src="item.img_url">
          <div class="mui-media-body">
            <h1>{{ item.author }}</h1>
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
      memlist: [] // 成员列表
    }
  },
  created() {
    this.getMemList()
  },
  methods: {
    getMemList() {
      // 获取列表
      this.$http.get('/getMemList').then(result => {
        if (result.body.status === 0) {
          // 如果没有失败，应该把数据保存到 data 上
          this.memlist = result.body.message
        } else {
          Toast('获取成员列表失败！')
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
// .mui-table-view {
//   li {
//     h1 {
//       font-size: 14px;
//     }
//     .mui-ellipsis {
//       font-size: 12px;
//       color: #226aff;
//       display: flex;
//       justify-content: space-between;
//     }
//   }
// }
.mui-table-view .mui-media-object{
height: 100px;
width: 100px;
max-width: 100px;
}
</style>
