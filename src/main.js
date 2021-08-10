import Vue from 'vue';
import Vuex from 'vuex';
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import axios from 'axios';
import VueAxios from 'vue-axios';
import Vuelidate from 'vuelidate';
import App from './App.vue';
import router from './router';
import store from './store';

// Import Bootstrap an BootstrapVue CSS files (order is important)
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

// scss style
import './assets/scss/index.scss';

Vue.config.productionTip = false;
Vue.use(Vuex);

Vue.use(Vuelidate);
// Make BootstrapVue available throughout your project
Vue.use(BootstrapVue);
// axios
Vue.use(VueAxios, axios);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
