<template>
  <div id="app">
    <Header></Header>
    <Container v-if="isRouterAlive"></Container>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from './components/header.vue'
import Container from './components/container.vue'
import Footer from './components/footer.vue'

export default {
  name: 'App',
  provide(){
    return{
      reload:this.reload
    }
  },
  components: {
    Header,
    Container,
    Footer
  },
  data(){
    return{
      isRouterAlive:true
    }
  },
  methods:{
    reload(){
      this.isRouterAlive=false
      this.$nextTick(function(){
        this.isRouterAlive=true
      })
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50
}
</style>
