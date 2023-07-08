<template>
<div class="sss fillcontain">
    <div class="register">
        <h1>欢迎来到注册界面</h1>
        <div class="kuang">
        <el-form>
            <el-form-item prop="username">
                        <el-input class="p1" v-model="msg_r.username" placeholder="用户名"></el-input>
                    </el-form-item>
                    <el-form-item prop="password">
                        <el-input class="p2" type="password" placeholder="密码" v-model="msg_r.password"></el-input>
                    </el-form-item>
                    <el-form-item prop="password">
                        <el-input class="p3" type="password" placeholder="请确认密码" v-model="password2"></el-input>
                    </el-form-item>
                    <el-form-item prop="password">
                        <el-input class="p5" placeholder="请输入邮箱" v-model="mail.email"></el-input>
                    </el-form-item>
                    <el-form-item prop="password">
                        <el-input class="p4" placeholder="请输入验证码" v-model="code2"></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button  type="primary" @click="re" class="submit_btn">确认</el-button>
                      </el-form-item>
        </el-form>
        </div>
    
    <el-button type="primary" class="btn4" @click="sendCode()"  :disabled="getCodeBtnDisable" icon="el-icon-s-promotion" circle></el-button>
    </div>
   
</div>
</template>
<script>
import axios from 'axios';
export default {
    data(){
        return{
            msg_r:{
                username:'',
                password:'',
                mail:'',

            },
            password2:'',
            ret_r:{
                status:'',
                message:'',
            },
            loginForm: {
            phoneNumber: '',
            verificationCode: '',
        },
        codeBtnWord: '获取验证码', // 获取验证码按钮文字
        waitTime:60, // 获取验证码按钮失效时间
        getCodeBtnDisable: false,
        mail: {
            email:'',
        },
        code:'',
        code2:'',
        }
    },
   
    methods:{
        re(){
            if(this.msg_r.password!=this.password2){
                alert('两次密码不一致，请重新输入！')
            }else if(this.code!=this.code2){
                alert('验证码不正确！')
                }else{
                const data=this.msg_r;
                const url = 'http://localhost:9090/api/register';
                axios.post(url, data)
                .then(response => {
                console.log(response.data);
                this.ret_r=response.data;
                if(this.ret_r.status==200){
            alert('注册成功')
            //console.log(response.data);
        }else{
            alert('用户已存在')
        }
                }
                )

        }
    },
    async sendCode () {
      try {
        const data=this.mail;
        const url='http://localhost:9090/api/email'
        //axios.post(url,data)
        axios.post(url, data).then(response => {
            //console.log(this.keywords.title+'nnn');
           
               this.code = response.data.code;
            
            });
        this.$message.success('验证码发送成功')
        let that = this
        that.waitTime--
        that.getCodeBtnDisable = true
        this.codeBtnWord = `${this.waitTime}s 后重新发送`
        let timer = setInterval(function () {
          if (that.waitTime > 1) {
            that.waitTime--
            that.codeBtnWord = `${that.waitTime}s 后重新发送`
          } else {
            clearInterval(timer)
            that.codeBtnWord = '获取验证码'
            that.getCodeBtnDisable = false
            that.waitTime = 60
          }
        }, 1000)
      } catch (res) {
        console.log(res)
      }
    },
    
}
}
</script>
<style scoped lang="less">
@import '../style/mixin';
.p1{
   width: 335px; 
}
.p2{
   width: 335px; 
}
.p3{
   width: 335px; 
}
.p4{
   width: 335px; 
}
.p5{
    width:250px;
    margin-left: -32px;
}
.register {
    .wh(330px, 400px);
    .ctp;
    padding: 25px;
    border-radius: 5px;
    //text-align: center;
    background-color:lightsteelblue;
    .submit_btn{
            width: 335px;
            font-size: 16px;
            margin-left: 0px;
        }
    .btn4{
        //padding-right: 80px;
        //padding-left: 100px;
        position: fixed;
        top:390px;
        left:890px;
        text-align: center;
        //font-size: 25px;
    }
}
.sss{
    background-color:#324057;
	  position:fixed;
	  top:0px;
}
h1{
    color:aliceblue
}
.kuang{
    width:280px;
}

</style>