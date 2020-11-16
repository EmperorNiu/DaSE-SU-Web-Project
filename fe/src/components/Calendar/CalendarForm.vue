<template>
  <el-dialog
    title="新增活动"
    :visible.sync="dialogFormVisible"
    :modal-append-to-body="false"
  >
    <el-form :model="form">
      <el-form-item
        label="活动名称"
        :label-width="formLabelWidth"
      >
        <el-input
          v-model="form.name"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        label="活动区域"
        :label-width="formLabelWidth"
      >
        <el-select
          v-model="form.region"
          placeholder="请选择活动区域"
          style="width:100%"
        >
          <el-option
            label="区域一"
            value="shanghai"
          />
          <el-option
            label="区域二"
            value="beijing"
          />
        </el-select>
      </el-form-item>
    </el-form>
    <div
      slot="footer"
      class="dialog-footer"
    >
      <el-button @click="dialogFormVisible = false">
        取 消
      </el-button>
      <el-button
        type="primary"
        @click="dialogFormVisible = false"
      >
        确 定
      </el-button>
    </div>
  </el-dialog>
</template>

<script>
export default {
  data() {
    return {
      dialogFormVisible: false,
      form: {
        name: '',
        region: ''
      },
      formLabelWidth: '80px'
    }
  }
}
</script>
