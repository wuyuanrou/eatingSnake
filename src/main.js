import Vue from 'vue'
import App from './App.vue'

//引入element-ui
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

//记得引入并注册vuex仓库！！！！！
import store from '@/store'

Vue.config.productionTip = false;
Vue.use(ElementUI);

new Vue({
  render: h => h(App),
  //注册全局事件总线
  beforeCreate() {
    Vue.prototype.$bus=this
  },
  store:store
}).$mount('#app')
