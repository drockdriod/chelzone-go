<template>
	<v-avatar
        :tile="true"
        :size="130">
			<img v-if="binary" :src="convertFromBinary('image/png',binary)" />
	</v-avatar>
</template>

<script>
    import { convertFromBinary } from '../utils/common'
	export default {
		props:["binary"],
		methods: {
            convertFromBinary
        }
	}
</script>