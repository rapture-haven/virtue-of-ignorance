import Vue from "vue";
import App from "./App.vue";
import store from "./store";

Vue.config.productionTip = false;

new Vue({
  store,
  render: (h) => h(App),
  created() {
    this.$client.init(this);
    this.$accessor.initialise();
  },
}).$mount("#virtue");
