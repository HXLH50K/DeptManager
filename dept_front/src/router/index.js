import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Dept from '../components/Dept.vue'
import DeptManager from '../components/DeptManager.vue'
import DeptAdd from '../components/DeptAdd.vue'
import App from '../App.vue'
import Index from '../components/Index.vue'
import DeptUpdate from '../components/DeptUpdate.vue'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: '部门管理',
      component: Index,
      redirect: '/Manager',
      children:[
        {
          path: '/Manager',
          name: '部门查询',
          component: DeptManager
        },{
          path: '/Add',
          name: '部门添加',
          component: DeptAdd
        },{
          path: '/Update',
          name: '部门编辑',
          component: DeptUpdate
        }
      ]
    }
  ]
})
