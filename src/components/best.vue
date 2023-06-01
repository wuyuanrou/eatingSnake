<template>
    <div>
        <div id='best'>
            <div class='best_title'>
                最佳战绩榜
            </div>
            <div class="show_btn">
                <el-button type="text" size="mini" @click="showHistory">点这里查看最近30次战绩</el-button>
            </div>
            <div>
                <el-table :data="bestData">
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
            </el-table>
            </div>
            <!-- <History v-if="showHistory"></History> -->
        </div>
        <div id='history' v-if="isShowHistory">
            <div>
                <div class='history_title'>历史战绩榜</div>
                <el-table :data="historyData" height="380">
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
            </el-table>
            </div>
            <div class='show_btn'>
                <el-button type="danger" @click="closeHistory" plain>关闭</el-button>
            </div>
        </div>
    </div>
</template>

<script>
//import History from './history.vue'
import {mapGetters} from 'vuex'

export default {
    name: 'BestComponent',
    data(){
        return {
            isShowHistory: false,
            // bestData: [{
            //     num:1,
            //     length: 1000,
            //     mode: '最佳-测试数据'
            //     }, {
            //     num:2,
            //     length: 800,
            //     mode: '最佳-测试数据'
            // }],
            // historyData: [{
            //     num:1,
            //     length: 200,
            //     mode: '测试数据'
            //     }, {
            //     num:2,
            //     length: 400,
            //     mode: '测试数据'
            // }],
        };
    },
    components:{
        //History,
    },
    computed:{
        ...mapGetters('history', ['historyData']),
        ...mapGetters('best', ['bestData'])
    },
    methods:{
        showHistory(){
            //console.log('发送')
            this.$store.dispatch('history/getRecentHistoryData')
            this.isShowHistory=true
        },
        closeHistory(){
            this.isShowHistory=false
        },
    }
}
</script>

<style>
    #best{
        border: lightblue solid 5px;
        border-radius: 10px;
        height: 500px;
        width: 300px;
        margin-left: 20px;
        background-color: white;
    }

    .best_title{
        height: 50px;
        line-height: 50px;
        text-align: center;
        font-weight: bold;
    }

    #history{
        z-index: 10;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        border: lightcoral solid 5px;
        border-radius: 10px;
        height: 500px;
        width: 300px;
        background-color: white;

        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-content: center;
    }

    .history_title{
        height: 50px;
        line-height: 50px;
        text-align: center;
        font-weight: bold;
    }

    .show_btn{
        text-align: center;
        margin: 10px;
    }
</style>