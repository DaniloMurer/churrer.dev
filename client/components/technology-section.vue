<script setup lang="ts">
const technologies = ref<Technology[]>([]);

onMounted(() => {
	$fetch<Technology[]>('/api/technology', {
		method: 'GET',
		headers: {
			'Accept': 'application/json',
		}
	}).then((data: Technology[]) => {
		technologies.value = data;
	});
});


</script>

<template>
	<div class="stats stats-vertical xl:stats-horizontal bg-base-300 p-2">
		<div class="stat" v-for="technology in technologies">
			<div class="stat-figure">
				<span class="iconify iconify-color text-8xl" :class="technology.logoClass" />
			</div>
			<div class="stat-title text-primary text-2xl">{{ technology.name }}</div>
			<div class="stat-value text-accent">{{ technology.experience }}</div>
			<div class="stat-desc text-wrap">{{ technology.description }}</div>
		</div>
	</div>
	<p class="mt-4">Of course there is more, but you can ask me about that personally!</p>
</template>

<style scoped></style>
