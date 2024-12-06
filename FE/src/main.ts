import './style.css'
import "preline/preline";

import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate';
import {MotionPlugin} from '@vueuse/motion'
const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);



const renderApp = () => {
  const app = createApp(App);
  
  app.use(pinia);
  app.use(MotionPlugin)
  app.use(router);
  app.mount("#app");
};
renderApp();
