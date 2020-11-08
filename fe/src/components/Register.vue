<template>
  <div class="register-container">
    <vue-particles color="#ccc" />
    <div class="register-box">
      <h1 class="title">
        注册
      </h1>
      <el-form
        ref="registerForm"
        :model="registerForm"
        :rules="rules"
        label-width="130px"
        class="register-form"
      >
        <el-form-item
          label="姓名"
          prop="name"
          class="changeLabel"
        >
          <el-input v-model="registerForm.name" />
        </el-form-item>
        <el-form-item
          label="密码"
          prop="password"
          class="changeLabel"
        >
          <el-input
            v-model="registerForm.password"
            type="password"
          />
        </el-form-item>
        <el-form-item
          label="再次输入密码"
          prop="checkPass"
          class="changeLabel"
        >
          <el-input
            v-model="registerForm.checkPass"
            type="password"
          />
        </el-form-item>
        <el-form-item
          label="性别"
          class="changeLabel"
        >
          <el-radio-group v-model="registerForm.sex">
            <el-radio label="男" />
            <el-radio label="女" />
          </el-radio-group>
        </el-form-item>
        <el-form-item
          label="邮箱"
          prop="mail"
          class="changeLabel"
        >
          <el-input v-model="registerForm.mail" />
        </el-form-item>
        <el-form-item
          label="入学年份"
          prop="grade"
          class="changeLabel"
        >
          <el-input
            v-model.number="registerForm.grade"
            type="number"
          />
        </el-form-item>
        <el-form-item
          label="专业"
          prop="major"
        >
          <el-input v-model="registerForm.major" />
        </el-form-item>
        <el-form-item class="btns">
          <el-button
            type="primary"
            class="btn1"
            @click="submitForm('ruleForm')"
          >
            注册
          </el-button>
          <el-button
            class="btn2"
            @click="resetForm('ruleForm')"
          >
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'))
      } else {
        if (this.ruleForm.checkPass !== '') {
          this.$refs.ruleForm.validateField('checkPass')
        }
        callback()
      }
    }
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.ruleForm.pass) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      registerForm: {
        name: '',
        password: '',
        checkPass: '',
        major: '',
        grade: '',
        sex: '',
        mail: ''
      },
      rules: {
        name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
        password: [{ validator: validatePass, trigger: 'blur' }],
        checkPass: [{ validator: validatePass2, trigger: 'blur' }],
        major: [{ required: true, message: '请输入专业', trigger: 'blur' }],
        grade: [
          { required: true, message: '年级不能为空' },
          { type: 'number', message: '必须输入年份' }
        ],
        mail: [
          { required: true, message: '请输入邮箱地址', trigger: 'blur' },
          {
            type: 'email',
            message: '请输入正确的邮箱地址',
            trigger: ['blur', 'change']
          }
        ]
      }
    }
  },
  methods: {
    register() {}
  }
}
</script>

<style lang="less">
.title {
  color: #ffffff;
  text-align: center;
}
.register-container {
  background-color: rgb(72, 72, 72);
  height: 100%;
  background-image: url("../assets/starry-sky-1654074_1920.jpg");
}
.register-box {
  width: 800px;
  height: 640px;
  position: absolute;
  border-radius: 20px;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  background-color: rgba(110, 110, 110, 0.678);
}
.register-form {
  position: absolute;
  bottom: 0;
  width: 90%;
  padding: 0 40px;
  box-sizing: border-box;
  .el-form-item__label {
    color: #fff !important;
  }
  .el-radio__label {
    color: #fff;
  }
}
.btns {
  display: flex;
  justify-content: flex-end;
}
.changeLabel {
  color: #fff;
}
</style>
