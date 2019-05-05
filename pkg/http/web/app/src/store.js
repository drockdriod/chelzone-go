import Vue from 'vue';
import Vuex from 'vuex';

import * as TeamGetters from './getters/TeamGetters'
import * as TeamActions from './actions/TeamActions'

import * as PlayerGetters from './getters/PlayerGetters'
import * as PlayerActions from './actions/PlayerActions'

import * as SocialGetters from './getters/SocialGetters'
import * as SocialActions from './actions/SocialActions'

Vue.use(Vuex);

const store = new Vuex.Store({
	state: {
		teams: [],
        leaders: [],
        tweets: []
    },
	mutations: {
        setTeams(state, teams) {
            state.teams = teams
        },
        setLeaders(state, leaders) {
            state.leaders = leaders
        },
        setTweets(state, tweets) {
            state.tweets = tweets
        }
    },
    getters: {
        ...TeamGetters,
        ...PlayerGetters,
        ...SocialGetters
    },
    actions: {
        ...TeamActions,
        ...PlayerActions,
        ...SocialActions
    }
})

export default store;