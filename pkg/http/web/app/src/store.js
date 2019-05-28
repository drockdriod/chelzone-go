import Vue from 'vue';
import Vuex from 'vuex';
import createLogger from 'vuex/dist/logger'

import * as TeamGetters from './getters/TeamGetters'
import * as TeamActions from './actions/TeamActions'

import * as PlayerGetters from './getters/PlayerGetters'
import * as PlayerActions from './actions/PlayerActions'

import * as SocialGetters from './getters/SocialGetters'
import * as SocialActions from './actions/SocialActions'

import * as GameGetters from './getters/GameGetters'
import * as GameActions from './actions/GameActions'

Vue.use(Vuex);

const store = new Vuex.Store({
    plugins: [createLogger()],
	state: {
		teams: [],
        leaders: [],
        tweets: [],
        youtubeitems: [],
        loading: false,
        socialMediaList: [],
        gameMilestones: []
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
        },
        appendTweets(state, tweets) {
            state.tweets.push(...tweets)
        },
        setYoutubeItems(state, items) {
            state.youtubeitems = items
        },
        appendYoutubeItems(state, items) {
            state.youtubeitems.push(...items)
        },
        setLoading(state, isLoading){
            state.loading = isLoading
        },
        appendSocialMediaList(state, items) {
            state.socialMediaList.push(...items)
        },
        setGameMilestones(state, milestones) {
            state.gameMilestones = milestones
        }
    },
    getters: {
        isLoading(state){
            return state.loading
        },
        ...TeamGetters,
        ...PlayerGetters,
        ...SocialGetters,
        ...GameGetters
    },
    actions: {
        setLoading({commit}, isLoading){
            commit('setLoading', isLoading)
        },
        ...TeamActions,
        ...PlayerActions,
        ...SocialActions,
        ...GameActions
    }
})

export default store;