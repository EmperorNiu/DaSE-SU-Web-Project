<template>
  <div>
    <div class="blog-intro-container">
    <div class="info">
      <div class="title">{{blogIntro.title}}</div>
      <div class="author">{{blogIntro.author}}</div>
    </div>
    <div class="recommend">{{blogIntro.introduction}}</div>
    <div class="statistics">
      <div class="comment-num">评论数：{{commentNums}}</div>
      <div class="view-num">观看人数： {{viewNums}}</div>
      <div class="view-time">预计观看时间：{{viewTime}}</div>
    </div>
    </div>

  </div>
</template>

<script>
export default {
  data() {
    return {
      blogIntro: {
        title: '深度学习优化器',
        author: 'Dylan Niu',
        introduction:
          '随机梯度下降 SGD、动量 Momentum、Momentum+SGD算法、牛顿动量 Nesterov、Nesterov+SGD算法、AdaGrad'
      }
    }
  },
  methods: {
    saveMd(value, render) {
      //   console.log('this is render' + render)
      window.sessionStorage.setItem('blog', render)
    }
  }
}
</script>

<style lang="less" scoped>
.blog-intro-container {
  width: 85%;
  height: 150px;
  // margin-left: 5%;
  background-color: rgb(247, 247, 247);
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);
  padding: 15px;
  margin-bottom: 15px;
}
.title {
  font-size: 20px;
  font-weight: 600;
}
.recommend {
  color: #7e7e7e
}
</style>
