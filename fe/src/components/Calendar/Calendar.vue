<template>
  <div>
    <el-calendar
      id="calendar"
      v-model="value"
    >
      <template
        slot="dateCell"
        slot-scope="{data}"
      >
        <div>
          <div class="calendar-day">
            {{ data.day.split('-').slice(2).join('-') }}
          </div>
          <div
            v-for="item in calendarData"
            :key="item"
          >
            <div v-if="(item.months).indexOf(data.day.split('-').slice(1)[0])!=-1">
              <div
                v-if="(item.days).indexOf(data.day.split('-').slice(2).join('-'))!=-1"
                @click="open"
              >
                <el-tooltip
                  class="item"
                  effect="dark"
                  :content="item.things"
                  placement="right"
                >
                  <div class="is-selected">
                    {{ item.things }}
                  </div>
                </el-tooltip>
              </div>
              <div v-else />
            </div>
            <div v-else />
          </div>
        </div>
      </template>
    </el-calendar>
  </div>
</template>

<script>
// import Highlighter from 'web-highlighter'
export default {
  data() {
    return {
      calendarData: [
        { months: ['09', '11'], days: ['15'], things: '看电影' },
        { months: ['10', '11'], days: ['02'], things: '去公园野炊' },
        { months: ['11'], days: ['02'], things: '看星星' },
        { months: ['11'], days: ['03'], things: '看月亮' }
      ],
      value: new Date()
    }
  },
  mounted() {
  },
  destroyed() {
    // highlighter.stop()
  },
  methods: {
    handleClose(done) {
      this.drawer = false
    },
    storeToJson() {
      const store = window.sessionStorage.getItem('hl')
      let sources
      try {
        sources = JSON.parse(store) || []
      } catch (e) {
        sources = []
      }
      return sources
    },
    open() {
      this.$alert('这是一段内容', '标题名称', {
        confirmButtonText: '确定',
        callback: action => {
          this.$message({
            type: 'info',
            message: 'action'
          })
        }
      })
    }
  }
}
</script>

<style lang="less" scoped>
.calendar-day{
  text-align: center;
  color: #202535;
  line-height: 30px;
  font-size: 12px;
}
.is-selected{
  color: #F8A535;
  font-size: 10px;
  margin-top: 5px;
  text-align: center;
}
#calendar .el-button-group>.el-button:not(:first-child):not(:last-child):after{
  content: '当月';
}
</style>
