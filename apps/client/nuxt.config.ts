// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  telemetry: { enabled: false },
  ssr: false,
  modules: ['@pinia/nuxt', '@nuxtjs/i18n'],
});
