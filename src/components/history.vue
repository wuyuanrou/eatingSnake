<template>
  <div id='history' v-if="showHistory">
    <!-- 代码已整体移到best组件，该组件已作废 -->
        <!-- <div class='history_title'>历史战绩榜</div>
        <el-table :data="historyData">
            <el-table-column
                prop="num"
                label="序号"
                width="100px">
            </el-table-column>
            <el-table-column
                prop="length"
                label="长度"
                width="100px">
            </el-table-column>
            <el-table-column
                prop="mode"
                label="模式">
            </el-table-column>
        </el-table> -->
    </div>
</template>

<script>
//import {mapGetters} from 'vuex'

export default {
    //name: 'HistoryComponent',
    // data(){
    //     return {
    //         historyData: [{
    //             num:1,
    //             length: 200,
    //             mode: '测试数据'
    //             }, {
    //             num:2,
    //             length: 400,
    //             mode: '测试数据'
    //         }],
    //     };
    // },
    // computed:{
    //     ...mapGetters('history', ['historyData'])
    // },
}
</script>

<style>
    /* #history{
        z-index: 10;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        border: lightblue solid 5px;
        border-radius: 10px;
        height: 500px;
        width: 300px;
        background-color: white;
    }

    .history_title{
        height: 50px;
        line-height: 50px;
        text-align: center;
        font-weight: bold;
    } */
</style>