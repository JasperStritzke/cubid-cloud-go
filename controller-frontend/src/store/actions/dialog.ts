import {ActionContext} from "vuex";
import {Dialog} from "@/store/types";

export function show_dialog(context: ActionContext<any, any>, dialog: Dialog) {
    return new Promise(resolve => {
        dialog.resolve = resolve
        context.commit("show_dialog", dialog)
    })
}
