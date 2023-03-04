import { createApp } from "vue";
import ArcoVue from '@arco-design/web-vue';
import "./style.css";
import App from "./App.vue";
import '@arco-design/web-vue/dist/arco.css';
//import VueRouter from 'vue-router'
import router from "./router";

const app = createApp(App);
app.use(ArcoVue).use(router);
app.mount('#app');
