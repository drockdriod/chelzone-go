import Vue from 'vue'
import './plugins/vuetify'
import './plugins/vueloader'
import App from './App.vue'
import store from './store'
import router from './routes'


Vue.config.productionTip = false

new Vue({
    store,
  	router,
    render: h => h(App),
}).$mount('#app')