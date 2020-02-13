<template>
  <el-container class="home-container">
    <el-header height="80px">
      <!-- <div style="width:100%"> -->
      <div class="left-part">
        <img src="#" />
        <span>DaSE SU</span>
      </div>

      <el-menu
        class="home-menu"
        mode="horizontal"
        background-color="#545c64"
        text-color="#fff"
        active-text-color="#ffd04b"
        router
      >
        <el-menu-item index="/home">Home</el-menu-item>
        <el-menu-item index="/blog">博客区</el-menu-item>
        <el-menu-item index="/achievement">风采区</el-menu-item>
        <el-menu-item index="/activity">学生会活动</el-menu-item>
        <el-menu-item index="/user">我的</el-menu-item>
        <el-submenu index="5">
          <template slot="title">待开发</template>
          <el-menu-item index="2-1">待开发</el-menu-item>
          <el-menu-item index="2-3">待开发</el-menu-item>
          <el-menu-item index="2-3">待开发</el-menu-item>
        </el-submenu>
      </el-menu>
      <!-- </div> -->
    </el-header>
    <el-container>
      <el-aside :width="isCollapse ? '64px' : '200px'">
        <div class="toggle-button" @click="toggleCollapse">|||</div>
        <!-- 侧边栏菜单区域 -->
        <el-menu
          background-color="#333744"
          text-color="#fff"
          active-text-color="#409EFF"
          unique-opened
          :collapse="isCollapse"
          :collapse-transition="false"
          router
          :default-active="activePath"
        >
          <!-- 一级菜单 -->
          <el-submenu :index="item.id + ''" v-for="item in menulist" :key="item.id">
            <!-- 一级菜单的模板区域 -->
            <template slot="title">
              <!-- 图标 -->
              <i :class="iconsObj[item.id]"></i>
              <!-- 文本 -->
              <span>{{item.authName}}</span>
            </template>

            <!-- 二级菜单 -->
            <!-- <el-menu-item
              :index="'/' + subItem.path"
              v-for="subItem in item.children"
              :key="subItem.id"
              @click="saveNavState('/' + subItem.path)"
            >
              <template slot="title">
                <i class="el-icon-menu"></i>
                <span>{{subItem.authName}}</span>
              </template>
            </el-menu-item>-->
          </el-submenu>
        </el-menu>
      </el-aside>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
  <!-- 顶栏导航 -->
</template>

<script>
export default {
  data() {
    return {
      menulist: [
        { id: '0', authName: '第一目录' },
        { id: '1', authName: '第二目录' },
        { id: '2', authName: '第三目录' }
      ],
      // 是否折叠
      isCollapse: false,
      // 被激活的链接地址
      activePath: 'xxx',
      iconsObj: {
        0: 'iconfont icon-user',
        1: 'iconfont icon-tijikongjian',
        2: 'iconfont icon-shangpin',
        3: 'iconfont icon-danju',
        4: 'iconfont icon-baobiao'
      }
    }
  },
  methods: {
    // 点击按钮，切换菜单的折叠与展开
    toggleCollapse() {
      this.isCollapse = !this.isCollapse
    }
  }
}
</script>

<style lang="less" scoped>
.el-header {
  background-color: #545c64;
  display: flex;
  padding-left: 0;
  justify-content: space-between;
  align-items: center;
  color: white;
  font-size: 18px;
  > div {
    display: flex;
    align-items: center;
    span {
      margin-left: 15px;
    }
  }
}
.el-aside {
  background-color: #333744;
  .el-menu {
    border-right: none;
  }
}
.el-main {
  background-color: #eaedf1;
}
.home-container {
  height: 100%;
}
// .el-menu {
//   // margin-left: 75px;
//   border-bottom: 0 !important;
// }
.home-menu {
  margin-left: 75px;
  border-bottom: 0 !important;
}
.el-aside {
  background-color: #333744;
  .el-menu {
    border-right: none;
  }
}
.toggle-button {
  background-color: #4a5064;
  font-size: 10px;
  line-height: 24px;
  color: #fff;
  text-align: center;
  letter-spacing: 0.2em;
  cursor: pointer;
}
</style>
