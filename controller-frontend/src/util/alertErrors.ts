import store from "@/store";

export function alertErrors(e: Error) {
    store.dispatch('trigger_alert', {type: "error", text: "An error occurred: " + e}).catch()
}

export function alertErrorsWithCustomMessage(message: string): () => void {
    return () => {
        store.dispatch('trigger_alert', {type: "error", text: message}).catch()
    }
}
