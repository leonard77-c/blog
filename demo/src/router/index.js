import Vue from 'vue'
import Router from 'vue-router'
import admin1 from '@/components/admin1'
import context from '@/components/context'
import function1 from '@/components/function1'
import function2 from '@/components/function2'
import function3 from '@/components/function3'
import function4 from '@/components/function4'
import function5 from '@/components/function5'
import reg from '@/components/reg'
import list from '@/components/list'
import info from '@/components/info'
Vue.use(Router)
export default new Router({
  routes: [
    {
      path: '/',
      name: 'admin1',
      component: admin1,
    },
    {
          path:"/context",
          name:'context',
          component: context,
          meta:{
            keepAlive:true
          }
    },
    
    {
      path:"/function1",
      name:'function1',
      component:function1,
      meta:{
        keepAlive:true
      }
    },
    {
      path:"/function2",
      name:'function2',
      component:function2,
      meta:{
        keepAlive:true
      }
    },
    {
      path:"/function3",
      name:'function3',
      component:function3,
      meta:{
        keepAlive:true
      }
    },
    {
      path:"/function4",
      name:'function4',
      component:function4,
    },
    {
      path:"/function5",
      name:'function5',
      component:function5,
    },
    {
      path:"/reg",
      name:reg,
      component:reg,
    },
    {
      path:"/list",
      component:list,
    },
    {
      path:"/info",
      component:info,
    }
      
    
    
      
    
  ]
})
