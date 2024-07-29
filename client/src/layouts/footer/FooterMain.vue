<script setup lang="ts">
import images from '@/assets/images'
import { RouterLink } from 'vue-router'
const footer = {
    logo: images.logo,
    about: ' Lorem, ipsum dolor sit amet consectetur adipisicing elit. Beatae ipsam natus veritatis optio? Delectus maxime obcaecati iste aperiam! Assumenda at minima ullam repellendus quisquam dolore corporis cumque ea quaerat corrupti.',
    nameCompany: 'Nguyễn Văn Ánh Shop',
    address: 'Số 1, Đại Cồ Việt, Hai Bà Trưng, Hà Nội',
    phone: '0123456789',
    email: 'hotro@gmail.com',
    footerMenu: [
        {
            id: 1,
            title: 'Company',
            items: [
                {
                    id: 1,
                    name: 'About',
                    link: '/about'
                },
                {
                    id: 2,
                    name: 'Contact',
                    link: '/contact'
                }
            ]
        },
        {
            id: 2,
            title: 'Support',
            items: [
                {
                    id: 1,
                    name: 'FAQ',
                    link: '/faq'
                },
                {
                    id: 2,
                    name: 'Policy',
                    link: '/policy'
                }
            ]
        }
    ]
}
</script>
<template>
    <div class="pt-12">
        <div class="border-t mt-8">
            <section class="px-4 max-w-screen-xl mx-auto">
                <div class="rounded-3xl pt-8 px-4">
                    <div class="grid grid-cols-5 py-6 gap-x-8">
                        <div class="col-span-2 pr-8">
                            <div class="flex items-center">
                                <img class="size-8 rounded-md" :src="footer.logo" alt="" />
                                <span class="text-lg font-semibold pl-3">{{
                                    footer.nameCompany
                                }}</span>
                            </div>
                            <div class="text-xs py-4 text-balance leading-5">
                                {{ footer.about }}
                            </div>
                        </div>
                        <div v-for="menu in footer.footerMenu" :key="menu.id" class="text-sm">
                            <p class="font-semibold py-2">{{ menu.title }}</p>
                            <ul class="*:py-2">
                                <li v-for="item in menu.items" :key="item.id">
                                    <RouterLink :to="item.link" class="">
                                        {{ item.name }}
                                    </RouterLink>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
            </section>
            <div class="bg-gray-50">
                <div
                    class="flex items-center justify-between py-6 text-sm px-8 max-w-screen-xl mx-auto"
                >
                    <div class="flex items-center">
                        <span> Copyright © 2024 Nguyễn Văn Ánh.</span>
                        <div class="ml-12 flex items-center gap-x-6 *:underlin">
                            <RouterLink to="/terms" class="hover:underline">
                                Privacy Policy
                            </RouterLink>
                            <RouterLink to="/terms" class="hover:underline">
                                Terms of Service
                            </RouterLink>
                        </div>
                    </div>
                    <div>
                        <p class="flex items-center gap-x-2 mt-2">
                            <span>{{ footer.address }}</span> -
                            <span>
                                {{ footer.phone }}
                            </span>
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
