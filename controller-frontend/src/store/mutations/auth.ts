import {State} from "@/store/types";
import setAuthToken from "@/plugins/axios";

export function setSetupMode(state: State, setupMode = false) {
    state.setup.setupMode = setupMode
}

export function view_totp(state: State, base64: string | undefined) {
    if (!base64) {
        state.setup.totpQRCode = undefined
        state.setup.totpTokenVisible = false
        state.setup.setupMode = false
        return
    }

    state.setup.totpQRCode = base64
    state.setup.totpTokenVisible = true
}

export function set_auth_token(state: State, {token = "", username = ""}) {
    if (token !== "" && username !== "") {
        state.auth = {
            sessionToken: token,
            userName: username
        }

        localStorage.setItem("auth", JSON.stringify(state.auth))
        setAuthToken(token)
        return
    }

    localStorage.removeItem("auth")
    state.auth = undefined
    setAuthToken(undefined)
}
