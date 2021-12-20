import Vue from 'vue'
import VueRouter, {RouteConfig} from 'vue-router'
import Dashboard from '../views/Dashboard.vue'
import store from "@/store";

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
    {
        path: '/',
        name: 'Dashboard',
        component: Dashboard,
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import("@/views/public/Login.vue"),
        meta: {
            public: true
        }
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.public)) {
        next();
        return;
    }

    //TODO implement isLoggedIn check
    if (store.getters.isAuthenticated) {
        next();
        return;
    }

    next("/login");
});


export default router
