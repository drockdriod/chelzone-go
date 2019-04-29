import Vue from 'vue';
import Vuex from 'vuex';

import * as TeamGetters from './getters/TeamGetters'
import * as TeamActions from './actions/TeamActions'

Vue.use(Vuex);

const store = new Vuex.Store({
	state: {
		teams: []
    },
	mutations: {
    	setTeams(state, teams) {
            state.teams = teams
        }
    },
    getters: {
        ...TeamGetters
    },
    actions: {
        ...TeamActions
    }
})

export default store;