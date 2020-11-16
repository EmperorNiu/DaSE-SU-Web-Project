<template>
  <div>
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/' }">
        首页
      </el-breadcrumb-item>
      <el-breadcrumb-item>
        <a href="/">活动管理</a>
      </el-breadcrumb-item>
      <el-breadcrumb-item>活动列表</el-breadcrumb-item>
      <el-breadcrumb-item>活动详情</el-breadcrumb-item>
    </el-breadcrumb>
    <div class="aLine" />
    <div class="intro">
      <div>
        <el-avatar
          :size="85"
          :src="circleUrl"
        />
      </div>
      <div class="text-intro">
        <div class="text-pattern1">
          {{ username }}
        </div>
        <div class="text-pattern1">
          {{ major }}
        </div>
      </div>
    </div>

    <div class="tags">
      <div>标签：</div>
      <el-tag>前端工程师</el-tag>
      <el-tag type="success">
        摄影师
      </el-tag>
      <el-tag type="info">
        PS达人
      </el-tag>
      <el-tag type="warning">
        Python
      </el-tag>
      <el-tag type="danger">
        javascript
      </el-tag>
    </div>
    <div
      id="my_chart"
      style="height:500px;"
    />
  </div>
</template>

<script>
// import echarts from 'echarts'
export default {
  data() {
    return {
      circleUrl:
        'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png',
      username: '开发者X7',
      major: '数据科学与工程学院'
    }
  },
  mounted() {
    this.initCharts()
  },
  methods: {
    // initCharts() {
    //   var chart = this.$echarts.init(document.getElementById('my_chart'))
  //     var option = {
  //       title: {
  //         text: 'ECharts 入门示例'
  //       },
  //       tooltip: {},
  //       xAxis: {
  //         data: ['衬衫', '羊毛衫', '雪纺衫', '裤子', '高跟鞋', '袜子']
  //       },
  //       yAxis: {},
  //       series: [
  //         {
  //           name: '销量',
  //           type: 'bar',
  //           data: [5, 20, 36, 10, 10, 20]
  //         }
  //       ]
  //     }

    //     chart.setOption(option)
    //   },
    //   setOptions() {
    //     this.chart.setOption({
    //       title: {
    //         text: 'ECharts 入门示例'
    //       },
    //       tooltip: {},
    //       xAxis: {
    //         data: ['衬衫', '羊毛衫', '雪纺衫', '裤子', '高跟鞋', '袜子']
    //       },
    //       yAxis: {},
    //       series: [
    //         {
    //           name: '销量',
    //           type: 'bar',
    //           data: [5, 20, 36, 10, 10, 20]
    //         }
    //       ]
    //     })
    //   }
    // }

    initCharts() {
      var myChart = this.$echarts.init(document.getElementById('my_chart'))
      var option = {
        legend: {},
        tooltip: {
          trigger: 'axis',
          showContent: false
        },
        dataset: {
          source: [
            ['product', '2020.3', '2020.4', '2020.5', '2020.6', '2020.7', '2020.8'],
            ['Blog', 41.1, 30.4, 65.1, 53.3, 83.8, 98.7],
            ['Project', 86.5, 92.1, 85.7, 83.1, 73.4, 55.1],
            ['Activity', 24.1, 67.2, 79.5, 86.4, 65.2, 82.5]
          ]
        },
        xAxis: { type: 'category' },
        yAxis: { gridIndex: 0 },
        grid: { top: '55%' },
        series: [
          { type: 'line', smooth: true, seriesLayoutBy: 'row' },
          { type: 'line', smooth: true, seriesLayoutBy: 'row' },
          { type: 'line', smooth: true, seriesLayoutBy: 'row' },
          // { type: 'line', smooth: true, seriesLayoutBy: 'row' },
          {
            type: 'pie',
            id: 'pie',
            radius: '30%',
            center: ['50%', '25%'],
            label: {
              formatter: '{b}: {@2012} ({d}%)'
            },
            encode: {
              itemName: 'product',
              value: '2012',
              tooltip: '2012'
            }
          }
        ]
      }

      myChart.on('updateAxisPointer', function (event) {
        var xAxisInfo = event.axesInfo[0]
        if (xAxisInfo) {
          var dimension = xAxisInfo.value + 1
          myChart.setOption({
            series: {
              id: 'pie',
              label: {
                formatter: '{b}: {@[' + dimension + ']} ({d}%)'
              },
              encode: {
                value: dimension,
                tooltip: dimension
              }
            }
          })
        }
      })
      myChart.setOption(option)
    }
  }
}
</script>

<style lang="less" scoped>
.aLine {
  background-color: rgb(146, 146, 146);
  height: 1px;
  margin-top: 15px;
  margin-bottom: 15px;
}
.intro {
  display: flex;
  flex-direction: row;
  margin-left: 10px;
}
.text-intro {
  margin-left: 50px;
  margin-top: 10px;
  display: flex;
  flex-direction: column;
}
.text-pattern1 {
  font-size: 18px;
  margin-bottom: 10px;
  word-spacing: 15px;
}
.tags {
  margin-top: 20px;
  margin-left: 20px;
}
.el-tag {
  margin: 5px;
}
</style>
