import { defineStore } from 'pinia';

export const useThemeStore = defineStore('theme', {
	state: () => ({
		isWhite: localStorage.getItem('isWhite') === 'true',
	}),
});
