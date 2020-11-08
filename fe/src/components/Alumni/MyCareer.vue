<template>
  <div>
    <el-card style="width:94%; margin-left:3%;">
      <h3 style="text-align:center;">我的毕业去向</h3>
      <!-- <div style="hei"></div> -->
      <el-divider></el-divider>
      <el-form
        :label-position="labelPosition"
        label-width="100px"
        :model="formLabelAlign"
        size="small"
      >
        <el-row>
          <el-col :span="12">
            <el-form-item label="姓名: ">
              <el-input v-model="formLabelAlign.name"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="毕业年份: ">
              <el-input v-model="formLabelAlign.name"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="手机: ">
              <el-input v-model="formLabelAlign.region" placeholder="选填"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="邮箱: ">
              <el-input v-model="formLabelAlign.region" placeholder="选填"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="毕业去向: ">
              <el-input v-model="formLabelAlign.type" placeholder="去往哪个城市"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="升学/工作: ">
              <el-input v-model="formLabelAlign.type"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="公司: ">
              <el-input v-model="formLabelAlign.type"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="研究生院校: ">
              <el-input v-model="formLabelAlign.type"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="寄语: ">
              <el-input v-model="formLabelAlign.type" type="textarea" placeholder="写下你想对学弟学妹说的吧！" :rows="5"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </el-card>

  </div>
</template>

<script>
export default {
  data() {
    return {
      labelPosition: 'right',
      formLabelAlign: {
        name: '',
        region: '',
        type: ''
      }
    }
  },
  methods: {}
}
</script>

<style lang="less" scoped>
</style>
