<script setup lang="ts">

import { definePageMeta, type Technology, type Experience } from '#imports';
import { ExperienceDto, TechnologyDto } from '~/common/types';

definePageMeta({
	middleware: ['auth'],
});

const experiences = ref<Experience[]>([]);
const technologies = ref<Technology[]>([]);
const experience = ref<Experience>(new ExperienceDto('', '', '', ''));
const technology = ref<Technology>(new TechnologyDto('', '', '', ''));

let authenticationToken: string | null = '';

onMounted(() => {
	authenticationToken = localStorage.getItem('token');
	$fetch<Technology[]>('/api/technology', {
		method: 'GET',
		headers: {
			'Accept': 'application/json',
		}
	}).then((data: Technology[]) => {
		technologies.value = data;
	});
	$fetch<Experience[]>('/api/experience', {
		method: 'GET',
		headers: {
			'Accept': 'application/json',
		}
	}).then((data: Experience[]) => {
		experiences.value = data;
	});
});

const saveExperience = function () {
	$fetch('/api/experience', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Accept': 'application/json',
			'Authorization': `Basic ${authenticationToken}`
		},
		body: JSON.stringify(experience.value)
	});
}

const saveTechnology = function () {
	$fetch('/api/technology', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Accept': 'application/json',
			'Authorization': `Basic ${authenticationToken}`
		},
		body: JSON.stringify(technology.value)
	})
}

const deleteTechnology = function (id: number) {
	$fetch(`/api/technology/${id}`, {
		method: 'DELETE',
		headers: {
			'Authorization': `Basic ${authenticationToken}`
		}
	});
}

const deleteExperience = function (id: number) {
	$fetch(`/api/experience/${id}`, {
		method: 'DELETE',
		headers: {
			'Authorization': `Basic ${authenticationToken}`
		}
	})
}
</script>

<template>
	<h1 class="text-2xl">Welcome to the admin page</h1>
	<Table :data="technologies" :tableTitle="'TechnologySection'">
		<template #table-header>
			<th>Name</th>
			<th>Experience</th>
			<th>Description</th>
			<th>Logo</th>
		</template>

		<template #table-row="{ item }">
			<td>{{ item.name }}</td>
			<td>{{ item.experience }}</td>
			<td>{{ item.description }}</td>
			<td>{{ item.logoClass }}</td>
			<td><button class="btn btn-error rounded" v-on:click="deleteTechnology(item.id)">Delete</button></td>
		</template>
	</Table>
	<Table :data="experiences" :tableTitle="'Experience'">
		<template #table-header>
			<th>Company</th>
			<th>Position</th>
			<th>TimeFrame</th>
			<th>Responsibilities</th>
		</template>

		<template #table-row="{ item }">
			<td>{{ item.company }}</td>
			<td>{{ item.position }}</td>
			<td>{{ item.timeFrame }}</td>
			<td>{{ item.responsibilities }}</td>
			<td><button class="btn btn-error rounded" @click="deleteExperience(item.id)">Delete</button></td>
		</template>
	</Table>
	<div class="flex flex-col items-center gap-10">
		<h2 class="font-bold text-4xl">Experience</h2>
		<div>
			<div class="label">
				<span class="label-text">Company</span>
			</div>
			<input type="text" class="input input-primary" v-model="experience.company" />
		</div>

		<div>
			<div class="label">
				<span class="label-text">Position</span>
			</div>
			<input type="text" class="input input-primary" v-model="experience.position" />
		</div>

		<div>
			<div class="label">
				<span class="label-text">TimeFrame</span>
			</div>
			<input type="text" class="input input-primary" v-model="experience.timeFrame" />
		</div>

		<div>
			<div class="label">
				<span class="label-text">Responsibilities</span>
			</div>
			<input type="text" class="input input-primary" v-model="experience.responsibilities" />
		</div>
		<button class="btn btn-secondary" v-on:click="saveExperience">Save</button>
	</div>
	<hr />
	<div class="flex flex-col items-center gap-10">
		<h2 class="font-bold text-4xl">Technology</h2>
		<div>
			<div class="label">
				<span class="label-text">Name</span>
			</div>
			<input type="text" class="input input-primary" v-model="technology.name" />
		</div>

		<div>
			<div class="label">
				<span class="label-text">Experience</span>
			</div>
			<input type="text" class="input input-primary" v-model="technology.experience" />
		</div>

		<div>
			<div class="label">
				<span class="label-text">Description</span>
			</div>
			<input type="text" class="input input-primary" v-model="technology.description" />
		</div>

		<div>
			<div class="label">
				<span class="label-text">Tailwind Logo</span>
			</div>
			<input type="text" class="input input-primary" v-model="technology.logoClass" />
		</div>
		<button class="btn btn-secondary" v-on:click="saveTechnology">Save</button>
	</div>
</template>
