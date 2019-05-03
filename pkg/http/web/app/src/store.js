import Vue from 'vue';
import Vuex from 'vuex';

import * as TeamGetters from './getters/TeamGetters'
import * as TeamActions from './actions/TeamActions'

import * as PlayerGetters from './getters/PlayerGetters'
import * as PlayerActions from './actions/PlayerActions'

Vue.use(Vuex);

const store = new Vuex.Store({
	state: {
		teams: [],
        leaders: []
    },
	mutations: {
        setTeams(state, teams) {
            state.teams = teams
        },
        setLeaders(state, leaders) {
            state.leaders = leaders
        }
    },
    getters: {
        ...TeamGetters,
        ...PlayerGetters
    },
    actions: {
        ...TeamActions,
        ...PlayerActions
    }
})

export default store;