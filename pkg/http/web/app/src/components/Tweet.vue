<template>
    <v-card color="#26c6da" dark class="socialMediaItem" @click="openLink(`https://twitter.com/${tweet.user.screen_name}/status/${tweet.id_str}`)">
        <v-card-title>
            <v-icon large left>
                $vuetify.icons.twitter
            </v-icon>
            <span class="title font-weight-light">Twitter</span>
        </v-card-title>
        <v-card-text class="body-1 font-weight-bold" v-html="tweet.full_text">
        </v-card-text>
        <v-card-actions :style="`background-image:url(${tweet.user.profile_banner_url});background-size:cover;text-shadow: -1px 0 black, 0 1px black, 1px 0 black, 0 -1px black;`">
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
</template>

<script>
    import { openLink } from '../utils/common'
    export default {
        props: ['tweet'],
        methods: {
            openLink
        },
    }
</script>