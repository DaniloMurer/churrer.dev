// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	compatibilityDate: '2024-04-03',
	devtools: {enabled: true},
	ssr: false,
	modules: ["@nuxtjs/tailwindcss", "nuxt-aos", "@pinia/nuxt"],
	telemetry: false,
	app: {
		baseURL: '/'
	},
	imports: {
		dirs: ['common'],
	},
	nitro: {
		devProxy: {
			'/api': {
				target: 'http://localhost:8080',
				changeOrigin: true
			}
		},
		output: {
			dir: '../server/spa',
			publicDir: '../server/spa'
		}
	}
})
