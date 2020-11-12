
<template>
  <div>
    <el-drawer
      title="2020-06-19 9:00~11:00 班次详情"
      :visible.sync="drawer"
      :before-close="handleClose"
      :modal-append-to-body="false"
      size="300px"
    >
      <div class="block">
        <el-timeline>
          <el-timeline-item
            v-for="(activity, index) in activities"
            :key="index"
            :icon="activity.icon"
            :type="activity.type"
            :color="activity.color"
            :size="activity.size"
            :timestamp="activity.timestamp"
          >
            <p>{{ activity.content }}</p>
            <div
              v-if="activity.url"
              class="demo-image__preview"
            >
              <el-image
                style="width: 100px; height: 100px"
                :src="activity.url"
                :preview-src-list="activity.srcList"
              />
            </div>
          </el-timeline-item>
        </el-timeline>
      </div>
    </el-drawer>
  </div>
</template>
<script>
export default {
  data() {
    return {
      drawer: false,

      activities: [
        // {
        //   content: '签到（张三）',
        //   timestamp: '2018-04-12 20:46',
        //   size: 'large',
        //   type: 'primary',
        //   icon: 'el-icon-more',
        //   url:
        //     'https://fuss10.elemecdn.com/8/27/f01c15bb73e1ef3793e64e6b7bbccjpeg.jpeg',
        //   srcList: [
        //     'https://fuss10.elemecdn.com/8/27/f01c15bb73e1ef3793e64e6b7bbccjpeg.jpeg',
        //     'https://fuss10.elemecdn.com/1/8e/aeffeb4de74e2fde4bd74fc7b4486jpeg.jpeg'
        //   ]
        // },
        {
          content: '数理统计考试',
          timestamp: '2020-11-06 8:00',
          size: 'large'
        },
        // {
        //   content: '岗位点打卡（李四）',
        //   timestamp: '2018-04-03 20:46',
        //   color: '#0bbd87'
        // },
        // {
        //   content: '久坐不动（张三）',
        //   timestamp: '2018-04-03 20:46',
        //   size: 'large'
        // },
        {
          content: '购物',
          timestamp: '2020-11-11 6:00'
        }
      ]
    }
  },
  mounted() {},
  methods: {
    handleClose() {
      this.drawer = false
    }
  }
}
</script>
