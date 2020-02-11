<template>
  <el-container class="home-container">
    <el-header height='80px'>
      <div style="width:100%">
        <img src="#" />
        <span>DaSE SU</span>
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
      </div>
    </el-header>
    <el-container>
      <el-aside width="150px">Aside</el-aside>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
  <!-- 顶栏导航 -->
</template>

<script>
export default {}
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
.el-menu {
    margin-left: 75px;
    border-bottom: 0 !important;
}
</style>
