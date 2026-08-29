import tailwindcss from '@tailwindcss/vite';

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  telemetry: { enabled: false },
  ssr: false,
  app: {
    head: {
      title: 'Danilo Jakob | Full-Stack Software Engineer',
      meta: [
        { name: 'description', content: 'Portfolio of Danilo Jakob - Full-Stack Software Engineer specialized in Java, Go, and Kubernetes. Relocating to Helsinki, Finland in 2027.' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { charset: 'utf-8' }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
        { rel: 'canonical', href: 'https://me.churrer.dev/' }
      ]
    }
  },
  site: {
	url: 'https://me.churrer.dev',
    name: 'Danilo Jakob'
  },
  modules: ['@nuxtjs/sitemap', '@pinia/nuxt', '@nuxtjs/i18n', 'shadcn-nuxt'],
  css: ['~/assets/css/tailwind.css'],
  vite: {
    plugins: [tailwindcss()],
  },
});
