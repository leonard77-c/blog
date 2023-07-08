<template>
<div class="zhuanti">
<h2 class="hometitle">特别推荐</h2>
<ul>
 <li style="cursor: pointer" @click="goToInfo()"> <i><img :src="imgurl"></i>
<p @click="goToInfo()" style="cursor: pointer"><span>阅读</span> </p>
 </li>
</ul>
</div>
</template>
<script>
export default {
  name: 'Recommend',
    data() {
    	return {
        slideList: [],
	    thirdData: [], //；一级推荐数据
        imgurl: require('../../assets/s.png'),

    	}
    },
    created() {
     
    },
    methods: {
      //跳转到文章详情【或推广链接】
      goToInfo() {
        let routeData = this.$router.resolve({
          path: "/info",
          query: {blogOid: 597,name:'hey'}
        });
        window.open(routeData.href, '_blank'); 
    },
      
    },

}
</script>

<style>
.hometitle { font-size: 18px; color: #282828; font-weight: 600; margin: 0; text-transform: uppercase; padding-bottom: 15px; margin-bottom: 25px; position: relative; }
.hometitle:after { content: ""; background-color: #282828; left: 0; width: 50px; height: 2px; bottom: 0; position: absolute; -webkit-transition: 0.5s; -moz-transition: 0.5s; -ms-transition: 0.5s; -o-transition: 0.5s; transition: 0.5s; }
.zhuanti { overflow: hidden;width:300px;height: 300px; }
.zhuanti ul { width: 100%; overflow: hidden; display: block; clear: both; }
.zhuanti li { margin: 0 0 20px 0; overflow: hidden; height: 200px; position: relative; }
.zhuanti li i { background: #000; display: block }
.zhuanti img { width: 100%; -moz-transition: all .5s ease; -webkit-transition: all .5s ease; -ms-transition: all .5s ease; -o-transition: all .5s ease; transition: all .5s ease; opacity: 0.5 }
.zhuanti p { position: absolute; top: 30%; left: 0; right: 0; color: #FFF; text-align: center; font-size: 15px; overflow: hidden; margin-top: 5px; padding: 0 40px; }
.zhuanti p a { color: #fff; }
.zhuanti span { width: 80px; margin: 10px auto; background: transparent; font-size: 12px; border: 1px solid #FFF; border-radius: 40px; padding: 4px 0; color:#000; display: block; clear: both; -webkit-transition: all .3s ease; -moz-transition: all .3s ease; -ms-transition: all .3s ease; -o-transition: all .3s ease; transition: all .3s ease; }
.zhuanti li:hover img { opacity: 0.6 }
.zhuanti li:hover span { background: #FFF; }
.zhuanti li:hover span a { color: #333 }
</style>
