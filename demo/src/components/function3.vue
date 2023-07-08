<template>
    <div>
        <sidebar1></sidebar1>
        <div class="box">
  <el-col :span="8" class="r1">
    <el-card shadow="hover" class="card">
  <el-input
    placeholder="请输入内容"
    prefix-icon="el-icon-search"
    v-model="value1"
    @keyup.enter.native="search">
  </el-input>
  <el-select v-model="value" placeholder="请选择" class="r2">
    <el-option
      v-for="item in options"
      :key="item.value"
      :label="item.label"
      :value="item.value">
    </el-option>
  </el-select>
    </el-card>
  </el-col>
    </div>
    </div>
</template>
<script>
import sidebar1 from "./common/sidebar1.vue";
import bus from "../utiles/bus.js"
export default {
    components:{
        sidebar1,
    },
    data(){
      return{
        options: [{
          value: 1,
          label: '根据类型搜索'
        }, {
          value: 2,
          label: '根据标题搜索'
        }, {
          value: 3,
          label: '根据作者搜索'
        }, 
        ],
        value: '',
        input2:'',
        input3:'',
        tableData:'',
        keyword:'',
        searchModel:0,
        uid:'',
        author:'',
        name:'',
        user:{},
        value1:'',

      }
    },
    created(){
      bus.$on('share',data=>{
          //console.log(val)
          this.user=data;
        });
    },
    methods:{
      search(){
        if(this.value==1){
          this.uid=this.value1;
          this.search1();
        }else if(this.value==2){
          this.keyword=this.value1;
          this.search2();
        }else if(this.value==3){
          this.author=this.value1;
          this.search3();
        }
      },
      search2: function () {
      if (this.keyword == "" || this.keyword.trim() == "") {
        this.$notify.error({
          title: '错误',
          message: "关键字不能为空",
          type: 'success',
          offset: 100
        });
        return;
      }
      this.$router.push({path: "/list", query: {keyword: this.keyword, model: this.searchModel,user:this.user}});
    },
    search1: function () {
      if (this.uid == "" || this.uid.trim() == "") {
        this.$notify.error({
          title: '错误',
          message: "关键字不能为空",
          type: 'success',
          offset: 100
        });
        return;
      }
      this.$router.push({path: "/list", query: {sortUid: this.uid,model: this.searchModel,user:this.user}});
    },
    search3: function () {
      if (this.author == "" || this.author.trim() == "") {
        this.$notify.error({
          title: '错误',
          message: "关键字不能为空",
          type: 'success',
          offset: 100
        });
        return;
      }
      this.$router.push({path: "/list", query: {author: this.author,model: this.searchModel,user:this.user}});
    },
    }
}
</script>
<style>
  .el-header {
    background-color: #B3C0D1;
    color: #333;
    line-height: 60px;
  }
  
  .el-aside {
    color: #333;
  }
  .r1{
    margin-left:100px;
  }
  .r2{
    left: 10px;
    top:0px;
  }
  .card{
    width: 800px;
    margin-top:100px;
    margin-left: 250px;
  }
  .el-input{
   position: relative;
 
    font-size: 14px;
 
    display: inline-block;
 
    width: 370px;
}
.box { width: 1600px;height: 10000px; margin: 20px auto; background-color:gainsboro;position:relative;}
</style>