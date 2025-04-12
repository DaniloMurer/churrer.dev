<script setup lang="ts">
import { useThemeStore } from '~/store/theme';
import { useAuthStore } from "~/store/auth";

const username = ref('');
const password = ref('');
const isWhite = ref(false);
const login = function () {
	const requestBody = {
		username: username.value,
		password: password.value,
		token: ''
	}
	$fetch('/api/authentication', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Accept': 'application/json',
		},
		body: JSON.stringify(requestBody)
	}).then((data: any) => {
		localStorage.setItem('token', data.token)
		useAuthStore().token = data.token;
		const loginModal = document.getElementById('loginModal');
		if (loginModal) {
			loginModal.close();
			navigateTo('/admin')
		}
	}).catch(err => {
		console.error(err)
	})
}

const toggleTheme = function () {
	localStorage.setItem("isWhite", String(!isWhite.value));
	useThemeStore().$patch({ isWhite: !isWhite.value });
}

onMounted(() => {
	useAuthStore().token = localStorage.getItem('token');
	window.addEventListener('keydown', (e) => {
		if (e.ctrlKey && e.altKey && e.shiftKey && e.key === 'L') {
			const token = localStorage.getItem('token');
			if (token) {
				navigateTo('/admin');
			} else {
				document.getElementById('loginModal').showModal();
			}
		}
	});
	isWhite.value = localStorage.getItem("isWhite") === 'true';
	useThemeStore().$subscribe((mutation, state) => {
		isWhite.value = state.isWhite;
	});
});
</script>
<template>
	<div class="p-3 flex flex-col h-lvh gap-16">
		<div class="navbar bg-base-300 rounded-2xl shadow-2xl">
			<div class="flex-1">
				<NuxtLink class="btn btn-ghost text-xl" to="/_bak/client/public">churrer.xyz</NuxtLink>
				<NuxtLink class="btn btn-ghost text-xl" to="/about">about this site</NuxtLink>
			</div>
			<div class="flex-none p-5">
				<label class="swap swap-rotate">
					<input type="checkbox" value="winter" v-bind:checked="isWhite" v-on:click="toggleTheme"
						class="theme-controller" />
					<span class="iconify carbon--moon text-2xl swap-off"></span>
					<span class="iconify carbon--sun text-2xl swap-on"></span>
				</label>
			</div>
		</div>
		<dialog id="loginModal" class="modal">
			<div class="modal-box p-5">
				<form method="dialog">
					<button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
				</form>
				<div class="flex flex-col align-items-center gap-16">
					<h3 class="text-lg font-bold">Log-In</h3>
					<div>
						<div class="label">
							<span class="label-text">Username</span>
						</div>
						<input class="input input-primary w-11/12" v-model="username" />
					</div>
					<div>
						<div class="label">
							<span class="label-text">Password</span>
						</div>
						<input type="password" class="input input-primary w-11/12" v-model="password" />
					</div>
				</div>
				<button class="btn btn-secondary w-24 float-right mt-16 mr-10" v-on:click="login()">Login</button>
			</div>
		</dialog>
		<NuxtRouteAnnouncer />
		<NuxtPage />
	</div>
</template>
<style>
html {
	font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;
}
</style>
