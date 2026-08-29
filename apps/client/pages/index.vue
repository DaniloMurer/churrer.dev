<script setup lang="ts">
import {BentoGrid, BentoGridItem,} from '@/components/ui/bento-grid';
import {CardBody, CardContainer, CardItem} from '@/components/ui/card-3d';
import {ParticlesBg} from '@/components/ui/particles-bg';
import {Button} from '@/components/ui/button';
import {IconCloud} from '@/components/ui/icon-cloud';
import {Dock, DockIcon} from '@/components/ui/dock';
import {Timeline} from '@/components/ui/timeline';
import {BorderBeam} from '@/components/ui/border-beam';
import {Cloud, Database, ExternalLink, Github, Layout, Linkedin, Mail, Server} from 'lucide-vue-next';

const experience = [
    {
      company: 'Future Opportunity',
      position: 'Software Engineer',
      timeFrame: 'Apr 2027 – Future',
      responsibilities:
        'Ready to jump into new challenges in wonderful Helsinki, Finland — hopefully with you! Bringing 8+ years of Java, Go, and Kubernetes expertise to the land of a thousand lakes. (Yes, I am serious about the move!)',
    },
    {
      company: 'Würth IT Switzerland AG',
      position: 'Software Engineer',
      timeFrame: 'Aug 2022 – Present',
      responsibilities:
        'Developing and maintaining backend services in Java (Spring Boot, Java EE, Vaadin), Go, and Kotlin. Building SAP UI5 / SAP ABAP and SAP Gateway applications for transportation management. Migrating and operating applications on OpenShift and Kubernetes, containerized with Docker. Building DevOps tooling and CI/CD pipelines (ArgoCD, Jenkins). Implementing observability with ELK Stack and APM.',
    },
    {
      company: 'Würth IT Switzerland AG',
      position: 'Software Developer Apprentice',
      timeFrame: 'Aug 2018 – Aug 2022',
      responsibilities:
        'Completed a 4-year apprenticeship (EFZ) as a software developer, gaining foundational experience across the full software development lifecycle, including Java-based application development and early exposure to DevOps practices.',
    },
  ];

  const skills = [
    { name: 'Java', icon: 'i-simple-icons-openjdk' },
    { name: 'Spring Boot', icon: 'i-simple-icons-springboot' },
    { name: 'Go', icon: 'i-simple-icons-go' },
    { name: 'Kotlin', icon: 'i-simple-icons-kotlin' },
    { name: 'TypeScript', icon: 'i-simple-icons-typescript' },
    { name: 'Vue', icon: 'i-simple-icons-vuedotjs' },
    { name: 'Nuxt', icon: 'i-simple-icons-nuxt' },
    { name: 'Kubernetes', icon: 'i-simple-icons-kubernetes' },
    { name: 'Docker', icon: 'i-simple-icons-docker' },
    { name: 'PostgreSQL', icon: 'i-simple-icons-postgresql' },
    { name: 'OpenShift', icon: 'i-simple-icons-redhatopenshift' },
    { name: 'Elasticsearch', icon: 'i-simple-icons-elasticsearch' },
    { name: 'Jenkins', icon: 'i-simple-icons-jenkins' },
    { name: 'ArgoCD', icon: 'i-simple-icons-argo' },
    { name: 'GitHub', icon: 'i-simple-icons-github' },
  ];

  const skillIcons = skills.map(
    (s) => `https://cdn.simpleicons.org/${s.icon.split('-').pop()}/white`,
  );

  const timelineItems = experience.map((exp, idx) => ({
    id: `exp-${idx}`,
    label: exp.timeFrame,
  }));

  const education = [
    {
      title: 'EFZ, Software Development',
      institution: 'Technische Berufsschule Zürich',
      timeFrame: 'Aug 2018 – Aug 2022',
      grade: 'Grade: 5.2',
    },
  ];

  const certifications = [
    {
      title: 'Cambridge C1 Advanced',
      institution: 'English Certification',
      timeFrame: '',
      grade: '',
    },
  ];

  const languages = [
    { name: 'German', level: 'Native', color: 'text-blue-500' },
    { name: 'Italian', level: 'Fluent', color: 'text-blue-500' },
    { name: 'English', level: 'C1 Advanced', color: 'text-blue-500' },
    { name: 'Finnish', level: 'Learning', color: 'text-neutral-400' },
  ];

  const interests = ['Fishing', 'Writing', 'Gaming'];

  const featuredProjects = [
    {
      title: 'churrers-cluster',
      description:
        'A comprehensive GitOps-driven Kubernetes cluster setup. Managing infrastructure as code with Kustomize, ArgoCD, and automated deployment pipelines.',
      link: 'https://github.com/DaniloMurer/churrers-cluster',
      tech: ['Kubernetes', 'Kustomize', 'ArgoCD', 'GitOps'],
      color: 'blue',
    },
    {
      title: 'space-simulations',
      description:
        'A collection of space-related simulations and visualizations, exploring physics and orbital mechanics in a 2D environment.',
      link: 'https://github.com/DaniloMurer/space-simulations',
      tech: ['C', 'Raylib', 'Physics', 'Orbital Mechanics'],
      color: 'purple',
    },
    {
      title: 'churrer.dev',
      description:
        'The source code of this very portfolio website. A modern, interactive developer profile built with Inspira UI and Nuxt 3.',
      link: 'https://github.com/DaniloMurer/churrer.dev',
      tech: ['Vue', 'Nuxt', 'Inspira UI', 'Tailwind CSS'],
      color: 'green',
    },
  ];

  useSeoMeta({
    title: 'Danilo Jakob - Full-Stack Software Engineer',
    ogTitle: 'Danilo Jakob - Full-Stack Software Engineer',
    description: 'Software Engineer at Würth IT Switzerland AG, specializing in Java, Go, and DevOps. Relocating to Helsinki, Finland in April 2027.',
    ogDescription: 'Software Engineer at Würth IT Switzerland AG, specializing in Java, Go, and DevOps. Relocating to Helsinki, Finland in April 2027.',
    twitterCard: 'summary_large_image',
  });

  useHead({
    script: [
      {
        type: 'application/ld+json',
        innerHTML: JSON.stringify({
          '@context': 'https://schema.org',
          '@type': 'Person',
          name: 'Danilo Jakob',
          url: 'https://me.churrer.dev',
          jobTitle: 'Full-Stack Software Engineer',
          worksFor: {
            '@type': 'Organization',
            name: 'Würth IT Switzerland AG',
          },
          address: {
            '@type': 'PostalAddress',
            addressLocality: 'Chur',
            addressCountry: 'CH',
          },
          sameAs: [
            'https://github.com/DaniloMurer',
            'https://linkedin.com/in/danilojakob', // Assuming standard LinkedIn URL
          ],
          description: 'Software Engineer specializing in Java, Go, and Kubernetes. Relocating to Helsinki, Finland in April 2027.',
        }),
      },
    ],
  });
</script>
<template>
  <div class="relative w-full min-h-screen bg-background overflow-x-hidden text-foreground">
    <!-- Global Background Particles -->
    <ParticlesBg
      class-name="size-full"
      class="fixed inset-0 z-0 pointer-events-none"
      :quantity="200"
      :ease="100"
      :staticity="20"
      refresh
    />

    <!-- Hero Section -->
    <section class="relative min-h-[80vh] flex items-center justify-center overflow-hidden py-20">
      <CardContainer class="z-10">
        <CardBody
          class="group/card relative size-auto rounded-xl border border-black/[0.1] p-6 sm:w-[35rem] dark:border-white/[0.2] bg-background/80 backdrop-blur-sm"
        >
          <CardItem :translate-z="50" as="h1" class="text-4xl font-bold text-neutral-800 dark:text-white">
            Danilo Jakob
          </CardItem>
          <CardItem
            as="p"
            translate-z="60"
            class="mt-2 text-xl text-neutral-600 dark:text-neutral-300 font-medium"
          >
            Full-Stack Software Engineer
            <span class="block text-sm text-neutral-500 mt-1">
              Chur, Switzerland · <span class="text-blue-500 font-semibold italic">Relocating to Helsinki, Finland in April 2027 (Hopefully with you!)</span>
            </span>
          </CardItem>
          <CardItem :translate-z="100" class="mt-6 w-full">
            <div
              class="h-64 w-full rounded-xl bg-gradient-to-br from-blue-500 to-purple-600 flex flex-col items-center justify-center text-white text-sm px-8 text-center leading-relaxed"
            >
              <div class="mb-4 px-3 py-1 rounded-full bg-white/20 backdrop-blur-md border border-white/30 text-[10px] uppercase tracking-widest font-bold">
                Relocation Goal: Helsinki 2027 🇫🇮
              </div>
              <p>
                8 years of experience building cloud-native enterprise applications. Specialized in
                Java, Go, and Kotlin backends, Kubernetes/OpenShift DevOps, and modern TypeScript
                frontends. Passionate about improving developer workflows, building robust systems,
                and apparently, moving to Finland.
              </p>
            </div>
          </CardItem>
          <div class="flex flex-wrap justify-between items-center mt-10 gap-4">
            <CardItem :translate-z="20" as="a" href="mailto:danilojakob@protonmail.com">
              <Button variant="ghost" class="text-xs font-normal dark:text-white">
                Contact Me
              </Button>
            </CardItem>
            <div class="flex gap-2">
              <CardItem :translate-z="20" as="a" href="/cv_danilo_jakob.pdf" download="cv_danilo_jakob.pdf">
                <Button variant="outline" class="px-4 py-2 rounded-xl text-xs font-bold">
                  Download CV
                </Button>
              </CardItem>
              <CardItem :translate-z="20" as="a" href="https://github.com/DaniloMurer" target="_blank">
                <Button class="px-4 py-2 rounded-xl text-xs font-bold bg-black dark:bg-white dark:text-black text-white">
                  GitHub
                </Button>
              </CardItem>
            </div>
          </div>
        </CardBody>
      </CardContainer>
    </section>

    <!-- Content Sections -->
    <div class="max-w-7xl mx-auto px-6 pb-20 space-y-16">
      <!-- Skills Section -->
      <section>
        <h2 class="text-3xl font-bold mb-10 text-center">Tech Stack</h2>
        <div class="relative flex size-full items-center justify-center overflow-hidden">
          <IconCloud :images="skillIcons" />
        </div>
      </section>

      <!-- Featured Projects -->
      <section>
        <h2 class="text-3xl font-bold mb-10 text-center">Featured Projects</h2>
        <BentoGrid class="max-w-4xl mx-auto">
          <BentoGridItem
            v-for="project in featuredProjects"
            :key="project.title"
            class="group relative overflow-hidden"
            :class="
              project.title === 'churrers-cluster' ? 'md:col-span-2' : 'md:col-span-1'
            "
          >
            <template #title>
              <div
                class="flex items-center gap-2"
                :class="{
                  'text-blue-500': project.color === 'blue',
                  'text-purple-500': project.color === 'purple',
                  'text-green-500': project.color === 'green',
                }"
              >
                <Github class="size-6" />
                <h3 class="text-xl font-bold">{{ project.title }}</h3>
              </div>
            </template>
            <template #description>
              <div class="space-y-4">
                <p class="text-sm">
                  {{ project.description }}
                </p>
                <div class="flex flex-wrap gap-2">
                  <span
                    v-for="t in project.tech"
                    :key="t"
                    class="px-2 py-1 rounded-md text-xs font-medium border"
                    :class="{
                      'bg-blue-500/10 text-blue-500 border-blue-500/20': project.color === 'blue',
                      'bg-purple-500/10 text-purple-500 border-purple-500/20':
                        project.color === 'purple',
                      'bg-green-500/10 text-green-500 border-green-500/20':
                        project.color === 'green',
                    }"
                  >
                    {{ t }}
                  </span>
                </div>
                <div class="pt-2">
                  <a
                    :href="project.link"
                    target="_blank"
                    class="inline-flex items-center text-xs font-bold hover:underline"
                  >
                    View on GitHub
                    <ExternalLink class="ml-1 size-3" />
                  </a>
                </div>
              </div>
            </template>
            <BorderBeam
              :size="150"
              :duration="15"
              :color-from="
                project.color === 'blue'
                  ? '#3b82f6'
                  : project.color === 'purple'
                    ? '#a855f7'
                    : '#22c55e'
              "
              :color-to="
                project.color === 'blue'
                  ? '#60a5fa'
                  : project.color === 'purple'
                    ? '#d8b4fe'
                    : '#86efac'
              "
            />
          </BentoGridItem>
        </BentoGrid>
      </section>

      <!-- Experience Section -->
      <section>
        <Timeline
          title="Professional Experience"
          description="My journey at Würth IT Switzerland AG, from apprentice to software engineer."
          :items="timelineItems"
        >
          <template v-for="(exp, idx) in experience" :key="idx" #[`exp-${idx}`]>
            <div
              class="relative w-full pb-2"
              :class="{ 'opacity-80 italic': exp.company === 'Future Opportunity' }"
            >
              <h3 class="text-2xl font-bold text-neutral-800 dark:text-white">
                {{ (exp as any).position }}
                <span
                  v-if="exp.company === 'Future Opportunity'"
                  class="ml-2 text-sm font-normal not-italic px-2 py-0.5 rounded-full bg-blue-500/20 text-blue-500 border border-blue-500/30"
                >
                  Future
                </span>
              </h3>
              <p class="text-neutral-500 dark:text-neutral-400 font-medium mb-4">
                {{ (exp as any).company }}
              </p>
              <p class="text-neutral-600 dark:text-neutral-300 leading-relaxed max-w-2xl">
                {{ (exp as any).responsibilities }}
              </p>
            </div>
          </template>
        </Timeline>
      </section>

      <!-- Core Competencies -->
      <section>
        <h2 class="text-3xl font-bold mb-10 text-center">Core Competencies</h2>
        <BentoGrid class="max-w-4xl mx-auto">
          <!-- Backend -->
          <BentoGridItem class="md:col-span-2 group relative overflow-hidden">
            <template #title>
              <div class="flex items-center gap-2 text-blue-500">
                <Server class="size-6" />
                <h3 class="text-xl font-bold">Backend Development</h3>
              </div>
            </template>
            <template #description>
              <div class="space-y-4">
                <p class="text-sm">
                  Designing robust and scalable server-side architectures using modern languages and frameworks.
                </p>
                <div class="flex flex-wrap gap-2">
                  <span
                    v-for="s in ['Java', 'Spring Boot', 'Go', 'Kotlin', 'Java EE', 'Vaadin']"
                    :key="s"
                    class="px-2 py-1 rounded-md bg-blue-500/10 text-blue-500 text-xs font-medium border border-blue-500/20"
                  >
                    {{ s }}
                  </span>
                </div>
              </div>
            </template>
            <BorderBeam :size="150" :duration="10" color-from="#3b82f6" color-to="#60a5fa" />
          </BentoGridItem>

          <!-- Frontend -->
          <BentoGridItem class="md:col-span-1 group relative overflow-hidden">
            <template #title>
              <div class="flex items-center gap-2 text-purple-500">
                <Layout class="size-6" />
                <h3 class="text-xl font-bold">Frontend</h3>
              </div>
            </template>
            <template #description>
              <div class="space-y-4">
                <p class="text-sm">Building reactive and accessible user interfaces.</p>
                <div class="flex flex-wrap gap-2">
                  <span
                    v-for="s in ['Vue', 'Nuxt', 'TypeScript', 'Vaadin']"
                    :key="s"
                    class="px-2 py-1 rounded-md bg-purple-500/10 text-purple-500 text-xs font-medium border border-purple-500/20"
                  >
                    {{ s }}
                  </span>
                </div>
              </div>
            </template>
            <BorderBeam :size="150" :duration="12" color-from="#a855f7" color-to="#d8b4fe" />
          </BentoGridItem>

          <!-- Cloud & DevOps -->
          <BentoGridItem class="md:col-span-1 group relative overflow-hidden">
            <template #title>
              <div class="flex items-center gap-2 text-green-500">
                <Cloud class="size-6" />
                <h3 class="text-xl font-bold">Cloud & DevOps</h3>
              </div>
            </template>
            <template #description>
              <div class="space-y-4">
                <p class="text-sm">Containerization and automated deployment pipelines.</p>
                <div class="flex flex-wrap gap-2">
                  <span
                    v-for="s in ['K8s', 'Docker', 'OpenShift', 'ArgoCD', 'Kustomize', 'Compose', 'Jenkins']"
                    :key="s"
                    class="px-2 py-1 rounded-md bg-green-500/10 text-green-500 text-xs font-medium border border-green-500/20"
                  >
                    {{ s }}
                  </span>
                </div>
              </div>
            </template>
            <BorderBeam :size="150" :duration="14" color-from="#22c55e" color-to="#86efac" />
          </BentoGridItem>

          <!-- SAP & Data -->
          <BentoGridItem class="md:col-span-2 group relative overflow-hidden">
            <template #title>
              <div class="flex items-center gap-2 text-yellow-500">
                <Database class="size-6" />
                <h3 class="text-xl font-bold">SAP & Data Solutions</h3>
              </div>
            </template>
            <template #description>
              <div class="space-y-4">
                <p class="text-sm">
                  Deep expertise in SAP ecosystems, integrating enterprise data with modern services.
                </p>
                <div class="flex flex-wrap gap-2">
                  <span
                    v-for="s in ['ABAP', 'SAP UI5', 'SAP Gateway', 'PostgreSQL', 'ELK']"
                    :key="s"
                    class="px-2 py-1 rounded-md bg-yellow-500/10 text-yellow-500 text-xs font-medium border border-yellow-500/20"
                  >
                    {{ s }}
                  </span>
                </div>
              </div>
            </template>
            <BorderBeam :size="150" :duration="16" color-from="#eab308" color-to="#fde047" />
          </BentoGridItem>
        </BentoGrid>
      </section>

      <!-- Education, Languages & Interests -->
      <section class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div
          class="p-8 rounded-2xl border border-neutral-200 dark:border-neutral-800 bg-neutral-50 dark:bg-neutral-900/50"
        >
          <h2 class="text-2xl font-bold mb-6">Education</h2>
          <div class="space-y-6">
            <div v-for="edu in education" :key="edu.title">
              <h3 class="font-bold">{{ edu.title }}</h3>
              <p class="text-sm text-neutral-500">{{ edu.institution }}</p>
              <p v-if="edu.timeFrame || edu.grade" class="text-xs mt-1 italic text-neutral-400">
                {{ [edu.timeFrame, edu.grade].filter(Boolean).join(' · ') }}
              </p>
            </div>
            <div>
              <h3 class="text-sm font-bold text-neutral-400 uppercase tracking-wider mb-3">
                Certifications
              </h3>
              <div v-for="cert in certifications" :key="cert.title">
                <h4 class="font-bold text-sm">{{ cert.title }}</h4>
                <p class="text-xs text-neutral-500">{{ cert.institution }}</p>
              </div>
            </div>
          </div>
        </div>

        <div
          class="p-8 rounded-2xl border border-neutral-200 dark:border-neutral-800 bg-neutral-50 dark:bg-neutral-900/50"
        >
          <h2 class="text-2xl font-bold mb-6">Languages</h2>
          <div class="space-y-4">
            <div
              v-for="(lang, idx) in languages"
              :key="lang.name"
              class="flex justify-between items-center"
              :class="{ 'border-t border-neutral-100 dark:border-neutral-800 pt-2': idx > 0 }"
            >
              <span>{{ lang.name }}</span>
              <span class="text-xs font-bold" :class="lang.color">{{ lang.level }}</span>
            </div>
          </div>
        </div>

        <div
          class="p-8 rounded-2xl border border-neutral-200 dark:border-neutral-800 bg-neutral-50 dark:bg-neutral-900/50"
        >
          <h2 class="text-2xl font-bold mb-6">Interests</h2>
          <div class="flex flex-wrap gap-3">
            <span
              v-for="interest in interests"
              :key="interest"
              class="px-4 py-2 rounded-full border border-neutral-200 dark:border-neutral-800 text-sm font-medium"
            >
              {{ interest }}
            </span>
          </div>
          <div class="mt-8">
            <h3 class="text-sm font-bold text-neutral-400 uppercase tracking-wider mb-3">
              Location
            </h3>
            <p class="text-sm">Chur, Graubünden, CH</p>
            <p class="text-xs text-neutral-500 mt-1 italic">
              Relocating to Helsinki, FI in April 2027 — See you there?
            </p>
          </div>
        </div>
      </section>
    </div>

    <!-- Footer -->
    <footer class="py-20 text-center border-t border-neutral-200 dark:border-neutral-800">
      <p class="text-neutral-500 text-sm">© 2026 Danilo Jakob · Chur, Switzerland</p>
    </footer>

    <!-- Floating Dock -->
    <div class="fixed bottom-6 inset-x-0 z-50 pointer-events-none flex justify-center">
      <div class="pointer-events-auto">
        <Dock>
          <DockIcon>
            <a
              href="https://github.com/DaniloMurer"
              target="_blank"
              class="flex size-full items-center justify-center text-foreground hover:text-blue-500 transition-colors"
            >
              <Github class="size-6" />
            </a>
          </DockIcon>
          <DockIcon>
            <a
              href="mailto:danilojakob@protonmail.com"
              class="flex size-full items-center justify-center text-foreground hover:text-blue-500 transition-colors"
            >
              <Mail class="size-6" />
            </a>
          </DockIcon>
          <DockIcon>
            <a
              href="https://linkedin.com"
              target="_blank"
              class="flex size-full items-center justify-center text-foreground hover:text-blue-500 transition-colors"
            >
              <Linkedin class="size-6" />
            </a>
          </DockIcon>
        </Dock>
      </div>
    </div>
  </div>
</template>
