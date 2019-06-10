<template>
    <!-- <v-container fluid grid-list-md> -->
        <v-layout row wrap align-end justify-end text-xs-left>
            <v-flex xs12 sm6 md6 lg6 v-for="(leader,index) in leaders" v-bind:key="index">
                <v-card color="#26c6da" dark max-width="400">
                    <v-card-title>
                        <span class="title font-weight-light">{{leader.player.measure.toUpperCase()}}</span>
                        <v-layout align-center justify-end>
                            <span class="headline font-weight-bold">
                                {{leader.player.stats.value.toFixed(1)}}
                            </span>
                        </v-layout>
                    </v-card-title>
                    <v-card-actions>
                        <v-list-tile class="grow">
                            <v-list-tile-avatar color="white">
                                <v-img class="elevation-6" v-bind:src="convertBinary(leader.player.team.logo.Data, 'svg+xml')"></v-img>
                            </v-list-tile-avatar>
                            <v-list-tile-content>
                                <v-list-tile-title>{{leader.player.person.fullName}}</v-list-tile-title>
                            </v-list-tile-content>
                        </v-list-tile>
                    </v-card-actions>
                </v-card>
            </v-flex>
        </v-layout>
    <!-- </v-container> -->
</template>
<script>
    export default {
        props: ["leaders"],
        methods: {
            convertBinary: (binary, type) => {
                return `data:image/${type};base64,${binary}`
            }
        }
    }
</script>