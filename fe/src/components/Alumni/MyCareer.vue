<template>
  <div>
    <el-card
      v-if="isShow"
      style="width:94%; margin-left:3%;"
    >
      <h3 style="text-align:center;">
        我的毕业去向
      </h3>
      <!-- <div style="hei"></div> -->
      <el-divider />
      <el-form
        :label-position="labelPosition"
        label-width="100px"
        :model="alumniForm"
        size="small"
      >
        <el-row>
          <el-col :span="12">
            <el-form-item label="姓名: ">
              <el-input v-model="alumniForm.name" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="毕业年份: ">
              <el-input v-model="alumniForm.graduateYear" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="手机: ">
              <el-input
                v-model="alumniForm.phone"
                placeholder="选填"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="邮箱: ">
              <el-input
                v-model="alumniForm.mail"
                placeholder="选填"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="毕业去向: ">
              <el-input
                v-model="alumniForm.city"
                placeholder="去往哪个城市"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="升学/工作: ">
              <el-input v-model="alumniForm.type" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="公司: ">
              <el-input v-model="alumniForm.company" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="研究生院校: ">
              <el-input v-model="alumniForm.schoolP" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="个人介绍: ">
              <el-input
                v-model="alumniForm.introduction"
                type="textarea"
                placeholder=""
                :rows="5"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="寄语: ">
              <el-input
                v-model="alumniForm.wishes"
                type="textarea"
                placeholder="写下你想对学弟学妹说的吧！"
                :rows="5"
              />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <el-row>
        <el-col
          :span="6"
          :offset="20"
        >
          <el-button
            type="primary"
            @click="uploadCareer"
          >
            提交
          </el-button>
          <el-button type="success">
            重置
          </el-button>
        </el-col>
      </el-row>
    </el-card>
    <el-card
      v-if="!isShow"
      style="width:94%; margin-left:3%;"
    >
      <h3 style="text-align:center;">
        我的毕业去向
      </h3>
      <!-- <div style="hei"></div> -->
      <el-divider />
      <el-row>
        <el-col :span="12">
          <div class="info">
            <div class="info-title">
              姓名:
            </div>
            <div class="info-content">
              {{ alumniInfo.name }}
            </div>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="info">
            <div class="info-title">
              毕业年份:
            </div>
            <div class="info-content">
              {{ alumniInfo.graduateYear }}
            </div>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <div class="info">
            <div class="info-title">
              手机:
            </div>
            <div class="info-content">
              {{ alumniInfo.phone }}
            </div>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="info">
            <div class="info-title">
              邮箱:
            </div>
            <div class="info-content">
              {{ alumniInfo.mail }}
            </div>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <div class="info">
            <div class="info-title">
              毕业去向:
            </div>
            <div class="info-content">
              {{ alumniInfo.city }}
            </div>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="info">
            <div class="info-title">
              升学/工作:
            </div>
            <div class="info-content">
              {{ alumniInfo.type }}
            </div>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <div class="info">
            <div class="info-title">
              公司:
            </div>
            <div class="info-content">
              {{ alumniInfo.company }}
            </div>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="info">
            <div class="info-title">
              研究生院校:
            </div>
            <div class="info-content">
              {{ alumniInfo.schoolP }}
            </div>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <div class="info">
            <div class="info-title">
              个人介绍:
            </div>
            <div class="info-content">
              {{ alumniInfo.introduction }}
            </div>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="info">
            <div class="info-title">
              寄语:
            </div>
            <div class="info-content">
              {{ alumniInfo.wishes }}
            </div>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col
          :span="6"
          :offset="22"
        >
          <el-button
            type="primary"
            @click="modification"
          >
            修改
          </el-button>
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    var now = new Date()
    return {
      username: '',
      labelPosition: 'right',
      alumniInfo: {
        name: '牛悦安',
        mail: 'niuyuean99@sina.com',
        city: '上海',
        phone: '15546845',
        type: '升学',
        graduateYear: now.getFullYear(),
        company: '无',
        schoolP: '华东师范大学',
        wishes: '本科阶段打好基础',
        introduction: '前端开发工程师'
      },
      isShow: false,
      alumniForm: {
        name: '',
        mail: '',
        city: '',
        phone: '',
        type: '',
        graduateYear: now.getFullYear(),
        company: '',
        schoolP: '',
        wishes: '',
        introduction: ''
      }
    }
  },
  mounted() {
    this.username = sessionStorage.getItem('user_name')
    this.alumniForm.name = this.username
  },
  methods: {
    getAlumniInfo() {
      var url = 'alumni/getInfo?username=' + this.username
      this.$http.get(url).then((result) => {
        this.alumniInfo = result.data
      })
    },
    uploadCareer() {
      var url = 'alumni/create'
      this.$http.post(url, this.alumniForm).then((result) => {
        this.$message({
          message: '上传成功',
          type: 'success'
        })
      })
    },
    updateCareer() {
      var url = 'alumni/update'
      this.$http.post(url, this.alumniForm).then((result) => {
        this.$message({
          message: '上传成功',
          type: 'success'
        })
      })
    },
    modification() {
      this.isShow = true
      this.alumniForm = this.alumniInfo
    }
  }
}
</script>

<style lang="less" scoped>
.info {
  display: flex;
  flex-direction: row;
  margin-top: 20px;
}
.info-title {
  font-weight: bold;
  margin-right: 8px;
}
.info-content {
}
</style>
