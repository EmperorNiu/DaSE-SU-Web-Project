<template>
  <div>
    <div v-html="ss"></div>
  </div>
</template>

<script>
export default {
  created () {
    var blog = window.sessionStorage.getItem('blog')
    this.ss = blog
  },
  data () {
    return {
      ss: '<h2>这是一个h2标签</h2>'
    }
  }
}
</script>
