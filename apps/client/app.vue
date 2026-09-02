<script setup lang="ts">
  import { navigationMenuTriggerStyle } from '~/components/ui/navigation-menu';

  useHead({
    htmlAttrs: { class: 'dark' },
  });
</script>
<template>
  <div class="size-full">
    <ui-navigation-menu class="p-5">
      <ui-navigation-menu-list>
        <ui-navigation-menu-item>
          <NuxtLink v-slot="{ isActive, href, navigate }" to="/" custom>
            <ui-navigation-menu-link
              :active="isActive"
              :href
              :class="navigationMenuTriggerStyle()"
              @click="navigate"
            >
              Home
            </ui-navigation-menu-link>
          </NuxtLink>
        </ui-navigation-menu-item>
      </ui-navigation-menu-list>
    </ui-navigation-menu>
    <NuxtRouteAnnouncer />
    <div class="p-9 h-full">
      <NuxtPage />
    </div>
  </div>
</template>
