<template>
  <div class="share">
    <p class="diggit" @click="praiseBlog(blogUid)">
      <a href="javascript:void(0);">很赞哦！</a>
      <span >
        (<b id="diggnum">{{praise}}</b>)
      </span>
    </p>
    <p class="dasbox">
      <a href="javascript:void(0)" @click="dashangToggle()" class="dashang" title="打赏，支持一下">打赏本站</a>
    </p>
    <div class="shang_box" v-if="showPay">
      <a class="shang_close" href="javascript:void(0)" @click="dashangToggle()" title="关闭">关闭</a>
      <div class="shang_tit">
        <p>感谢您的支持，我会继续努力的!</p>
      </div>
      <div class="shang_payimg">
        <img :src="payCode" alt="扫码支持" title="扫一扫">
      </div>
      <div class="pay_explain">扫码打赏，你说多少就多少</div>
      <div class="shang_payselect">
        <div
          class="pay_item"
          :class="[this.payMethod==1?'checked':'']"
          data-id="alipay"
          @click="choosePay('1')"
        >
          <span class="radiobox"></span>
          <span class="pay_logo">
            <img src="../../../static/images/alipay.jpg" alt="支付宝">
          </span>
        </div>
        <div
          class="pay_item"
          :class="[this.payMethod==2?'checked':'']"
          data-id="weipay"
          @click="choosePay('2')"
        >
          <span class="radiobox"></span>
          <span class="pay_logo">
            <img src="../../../static/images/wechat.jpg" alt="微信">
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {
  praiseBlogByUid,
  getBlogPraiseCountByUid
} from "../../api/blogContent";
import axios from 'axios';
import aliPayPhoto from"../../../static/images/2.jpg"
import weixinPayPhoto from"../../../static/images/1.jpg"
export default {
  name: "PayCode",
  props: {
    praise_count: {
      type: Number,
      default: 0
    },
    blogUid: {
      type: Number
    },
    username: {
        type:String
    }
  },
  data() {
    return {
      webConfigData: [],
      showPay:false, //是否显示支付
      payMethod: 1, // 1: 支付宝  2：微信
      payCode: "", //支付码图片
      praise:this.praise_count,
      blog:{
        uid:this.blogUid,
      },
      pic:{
        aliPayPhoto,
        weixinPayPhoto
      }
    };
  },
  created() {
    
  },
  methods: {
    //拿到vuex中的写的方法
    
    dashangToggle: function() {
      this.showPay = !this.showPay;
    },
    // 支付方式
    choosePay: function(type) {
      this.payMethod = type;
      if (type == 1) {
        this.payCode = this.pic.aliPayPhoto;
      } else {
        this.payCode = this.pic.weixinPayPhoto;
      }
    },
    //博客点赞
    praiseBlog: function(uid) {
      // 判断用户是否登录
      const url = 'http://localhost:9090/api/article/praise';
      const data = this.blog;
      axios.post(url,data).then(response => {
          this.$notify({
            title: '成功',
            message: "点赞成功",
            type: 'success',
            offset: 100
          });
          this.praise = response.data.praise_count;
          this.$emit('update:praiseCount',response.data);

      });
    },
    //获取点赞数
    getPraiseCount: function(uid) {
      var params = new URLSearchParams();
      params.append("uid", uid);
      getBlogPraiseCountByUid(params).then(response => {
        if (response.code == this.$ECode.SUCCESS) {
          this.praiseCount = response.data;
        }
      });
    }
  }
};
</script>

<style>
  .diggit {
    cursor: pointer;
  }
</style>