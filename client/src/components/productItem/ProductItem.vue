<script setup lang="ts">
import type { ProductProps } from '@/types/product'
import { Icon } from '@iconify/vue'
import { ref } from 'vue'
import { RouterLink } from 'vue-router'
import { priceFormat } from '@/utils/helper'
const props = defineProps<{
    product: ProductProps
}>()
const like = ref(props.product.liked)
const handleLike = () => {
    like.value = !like.value
}
</script>
<template>
    <div class="rounded-xl">
        <RouterLink :to="`/product/${product.slug}`">
            <div class="py-2 px-2 flex justify-between items-center">
                <div>
                    <span
                        v-if="props.product.discount"
                        class="text-xs ml-2 px-2 py-1.5 font-semibold bg-rose-500 rounded-full text-white"
                        >{{ props.product.discount }}%</span
                    >
                </div>
                <span
                    @click.prevent="handleLike"
                    class="p-1.5 rounded-lg transition-all hover:bg-gray-50 cursor-pointer"
                >
                    <Icon
                        icon="solar:heart-bold"
                        :class="{
                            'text-rose-600': like,
                            'text-gray-300': !like
                        }"
                        class="size-6 transition-all"
                    />
                </span>
            </div>
            <div class="px-6 select-none">
                <div class="w-full aspect-square">
                    <img
                        class="w-full h-full object-contain"
                        :src="props.product.thumbnail"
                        :alt="props.product.name"
                    />
                </div>
            </div>
            <div class="p-4">
                <p v-if="product.priceOrigin" class="text-sm py-1 text-gray-500 line-through">
                    {{ priceFormat(product.priceOrigin) }}
                </p>
                <h2 class="font-semibold text-pretty text-sm">
                    {{ props.product.name }}
                </h2>
                <p class="font-bold text-rose-600 py-1">{{ priceFormat(product.price) }}</p>
                <div
                    v-if="product.star || product.sales"
                    class="flex items-center justify-between mt-3"
                >
                    <p v-if="product.star" class="flex items-center text-sm gap-x-1 font-semibold">
                        <span> {{ product.star }}</span>
                        <Icon icon="iconoir:star-solid" class="text-orange-500" />
                    </p>
                    <p v-if="product.sales" class="text-xs text-gray-500">
                        Đã bán {{ product.sales }}
                    </p>
                </div>
            </div>
        </RouterLink>
    </div>
</template>
