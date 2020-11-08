<template>
  <div>
    <el-table :data="tableData" style="width: 100%" ref="filterTable">
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="姓名">
              <span>{{ props.row.name }}</span>
            </el-form-item>
            <el-form-item label="籍贯">
              <span>{{ props.row.shop }}</span>
            </el-form-item>
            <el-form-item label="毕业年份">
              <span>{{ props.row.id }}</span>
            </el-form-item>
            <el-form-item label="毕业去向">
              <span>{{ props.row.category }}</span>
            </el-form-item>
            <el-form-item label="详细去向">
              <span>{{ props.row.address }}</span>
            </el-form-item>
            <el-form-item label="寄语">
              <span>{{ props.row.desc }}</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column label="毕业年份" prop="id" sortable
        :filters="[{text: '2020', value: '2020'}, {text: '2021', value: '2021'}]"
        :filter-method="filterHandler"> </el-table-column>
      <el-table-column label="姓名" prop="name"> </el-table-column>
      <el-table-column label="毕业去向(工作/升学)" prop="desc"> </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      tableData: [{
        id: '2020',
        name: '韩小聪',
        category: '江浙小吃、小吃零食',
        desc: '升学',
        address: '上海市普陀区真北路',
        shop: '王小虎夫妻店'
      }, {
        id: '2020',
        name: '江帆帆',
        category: '江浙小吃、小吃零食',
        desc: '升学',
        address: '上海市普陀区真北路',
        shop: '王小虎夫妻店'
      }, {
        id: '2020',
        name: '杜云海',
        category: '江浙小吃、小吃零食',
        desc: '升学',
        address: '上海市普陀区真北路',
        shop: '王小虎夫妻店'
      }, {
        id: '2020',
        name: '李琨',
        category: '江浙小吃、小吃零食',
        desc: '升学',
        address: '上海市普陀区真北路',
        shop: '王小虎夫妻店'
      }, {
        id: '2020',
        name: '郁思思',
        category: '江浙小吃、小吃零食',
        desc: '升学',
        address: '上海市普陀区真北路',
        shop: '王小虎夫妻店'
      }, {
        id: '2021',
        name: '安悦',
        category: '江浙小吃、小吃零食',
        desc: '工作',
        address: '上海市普陀区真北路',
        shop: '王小虎夫妻店'
      }, {
        id: '2021',
        name: '徐珍珍',
        category: '江浙小吃、小吃零食',
        desc: '升学',
        address: '上海市普陀区真北路',
        shop: '王小虎夫妻店'
      }, {
        id: '2021',
        name: 'xxx',
        category: '江浙小吃、小吃零食',
        desc: '升学',
        address: '上海市普陀区真北路',
        shop: '王小虎夫妻店'
      }]
    }
  },
  methods: {
    filterHandler(value, row, column) {
      const property = column.property
      return row[property] === value
    }
  }
}
</script>

<style lang="less" scoped>
.demo-table-expand {
  font-size: 0;
}
.demo-table-expand label {
  width: 90px;
  color: #99a9bf;
}
.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}
</style>
