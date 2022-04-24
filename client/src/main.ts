import './assets/styles/main.scss'

import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import Notifications from 'vue-notification'
import Client from './plugins/virtue-of-ignorance'
import Logger from './plugins/log'
import Swal from './plugins/swal'

Vue.config.productionTip = false;

Vue.use(Notifications)
Vue.use(Logger)
Vue.use(Swal)
Vue.use(Client)

new Vue({
  router,
  store,
  render: (h) => h(App),
  created() {
    this.$client.init(this)
    this.$accessor.initialise()
  },
}).$mount("#app");
