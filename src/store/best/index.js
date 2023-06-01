import { requestBestHistoryDataByMode, requestBestHistoryData} from "@/api/index"

const state={
    bestData:[]
}

const actions={
    async getBestHistoryData({commit}){
        //console.log('actions')
        let result=await requestBestHistoryData()
        //console.log(result)
        if(result.code==200){
            //console.log("最佳数据获取成功")
            commit('GetBestHistoryData',result.data)
        }
    },
    async getBestHistoryDataByMode({commit}, mode){
        //console.log('actions')
        let result=await requestBestHistoryDataByMode(mode)
        //console.log(result)
        if(result.code==200){
            //console.log("最佳数据获取成功")
            commit('GetBestHistoryDataByMode',result.data)
        }
    },
}

const mutations={
    GetBestHistoryData(state, bestData){
        state.bestData=bestData
    },
    GetBestHistoryDataByMode(state, bestData){
        state.bestData=bestData
    }
}

const getters={
    bestData(state){return state.bestData||[]},
}

export default{
    namespaced: true,
    state, 
    actions, 
    mutations, 
    getters
}