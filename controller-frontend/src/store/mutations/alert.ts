import store from "@/store";
import {State} from "@/store/types";

export function trigger_alert(state: State, {type = "warning", duration = 5000, text = "No text provided", id = -1}) {
    if (id < 0) return

    state.alerts.push({
        type,
        duration,
        text,
        id
    })

    setTimeout(() => store.commit('remove_alert', id), duration)
}

export function remove_alert(state: State, id: number) {
    state.alerts = state.alerts.filter(alert => alert.id !== id)
}
