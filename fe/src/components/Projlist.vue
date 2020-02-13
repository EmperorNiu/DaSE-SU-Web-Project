<template>
  <div>
    <ul class="mui-table-view">
      <li class="mui-table-view-cell mui-media" v-for="item in projlist" :key="item.id">
        <router-link :to="'/home/projinfo/' + item.id">
          <img class="mui-media-object mui-pull-left" :src="item.img_url">  
		  <!--img_url是该项目展示图标的图片路径-->
          <div class="mui-media-body">
            <h1>{{ item.title }}</h1>
          </div>
        </router-link>
      </li>
    </ul>

  </div>
</template>
<script>
import { Toast } from "mint-ui";

export default {
  data() {
    return {
      projlist: [] // 项目列表
    };
  },
  created() {
    this.getProjList();
  },
  methods: {
    getProjList() {
      // 获取项目列表
      this.$http.get("api/getprojlist").then(result => {
        if (result.body.status === 0) {
          // 如果没有失败，应该把数据保存到 data 上
          this.projlist = result.body.message;
        } else {
          Toast("获取项目列表失败！");
        }
      });
    }
  }
};
</script>

<style lang="scss" scoped>
.mui-table-view {
  li {
    h1 {
      font-size: 14px;
    }
  }
}
</style>
