<template>
  <div>
    <el-card class="_calendar">
      <el-container>
        <el-main>
          <el-card>
            <el-calendar v-model="value" :first-day-of-week="7">
              <template slot="dateCell" slot-scope="{data}">
                <div slot="reference" class="div-Calendar" @dblclick="saveOnClick">
                  <p :class="data.isSelected ? 'is-selected' : ''">
                    {{ data.day.split('-').slice(1).join('-') }}
                    <i
                      :class="[data.isSelected ?'el-icon-check':'']"
                    ></i>
                    <!-- <i v-if="_.indexOf(isArrange, data.day)>0" class="el-icon-s-claim"></i> -->
                    <!-- <i class="el-icon-coffee-cup"></i> -->
                  </p>
                </div>
              </template>
            </el-calendar>
          </el-card>
        </el-main>
        <el-aside width="40%" style="overflow: hidden;">
          <el-card>
            <div class="el-calendar__header">
              <div class="el-calendar__title">日程详情</div>
              <div class="el-calendar__button-group">
                <div class="el-button-group">
                  <button
                    type="button"
                    class="el-button el-button--plain el-button--mini"
                    @click="saveOnClick"
                  >
                    <span>新增</span>
                  </button>
                </div>
              </div>
            </div>
            <div class="calendar-info">
              <div style="padding: 15px;">
                <div role="alert" class="el-alert el-alert--success is-dark" @click="infoOnClick">
                  <!-- <i class="el-alert__icon el-icon-success is-big"></i> -->
                  <div class="el-alert__content">
                    <span class="el-alert__title is-bold">2020-06-19 9:00~11:00</span>
                    <p class="el-alert__description">算法考试</p>
                    <i class="el-alert__closebtn el-icon-close" @click.stop="infoDel"></i>
                  </div>
                </div>
                <!-- <div role="alert" class="el-alert el-alert--info is-dark" @click="infoOnClick">
                  <i class="el-alert__icon el-icon-info is-big"></i>
                  <div class="el-alert__content">
                    <span class="el-alert__title is-bold">2020-06-19 9:00~11:00（待审核）</span>
                    <p class="el-alert__description">值班人员：张三、李四、王五</p>
                    <p class="el-alert__description">携带设备：123547、985431、745481</p>
                    <i class="el-alert__closebtn el-icon-close"></i>
                  </div>
                </div>
                <div role="alert" class="el-alert el-alert--warning is-dark" @click="infoOnClick">
                  <i class="el-alert__icon el-icon-warning is-big"></i>
                  <div class="el-alert__content">
                    <span class="el-alert__title is-bold">警告提示的文案</span>
                    <p class="el-alert__description">文字说明文字说明文字说明文字说明文字说明文字说明</p>
                    <i class="el-alert__closebtn el-icon-close"></i>
                  </div>
                </div>
                <div role="alert" class="el-alert el-alert--error is-dark" @click="infoOnClick">
                  <i class="el-alert__icon el-icon-error is-big"></i>
                  <div class="el-alert__content">
                    <span class="el-alert__title is-bold">错误提示的文案</span>
                    <p class="el-alert__description">文字说明文字说明文字说明文字说明文字说明文字说明</p>
                    <i class="el-alert__closebtn el-icon-close"></i>
                  </div>
                </div> -->
              </div>
            </div>
          </el-card>
        </el-aside>
      </el-container>
    </el-card>
    <calendarDrawer ref="calendarDrawer"></calendarDrawer>
    <calendarForm ref="calendarForm"></calendarForm>
  </div>
</template>

<script>
import calendarDrawer from './CalendarDrawer.vue'
import calendarForm from './CalendarForm.vue'
export default {
  components: { calendarDrawer, calendarForm },
  data: function() {
    return {
      value: new Date(),
      isArrange: [
        '2020-06-08',
        '2020-06-09',
        '2020-06-10',
        '2020-06-11',
        '2020-06-17',
        '2020-06-18'
      ]
    }
  },
  created: function() {
    this.$nextTick(() => {
      // 点击前一个月
      var prevBtn = document.querySelector(
        '.el-calendar__button-group .el-button-group>button:nth-child(1)'
      )
      prevBtn.addEventListener('click', e => {
        console.log(this.data)
        this.$notify({
          title: '上一月',
          message: '上个月头天：' + this.value,
          type: 'success',
          position: 'top-left'
        })
      })
      // 点击下一个月
      var nextBtn = document.querySelector(
        '.el-calendar__button-group .el-button-group>button:nth-child(3)'
      )
      nextBtn.addEventListener('click', () => {
        console.log(this.value)
        this.$notify({
          title: '下一月',
          message: '下个月头天：' + this.value,
          type: 'warning',
          position: 'top-left'
        })
      })
      // 点击今天
      var todayBtn = document.querySelector(
        '.el-calendar__button-group .el-button-group>button:nth-child(2)'
      )
      todayBtn.addEventListener('click', () => {
        this.$notify.info({
          title: '今天',
          message: '今天：' + this.value,
          position: 'top-left'
        })
      })
    })
  },
  mounted: function() {},
  methods: {
    // 点击日期块
    calendarOnClick(e) {
      console.log(e)
      // this.isArrange.push('2020-06-19');
      this.$notify.error({
        title: '日历块点击',
        message: e.day,
        position: 'top-left'
      })
    },
    // 点击信息块
    infoOnClick() {
      this.$refs.calendarDrawer.drawer = true
    },
    // 新增排班
    saveOnClick() {
      this.$refs.calendarForm.dialogFormVisible = true
    },
    // 删除排班
    infoDel() {
      this.$confirm('此操作将永久删除该排班, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    }
  }
}
</script>

<style scoped>
.is-selected {
  color: #1989fa;
}
.p-popover {
  text-align: center;
}
._calendar {
  height: 750px;
  width: 100%;
}
.el-main {
  padding: 0px;
  overflow: hidden;
}
.calendar-info {
  height: 94%;
  overflow: auto;
}
.el-alert {
  margin: 15px 0px;
}
.el-alert:hover {
  transform: scale(1.02);
}
.el-alert:active {
  transform: scale(1.01);
}
</style>
<style >
._calendar .div-Calendar {
  height: 125px;
  box-sizing: border-box;
  padding: 8px;
}
._calendar .el-calendar-table .el-calendar-day {
  padding: 0px;
  height: 125px;
}
._calendar .el-container,
._calendar .el-calendar {
  height: 100%;
}
._calendar .el-card__body {
  padding: 0px;
}
</style>
