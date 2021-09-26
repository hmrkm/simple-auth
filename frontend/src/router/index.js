import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router);

export const router = new Router({
    scrollBehavior() {
        return window.scrollTo({ top: 0, behavior: 'smooth' });
    },
    routes: [        
        {
            path: '/',
            name: 'login',
            component: () => import('../views/Home.vue'),
        }
    ]
})