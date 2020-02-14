<template>
  <div>
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>
        <a href="/">活动管理</a>
      </el-breadcrumb-item>
      <el-breadcrumb-item>活动列表</el-breadcrumb-item>
      <el-breadcrumb-item>活动详情</el-breadcrumb-item>
    </el-breadcrumb>
    <div class="aLine"></div>
    <div class="intro">
      <div>
        <el-avatar :size="85" :src="circleUrl"></el-avatar>
      </div>
      <div class="text-intro">
        <div class="text-pattern1">{{username}}</div>
        <div class="text-pattern1">{{major}}</div>
      </div>
    </div>

    <div class="tags">
      <div>标签：</div>
      <el-tag>前端工程师</el-tag>
      <el-tag type="success">摄影师</el-tag>
      <el-tag type="info">PS达人</el-tag>
      <el-tag type="warning">Python</el-tag>
      <el-tag type="danger">javascript</el-tag>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      circleUrl:
        'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png',
      username: '开发者X7',
      major: '数据科学与工程学院'
    }
  }
}
</script>

<style lang="less" scoped>
.aLine {
  background-color: rgb(146, 146, 146);
  height: 1px;
  margin-top: 15px;
  margin-bottom: 15px;
}
.intro {
  display: flex;
  flex-direction: row;
  margin-left: 10px;
}
.text-intro {
  margin-left: 50px;
  margin-top: 10px;
  display: flex;
  flex-direction: column;
}
.text-pattern1 {
  font-size: 18px;
  margin-bottom: 10px;
  word-spacing: 15px;
}
.tags {
    margin-top: 20px;
    margin-left: 20px;
}
.el-tag {
    margin: 5px;
}
</style>
