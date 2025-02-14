<script setup lang="ts">
const experiences = ref<Experience[]>([]);
onMounted(() => {
	$fetch<Experience[]>('/api/experience', {
		method: 'GET',
		headers: {
			'Accept': 'application/json',
		}
	}).then((data: Experience[]) => {
		experiences.value = data;
		console.log(experiences)
	}).catch((err) => {
		console.error('Something went wrong. check backend logs....', err);
	})
})
</script>

<template>
	<ul class="timeline timeline-vertical timeline-snap-icon">
		<li v-for="(experience, index) in experiences">
			<div class="mb-10" :class="index % 2 == 0 ? 'timeline-start' : 'timeline-end'">
				<WorkCard :class="index % 2 == 0 ? 'lg:mr-10' : 'lg:ml-10'"
					:data-aos="index % 2 == 0 ? 'fade-right' : 'fade-left'" data-aos-duration="1000"
					:title="experience.position" :company="experience.company" :timeframe="experience.timeFrame"
					:text="experience.responsibilities" />
			</div>
			<div class="timeline-middle">
				<span class="iconify carbon--time text-xl"></span>
			</div>
			<!--	fixme: https://github.com/DaniloMurer/churrer.xyz/issues/8		-->
			<hr v-if="index % 2 == 0" />
		</li>
	</ul>
</template>

<style scoped></style>
