<script lang="ts" setup>
import type { HTMLAttributes } from "vue";
import { Motion, useScroll, useTransform } from "motion-v";
import { ref, onMounted, nextTick, watch } from "vue";

interface Props {
  containerClass?: HTMLAttributes["class"];
  class?: HTMLAttributes["class"];
  items?: {
    id: string;
    label: string;
  }[];
  title?: string;
  description?: string;
}

const props = withDefaults(defineProps<Props>(), {
  items: () => [],
});

const timelineContainerRef = ref<HTMLElement | null>(null);
const timelineRef = ref<HTMLElement | null>(null);
const height = ref(0);

onMounted(async () => {
  await nextTick();
  if (timelineRef.value) {
    const rect = timelineRef.value.getBoundingClientRect();
    height.value = rect.height;
  }
});

const { scrollYProgress } = useScroll({
  target: timelineRef,
  offset: ["start 10%", "end 50%"],
});

const opacityTransform = useTransform(scrollYProgress, [0, 0.1], [0, 1]);
const heightTransform = ref(useTransform(scrollYProgress, [0, 1], [0, 0]));

watch(height, (newHeight) => {
  heightTransform.value = useTransform(scrollYProgress, [0, 1], [0, newHeight]);
});
</script>

<template>
  <div
    ref="timelineContainerRef"
    class="w-full bg-transparent font-sans md:px-10"
  >
    <div class="mx-auto max-w-7xl px-2 py-20 md:px-8 lg:px-10">
      <h2 class="mb-4 max-w-4xl text-lg font-bold md:text-4xl">
        {{ title }}
      </h2>
      <p class="max-w-sm text-sm text-neutral-500 md:text-base">
        {{ description }}
      </p>
    </div>

    <div
      ref="timelineRef"
      class="relative z-0 mx-auto max-w-7xl pb-20"
    >
      <div
        v-for="(item, index) in props.items"
        :key="item.id + index"
        class="flex justify-start pt-10 md:gap-10 md:pt-20"
      >
        <div
          class="sticky top-40 z-40 flex max-w-xs flex-col items-center self-start md:w-full md:flex-row lg:max-w-sm"
        >
          <div
            class="absolute left-3 flex size-10 items-center justify-center rounded-full bg-background border border-neutral-200 dark:border-neutral-800 md:left-3"
          >
            <div
              class="size-4 rounded-full border border-neutral-300 bg-neutral-200 p-2 dark:border-neutral-700 dark:bg-neutral-800 shadow-sm"
            />
          </div>
          <h3
            class="text-xl font-bold text-neutral-500 pl-16 md:pl-20 md:text-5xl dark:text-neutral-500"
          >
            {{ item.label }}
          </h3>
        </div>
        <div class="relative w-full pl-4 pr-4 md:pl-4">
          <slot :name="item.id" />
        </div>
      </div>
      <div
        :style="{
          height: `${height}px`,
        }"
        class="absolute top-0 left-8 w-[2px] overflow-hidden bg-[linear-gradient(to_bottom,var(--tw-gradient-stops))] from-transparent from-0% via-neutral-200 to-transparent to-99% mask-[linear-gradient(to_bottom,transparent_0%,black_10%,black_90%,transparent_100%)] md:left-8 dark:via-neutral-700"
      >
        <Motion
          as="div"
          :style="{
            height: heightTransform as any,
            opacity: opacityTransform,
          }"
          class="absolute inset-x-0 top-0 w-[2px] rounded-full bg-linear-to-t from-purple-500 from-0% via-blue-500 via-10% to-transparent"
        />
      </div>
    </div>
  </div>
</template>
