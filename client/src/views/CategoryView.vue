<script setup lang="ts">
import { Sort, Filter } from '@/components/filter'
import ProductItem from '@/components/productItem'
import { Icon } from '@iconify/vue/dist/iconify.js'
const products = [
    {
        id: 1,
        name: 'MacBook Air 13 inch M2 | Chính hãng',
        price: 18900000,
        priceOrigin: 20990000,
        sales: 4,
        star: '0',
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
        star: '3',
        liked: false,
        thumbnail:
            'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        discount: 12,
        slug: 'macbook-air-13-inch-m2-chinh-hang3'
    },
    {
        id: 3,
        name: 'MacBook Air 13 inch M2 | Chính hãng',
        price: 19900000,
        priceOrigin: 24990000,
        sales: 4,
        star: '3',
        liked: false,
        thumbnail:
            'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        discount: 12,
        slug: 'macbook-air-13-inch-m2-chinh-hang3'
    },
    {
        id: 5,
        name: 'MacBook Air 13 inch M2 | Chính hãng',
        price: 18900000,
        priceOrigin: 20990000,
        sales: 4,
        star: '0',
        liked: false,
        thumbnail:
            'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        discount: 0,
        slug: 'macbook-air-13-inch-m2-chinh-hang'
    },
    {
        id: 4,
        name: 'MacBook Air 13 inch M2 | Chính hãng',
        price: 19900000,
        priceOrigin: 24990000,
        sales: 4,
        star: '3',
        liked: false,
        thumbnail:
            'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        discount: 12,
        slug: 'macbook-air-13-inch-m2-chinh-hang3'
    }
]
</script>

<template>
    <main class="max-w-screen-xl mx-auto px-4 grid grid-cols-12 my-8">
        <section class="col-span-3 mr-6 pr-4 bg-gray- rounded-xl px-2 mt-2">
            <Filter />
        </section>
        <section class="col-span-9">
            <div class="flex items-center justify-between">
                <p class="text-sm flex items-center gap-x-2">
                    <span class="font-semibold">Home</span>
                    <Icon icon="ph:caret-right" class="size-3" />
                    <span class="text-gray-500">Laptop</span>
                </p>
                <Sort />
            </div>
            <div class="grid grid-cols-4 gap-x-2 gap-y-4 mt-6">
                <ProductItem
                    class="bg-gray-100"
                    v-for="product in products"
                    :key="product.id"
                    :product="product"
                />
            </div>
        </section>
    </main>
</template>
