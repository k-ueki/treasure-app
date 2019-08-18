<template>
  <div id="app">
	  <!-- <Idea></Idea> -->
	  <!-- <IdeasPreview></IdeasPreview> -->
	  <!-- <Auth></Auth> -->
	  <router-view></router-view>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Idea from './components/pages/Idea.vue'
import IdeasPreview from './components/pages/IdeasPreview.vue'
import Auth from './components/pages/Auth.vue'

export default {
  name: 'app',
	components: {
		HelloWorld,
		Idea,
		IdeasPreview,
		Auth
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
