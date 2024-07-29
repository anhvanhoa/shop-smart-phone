<script setup lang="ts">
import ProductItem from './productItem'
const props = defineProps<{
    title: string
}>()
const products = [
    {
        id: 1,
        name: 'MacBook Air 13 inch M2 | Chính hãng',
        price: 18900000,
        priceOrigin: 20990000,
        sales: 4,
        star: "0",
        liked: false,
        thumbnail:
            'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        discount: 0,
        slug: 'macbook-air-13-inch-m2-chinh-hang'
    },
    {
        id: 2,
        name: 'MacBook Air 13 inch M2 | Chính hãng',
        price: 19900000,
        priceOrigin: 24990000,
        sales: 4,
        star: "3",
        liked: false,
        thumbnail:
            'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        discount: 12,
        slug: 'macbook-air-13-inch-m2-chinh-hang3'
    }
]
</script>
<template>
    <div class="max-w-screen-xl mx-auto px-4 py-8">
        <h2 class="text-2xl font-bold">{{props.title}}</h2>
        <div class="grid grid-cols-5 gap-4 mt-6">
            <ProductItem v-for="product in products" :product="product" :key="product.id" class="bg-gray-50" />
        </div>
    </div>
</template>
