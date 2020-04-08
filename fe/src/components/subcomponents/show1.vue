<template>
    <form  method="post" class="smart-green">
    <label><span>页面标题 :</span>
        <input id="title" type="text" name="title"  />
        <div class="error-msg"></div>
    </label>
   <!-- <label><span>邮箱 :</span>
        <input id="email" type="email" value="" name="email" /> &nbsp; &nbsp; &nbsp; &nbsp;
        <div class="error-msg"></div></label>-->
    <!-- <label><span>关键字 :</span>
        <input id="keywords" type="text" value="" name="keywords" />
        <div class="error-msg"></div></label> -->
    <!-- <label><span>页面描述 :</span> -->
    <!-- <div>
        <quill-editor ref="text" v-model="content" class="myQuillEditor" :options="editorOption" />
        <el-button type="primary" @click="submit">提交</el-button>
        </div>
        // eslint-disable-next-line vue/no-parsing-error
        <textarea id="description" name="description">
        <div>
        <quill-editor ref="text" v-model="content" class="myQuillEditor" :options="editorOption" />
        <el-button type="primary" @click="submit">提交</el-button>
        </div>
        </textarea> -->
        <!-- <div class="error-msg"></div> &nbsp; &nbsp; </label>
    <div class="success-msg"></div>
    <label><span>&nbsp;</span><input type="submit" class="button" value="确定" /></label>
    <input type="hidden" name="csrfmiddlewaretoken" value="SfHkbL4feo1G00sJQtbO7TtLN4c2BUwa" /> -->
    <label><span>页面描述 :</span></label>
    <div>
        <quill-editor ref="text" v-model="content" class="myQuillEditor" :options="editorOption" />
        <el-button type="primary" @click="submit">提交</el-button>
    </div>
</form>
</template>

<script>
import { quillEditor } from 'vue-quill-editor'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'

export default {
  data() {
    return {
      content: '',
      editorOption: {}
    }
  },
  methods: {
    submit () {
      console.log(this.$refs.text.value)
    }
  },
  components: {
    quillEditor
  }
}
</script>

<style>
.smart-green {
/*margin-left: auto;*/
margin-right: auto;
max-width: 500px;
/*background: #F8F8F8;*/
padding: 30px 30px 20px 30px;
font: 12px Arial, Helvetica, sans-serif;
color: #666;
border-radius: 5px;
-webkit-border-radius: 5px;
-moz-border-radius: 5px;
}
.smart-green h1 {
font: 24px "Trebuchet MS", Arial, Helvetica, sans-serif;
padding: 20px 0px 20px 40px;
display: block;
margin: -30px -30px 10px -30px;
color: #FFF;
background: #9DC45F;
text-shadow: 1px 1px 1px #949494;
border-radius: 5px 5px 0px 0px;
-webkit-border-radius: 5px 5px 0px 0px;
-moz-border-radius: 5px 5px 0px 0px;
 border-bottom: 1px solid #89AF4C;
}
.smart-green h1 > span {
 display: block;
 font-size: 11px;
color: #FFF;
}
.smart-green label {
display: block;
margin: 0px 0px 5px;
}
.smart-green label > span {
float: left;
margin-top: 10px;
color: #5E5E5E;
}
.smart-green input[type="text"], .smart-green input[type="email"],
.smart-green textarea, .smart-green select {
color: #555;
height: 30px;
line-height: 15px;
width: 100%;
padding: 0px 0px 0px 10px;
margin-top: 2px;
border: 1px solid #E5E5E5;
background: #FBFBFB;
outline: 0;
-webkit-box-shadow: inset 1px 1px 2px rgba(238, 238, 238, 0.2);
box-shadow: inset 1px 1px 2px rgba(238, 238, 238, 0.2);
font: normal 14px/14px Arial, Helvetica, sans-serif;
}
.smart-green textarea {
height: 100px;
padding-top: 10px;
}
.smart-green .button {
background-color: #9DC45F;
border-radius: 5px;
-webkit-border-radius: 5px;
-moz-border-border-radius: 5px;
border: none;
padding: 10px 25px 10px 25px;
color: #FFF;
text-shadow: 1px 1px 1px #949494;
}
.smart-green .button:hover {
background-color: #80A24A;
 }
.error-msg{
color: red;
margin-top: 10px;
}
.success-msg{
color: #80A24A;
margin-top: 10px;
margin-bottom: 10px;
}
</style>
