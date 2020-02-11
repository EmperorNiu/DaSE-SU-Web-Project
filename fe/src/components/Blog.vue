<template>
  <div>
    <mavon-editor v-model="value" :ishljs="true" @save="saveMd" />
  </div>
</template>

<script>
export default {
  data () {
    return {
      md: ''
    }
  },
  methods: {
    saveMd (value, render) {
      console.log('this is render' + render)
      window.sessionStorage.setItem('blog', render)
    }
  }
}
</script>
