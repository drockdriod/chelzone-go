<template>
	<v-carousel hide-delimiters :cycle="carouselCycle">
        <v-carousel-item
          v-for="(item,i) in topGameMilestones"
          :key="i"
        >
            <video width=564 height=450 preload controls
                @canplay="updatePaused" @playing="updatePaused" @pause="updatePaused">
                <source :src="item.highlight.playbacks[item.highlight.playbacks.length - 1].url" type="video/mp4" />
                Your browser does not support the video tag.
            </video>
        </v-carousel-item>
    </v-carousel>
</template>
<script>
	export default {
		
	}
	
</script>