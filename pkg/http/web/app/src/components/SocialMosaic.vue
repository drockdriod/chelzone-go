<template>
    <div>
        <v-layout align-end justify-end row>
            <v-item-group
            multiple
            >
                <v-item>
                    <v-btn fab dark medium color="primary" v-on:click="lastPage()" v-if="currentIndex !== 0">
                      <v-icon dark>navigate_before</v-icon>
                    </v-btn>
                </v-item>
                <v-item>
                    <v-btn fab dark medium color="primary" v-on:click="nextPage()" v-if="socialMediaList.length !== currentIndex+1">
                      <v-icon dark>navigate_next</v-icon>
                    </v-btn>
                </v-item>
            </v-item-group>
        </v-layout>
        <v-layout text-xs-left wrap class="socialMediaContainer">
                <v-layout v-for="(socialmedia,index) in socialMediaList" v-bind:key="index" row v-show="index === currentIndex" wrap>
                    <v-flex xs6 sm4 md3 lg3 v-for="(item,index2) in socialmedia" v-bind:key="index2">
                        <transition tag="div"
                            v-on:before-enter="beforeEnter"
                            v-on:enter="enter"
                            v-bind:css="false" appear
                        >

                            <Tweet :tweet="item" v-show="index === currentIndex" v-if="item.user && item.user.screen_name" />
                            <YoutubeItem :youtubeitem="item" v-show="index === currentIndex" v-if="item.id && item.id.videoId" />
                            
                        </transition>
                    </v-flex>
                </v-layout>

        </v-layout>
    </div>
</template>
<script>
    import { mapGetters, mapActions } from 'vuex'

    import Velocity from 'velocity-animate'
    import Tweet from './Tweet'
    import YoutubeItem from './YoutubeItem'

    import {SOCIAL_MEDIA_SKIP} from '../utils/common'

    export default {
        components: { Tweet, YoutubeItem },
        computed: mapGetters(['socialMediaList']),
        data: function(){
            return {
                currentIndex: 0,
                skipBy: 1
            }
        },
        methods:{
            ...mapActions(['getSocialMedia']),
            lastPage: function(){
                this.currentIndex --
            },
            nextPage: function(){
                this.currentIndex ++

                if(this.socialMediaList.length - 2 === this.currentIndex){
                    console.log("RENDERING", (SOCIAL_MEDIA_SKIP * this.skipBy))
                    // Append SocialMediaList via vuex action call here
                    this.$store.dispatch('getSocialMedia', (SOCIAL_MEDIA_SKIP * this.skipBy))
                    this.skipBy++
                }
            },
            beforeEnter: function (el) {
                el.style.opacity = 0
            },
            enter: function (el, done) {
                Velocity(el, { opacity: 1, fontSize: '1.4em' }, { duration: 300 })
                Velocity(el, { fontSize: '1em' }, { complete: done })
            }
        }
    }
</script>