/* eslint-disable no-tabs */
<template>
  <div class="login-container">
    <!-- 背景粒子效果 -->
    <vue-particles color="#ccc"></vue-particles>
    <div class="login-box">
      <!-- 图标区域 -->
      <!-- <div class="avatar-box">
        <img src="../assets/daseLogo.png" />
      </div> -->
      <div class='title'>
        登录
      </div>
      <!-- 表单区域 -->
      <el-form
        label-width="0"
        class="login_form"
        :model="loginForm"
        :rules="loginFormRules"
        ref="loginFormRef"
      >
        <el-form-item prop="username">
          <el-input
            placeholder="请输入账号"
            prefix-icon="iconfont icon-user"
            v-model="loginForm.username"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            placeholder="请输入密码"
            type="password"
            prefix-icon="iconfont icon-3702mima"
            v-model="loginForm.password"
          ></el-input>
        </el-form-item>
        <el-form-item class="btns">
          <el-button type="primary" round @click="login">登录</el-button>
          <el-button type="cancel" round @click="resetForm">注册</el-button>
          <el-button type="warning" circle>忘记密码</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loginForm: {
        username: 'admin',
        password: '123456'
      },
      loginFormRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    resetForm() {
      this.$router.push('/register')
    },
    login() {
      this.$refs.loginFormRef.validate(async valid => {
        if (!valid) return
        await this.$http
          .post('/auth/login', this.loginForm)
          .then(res => {
            this.$message({
              message: '登陆成功',
              type: 'success'
            })
            window.sessionStorage.setItem('token', res.data.token)
            this.$router.push('/home')
          })
          .catch(err => {
            this.$message('登录失败')
            console.log(err)
          })
      })
    }
  }
}
</script>

<style lang="less" scoped>
.login-container {
  background-color: #2b4b6b;
  height: auto;
  background-image: url(../assets/starry-sky-1654074_1920.jpg);
}
.login-box {
  width: 500px;
  height: 360px;
  position: absolute;
  background-color: #D3D3D3;
  opacity: 0.9;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  border-radius: 30px;

  .avatar-box {
    height: 160px;
    width: 160px;
    // border: 1px solid #ccc;
    // border-radius: 50%;
    padding: 10px;
    // box-shadow: 0 0 10px #ddd;
    position: absolute;
    // left: 50%;
    // background-color: #fff;
    transform: translate(-50%, -50%);
    img {
      width: 100%;
      height: 100%;
      // border-radius: 50%;
      // background-color: #eee;
    }
  }
}
.title {
  text-align: center;
  margin-top: 50px;
  font-size: 40px;
  color: rgb(255, 254, 184)
}
.login_form {
  position: absolute;
  bottom: 0;
  width: 100%;
  padding: 0 20px;
  box-sizing: border-box;
}
.btns {
  display: flex;
  justify-content: flex-end;
}
</style>
