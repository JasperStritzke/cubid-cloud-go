import Vue from 'vue'
import Vuex from 'vuex'
import mutations from './mutations'
import actions from './actions'
import getters from './getters'

import {State} from "@/store/types";

Vue.use(Vuex)

const defaultState: State = {
    alerts: [],
    snackbar: {
        queue: [],
        value: false,
        text: undefined,
    },
    auth: undefined
}

const store = new Vuex.Store({
    state: defaultState,
    mutations: mutations,
    actions: actions,
    getters: getters,
    modules: {}
})


export default store;
