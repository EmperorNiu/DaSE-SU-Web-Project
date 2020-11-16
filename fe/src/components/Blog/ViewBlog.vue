/* eslint-disable */
<template>
<div class="blog-container">
  <div class="title">{{blog.title}}</div>
  <div class="info-box">
    <div class="top-bar">
    <a class="author">{{blog.author_name}}</a>
    <i class="el-icon-time" style="float:left;"></i>
    <span class="time">{{blog.created_at}}</span>
    <i class="el-icon-view" ></i>
    <a class="browse-nums">{{blog.read_times}}</a>
    </div>
    <div class="labels">
      <span  class ="tag">标签:</span>

      <el-tag type="info" v-for="tag in  labels " v-bind:key="tag"> {{tag}}</el-tag>
      <el-tag type="info">{{labels}}</el-tag>
      <!-- <el-tag type="info">标签二</el-tag> -->
      <!-- <el-tag type="info">标签三</el-tag> -->
    </div>
  </div>
  <div class="content">
    <VueMarkdown :source="blog.content_md" ></VueMarkdown>

  </div>
  <div class="comment">
       <show1></show1>
    <comment-box :id="blog.blog_id"></comment-box>
  </div>
  <!-- <div class="favor"></div> -->
  <!-- 收藏 -->
  <!-- <div class="add"> </div> -->
</div>

</template>

<script>
import VueMarkdown from 'vue-markdown'
import comment from '../subcomponents/comment.vue'
import show1 from '../subcomponents/show1.vue'
// 这种方式引入 数据可直接使用
// import markdownData from './test.md';
export default {
  components: {
    VueMarkdown
    // 'comment-box': comment,
    // show1
  },
  created() {
    this.initBlog()
    this.addBrowseNums()
    this.str2list()
  },
  data() {
    return {
        blog:{
            blog_id: 0,
            created_at: "2020-04-11T18:56:33Z",
            title: "测试用文章",
            author_id: 0,
            author_name: "admin",
            content_html: "",
            content_md: "## 测试",
            labels:"tag tag tag",
            read_times: 0,
            star_times: 0,
            thumbs_times: 0,
            change_times: 0,
            comments: null,
            data_statistics: null
        },
        labels:['tag1','tag2']
    }
  },
  methods:{
      initBlog(){
          this.$http
            .get('blog/getBlog',{params:{blog_id:this.$route.query.id}})
            .then(
                result => {
                    this.blog = result.data.project
          }
        )
      },
      str2list(){
        // str = this.blog.labels

        this.labels = this.blog.labels.split(' ')
      },
      addBrowseNums(){
          //后端数据库中browseNums字段自增
           this.$http
          .post('blog/addBrowse',{blog_id:this.blog.blog_id})
          .then(result =>{
            console.log(result)

          }
          )
          .catch(err=>{
            console.log(err)
          } 
          )
      }
  }
}
</script>

<style lang="less" scoped>
.blog-container{
width: 95%;
height: 95%;
.title{
    font-size: 28px;
    word-wrap: break-word;
    color: #222226;
    font-weight: 600;
    margin: 0;
    word-break: break-all;
    margin-bottom: 8px;

}
.info-box{
    position: relative;
    background: #d6d6e4;
    border-radius: 4px;
    .top-bar{
        padding: 10px;
        font-size: 20px;
        .author{
            float: left;
            width:33%;
            margin-left: 20px;
            color: #5893c2;
            margin-right: 20px;
        }
        .time{
            float:left;
            width: 33%;
        }
    }

    .labels{
        padding : 10px;
        .tag{
            margin-left: 20px;
            display: inline-block;
            color: #999aaa;
        }
    }

    .labels .el-tag{
            margin-left: 5px;
            margin-right: 5px;
    }
}

.content{

    margin:20px;
    padding: 20px;
}
.comment
{

}
.favor{

}
.button{
  display: flex;
  justify-content: flex-end;
  padding: 20px 0px;
}
}
</style>
