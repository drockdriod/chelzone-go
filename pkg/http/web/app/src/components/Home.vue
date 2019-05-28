<template>
    <div>
        <v-parallax :src="rink">
            <v-container grid-list-md>
                <v-layout text-xs-center wrap>
                    <v-flex mb-4>
                        <h1 class="display-2 font-weight-bold mb-3">
                            Welcome to Chelzone
                        </h1>
                        <p class="subheading font-weight-regular">
                            Get up to date stats, standings, in game action and more!
                        </p>
                        <v-carousel hide-delimiters :cycle="carouselCycle">
                            <v-carousel-item
                              v-for="(item,i) in topGameMilestones"
                              :key="i"
                            >
                                <video width=564 height=450 preload controls :poster="item.imageUrl"
                                    @canplay="updatePaused" @playing="updatePaused" @pause="updatePaused">
                                    <source :src="item.videoUrl" type="video/mp4" />
                                    Your browser does not support the video tag.
                                </video>
                            </v-carousel-item>
                        </v-carousel>
                    </v-flex>

                    <v-flex xs12 sm12 md6 lg6 grid-list-md>
                        <LeaderStats v-bind:leaders="leaders" v-show="!isLoading"/>
                    </v-flex>
                    <v-flex xs12 sm12 md12 lg12 grid-list-md>
                        <Divisions v-bind:teams="teamsByDivision" v-show="!isLoading" />
                    </v-flex>
                </v-layout>
            </v-container>
        </v-parallax>
        <v-container grid-list-md v-show="!isLoading">
                <SocialMosaic />
        </v-container>
    </div>
</template>
<script>
    import { mapGetters } from 'vuex'

    import Divisions from './Divisions'
    import SocialMosaic from './SocialMosaic'
    import LeaderStats from './LeaderStats'
    
    import rink from "../assets/images/rink.jpg"

    export default {
        components: { Divisions, SocialMosaic, LeaderStats },
        computed: mapGetters(['teams', 'teamsByDivision', 'leaders', 'isLoading', 'topGameMilestones']),
        created(){
            this.loader = this.$loading.show()
            if(!this.isLoading){
                this.loader.hide()
            }
        },
        data () {
            return {
                publicPath: process.env.BASE_URL,
                carouselCycle: true,
                videoElement: null,
                paused: null,
                rink
            }
        },
        methods: {
            updatePaused(event) {
                console.log(event.target.paused)
                this.videoElement = event.target
                this.paused = event.target.paused

                this.carouselCycle = this.paused
            },
            play() {
                this.videoElement.play()
            },
            pause() {
                this.videoElement.pause()
            }
        },
        beforeUpdate(){
            if(!this.isLoading){
                this.loader.hide()
            }

        },
    }
</script>
<style>
    .v-parallax{
        height: auto !important
    }
</style>