import { createApp } from "vue";
import ArcoVue from '@arco-design/web-vue';
import "./style.css";
import App from "./App.vue";
import '@arco-design/web-vue/dist/arco.css';
//import VueRouter from 'vue-router'

const app = createApp(App);
app.use(ArcoVue)//.use(VueRouter);
app.mount('#app');
