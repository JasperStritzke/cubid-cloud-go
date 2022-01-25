export interface Alert {
    type: string,
    duration: number,
    text: string,
    id: number,
}

export interface Dialog {
    title: string,
    text: string | undefined,
    type: "success" | "warning" | "error",
    confirmButton: string | undefined,
    cancelText: string | undefined,
    danger: boolean | undefined,
    resolve: (value: boolean) => void | undefined,
    id: string,
    waitTime: number
}

export interface AuthState {
    userName: string,
    sessionToken: string,
}

export interface SetupState {
    setupMode: boolean,
    totpTokenVisible: boolean,
    totpQRCode: string | undefined
}

export interface State {
    alerts: Alert[],
    snackbar: {
        queue: string[],
        value: boolean,
        text: string | undefined
    },
    dialogs: Dialog[],
    auth: undefined | AuthState,
    setup: SetupState,
}

