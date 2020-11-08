<template>
  <el-carousel
    :interval="4000"
    type="card"
    height="500px"
  >
    <el-carousel-item
      v-for="(img,index) in imgList"
      :key="index"
    >
      <!-- <a href=img.address> -->
      <img
        :src="img.url"
        @click="skip(img.address)"
      >
    <!-- </a> -->
    </el-carousel-item>
  </el-carousel>
</template>

<script>
export default {
  data() {
    return {
      imgList: [
        { url: require('../assets/activity1.jpg'), address: 'https://mp.weixin.qq.com/s/qMtngPkoRIJMTulTe97UeA' },
        { url: require('../assets/activity2.jpg'), address: 'https://mp.weixin.qq.com/s/lAzAzcJUcF6LDxC0fOumuw' },
        { url: require('../assets/activity3.jpg'), address: 'https://mp.weixin.qq.com/s/8R_igg4IUJFDxhUZwsYtdA' }
      ]
    }
  },
  methods: {
    skip (url) {
      location.href = url
    }
    // link_to() {
    //   this.$router.push(this.url)
    // }
  }
}
</script>

<style>
  .el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 300px;
    margin: 0;
  }

  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }

  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }
  img{
width:100%;
height:100%;
}
</style>
