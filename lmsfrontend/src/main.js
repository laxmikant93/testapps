// src/main.js
import Vue from 'vue';
import App from './App.vue';
import router from './router';  // Import router.js

Vue.config.productionTip = false;

new Vue({
  render: h => h(App),
  router,  // Add the router to your Vue instance
}).$mount('#app');
