import {State} from "@/store/types";
import store from "@/store";

export function add_snackbar(state: State, payload: string) {
    state.snackbar.queue.push(payload)

    if (!state.snackbar.value) {
        store.dispatch("next_snackbar").catch()
    }
}

export function set_snackbar(state: State, payload: string) {
    return new Promise(async (resolve: any) => {
        await hide_snackbar(state)

        setTimeout(() => {
            state.snackbar.value = true;
            state.snackbar.text = payload;

            state.snackbar.queue.pop()

            resolve()
        }, 100)
    })
}

export async function hide_snackbar(state: State) {
    state.snackbar.value = false;
    state.snackbar.text = undefined;
}
