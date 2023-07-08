<template>
  <article>
    <div class="infosbox">
      <div class="newsview">
        <h3 class="news_title" v-if="blogData.data">{{blogData.data.title}}</h3>
        <div class="bloginfo" >
          <ul>
            <li class="author" v-if="blogData.data">
              <a href="javascript:void(0);" @click="goToAuthor(blogData.author)">{{blogData.data.User.username}}</a>
            </li>
            <li class="lmname" v-if="blogData.data">
              <a
                href="javascript:void(0);"
                @click="goToSortList(blogData.blogSort.uid)"
              >{{blogData.data.Category.name}}</a>
            </li>
            <li class="createTime" v-if="blogData.data">
              {{blogData.data.CreatedAt}}
            </li>
           </ul>
        </div>
        <div class="news_con ck-content" v-if="blogData.data" v-html="blogData.data.content"></div> 
        </div>
        <PayCode v-if="blogData.data" :blogUid="blogData.data.ID" :praise_count.sync="blogData.data.praise_count"></PayCode>
      </div>
      <div class="news_pl" :style="opemCommentCss">
        <h2>文章评论</h2>
        <ul>
            <CommentBox
              :userInfo="userInfo"
              :commentInfo="commentInfo"
              @submit-box="submitBox"
              :showCancel="showCancel"
              :id="blogOid.id"
              :username="user"
            ></CommentBox>
          <div class="message_infos">
            <CommentList :comments="comments" :commentInfo="commentInfo" :maxReplyLevel="4"></CommentList>
            <div class="noComment" v-if="this.total==0">还没有评论，快来抢沙发吧！</div>
          </div>
        </ul>
      </div>
  </article>
</template>

<script>
    import CommentList from "./CommentList";
    import CommentBox from "./CommentBox";
    // vuex中有mapState方法，相当于我们能够使用它的getset方法
    import { addComment, getCommentList } from "../api/comment";
    import { getBlogByUid, getSameBlogByBlogUid } from "../api/blogContent";
    import axios from 'axios';
    import PayCode from "../components/PayCode";
    export default {
        name: "info",
        data() {
            return {
                // 目录列表数
                user:'',
                name:'',
                catalogSum: 0,
                showStickyTop: false,
                showSideCatalog: true,
                showSidebar: true, //是否显示侧边栏
                blogContent: "",
                catalogProps: {
                  // 内容容器selector(必需)
                  container: '.ck-content',
                  watch: true,
                  levelList: ["h2", "h3"],
                },
                loadingInstance: null, // loading对象
                showCancel: false,
                submitting: false,
                comments: [],
                commentInfo: {
                    // 评论来源： MESSAGE_BOARD，ABOUT，BLOG_INFO 等 代表来自某些页面的评论
                    source: "BLOG_INFO",
                    blogUid: this.$route.query.blogUid
                },
                currentPage: 1,
                pageSize: 10,
                total: 0, //总数量
                toInfo: {},
                userInfo: {},
                blogUid: null, //传递过来的博客uid
                blogOid: {
                    id:0,
                }, // 传递过来的博客oid
                blogData:{
                    blogSort: {}
                },
                sameBlogData: [], //相关文章
                dialogPictureVisible: false,
                dialogImageUrl: "",
                openComment: "0", // 开启评论
                openAdmiration: "0", // 开启赞赏
            };
        },
        computed: {
          vueCategory: function () {
            if(!this.showStickyTop && this.showSideCatalog) {
              return 'catalog'
            }
            if(!this.showStickyTop && !this.showSideCatalog) {
              return 'catalog'
            }
            if(this.showStickyTop && this.showSideCatalog) {
              return 'catalog3'
            }
            if(this.showStickyTop && !this.showSideCatalog) {
              return 'catalog2'
            }
          },
          opemCommentCss: function () {
            if(this.openComment == 0) {
              return {
                "min-height": "10px"
              }
            }
          }
        },
        components: {
            //注册组件
            CommentList,
            CommentBox,
            PayCode,
        },
      watch: {
        $route(to, from) {
          location.reload()
        }
      },
        created() {
            this.blogOid.id = this.$route.query.blogOid;
            this.user=this.$route.query.name;
            this.blogOid.id = parseInt(this.blogOid.id)
            const url = 'http://localhost:9090/api/admin/article/info';
            const data = this.blogOid;
            axios.post(url, data).then(response => {
            //console.log(this.keywords.title+'nnn');
           
                this.blogData = response.data;
            
            });
           this.getCommentDataList();
        },
        methods: {
            
            //拿到vuex中的写的两个方法
            handleCurrentChange: function(val) {
                this.currentPage = val;
                this.getCommentDataList();
            },
            getSameBlog() {
              var blogParams = new URLSearchParams();
              blogParams.append("blogUid", this.blogUid);
              getSameBlogByBlogUid(blogParams).then(response => {
                if (response.code == this.$ECode.SUCCESS) {
                  this.sameBlogData = response.data.records;
                }
              });
            },
            // 设置是否开启评论和赞赏
            setCommentAndAdmiration() {
              //let webConfigData = this.$store.state.app.webConfigData
              if(webConfigData.createTime) {
                this.openAdmiration = webConfigData.openAdmiration
                this.openComment = webConfigData.openComment
              } else {
                getWebConfig().then(response => {
                  if (response.code == this.$ECode.SUCCESS) {
                    webConfigData = response.data;
                    // 存储在Vuex中
                    this.setWebConfigData(response.data)
                    this.openAdmiration = webConfigData.openAdmiration
                    this.openComment = webConfigData.openComment
                  }
                });
              }
            },
            submitBox(e) {
                let params = {};
                e.username1 = parseInt(e.username1)
                params.username1 = e.username1;
                params.value = e.value;
                params.id1 = e.id1;
                const url = 'http://localhost:9090/api/comment/add';
            axios.post(url,params).then(response => {
            console.log(response.data);
            this.$message.success('评论发布成功！')
            this.getCommentDataList();
            });
            },
            getCommentDataList: function() {
                let params = {};
                params.ID = this.blogOid.id;
                const url = 'http://localhost:9090/api/admin/article/comment/search';
                 axios.post(url, params).then(response => {
                        this.comments = response.data.commentList.Comment;
                        //this.setCommentList(this.comments);
                        //this.currentPage = response.data.current;
                        //this.pageSize = response.data.size;
                        this.total = response.data.commentList.Total;
                });
            },
            //跳转到文章详情
            goToInfo(item) {
                let routeData = this.$router.resolve({
                    path: "/info",
                    query: { blogUid: item }
                });
                window.open(routeData.href, "_blank");
            },
            //跳转到搜索详情页
            goToList(uid) {
                let routeData = this.$router.resolve({
                    path: "/list",
                    query: { tagUid: uid }
                });
                window.open(routeData.href, "_blank");
            },
            //跳转到搜索详情页
            goToSortList(uid) {
                let routeData = this.$router.resolve({
                    path: "/list",
                    query: { sortUid: uid }
                });
                window.open(routeData.href, "_blank");
            },
            //跳转到搜索详情页
            goToAuthor(author) {
                let routeData = this.$router.resolve({
                    path: "/list",
                    query: { author: author }
                });
                window.open(routeData.href, "_blank");
            },

            imageChange: function(e) {
                //首先需要判断点击的是否是图片
                var type = e.target.localName;
                if (type == "img") {
                    // window.open(e.target.currentSrc);
                    this.dialogPictureVisible = true;
                    this.dialogImageUrl = e.target.currentSrc;
                }
            },
            //切割字符串
            subText: function(str, index) {
                if (str.length < index) {
                    return str;
                }
                return str.substring(0, index) + "...";
            }
        }
    };
</script>

<style>
  .newsview {
    background-image: linear-gradient(90deg,rgba(50,0,0,.05) 3%,transparent 0),linear-gradient(1turn,rgba(50,0,0,.05) 3%,transparent 0);
    background-size: 20px 20px;
    background-position: 50%;
  }
  .emoji-panel-wrap {
    box-sizing: border-box;
    border: 1px solid #cccccc;
    border-radius: 5px;
    background-color: #ffffff;
    width: 470px;
    height: 190px;
    position: absolute;
    z-index: 999;
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
  .iconfont {
    font-size: 14px;
    margin-right: 3px;
  }
  .message_infos {
    width: 96%;
    margin-left: 10px;
  }
  .noComment {
    width: 100%;
    text-align: center;
  }
  .catalog {
    position: fixed;
    margin-left: 20px;
    /*max-height: 700px*/
  }
  .catalog2 {
    position: fixed;
    margin-left: 20px;
    top: 70px;
  }
  .catalog3 {
    position: fixed;
    margin-left: 20px;
    top: 20px;
  }
  .line-style {
    display: inline-block;
    height: 20px;
    width: 3px;
    background: transparent;
  }
  .line-style--active {
    background: currentColor;
  }
  .infosbox { float: left; width: 74%; overflow: hidden; background: #FFF; }
  .infosbox .news_con { line-height: 1.8; font-size: 16px; text-align: justify; }
.infosbox .news_con p { margin-bottom: 25x }
.infosbox .news_con img { max-width: 740px; height: auto; }
 .infosbox .news_con video { max-width: 740px; height: auto; }
.infosbox .share { padding: 20px; }
.infosbox { width: 100% }
.infosbox .news_con img {
  vertical-align:middle;
  width: 270px;
  text-align:center;
  border: 1px solid #eee;
}
</style>
