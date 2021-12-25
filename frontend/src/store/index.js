import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios'
import VueAxios from 'vue-axios'

import { login } from './login.module';
import { message } from './message.module';

Vue.use(Vuex);
Vue.use(VueAxios, axios)

export const store = new Vuex.Store({
    modules: {
        login,
        message
    }
});
