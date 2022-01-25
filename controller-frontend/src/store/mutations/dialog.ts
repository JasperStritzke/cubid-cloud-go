import {Dialog, State} from "@/store/types";
import {randomInsecureString} from "@/util/random";

export function show_dialog(state: State, dialog: Dialog) {
    dialog.id = randomInsecureString(6)

    state.dialogs.push(dialog)
}

export function hide_dialog(state: State, dialogId: string) {
    const index = state.dialogs.map(({id}) => id).indexOf(dialogId)

    state.dialogs.splice(index, 1)
}
