import {createApp} from 'vue'
import App from './App.vue'
import {createRouter, createWebHistory} from 'vue-router'
import Home from './views/Home';
import Stock from './views/Stock';

const routers = [
    {path: '/', component: Home},
    {name: 'stock', path: '/stock/:ticker', component: Stock}
];

const router = createRouter({
    history: createWebHistory(''),
    routes: routers,
})

createApp(App).use(router).mount('#app')
