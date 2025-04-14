<script setup lang="ts"></script>
<template>
  <div
    class="relative flex w-full flex-col items-center justify-center overflow-hidden rounded-lg h-full"
  >
    <span
      class="pointer-events-none whitespace-pre-wrap bg-gradient-to-b from-black to-gray-300/80 bg-clip-text text-center text-8xl font-semibold leading-none text-transparent dark:from-white dark:to-slate-900/10"
    >
    </span>
    <ui-particles-bg
      class-name="size-full"
      class="absolute inset-0"
      :quantity="500"
      :ease="100"
      :staticity="20"
      refresh
    />
    <ui-card-container>
      <ui-card-body
        class="group/card relative size-auto rounded-xl border border-black/[0.1] p-6 sm:w-[30rem] dark:border-white/[0.2] bg-background"
      >
        <ui-card-item :translate-z="50" class="text-xl font-bold text-neutral-600 dark:text-white">
          Danilo Jakob
        </ui-card-item>
        <ui-card-item
          as="p"
          translate-z="60"
          class="mt-2 max-w-sm text-sm text-neutral-500 dark:text-neutral-300"
        >
          The developer you were looking for all along
        </ui-card-item>
        <ui-card-item :translate-z="100" class="mt-4 w-full">
          <img
            src="https://images.unsplash.com/photo-1441974231531-c6227db76b6e?q=80&w=2560&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
            height="1000"
            width="1000"
            class="h-60 w-full rounded-xl object-cover group-hover/card:shadow-xl"
            alt="thumbnail"
          />
        </ui-card-item>
      </ui-card-body>
    </ui-card-container>
  </div>
</template>
