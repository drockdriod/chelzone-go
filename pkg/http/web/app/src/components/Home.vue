<template>
    <div>
        <v-parallax :src="rink" height="3000">
            <v-container grid-list-md>
                <v-layout text-xs-center wrap>
                    <v-flex mb-4>
                        <h1 class="display-2 font-weight-bold mb-3">
                            Welcome to Chelzone
                        </h1>
                        <p class="subheading font-weight-regular">
                            Get up to date stats, standings, in game action and more!
                        </p>
                        <VideoCarousel />
                    </v-flex>

                    <v-flex xs12 sm12 md6 lg6 grid-list-md style="margin-top:103px">
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
    import VideoCarousel from './VideoCarousel'
    
    import rink from "../assets/images/rink.jpg"

    export default {
        components: { Divisions, SocialMosaic, LeaderStats, VideoCarousel },
        computed: mapGetters(['teams', 'teamsByDivision', 'leaders', 'isLoading']),
        created(){
            this.loader = this.$loading.show()
            if(!this.isLoading){
                this.loader.hide()
            }
        },
        data () {
            return {
                publicPath: process.env.BASE_URL,
                rink
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