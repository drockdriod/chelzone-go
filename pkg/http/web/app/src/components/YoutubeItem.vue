<template>
    <v-card color="red" dark @click="openLink(`https://www.youtube.com/watch?v=${youtubeitem.id.videoId}`)" class="socialMediaItem">
        <v-card-title>
            <v-icon large left>
                $vuetify.icons.youtube
            </v-icon>
            <span class="title font-weight-light">Youtube</span>
        </v-card-title>
        <v-card-text class="body-1 font-weight-bold" v-html="youtubeitem.snippet.title" :style="`background-image:url(${youtubeitem.snippet.thumbnails.medium.url});background-size:cover;height:${parseInt(youtubeitem.snippet.thumbnails.medium.height)}px;
    text-shadow: -1px 0 black, 0 1px black, 1px 0 black, 0 -1px black;`">
        </v-card-text>
        <v-card-actions :style="'text-shadow: -1px 0 black, 0 1px black, 1px 0 black, 0 -1px black;'">
            <v-list-tile class="grow">
                <v-list-tile-avatar color="grey darken-3">
                    <!-- <v-img class="elevation-6" v-bind:src="youtubeitem.snippet.thumbnails.default.url"></v-img> -->
                </v-list-tile-avatar>
                <v-list-tile-content>
                    <v-list-tile-title>{{youtubeitem.snippet.channelTitle}}</v-list-tile-title>
                </v-list-tile-content>
            </v-list-tile>
        </v-card-actions>
    </v-card>
</template>

<script>
    import loader from "../assets/images/rink.jpg"
    import { openLink } from '../utils/common'

    export default {
        props: ['youtubeitem'],
        methods: {
            openLink
        },
        data () {
            return {
                loader
            }
        }
    }
</script>