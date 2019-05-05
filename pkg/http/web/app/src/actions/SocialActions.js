import ApiClient from '../ApiClient'

export const getTweets = async({commit}) => {
	const results = await ApiClient.perform('get', '/twitter/tweets')

	commit('setTweets', results.tweets)
}