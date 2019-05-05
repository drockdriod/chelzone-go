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
                    <v-btn fab dark medium color="primary" v-on:click="nextPage()" v-if="itemseries.length !== currentIndex+1">
                      <v-icon dark>navigate_next</v-icon>
                    </v-btn>
                </v-item>
            </v-item-group>
        </v-layout>
        <v-layout text-xs-left wrap>
                <v-layout v-for="(tweets,index) in itemseries" v-bind:key="index" row v-show="index === currentIndex" wrap>
                    <v-flex xs6 sm4 md3 lg3 v-for="(tweet,index2) in tweets" v-bind:key="index2">
                        <transition tag="div"
                            v-on:before-enter="beforeEnter"
                            v-on:enter="enter"
                            v-bind:css="false" appear
                        >
                            <v-card color="#26c6da" dark v-show="index === currentIndex">
                                <v-card-title>
                                    <v-icon large left>
                                        mdi-twitter
                                    </v-icon>
                                    <span class="title font-weight-light">Twitter</span>
                                </v-card-title>
                                <v-card-text class="body-1 font-weight-bold">
                                    {{ tweet.full_text }}
                                </v-card-text>
                                <v-card-actions>
                                    <v-list-tile class="grow">
                                        <v-list-tile-avatar color="grey darken-3">
                                            <v-img class="elevation-6" v-bind:src="tweet.user.profile_image_url"></v-img>
                                        </v-list-tile-avatar>
                                        <v-list-tile-content>
                                            <v-list-tile-title>{{tweet.user.screen_name}}</v-list-tile-title>
                                        </v-list-tile-content>
                                        <v-layout align-center justify-end>
                                            <v-icon class="mr-1">mdi-share-variant</v-icon>
                                            <span class="subheading">{{tweet.retweet_count}}</span>
                                            <span class="mr-1">·</span>
                                            <v-icon class="mr-1">mdi-heart</v-icon>
                                            <span class="subheading mr-2">{{tweet.favorite_count}}</span>
                                        </v-layout>
                                    </v-list-tile>
                                </v-card-actions>
                            </v-card>
                        </transition>
                    </v-flex>
                </v-layout>

        </v-layout>
    </div>
</template>
<script>
    import Velocity from 'velocity-animate'
    // import { mapGetters, mapActions } from 'vuex'
    export default {
        props: ['itemseries'],
        data: function(){
            return {
                currentIndex: 0
            }
        },
        watch: {
            itemseries (n, o) {
                console.log(n, o) // n is the new value, o is the old value.
            }
        },
        methods:{
            lastPage: function(){
                this.currentIndex --
            },
            nextPage: function(){
                this.currentIndex ++
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