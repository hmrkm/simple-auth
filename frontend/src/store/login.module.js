import { loginService } from '../services/login.service';

const state = {
    token:null
};

const actions = {
    login({commit}, { email, password }) {
        commit('loginRequest', { email });

        return loginService.login(email, password).then(result =>{
            return result.success;
        });
    },
};

const mutations = {
    loginRequest(state) {
        state.token=null;
    },
    loginSuccess(state, token) {
        state.token=token;
    },
};

export const login = {
    namespaced: true,
    state,
    actions,
    mutations
};
