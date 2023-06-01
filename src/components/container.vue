<template>
  <div id='container'>
    <div id='content'>
        <div id='center'>
            <div class='left'>
                <div id='mode'>
                    <el-radio-group v-model="radio1">
                        <el-radio-button label="普通模式" @click.native="changeToNormalMode"></el-radio-button>
                        <el-radio-button label="困难模式" @click.native="changeToHardMode"></el-radio-button>
                    </el-radio-group>
                </div>
                <Screen :mode="mode"></Screen>
                <!-- <div id='board'>
                    <div>
                        <el-statistic group-separator="," :precision="2" decimal-separator="." :value="length" :title="title">
                            <template slot="prefix">
                            <i class="el-icon-s-flag" style="color: red"></i>
                            </template>
                            <template slot="suffix">
                            <i class="el-icon-s-flag" style="color: blue"></i>
                            </template>
                        </el-statistic>
                    </div>
                </div> -->
            </div>
            <Best></Best>
        </div>
    </div>
  </div>
</template>

<script>
import Best from './best.vue'
import Screen from './screen.vue'

export default {
    name:'ContainerComponent',
    inject:['reload'],
    provide(){
        return {reload:this.reload}
    },
    data () {
      return {
        radio1: '普通模式',
        mode:'normal',
        // length: 4154.564,
        // title: '长度',
      };
    },
    components:{
        Best,
        Screen
    },
    methods:{
        async getHistoryData_Normal(){
            try{
                //console.log('请求普通模式战绩')
                await this.$store.dispatch('best/getBestHistoryDataByMode', 'normal')
            }catch(error){
                console.log(error.message)
            }
        },
        async getHistoryData_Hard(){
            try{
                //console.log('请求困难模式战绩')
                await this.$store.dispatch('best/getBestHistoryDataByMode', 'hard')
            }catch(error){
                console.log(error.message)
            }
        },
        changeToNormalMode(){
            this.mode='normal'
            this.getHistoryData_Normal()
        },
        changeToHardMode(){
            this.mode='hard'
            this.getHistoryData_Hard()
        },
    },
    mounted(){
        this.$store.dispatch('best/getBestHistoryDataByMode', 'normal')
    }
}
</script>

<style>
    #container{
        padding-top: 20px;
        padding-bottom: 20px;
    }

    #content{
        margin-top: 10px;
        margin-bottom: 10px;
    }

    #mode{
        height: 50px;
        width: 750px;
        margin-bottom:10px;
    }

    #center{
        display: flex;
        justify-content: center;
        align-items:flex-end;
    }

    .left{
        display: flex;
        flex-direction: column;
        justify-content: center;
    }

    /* #board{
        height: 50px;
        width: 800px;
        margin-top: 10px;
        display: flex;
        justify-content: center;
    } */
</style>