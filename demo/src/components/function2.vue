<template>
  <div>
    <sidebar1></sidebar1>
    <div class="container">
      <!--blogsbox begin-->
      <br/>
      <h1>欢迎查看笔记</h1>
      <div class="blogsbox">
        <div
          v-for="item in blogData"
          :key="item.ID"
          class="blogs"
          data-scroll-reveal="enter bottom over 1s"
        >
          <div class="blogtitle">
            <a
              href="javascript:void(0);"
              @click="goToInfo(item.ID)"
              v-html="item.title"
            ></a>
          </div>
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
        </div>

        
      </div>
      <div class="sidebar">
        <Recommend class="re"></Recommend>
        <TagCloud></TagCloud>
        <FollowUs></FollowUs>
      </div>
    </div>
    <template>
  <el-backtop></el-backtop>
    </template>
  </div>
</template>

<script>


//import {getBlogByUid} from "../api/blogContent";
import axios from 'axios';
import sidebar1 from "./common/sidebar1.vue";
import Recommend from "../components/Recommend";
import TagCloud from "../components/TagCloud";
import FollowUs from "../components/FollowUs";
export default {
  data() {
    return {
      searchModel: 0, //搜索模式 0:SQL搜索、1:ES搜索、2:Solr搜索
      blogData: [],
      keywords:{
        title:'',
      },
      //currentPage: 1,
      //totalPages: 0,
      //pageSize: 10,
      total: 0, //总数量
      //tagUid: "",
      searchBlogData: [], //搜索出来的文章
      sortUid:{
        name:'',
      },
      //isEnd: false, //是否到底底部了
      //loading: false //内容是否正在加载
      author:{
        username:'',
      },
      user:{

      }
    };
  },
  components: {
    sidebar1,
    Recommend,
    TagCloud,
    FollowUs,
  },
  created() {
    this.search();
  },
  mounted() {
    
  },
  methods: {

    //跳转到文章详情
    goToInfo(blog) {
        let routeData = this.$router.resolve({
          path: "/info",
          query: {blogOid: blog,name:this.user.uid}
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
    // 加载内容
    loadContent: function() {
      this.currentPage = this.currentPage + 1;
      this.search();
    },
    convertSearchData: function (that, response) {
      if (response.data.status==200) {
        //that.isEnd = false;
        //获取总页数
        //that.totalPages = response.data.blogList.length;
        that.total = response.data.blogList.total;
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
      let searchModel = this.searchModel
      let that = this;
      //that.loading = true;
          const url = 'http://localhost:9090/api/article/search';
          const data = this.keywords;
          axios.post(url, data).then(response => {
            //console.log(this.keywords.title+'nnn');
            that.convertSearchData(that, response)
          });  
    }
        
}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less">
@import '../style/mixin';
.blogsbox { width: 66%; overflow: hidden; float: left;margin-top: 50px; }
.blogs { overflow: hidden; margin-bottom: 20px; padding: 20px; background:white; -webkit-box-shadow: 0 2px 5px 0 rgba(146, 146, 146, .1); -moz-box-shadow: 0 2px 5px 0 rgba(146, 146, 146, .1); box-shadow: 0 2px 5px 0 rgba(146, 146, 146, .1); -webkit-transition: all 0.6s ease; -moz-transition: all 0.6s ease; -o-transition: all 0.6s ease; transition: all 0.6s ease; }
.blogs .blogtitle { margin: 0 0 10px 0; font-size: 20px; overflow: hidden;}
.blogs .blogtitle a:hover { color:black; }
.blogs .blogtext { font-size: 14px; color: #566573; overflow: hidden; text-overflow: ellipsis; -webkit-box-orient: vertical; display: -webkit-box; -webkit-line-clamp: 3; margin-top: 20px }
.bloginfo { overflow: hidden; margin-top: 30px }
.bloginfo ul li {font-size: 12px; padding: 0 0 0 20px; color: #748594; line-height: 1.5; display: inline-block; }
.bloginfo ul li a { color: #748594; }
.bloginfo ul li a:hover { color: #000 }
.bloginfo .author { background: url(../images/auicon.jpg) no-repeat 0 0 }
.bloginfo .lmname { background: url(../images/auicon.jpg) no-repeat top -23px left; }
.bloginfo .createTime { background: url(../images/auicon.jpg) no-repeat top -44px left; }
.container { width: 1600px;height: 15000px; margin: 20px auto; background-color:gainsboro;position:relative;}
a{
  color: cornflowerblue;
}
.sidebar {width: 32%; overflow: hidden; float: right;margin-top:50px}
.sidebar div { padding: 35px 30px; margin-bottom: 20px; background: #FFF; -webkit-box-shadow: 0 2px 5px 0 rgba(146, 146, 146, .1); -moz-box-shadow: 0 2px 5px 0 rgba(146, 146, 146, .1); box-shadow: 0 2px 5px 0 rgba(146, 146, 146, .1); }
.sidebar div:hover ::after { width: 70px; }
.re{
  width:300px;
}
</style>


