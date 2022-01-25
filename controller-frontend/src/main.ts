import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import axios from "axios";
import store from "@/store";

Vue.config.productionTip = false

axios.defaults.baseURL = "http://192.168.178.24:6080/api"
//DEV: axios.defaults.baseURL = "/api"
axios.defaults.headers.common["Content-Type"] = "application/json";

new Vue({
    router,
    store,
    vuetify,
    render: h => h(App)
}).$mount('#app')

axios.get("/auth/checkSetupMode")
    .then(result => result.data)
    .then(({setupMode, authValid}) => {
        store.commit('setSetupMode', setupMode)

        if (!authValid && store.state.auth) {
            store.commit('set_auth_token', {token: "", username: ""})
            store.dispatch('trigger_alert', {
                type: "warning",
                text: "Session timed out"
            }).catch()

            router.push("/login").catch()
        }
    })
    .catch(() => {
        store.commit('set_auth_token', {token: "", username: ""})
        store.dispatch("trigger_alert", {type: "error", text: "Backend is unavailable."}).catch()
        router.push("/login").catch()
    })
