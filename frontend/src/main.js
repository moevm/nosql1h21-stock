import {createApp} from 'vue'
import App from './App.vue'
import {createRouter, createWebHistory} from 'vue-router'
import Home from './views/Home';
import Stock from './views/Stock';
import Comparator from './views/Comparator';

const routes = [
    {path: '/', component: Home},
    {name: 'stock', path: '/stock/:ticker', component: Stock},
    {path: '/comparator', component: Comparator},
];

const router = createRouter({
    history: createWebHistory(''),
    routes: routes,
})

createApp(App).use(router).mount('#app')
