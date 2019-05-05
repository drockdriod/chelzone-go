export const tweets = (state) =>{
	return state.tweets
}

const SOCIAL_MEDIA_LIMIT = 10

export const tweetsByLimit = state => 
	Array(Math.ceil(state.tweets.length / SOCIAL_MEDIA_LIMIT)).fill().map(
		(_, index) => index * SOCIAL_MEDIA_LIMIT).map(begin => state.tweets.slice(begin, begin + SOCIAL_MEDIA_LIMIT));
