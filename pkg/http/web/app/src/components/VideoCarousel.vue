<template>
	<v-carousel hide-delimiters :cycle="carouselCycle">
        <v-carousel-item
          v-for="(item,i) in topGameMilestones"
          :key="i"
        >
            <div class="video-container">
                <video width=564 height=450 controls :poster="item.imageUrl"
                    @canplay="updatePaused" @playing="updatePaused" @pause="updatePaused">
                    <source :src="item.videoUrl" type="video/mp4" />
                    Your browser does not support the video tag.
                </video>
                <label>{{item.highlight.description}}</label>
            </div>
        </v-carousel-item>
    </v-carousel>
</template>
<script>
    import { mapGetters } from 'vuex'
	export default {
		computed: mapGetters(['topGameMilestones']),
        data () {
            return {
                carouselCycle: true,
                videoElement: null,
                paused: null
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
	}
	
</script>