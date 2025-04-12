<script setup lang="ts">
import { useThemeStore } from '~/store/theme';

const isWhite = ref(false);

onMounted(() => {
	isWhite.value = localStorage.getItem('isWhite') === 'true';
	useThemeStore().$subscribe((mutation, state) => {
		isWhite.value = state.isWhite;
	});
});
</script>

<template>
	<div class="w-full h-full flex flex-col items-center justify-items-center">
		<h1 class="font-bold text-4xl">about churrer.xyz</h1>
		<div class="card">
			<div class="card-body text-center">
				<p class="card-text max-w-6xl">
					churrer.xyz is my portfolio website. I'm something like a full stack engineer, but I'm also a
					'kubernaut'. Meaning I like the DevOps and automation side of things as well.
				</p>
				<p class="card-text max-w-6xl">
					everything is powered by a managed kubernetes cluster on digital ocean. which is also being
					monitored by prometheus and grafana.
				</p>
				<p class="card-text max-w-6xl">
					I also got CI/CD going, with github actions I'm building the needed docker images and pushing them
					to my container registry. With Argo CD I'm deploying the containers to the cluster.
				</p>
				<p class="card-text max-w-6xl">
					the frontend of this website is built with nuxt, tailwindcss and daisyui. the backend is made with
					Go...it's like the best thing I stumbled upon.
				</p>
				<p class="card-text max-w-6xl">
					I'm also a big fan of the open source community. Which is also why I decided to open-source the code
					for this site. so you don't have to take my word for what i used as technologies : )
				</p>
			</div>
		</div>
		<img v-if="!isWhite" class="size-2/4" src="/assets/churrer_xyz.svg"
			alt="churrer.xyz diagram that shows the different components of the website" />
		<img v-if="isWhite" class="size-2/4" src="/assets/churrer_xyz_white.svg"
			alt="churrer.xyz diagram that shows the different components of the website" />
		<span class="mt-5">very nice diagram of this website</span>
	</div>
</template>
