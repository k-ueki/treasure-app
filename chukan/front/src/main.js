import Vue from 'vue'
import App from './App.vue'
import firebase from 'firebase/app'

import router from './router.js'

Vue.config.productionTip = false

console.log("test",process.env.VUE_APP_TEST);

new Vue({
	router,
  render: h => h(App),
}).$mount('#app')
