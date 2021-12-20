export function add_snackbar(context: any, message: string) {
    context.commit('add_snackbar', message)
}

export function next_snackbar(context: any) {
    if(context.state.snackbar.queue.length > 0) {
        context.commit('set_snackbar', context.state.snackbar.queue[0])
        return;
    }

    context.commit("hide_snackbar")
}
