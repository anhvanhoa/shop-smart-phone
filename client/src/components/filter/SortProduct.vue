<script setup lang="ts">
import { Icon } from '@iconify/vue/dist/iconify.js';
import { RouterLink } from 'vue-router'
import { Tippy } from 'vue-tippy'
const sorts = [
    {
        name: 'Mặc định',
        value: ''
    },
    {
        name: 'Mới nhất',
        value: 'newest'
    },
    {
        name: 'Giá thấp đến cao',
        value: 'price-asc'
    },
    {
        name: 'Giá cao đến thấp',
        value: 'price-desc'
    }
]
</script>

<template>
    <Tippy trigger="click" placement="bottom-end" interactive>
        <button
            class="cursor-pointer border py-1.5 px-3 text-sm rounded-lg flex items-center gap-x-1"
        >
            <span> Sắp xếp </span>
            <span>
                <Icon icon="gg:sort-az" class="size-5" />
            </span>
        </button>
        <template #content>
            <div class="py-3 px-4 bg-white *:whitespace-nowrap text-sm rounded-lg shadow-md">
                <RouterLink
                    v-for="sort in sorts"
                    :key="sort.value"
                    :to="`?sort=${sort.value}`"
                    class="cursor-pointer hover:underline py-1 block"
                    >{{ sort.name }}</RouterLink
                >
            </div>
        </template>
    </Tippy>
</template>
