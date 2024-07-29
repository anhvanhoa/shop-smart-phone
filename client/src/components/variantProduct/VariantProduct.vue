<script setup lang="ts">
import { Icon } from '@iconify/vue';
import VariantItem from './VariantItem.vue'

const data = [
    {
        title: 'Color',
        variants: [
            { name: 'Màu đen', value: 'black', id: 1 },
            { name: 'Màu đỏ', value: 'red', id: 2 },
            { name: 'Màu xanh', value: 'blue', id: 3 }
        ]
    },
    {
        title: 'Ram',
        variants: [
            { name: '8GB', value: '8gb', id: 4 },
            { name: '16GB', value: '16gb', id: 5 },
            { name: '24GB', value: '24gb', id: 6 }
        ]
    }
]
</script>
<template>
    <div class="mb-12">
        <VariantItem v-for="item in data" :key="item.title" v-bind="item" />
        <div class="flex items-center gap-x-4 pt-4">
            <button class="bg-gray-100 p-2 rounded-lg text-gray-400">
                <Icon icon="solar:heart-bold" class="size-6" />
            </button>
            <button class="">
                <span
                    class="border border-rose-500 py-2.5 px-3 text-rose-500 rounded-lg text-sm font-semibold"
                    >Thêm vào giỏ hàng</span
                >
            </button>
            <button class="bg-rose-500 text-white py-2.5 px-14 rounded-lg text-sm font-semibold">
                <span class="text-sm font-semibold">Mua ngay</span>
            </button>
        </div>
    </div>
</template>
