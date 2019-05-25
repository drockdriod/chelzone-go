import ApiClient from '../ApiClient'
import * as common from '../utils/common'

export const getTweets = async({commit}) => {

	const results = await ApiClient.perform('get', '/twitter/tweets/')

	commit('setTweets', results.tweets)
}

export const getYoutubeItems = async({commit}) => {
	const results = await ApiClient.perform('get', '/youtube/items/')
	commit('setYoutubeItems', results.youtubeitems)
}

export const getSocialMedia = async({commit}, skip) => {
	const tw = await ApiClient.perform('get', `/twitter/tweets/${skip}`)
	const yt = await ApiClient.perform('get', `/youtube/items/${skip}`)


	const combined = common.chunk(common.shuffleArray(tw.tweets.concat(yt.youtubeitems)), common.SOCIAL_MEDIA_LIMIT)

	console.log("combined", combined)
	commit('appendYoutubeItems', yt.youtubeitems)
	commit('appendTweets', tw.tweets)
	commit('appendSocialMediaList', combined)
}