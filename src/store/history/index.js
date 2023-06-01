import { requestRecentHistoryData} from "@/api/index"

const state={
    historyData:[],
}

const actions={
    async getRecentHistoryData({commit}){
        //console.log('actions')
        let result=await requestRecentHistoryData()
        //console.log(result)
        if(result.code==200){
            //console.log("历史数据获取成功")
            commit('GetRecentHistoryData',result.data)
        }
    },
}

const mutations={
    GetRecentHistoryData(state, historyData){
        state.historyData=historyData
    }
}

const getters={
    //如果没拿到数据，就会返回undefined，模板里读取undefined里的属性就会报错。但是读取空对象里的属性，虽然是undefined但不会报错
    //拆分state里面的数据
    historyData(state){return state.historyData||[]},
}

export default{
    namespaced: true,
    state, 
    actions, 
    mutations, 
    getters
}