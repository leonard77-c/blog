<template>
<div>
    <sidebar1></sidebar1>
    <div class="box">
      <el-button class="btn2" type="primary" @click="sho">点击刷新</el-button>
    <div v-if="flag" class="blogsbox">
        <div
          v-for="item in blogData"
          :key="item.ID"
          class="blogs"
          data-scroll-reveal="enter bottom over 1s"
        >
          <h3 class="blogtitle">
            <a
              href="javascript:void(0);"
              @click="goToInfo(item.ID)"
              v-html="item.title"
            ></a>
          </h3>
          <p class="blogtext" v-html="item.desc"></p>
          <div class="bloginfo">
            <ul>
              <li class="author">
                
                <a href="javascript:void(0);" @click="goToAuthor(item.User.username)">{{item.User.username}}</a>
              </li>
              <li class="lmname" >
                
                <a href="javascript:void(0);" @click="goToList(item.Category.name)">{{item.Category.name}}</a>
              </li>
              <li class="createTime">{{item.CreatedAt}}</li>
            </ul>
          </div>
          <div class="delete">
            <el-button type="danger" icon="el-icon-delete" circle @click="del(item.ID)"></el-button>
          </div>
        </div>
    </div>
      <el-card class="person">
  <div slot="header" class="clearfix">
    <span>个人信息</span>
    <el-button style="float: right" type="success" @click="save">保存</el-button>
  </div>
  <div>
    <i class="el-icon-s-custom">账户名称</i>
    <el-input v-model="user.username" placeholder="请输入内容"></el-input>
    <div></div><br/>
    <i class="el-icon-finished">个性签名</i>
    <el-input v-model="desc" placeholder="请输入内容"></el-input>
    <div></div><br/>
    <i class="el-icon-chat-dot-round">微信</i>
    <el-input v-model="wechat" placeholder="请输入内容"></el-input>
    <div></div><br/>
    <i class="el-icon-message">邮箱</i>
    <el-input v-model="email" placeholder="请输入内容"></el-input>
    <div></div><br/>
    <i class="el-icon-key">修改密码</i>
    <el-input v-model="code" placeholder="请输入内容" @keyup.enter.native="rewrite"></el-input>


  </div>
</el-card>
    </div>
    <template>
  <el-backtop></el-backtop>
    </template>
</div>
</template>
<script>
import sidebar1 from "./common/sidebar1.vue";
import axios from 'axios';
import bus from "../utiles/bus.js";
import CommentBox from "./CommentBox";
export default {
    components:{
        sidebar1,
        CommentBox,
    },
    created() {
        this.user.username = this.$route.query.username;
        this.user.uid = this.$route.query.uid;
        this.desc=this.$route.query.desc;
        this.wechat=this.$route.query.wechat;
        this.email=this.$route.query.email;
        this.code='';
        //sessionStorage.setItem("user",JSON.stringify(user));
        setTimeout(()=>{
          bus.$emit('share',this.user);
          //console.log(this.user.username);
        },15000);
        this.search();
    },
    data(){
        return{
          code:'',
          desc:'',
          wechat:'',
          email:'',
          size:'',
            user:{
                username:'',
                uid:0,
            },
            blogData: [],
            total: 0,
            searchBlogData: [],
            flag: true,
            items:[],
        }
    },
    methods: {
      rewrite(){
        let params = {};
                this.user.uid = parseInt(this.user.uid)
                params.id = this.user.uid
                params.password = this.code;
                const url = 'http://localhost:9090/api/password_change';
            axios.post(url,params).then(response => {
            console.log(response.data);
            this.$message.success('修改密码成功！')
            });
      },
      save(){
         let params = {};
                params.id = this.user.uid
                params.name = this.user.username;
                params.desc = this.desc;
                params.qq_chat = '';
                params.wechat = this.wechat;
                params.webo = '';
                params.bili = '';
                params.email = this.email;
                params.img = '';
                params.icp_record = '';
                const url = 'http://localhost:9090/api/updateinfo';
            axios.post(url,params).then(response => {
            console.log(response.data);
            this.$message.success('保存成功！')
            });

      },
        convertSearchData: function (that, response) {
      if (response.data.status==200) {
        //that.isEnd = false;
        //获取总页数
        //that.totalPages = response.data.blogList.length;
        that.total = response.data.blogList.Total;
        
        //that.pageSize = response.data.pageSize;
        //that.currentPage = response.data.currentPage;
        var blogData = response.data.blogList.ArticleList;

        // 判断搜索的博客是否有内容
        if(response.data.total <= 0) {
          that.isEnd = true;
          that.loading = false;
          this.blogData = []
          return;
        }
        //全部加载完毕
        //if (blogData.length < that.pageSize) {
        //  that.isEnd = true;
        //}
        blogData = that.searchBlogData.concat(blogData);
        that.searchBlogData = blogData;
        this.blogData = blogData;
      } else {
        //that.isEnd = true;
      }
      //that.loading = false;
    },
    search: function() {
      let that = this;
      //const user = JSON.parse(sessionStorage.getItem("user"));
      if (this.user.username != undefined) {
          const url = 'http://localhost:9090/api/admin/article/search';
          const data = this.user;
          axios.post(url, data).then(response => {
            //console.log(this.keywords.title+'nnn');
            that.convertSearchData(that, response)
          });
        
        
      } 
    },
    sho(){
      this.blogData.splice(0,this.total);
      this.search();
      
    },
    goToInfo(blog) {
        let routeData = this.$router.resolve({
          path: "/info",
          query: {blogOid: blog}
        });
        window.open(routeData.href, '_blank'); 
    },
    //点击了分类
    goToList(uid) {
      let routeData = this.$router.resolve({
        path: "/list",
        query: { sortUid: uid }
      });
      window.open(routeData.href, '_blank');
    },
    goToAuthor(author) {
      let routeData = this.$router.resolve({
        path: "/list",
        query: {author: author}
      });
      window.open(routeData.href, '_blank');
    },
    del(id){
      const data={
                    aid:id,
                  };
      const url = 'http://localhost:9090/api/article/delete';
      axios.post(url, data).then(response => {
        console.log(data.ID);
                });
      this.blogData.splice(0,1);
    } 
    },
    
}
</script>
<style>
.container { width: 1600px;height: 1000px; margin: 20px auto; background-color:gainsboro;position:relative;}
.blogsbox { width: 66%; overflow: hidden; float: left;margin-top: 50px; }
.blogs { overflow: hidden; margin-bottom: 20px; padding: 20px; background:white; -webkit-box-shadow: 0 2px 5px 0 rgba(146, 146, 146, .1); -moz-box-shadow: 0 2px 5px 0 rgba(146, 146, 146, .1); box-shadow: 0 2px 5px 0 rgba(146, 146, 146, .1); -webkit-transition: all 0.6s ease; -moz-transition: all 0.6s ease; -o-transition: all 0.6s ease; transition: all 0.6s ease; }
.blogs .blogtitle { margin: 0 0 10px 0; font-size: 20px; overflow: hidden;}
.blogs .blogtitle a:hover { color:black; }
.blogs .blogtext { font-size: 14px; color: #566573; overflow: hidden; text-overflow: ellipsis; -webkit-box-orient: vertical; display: -webkit-box; -webkit-line-clamp: 3; margin-top: 20px }
.bloginfo { overflow: hidden; margin-top: 30px }
.bloginfo ul li {  font-size: 12px; padding: 0 0 0 20px; color: #748594; line-height: 1.5; display: inline-block; }
.bloginfo ul li a { color: #748594; }
.bloginfo ul li a:hover { color: #000 }
.bloginfo .author { background: url(../images/auicon.jpg) no-repeat 0 0 }
.bloginfo .lmname { background: url(../images/auicon.jpg) no-repeat top -23px left; }
.bloginfo .createTime { background: url(../images/auicon.jpg) no-repeat top -44px left; }
a{
  color: cornflowerblue;
}
.btn2{
  position:relative;
  top: 0px;
  left:-800px;
}
.box { width: 1600px;height: 10000px; margin: 20px auto; background-color:gainsboro;position:relative;}
.text {
    font-size: 14px;
  }

  .item {
    margin-bottom: 18px;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
    content: "";
  }
  .clearfix:after {
    clear: both
  }

  .box-card {
    width: 480px;
  }
  .person {
    width:300px;
    margin-top: 0px;
    margin-left:1100px;
    height: 500px;
  }
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
</style>