<script setup lang="ts">
import { Icon } from '@iconify/vue/dist/iconify.js'
import ProductItem from '../productItem'
import CountDown from './CountDown.vue'
import CarouselCom from '../carousel/CarouselCom.vue'

const products = [
    {
        id: 1,
        name: 'MacBook Air 13 inch M2 | Chính hãng',
        price: 18900000,
        priceOrigin: 20990000,
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
        liked: false,
        thumbnail:
            'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        discount: 12,
        slug: 'macbook-air-13-inch-m2-chinh-hang-4'
    },
    {
        id: 3,
        name: 'MacBook Air 13 inch M2 | Chính hãng',
        price: 19900000,
        priceOrigin: 24990000,
        liked: false,
        thumbnail:
            'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        discount: 12,
        slug: 'macbook-air-13-inch-m2-chinh-hang-4'
    },
    {
        id: 4,
        name: 'MacBook Air 13 inch M2 | Chính hãng',
        price: 19900000,
        priceOrigin: 24990000,
        liked: false,
        thumbnail:
            'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        discount: 11,
        slug: 'macbook-air-13-inch-m2-chinh-hang-4'
    },
    {
        id: 5,
        name: 'MacBook Air 13 inch M2 | Chính hãng',
        price: 19900000,
        priceOrigin: 24990000,
        liked: false,
        thumbnail:
            'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        discount: 12,
        slug: 'macbook-air-13-inch-m2-chinh-hang-4'
    },
    {
        id: 6,
        name: 'MacBook Air 13 inch M2 | Chính hãng',
        price: 19900000,
        priceOrigin: 24990000,
        liked: false,
        thumbnail:
            'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        discount: 12,
        slug: 'macbook-air-13-inch-m2-chinh-hang-4'
    },
    {
        id: 7,
        name: 'MacBook Air 13 inch M2 | Chính hãng',
        price: 19900000,
        priceOrigin: 24990000,
        liked: false,
        thumbnail:
            'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        discount: 12,
        slug: 'macbook-air-13-inch-m2-chinh-hang-4'
    },
    {
        id: 8,
        name: 'MacBook Air 13 inch M2 | Chính hãng',
        price: 19900000,
        priceOrigin: 24990000,
        liked: false,
        thumbnail:
            'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        discount: 5,
        slug: 'macbook-air-13-inch-m2-chinh-hang-4'
    }
]
</script>

<template>
    <section class="max-w-screen-xl mx-auto px-8 py-8">
        <div class="flex items-center justify-between">
            <div class="flex items-center text-orange-600">
                <Icon icon="typcn:flash-outline" class="size-8" />
                <h3 class="text-2xl uppercase font-semibold">Flash Sale</h3>
            </div>
            <CountDown />
        </div>
        <CarouselCom :item-on-row="4" class="bg-gray-100 rounded-xl px-8 py-8 mt-6 -mx-4">
            <div class="px-4" v-for="product in products" :key="product.id">
                <ProductItem :product="product" class="!bg-white" />
            </div>
        </CarouselCom>
    </section>
</template>
