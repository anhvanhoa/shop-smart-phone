<script setup lang="ts">
import { Icon } from '@iconify/vue'
import type { RouterLink } from 'vue-router'
import { Tippy } from 'vue-tippy'
const icons = [
    // {
    //     id: 1,
    //     name: 'solar:heart-linear',
    //     path: '/wishlist'
    // },
    {
        id: 2,
        name: 'hugeicons:notification-03',
        path: '/account/notification'
    },
    {
        id: 3,
        name: 'solar:cart-large-minimalistic-linear',
        path: '/cart'
    }
]
</script>
<template>
    <div class="flex items-center">
        <div class="flex items-center gap-3">
            <Tippy :delay="[500, 0]" placement="bottom" v-for="icon in icons" :key="icon.id">
                <RouterLink
                    :to="icon.path"
                    class="p-1.5 rounded-lg transition-all hover:bg-gray-100 block relative"
                >
                    <span
                        class="absolute size-1.5 right-0.5 top-0.5 bg-rose-500 rounded-full"
                    ></span>
                    <Icon :icon="icon.name" class="size-5" />
                </RouterLink>
                <template #content>
                    <div class="p-3 bg-white shadow-xl rounded-lg border">
                        <p class="text-sm">Bạn có 1 thông báo mới</p>
                    </div>
                </template>
            </Tippy>
        </div>
        <slot></slot>
    </div>
</template>
