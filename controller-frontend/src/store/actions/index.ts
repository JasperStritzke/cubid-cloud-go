import trigger_alert from "@/store/actions/alert";
import {next_snackbar, add_snackbar} from "@/store/actions/snackbar";
import logout from "@/store/actions/auth";
import {show_dialog} from "@/store/actions/dialog";

export default {
    trigger_alert,
    next_snackbar, add_snackbar, logout,
    show_dialog
}
