<template>
  <div>
    <div class="commentBox">
    
    <span class="right">
      <textarea id="textpanel" class="textArea" placeholder="既然来了，那就留下些什么吧~ 【评论支持Markdown语法】" v-model="mes.value" @click="hideEmojiPanel" @input="vaildCount"></textarea>
    </span>

    </div>
    <div class="bottom">
      <el-button class="submit p2" type="primary"  @click="handleSubmit">发送评论</el-button>
      <el-button class="cancel p2" type="info" @click="handleCancle">取消评论</el-button>
<div class="tableTitle"></div>
      <el-popover
        placement="top"
        v-model="isShowEmojiPanel">
        

        <div class="emoji-panel-btn p2" @click="showEmojiPanel" slot="reference">
        </div>

      </el-popover>



      <span class="allow p2" v-if="isShowAvatar">还能输入{{count}}个字符</span>

    </div>
  </div>

</template>

<script>
import axios from 'axios';
  export default {
    name: 'CommentBox',
    props: {
      userInfo: {
        type: Object
      },
      // 回复的对象
      toInfo: {
        type: Object
      },
      // 博客信息
      commentInfo: {
        type: Object
      },
      showCancel: {
        type: Boolean,
        default: true
      },
      username: {
        type:String
      },
      id: {
        type:Number

      }
    },
    components: {
      
    },
    data() {
      return {
        comments: [],
        submitting: false,
        mes:{
          value:'',
          username1:this.username,
          id1:this.id,

        },
        user: {},
        count: 1024,
        isShowEmojiPanel: false, // 是否显示表情面板
        isShowAvatar: true, // 是否显示头像
      }
    },
    computed: {
      
    },
    mounted() {
      // 页面加载的时候调用监听
      //this.hideEmojiPanel()

      // 获取宽高
      window.onresize = () => {
        return (() => {
          this.resizeWin();
        })();
      };
      this.resizeWin();
    },
    methods: {
      //拿到vuex中的写的方法
      
      vaildCount: function() {
        var count = 1024 - this.mes.value.length;
        if(count <= 0) {
          this.count = 0
        } else {
          this.count = count;
        }
      },
      handleSubmit() {
        //let info = this.$store.state.user.userInfo
        //let isLogin = this.$store.state.user.isLogin
        

        if(this.value =="") {
          this.$notify.error({
            title: '警告',
            message: '评论内容不能为空哦~',
            offset: 100
          });
          return;
        }
        this.$emit("submit-box", this.mes)
       

        
      },
      emoji(word) {
        // 生成html
        const type = word.substring(1, word.length - 1);
        return `<span class="emoji-item-common emoji-${type} emoji-size-small"></span>`;
      },
      handleCancle() {
        this.value = '';
        this.count = 1024;
        this.isShowEmojiPanel = false
        if(this.toInfo) {
          this.$emit("cancel-box", this.toInfo.commentUid)
        }
        //this.hideEmojiPanel()
      },
      showEmojiPanel() {
        this.isShowEmojiPanel = !this.isShowEmojiPanel;
      },
      hideEmojiPanel() {
        this.isShowEmojiPanel = false
      },
      appendEmoji(text) {
        // const el = document.getElementById("textpanel");
        this.value = this.value + ":" + text + ":";
      },
      resizeWin() {
        //当前window 宽
        let centerWidth = document.documentElement.scrollWidth;
        if (centerWidth > 800) {
          this.isShowAvatar = true
        } else {
          // 是否显示头像框
          this.isShowAvatar = false
        }
      },
    },
  };
</script>


<style>
  .tableTitle {
    position: relative;
    margin: 0 auto;
    width: 1500px;
    height: 2px;
    background-color:black;
    text-align: center;
    font-size: 16px;
    color:black;
  }

  .el-popover {
    height: 135px;
    width: 420px;
  }

  .emoji-panel-wrap {
    box-sizing: border-box;
    border: 1px solid #cccccc;
    border-radius: 5px;
    background-color: #ffffff;
    width: 423px;
    height: 145px;
    position: absolute;
    z-index: 99;
    top: 10px;
  }
  .emoji-size-small {
    zoom: 0.3;
    margin: 5px;
    vertical-align: middle;
  }
  .emoji-size-large {
    zoom: 0.5; /* emojipanel表情大小 */
    margin: 5px;
  }
  .commentBox {
    /*min-width: 700px;*/
    width: 100%;
    height: 100px;
    margin: 0 auto;
  }
  .commentBox .left {
    display: inline-block;
    width: 4%;
    height: 100%;
    padding-top: 3px;
  }
  .commentBox .left img {
    cursor: pointer;
    margin: 0 auto;
    width: 90%;
    border-radius: 50%;
  }
  .commentBox .right {
    display: inline-block;
    margin: 5px 2px 0 0;
    width: 95%;
    height: 100%;
  }
  textarea::-webkit-input-placeholder {
    color: #909399;
  }
  .commentBox .right textarea {
    color: #606266;
    padding:10px 5px 5px 10px;
    resize: none;
    width: 98%;
    height: 100%;
  }
  .bottom {
    position: relative;
    width: 98%;
    height: 60px;
    line-height: 40px;
    margin-top: 30px;
  }
  .bottom .p2 {
    float: right;
    margin-top: 5px;
    margin-right: 10px;
  }
  .emoji-panel-btn img{
    height: 35px;
    width: 35px;
  }
  .emoji-panel-btn:hover {
    cursor: pointer;
    opacity: 0.8;
  }
  .emoji-item-common {
    
    display: inline-block;
  }
  .emoji-item-common:hover {
     cursor: pointer;
   }
  .emoji-size-small {
    /* 表情大小 */
    zoom: 0.3;
  }

  @media only screen and (min-width: 300px) and (max-width: 767px) {
    .commentBox .right {
      display: inline-block;
      margin: 5px 2px 0 0;
      width: 99%;
      height: 100%;
    }

    .emoji-panel-wrap {
      box-sizing: border-box;
      border: 1px solid #cccccc;
      border-radius: 5px;
      background-color: #ffffff;
      width: 300px;
      height: 220px;
      position: absolute;
      z-index: 99;
      top: 10px;
    }

    .el-popover {
      height: 220px;
      width: 300px;
    }

  }

</style>
