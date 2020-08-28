<template>
  <div id="app">
    <Questions />
  </div>
</template>

<script>
import Questions from './components/Questions.vue'

export default {
  name: 'App',
  components: {
    Questions
  }
}
</script>

<style>
body {
  background-color: darkgray;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 20px;
}

.btn {
  background-color: dimgrey;
  border: none;
  color: black;
  padding: 15px 15px;
  margin: 10px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  text-transform: uppercase;
}
</style>
