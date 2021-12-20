export interface Alert {
    type: string,
    duration: number,
    text: string,
    id: number,
}

export interface AuthState {
    loggedIn: boolean,
    userName: string,
    sessionToken: string
}

export interface State {
    alerts: Alert[],
    snackbar: {
        queue: string[],
        value: boolean,
        text: string | undefined
    },
    auth: undefined | AuthState
}

