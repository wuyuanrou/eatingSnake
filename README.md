# eating_snake

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

axios前后端通信，vuex任意组件间通信
HistoryData是一维数组，数据元素是对象，对象成员有序号num+长度length+模式mode，即[{num:, length:, mode:}, {num:, length:, mode:}]
mysql表存放id+长度，取后30次成绩（30个对象），服务器处理好以数组形式返回给前端，接口地址为'/api/history/getHistoryData/normal或hard',方式为GET
模式不同获取的数据不同