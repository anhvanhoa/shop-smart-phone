<script setup lang="ts">
import { Icon } from '@iconify/vue/dist/iconify.js'
import type { InformationProductProps } from '@/types/product'
import { priceFormat } from '@/utils/helper'
const props = defineProps<{
    product: InformationProductProps
}>()
const stars = Array.from({ length: 5 }, (_, index) => index + 1)
</script>

<template>
    <div class="border-b pb-8 mb-8">
        <div class="flex items-center justify-between">
            <div class="flex items-center gap-x-3">
                <span
                    class="bg-blue-500/10 text-blue-500 rounded-full px-3 py-1.5 text-sm font-semibold"
                >
                    Hàng mới
                </span>
                <span
                    v-if="props.product.discount"
                    class="px-2 py-0.5 rounded-full bg-rose-500/10 text-rose-500 text-sm font-semibold"
                >
                    {{ props.product.discount }}%
                </span>
            </div>
            <span class="px-1 py-0.5 bg-gray-100 rounded-xl cursor-pointer">
                <Icon icon="solar:menu-dots-bold" class="size-5" />
            </span>
        </div>
        <div>
            <h1 class="text-[32px] py-4 font-semibold leading-[48px] text-pretty">
                {{ props.product.name }}
            </h1>
        </div>
        <div class="flex items-center gap-x-3">
            <div class="flex items-center gap-x-1">
                {{ props.product.star }}
                <p class="flex items-center gap-x-0.5">
                    <span v-for="star in stars" :key="star">
                        <Icon
                            v-if="Number(props.product.star) >= star"
                            icon="iconoir:star-solid"
                            class="text-orange-500"
                        />
                        <Icon
                            v-if="Number(props.product.star) < star"
                            icon="ph:star-light"
                            class="text-orange-500"
                        />
                    </span>
                </p>
            </div>
            <span class="w-0.5 bg-gray-300 h-3 rounded-full"></span>
            <p>{{ props.product.countComment }} đánh giá</p>
            <span class="w-0.5 bg-gray-300 h-3 rounded-full"></span>
            <p>{{ props.product.quantity }} sản phẩm</p>
        </div>
        <div class="flex items-center mt-4 gap-x-4">
            <p v-if="props.product.priceOrigin" class="font-semibold line-through text-lg leading-4 text-gray-500">
                {{ priceFormat(props.product.priceOrigin) }}
            </p>
            <p class="font-semibold text-2xl text-rose-500">{{ priceFormat(props.product.price) }}</p>
        </div>
    </div>
</template>
