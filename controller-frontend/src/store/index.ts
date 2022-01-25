import Vue from 'vue'
import Vuex from 'vuex'
import mutations from './mutations'
import actions from './actions'
import getters from './getters'

import {AuthState, State} from "@/store/types";
import setAuthToken from "@/plugins/axios";

Vue.use(Vuex)

let authState: AuthState | undefined = undefined

if (localStorage.getItem("auth") !== null) {
    const s: string = <string>localStorage.getItem("auth")
    authState = JSON.parse(s)
}

if (authState) {
    setAuthToken(authState.sessionToken)
}

const defaultState: State = {
    alerts: [],
    snackbar: {
        queue: [],
        value: false,
        text: undefined,
    },
    dialogs: [],
    auth: authState,
    setup: {
        setupMode: false,
        totpTokenVisible: false,
        totpQRCode: undefined,
    }
}

const store = new Vuex.Store({
    state: defaultState,
    mutations: mutations,
    actions: actions,
    getters: getters,
    modules: {}
})

export default store;
