import "./assets/styles/main.scss";

import Vue from "vue";
import App from "./App.vue";
import store from "./store";

import { i18n } from "@/plugins/i18n";
import Notifications from "vue-notification";
import Client from "./plugins/virtue-of-ignorance";
import Logger from "./plugins/log";
import Swal from "./plugins/swal";
import Axios from "./plugins/axios";
import ToolTip from "v-tooltip";

Vue.config.productionTip = false;

Vue.use(Notifications);
Vue.use(ToolTip);
Vue.use(Logger);
Vue.use(Swal);
Vue.use(Client);
Vue.use(Axios);

new Vue({
  i18n,
  store,
  render: (h) => h(App),
  created() {
    this.$client.init(this);
    this.$accessor.initialise();
  },
}).$mount("#virtue");
