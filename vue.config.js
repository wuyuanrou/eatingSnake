// const { defineConfig } = require('@vue/cli-service')
// module.exports = defineConfig({
//   transpileDependencies: true
// })

module.exports = {
    // devServer: {
    //     host: '0.0.0.0',
    //     // https:true,
    //     port: 6103,
    //     client: {
    //       webSocketURL: 'ws://0.0.0.0:6103/ws',
    //     },
    //     headers: {
    //       'Access-Control-Allow-Origin': '*',
    //     }
    // },
    // transpileDependencies: true,
    devServer:{
      proxy:{
        '/api':{
            target:`http://127.0.0.1:8000`, //请求后台接口
            changeOrigin:true, // 允许跨域
        }
      }
    }
}