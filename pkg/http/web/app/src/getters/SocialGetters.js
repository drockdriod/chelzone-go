import * as common from '../utils/common'
export const tweets = (state) =>{
	return state.tweets
}

export const socialMediaByLimit = state => common.chunk(common.shuffleArray(state.tweets.concat(state.youtubeitems)), common.SOCIAL_MEDIA_LIMIT)

export const socialMediaList = state => state.socialMediaList