import Vue from 'vue'
import Vuex from 'vuex'

//引入小仓库
import history from './history'
import best from './best'
import screen from './screen'

Vue.use(Vuex)

export default new Vuex.Store({
    //合并小仓库
    modules:{
        history,
        best,
        screen
    }
})