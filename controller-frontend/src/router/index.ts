import Vue from 'vue'
import VueRouter, {RouteConfig} from 'vue-router'
import store from "@/store";
import Login from "@/views/public/Login.vue"

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
    {
        path: '/',
        name: 'Dashboard',
        component: () => import("@/views/Dashboard.vue"),
    },
    {
        path: '/login',
        name: 'Login',
        component: Login,
        meta: {
            public: true
        }
    },
    {
        path: "/executors",
        name: "Executors",
        component: () => import("@/views/Executors.vue"),
    },
    {
        path: "/executors/:name",
        name: "Executors Detail",
    },
    {
        path: "/*",
        name: "404",
        beforeEnter() {
            store.dispatch('trigger_alert', {type: "warning", text: "404 - Route not found"}).catch()

            if(router.currentRoute.path === "/") return
            router.push("/").catch()
        }
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

router.afterEach(to => {
    document.title = "cubid.cloud | " + to.name
})

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.public)) {
        if (to.name === "Login") {
            if (store.getters.isAuthenticated) {
                next("/")
                return
            }
        }

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
