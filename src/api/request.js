///二次封装axios
import axios from 'axios';

//引入进度条
import nprogress from 'nprogress';
//是一个对象，有start和done方法
//一定要引入样式，不然没效果。修改样式到nprogress.css改
import 'nprogress/nprogress.css'

//创建并配置axios实例
const request=axios.create({
    baseURL: "/api",
    timeout: 5000,//超过5秒就算是请求失败
});

//请求拦截器。在发请求之前，可以使用拦截器对请求做一些统一的处理，比如加token
request.interceptors.request.use((config)=>{
    nprogress.start();
    return config;
});

//拦截器收到服务器响应时统一进行处理后再返回响应的数据
request.interceptors.response.use((res)=>{
    nprogress.done();
    return res.data;
}, (error)=>{
    return Promise.reject(new Error('failed to interceptors', error));
});

//记得要暴露
export default request;