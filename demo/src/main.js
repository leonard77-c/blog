// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import constant from './utils/constant'
import axios from 'axios'
import VueAxios from 'vue-axios'
import '../static/css/index.css';
//axios.defaults.baseURL = '/api'
Vue.use(VueAxios,axios)
Vue.use(ElementUI)
Vue.prototype.$constant = constant
Vue.config.productionTip = false
Vue.prototype.$axios = axios
//main.js
//main.js
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
