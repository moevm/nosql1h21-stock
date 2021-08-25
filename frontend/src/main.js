import { createApp } from 'vue'
import App from './App.vue'
import { createRouter, createWebHistory } from 'vue-router'
import Home from './views/Home';
import Todos from './views/Todos';

const routers = [
    {path: '/', component: Home},
    {path: '/todos', component: Todos}
];

const router = createRouter({
    history: createWebHistory(''),
    routes: routers,
})

createApp(App).use(router).mount('#app')
