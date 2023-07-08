<template>
<div>
    <div class="xxx fillcontain">
            <div class="form_contianer" v-show="showLogin">
                  <div class="manage_tip">
                      <p>博客系统demo</p>
                  </div>
                <el-form >
                    <el-form-item prop="username">
                        
                        <i class="el-icon-user-solid"></i>
                        <el-input  v-model="msg.username" placeholder="用户名"><span>dsfsf</span></el-input>
                    </el-form-item>
                    <el-form-item prop="password">
                        <i class="el-icon-lock"></i>
                        <el-input type="password" placeholder="密码" v-model="msg.password"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="show" class="submit_btn">登录</el-button>
                      </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="reg" class="submit_btn">注册</el-button>
                      </el-form-item>
                </el-form>
                <!--p class="tip">温馨提示：</p>
                <p class="tip">未登录过的新用户，自动注册</p>
                <p class="tip">注册过的用户可凭账号密码登录</p-->
            </div>
		
    </div>
</div>
</template>
<script>
import axios from 'axios';
import sidebar1 from "./common/sidebar1.vue";
export default{
    components:{
        sidebar1,
        
    },
    data(){
        return{
            msg:{
                username:'',
                password:'',
            },
            ret:{
                status:'',
                data:'',
                id:'',
                message:'',
            },
            showLogin:true,
            desc:'',
            wechat:'',
            email:'',
        }
    },
    methods:{
        show(){
        const data = this.msg;
        const url = 'http://localhost:9090/api/login';
        axios.post(url, data)
        .then(response => {
        console.log(response.data);
        this.ret=response.data;
         let params = {};
                params.username = this.msg.username
                const url = 'http://localhost:9090/api/getpro';
            axios.post(url,params).then(response => {
            console.log(response.data);
            this.desc = response.data.data.desc;
            this.wechat = response.data.data.wechat;
            this.email = response.data.data.email;
            if(this.ret.status==200){
            this.$router.push({path: '/context',query: {username: this.msg.username,uid:this.ret.id,desc:this.desc,wechat:this.wechat,email:this.email}})
            
        }else{
            if(this.ret.status==1003){
                alert('登录失败！,用户不存在！')
            }
            if(this.ret.status==1002){
                alert('登录失败！,密码错误！')
            }
            

        }
            });

        })
        .catch(error => {
        console.log(error.response.data);
        });
        
            
        },
		reg(){
			this.$router.push({path: '/reg'})
		}
    }

}
</script>

<style scoped lang="less">
@import '../style/mixin';
.el-icon-lock{
     font-size: 25px;
}
.el-icon-user-solid{
    //display: inline-block;
    //width:150px;
    //text-align: right;
    font-size: 25px;
}
.xxx{
      background-color: #324057;
	  position:fixed;
	  top:0px;
}
.manage_tip{
        position: absolute;
        width: 100%;
        top: -100px;
        left: 0;
        p{
            font-size: 34px;
            color:aqua;
        }
    }
.form_contianer{
        .wh(320px, 210px);
        .ctp;
        padding: 25px;
        border-radius: 5px;
        text-align: center;
        background-color:white;
        .submit_btn{
            width: 100%;
            font-size: 16px;
        }
    }
.tip{
        font-size: 12px;
        color: red;
    }	
.form-fade-enter-active, .form-fade-leave-active {
	  	transition: all 1s;
	}
	.form-fade-enter, .form-fade-leave-active {
	  	transform: translate3d(0, -50px, 0);
	  	opacity: 0;
	}
.el-input {
 
    position: relative;
 
    font-size: 14px;
 
    display: inline-block;
 
    width: 270px;
 
}
</style>