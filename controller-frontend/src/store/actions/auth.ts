import {ActionContext} from "vuex";
import router from "@/router";

export default function logout(context: ActionContext<any, any>) {
    context.commit('set_auth_token', {token: "", username: ""})

    context.dispatch('trigger_alert', {
        type: "success",
        text: "Logged out successfully."
    }).catch()

    router.push("/login").catch()
}
