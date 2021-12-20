import store from "@/store";

export default function alertErrors(e: Error) {
    store.dispatch('trigger_alert', {type: "error", text: "An error occurred: " + e}).catch()
}
