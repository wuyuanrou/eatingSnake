//统一管理api
//引入配置好封装好的axios
import request from './request.js'

//封装请求/收到响应的操作
export const requestBestHistoryDataByMode=(mode)=>request({method: 'GET',url: `/history/getBestHistoryDataByMode/${mode}`})
export const requestBestHistoryData=()=>request({method: 'GET',url: '/history/getBestHistoryData'})
export const requestRecentHistoryData=()=>request({method: 'GET',url: '/history/getRecentHistoryData'})
//不带请求参数，至少要传一个空对象进去，否则请求会失败
export const reqInsertHistoryData=(params)=>request({method:'POST', url: '/history/insertHistoryData', 
    data: params,headers: {
    'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
  }})
  //这里post不加head，后端无法正常读取到参数。因为后端用的postform要求数据是urlencoded
/*
export const reqGoodsInfo=(skuId)=>request({method:'GET',url:`/item/${skuId}`})
*/