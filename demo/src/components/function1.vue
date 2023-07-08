<template>
    <div>
        <sidebar1></sidebar1>
        <div class="box">
          
        <el-form>
          <el-form-item>
          <i class="el-icon-coordinate"></i>
        <el-input
  class="form1"
  type="text"
  placeholder="请输入标题"
  v-model="msg_t.title"
  maxlength="10"
  show-word-limit
>
</el-input>
          </el-form-item>
          <el-form-item>
<i class="el-icon-magic-stick"></i>
<el-input
  class="form2"
  type="text"
  placeholder="请输入概要"
  v-model="msg_t.desc"
  maxlength="10"
  show-word-limit
>
</el-input>
          </el-form-item>
        </el-form>     
<i class="el-icon-edit-outline"></i>
<el-input
  class="form3"
  type="textarea"
  autosize
  placeholder="请输入内容"
  v-model="msg_t.content">
</el-input>
<div style="margin: 20px 0;"></div>
<i class="el-icon-connection"></i>
<div class="desc">
    <el-radio-group v-model="radio1">
      <el-radio-button label="学习"></el-radio-button>
      <el-radio-button label="旅游"></el-radio-button>
      <el-radio-button label="电影"></el-radio-button>
      <el-radio-button label="音乐"></el-radio-button>
      <el-radio-button label="科技"></el-radio-button>
      <el-radio-button label="时尚"></el-radio-button>
      <el-radio-button label="体育"></el-radio-button>
      <el-radio-button label="生活"></el-radio-button>
      <el-radio-button label="娱乐"></el-radio-button>
      <el-radio-button label="汽车"></el-radio-button>
    </el-radio-group>
  </div>
<div class="button">
<el-button type="primary" @click="update">上传<i class="el-icon-upload el-icon--right"></i></el-button>
</div>



</div>
    </div>
</template>
<script>
import sidebar1 from "./common/sidebar1.vue";
import axios from 'axios';
import bus from "../utiles/bus.js"
export default {
    components:{
        sidebar1,
    },
    data() {
    return {
      text_date: '',
      hono:true,
      msg_t:{
      content: '',
      title:'',
      desc:'',
      img:'',
      comment_count:0,
      read_count:0,
      category:{
        id:1,
        name:'cyf',
      }
      },
      ret_t:{
        status: '',
        data: '',
        message:'',
      },
      use:{},
      cid:0,
      dis1:false,
      dis2:false,
      dis3:false,
      dis4:false,
      dis5:false,
      radio1:'',

    }
  },
  created(){
    bus.$on('share',data=>{
          //console.log(val)
          this.use=data;
        });
  },
  methods:{
    update(){
      this.use.uid = parseInt(this.use.uid)
      if(this.radio1=='学习'){
        this.cid=1;
      }else if(this.radio1=='旅游'){
        this.cid=2;
      }else if(this.radio1=='电影'){
        this.cid=3;
      }else if(this.radio1=='音乐'){
        this.cid=4;
      }else if(this.radio1=='科技'){
        this.cid=5;
      }else if(this.radio1=='时尚'){
        this.cid=6;
      }else if(this.radio1=='体育'){
        this.cid=7;
      }else if(this.radio1=='生活'){
        this.cid=8;
      }else if(this.radio1=='娱乐'){
        this.cid=9;
      }else if(this.radio1=='汽车'){
        this.cid=10;
      }
      const data={
                    uid:this.use.uid,
                    title:this.msg_t.title,
                    cid:this.cid,
                    desc:this.msg_t.desc,
                    content:this.msg_t.content,
                    img:''
                  };
                console.log(data);
                const url = 'http://localhost:9090/api/article/add';
                axios.post(url, data)
                .then(response => {
                console.log(response.data);
                this.ret_t=response.data;
                if(this.ret_t.status==200){
            alert('上传成功')
            //console.log(response.data);
        }else{
            alert('上传失败')
        }
                }
                )

    },
    p1(){
      this.cid=1;
      this.dis1=!this.dis1;
    },
    p2(){
      this.cid=2;
      this.dis2=!this.dis2;
    },
    p3(){
      this.cid=3;
      this.dis3=!this.dis3;
    },
    p4(){
      this.cid=4;
      this.dis4=!this.dis4;
    },
    p5(){
      this.cid=5;
      this.dis5=!this.dis5;
    },


  }
}
</script>
<style scoped lang="less">
@import '../style/mixin';
.text1{
    width:250px;
    .ctp;
    top:100px;
}
.text2{
    width:200px;
    .ctp;
    top:200px;
}
.text3{
    width:300px;
    .ctp;
    top:300px;
}
.button{
    .ctp;
    margin-top:50px;
    position: relative;
}
.desc{
  .ctp;
  top:0px;
  left:0px;
  position: relative;
 
    font-size: 14px;
 
   display: inline-block;
 
    width: 370px;

}
.el-icon-coordinate{
     font-size: 25px;
     margin-top:70px;
     
}
.el-icon-magic-stick{
  font-size: 25px;
}
.el-icon-edit-outline{
  font-size: 25px;
}
.el-icon-connection{
  font-size: 25px;
  //margin-top:200px;
}
.form1{
   position: relative;
 
    font-size: 14px;
 
    display: inline-block;
 
    width: 370px;
    top:0px;
}
.form2{
   position: relative;
 
    font-size: 14px;
 
    display: inline-block;
 
    width: 370px;
    top:0px;
}
.form3{
   position: relative;
 
    font-size: 14px;
 
    display: inline-block;
 
    width: 370px;
    top:0px;
}
.s3{
  position: relative;
 
    font-size: 14px;
 
    display: inline-block;
 
    width: 370px;

}
.box { width: 1600px;height: 10000px; margin: 20px auto; background-color:gainsboro;position:relative;}
</style>