<script setup lang="ts">
const props = defineProps(['title', 'text', 'timeframe', 'company'])
</script>

<template>
	<div class="card bg-base-300 lg:w-96">
		<div class="card-body">
			<h2 class="card-title">{{ props.title }}</h2>
			<sub><b>{{ props.company }} | {{ props.timeframe }}</b></sub>
			<p class="mt-5">{{ props.text }}</p>
		</div>
	</div>
</template>

<style scoped></style>
