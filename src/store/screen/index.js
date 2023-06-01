import { reqInsertHistoryData} from "@/api/index"

const state={
    
}

const actions={
    async insertHistoryData(_, params){
        //console.log('插入actions')
        //console.log(params)
        let result=await reqInsertHistoryData(params)
        //console.log(result)
        if(result.code==200){
            console.log("插入成功")
        }
    },
}

const mutations={
    
}

const getters={
    
}

export default{
    namespaced: true,
    state, 
    actions, 
    mutations, 
    getters
}