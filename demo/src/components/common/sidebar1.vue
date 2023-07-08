
<template>
  <el-breadcrumb separator="/">
  <el-breadcrumb-item><router-link to="/context">关于我</router-link></el-breadcrumb-item>
  <el-breadcrumb-item><router-link to="/function2">大厅</router-link></el-breadcrumb-item>
  <el-breadcrumb-item><router-link to="/function1">发布笔记</router-link></el-breadcrumb-item>
  <el-breadcrumb-item><router-link to="/function3">查找笔记</router-link></el-breadcrumb-item>
  <el-button class="btn1" type="danger" @click="out">退出登录</el-button>
  <el-breadcrumb-item></el-breadcrumb-item>

  </el-breadcrumb>
</template>
<script>
  export default {
    data() {
      return {
        activeIndex: '1',
        activeIndex2: '1',
        activeName: 'second'
      };
    },
    methods: {
      handleSelect(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClick(tab, event) {
        //console.log(tab, event);
        if(tab.index==0){
          this.$router.push('/context')
        }else if(tab.index==1){
          this.$router.push('/function2')
        }else if(tab.index==2){
          this.$router.push('/function1')
        }else{
          this.$router.push('/function3')
        }
      },
      out(){
        this.$router.push('/')
      }
      
    }
  }
</script>
<style scoped>
.router-link-active {
  text-decoration: none;
  color:brown;
}
a{
  text-decoration: none;
  color:darkslateblue;
}
.btn1{
  margin-left:1000px;
  margin-top:0px;
}
</style>
