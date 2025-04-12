import { useAuthStore } from '~/store/auth';

export default defineNuxtRouteMiddleware((to, from) => {
	const token = useAuthStore().token;
	if (!token) {
		return navigateTo('/')
	}
});
