import {trigger_alert, remove_alert} from "@/store/mutations/alert";
import {hide_snackbar, set_snackbar, add_snackbar} from "@/store/mutations/snackbar";
import {setSetupMode, view_totp, set_auth_token} from "@/store/mutations/auth";
import {show_dialog, hide_dialog} from "@/store/mutations/dialog";

export default {
    trigger_alert, remove_alert,
    hide_snackbar, set_snackbar, add_snackbar,
    setSetupMode, view_totp, set_auth_token,
    show_dialog, hide_dialog
}

