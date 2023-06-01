<template>
  <div id='header'>
    <span id='text'>进击吧，小蛇蛇！</span>
  </div>
</template>

<script>
export default {
  name: 'HeaderComponent'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  #header{
    text-align: center;
    height: 100px; 
    background-color: skyblue;
  }

  #text{
    line-height: 100px;
    font-size: 30px;
  }
</style>
