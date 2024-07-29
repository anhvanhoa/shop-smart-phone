<script setup lang="ts">
import FilterCheck from './FilterCheck.vue'
import FilterItem from './FilterItem.vue'
import FilterPrice from './FilterPrice.vue'

const filters = [
    {
        id: 1,
        title: 'Giá',
        type: 'price'
    },
    {
        id: 2,
        title: 'Hãng',
        type: 'check',
        data: [
            {
                id: 1,
                title: 'Apple',
                value: 'apple'
            },
            {
                id: 2,
                title: 'Samsung',
                value: 'samsung'
            },
            {
                id: 3,
                title: 'Xiaomi',
                value: 'xiaomi'
            },
            {
                id: 4,
                title: 'Oppo',
                value: 'oppo'
            },
            {
                id: 5,
                title: 'Vivo',
                value: 'vivo'
            },
            {
                id: 6,
                title: 'Realme',
                value: 'realme'
            },
            {
                id: 7,
                title: 'Nokia',
                value: 'nokia'
            },
            {
                id: 8,
                title: 'Huawei',
                value: 'huawei'
            },
            {
                id: 9,
                title: 'Vsmart',
                value: 'vsmart'
            },
            {
                id: 10,
                title: 'OnePlus',
                value: 'oneplus'
            },
            {
                id: 11,
                title: 'Sony',
                value: 'sony'
            },
            {
                id: 12,
                title: 'Asus',
                value: 'asus'
            },
            {
                id: 13,
                title: 'Google',
                value: 'google'
            },
            {
                id: 14,
                title: 'LG',
                value: 'lg'
            },
            {
                id: 15,
                title: 'HTC',
                value: 'htc'
            },
            {
                id: 16,
                title: 'Motorola',
                value: 'motorola'
            },
            {
                id: 17,
                title: 'Lenovo',
                value: 'lenovo'
            },
            {
                id: 18,
                title: 'BlackBerry',
                value: 'blackberry'
            },
            {
                id: 19,
                title: 'Meizu',
                value: 'meizu'
            },
            {
                id: 20,
                title: 'Coolpad',
                value: 'coolpad'
            },
            {
                id: 21,
                title: 'Wiko',
                value: 'wiko'
            },
            {
                id: 22,
                title: 'Mobell',
                value: 'mobell'
            },
            {
                id: 23,
                title: 'Masstel',
                value: 'masstel'
            },
            {
                id: 24,
                title: 'Itel',
                value: 'itel'
            },
            {
                id: 25,
                title: 'Black Shark',
                value: 'black-shark'
            },
            {
                id: 26,
                title: 'Poco',
                value: 'poco'
            },
            {
                id: 27,
                title: 'IQOO',
                value: 'iqoo'
            },
            {
                id: 28,
                title: 'Nubia',
                value: 'nubia'
            },
            {
                id: 29,
                title: 'ZTE',
                value: 'zte'
            },
            {
                id: 30,
                title: 'Sharp',
                value: 'sharp'
            },
            {
                id: 31,
                title: 'Fujitsu',
                value: 'fujitsu'
            },
            {
                id: 32,
                title: 'Panasonic',
                value: 'panasonic'
            },
            {
                id: 33,
                title: 'Acer',
                value: 'acer'
            },
            {
                id: 34,
                title: 'TCL',
                value: 'tcl'
            },
            {
                id: 35,
                title: 'Infinix',
                value: 'infinix'
            },
            {
                id: 36,
                title: 'Tecno',
                value: 'tecno'
            },
            {
                id: 37,
                title: 'Gionee',
                value: 'gionee'
            },
            {
                id: 38,
                title: 'Symphony',
                value: 'symphony'
            },
            {
                id: 39,
                title: 'Micromax',
                value: 'micromax'
            },
            {
                id: 40,
                title: 'Lava',
                value: 'lava'
            }
        ]
    },
    {
        id: 3,
        title: 'Ram',
        type: 'check',
        data: [
            {
                id: 1,
                title: 'Apple',
                value: 'apple'
            },
            {
                id: 2,
                title: 'Samsung',
                value: 'samsung'
            },
            {
                id: 3,
                title: 'Xiaomi',
                value: 'xiaomi'
            },
            {
                id: 4,
                title: 'Oppo',
                value: 'oppo'
            },
            {
                id: 5,
                title: 'Vivo',
                value: 'vivo'
            },
            {
                id: 6,
                title: 'Realme',
                value: 'realme'
            },
            {
                id: 7,
                title: 'Nokia',
                value: 'nokia'
            },
        ]
    },
]
</script>

<template>
    <FilterItem v-for="filter in filters" :key="filter.id" :title="filter.title">
        <FilterPrice v-if="filter.type === 'price'" />
        <FilterCheck v-if="filter.type === 'check' && filter.data" :data="filter.data" />
    </FilterItem>
    <div class="text-center">
        <button class="rounded-lg text-rose-500 hover:underline underline-offset-2">
            <span class="text-sm font-semibold">Xóa tất cả</span>
        </button>
    </div>
</template>
