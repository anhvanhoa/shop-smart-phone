<script setup lang="ts">
import { reactive } from 'vue'
import { RouterLink } from 'vue-router'
import CarouselCom, { NextCarousel, PrevCarousel } from '@/components/carousel'
import { Icon } from '@iconify/vue/dist/iconify.js'
const banners = reactive([
    {
        id: 1,
        src: 'https://cdn.hoanghamobile.com/i/home/Uploads/2024/07/01/thu-cu-t7-web_638554214841757467.png',
        alt: 'banner 1',
        link: '/png'
    },
    {
        id: 2,
        src: 'https://cdn.hoanghamobile.com/i/home/Uploads/2024/06/27/web-ss.png',
        alt: 'banner 2',
        link: '/'
    },
    {
        id: 3,
        src: 'https://cdn.hoanghamobile.com/i/home/Uploads/2024/06/08/xiaomi-redmi-12-web.png',
        alt: 'banner 3',
        link: '/'
    }
])
</script>

<template>
    <div class="max-w-screen-xl mx-auto px-4 py-8 relative select-none">
        <!-- Carousel -->
        <CarouselCom :itemOnRow="1" class="col-span-4">
            <template v-for="banner in banners" :key="banner.id">
                <RouterLink :to="banner.link">
                    <img
                        class="aspect-[16/5] w-full h-full object-cover"
                        :src="banner.src"
                        :alt="banner.alt"
                    />
                </RouterLink>
            </template>
            <template #next>
                <NextCarousel class="mr-4">
                    <Icon icon="mingcute:right-line" class="size-8 text-white"
                /></NextCarousel>
            </template>
            <template #prev>
                <PrevCarousel class="ml-4">
                    <Icon icon="mingcute:left-line" class="size-8 text-white" />
                </PrevCarousel>
            </template>
        </CarouselCom>
    </div>
</template>
