<script setup lang="ts">
import ImgProduct from '@/components/imgProduct'
import {
    InformationProduct,
    DetailProduct,
    DesciptionProduct
} from '@/components/informationProduct'
import ProductItem from '@/components/productItem'
import ReviewsProduct from '@/components/reviewsProduct'
import VariantProduct from '@/components/variantProduct'
import { Icon } from '@iconify/vue'
const images = [
    {
        id: 1,
        src: 'https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg',
        name: 'Macbook m1'
    },
    {
        id: 2,
        src: 'https://down-vn.img.susercontent.com/file/sg-11134301-7rcc9-ltm7wr6f2v3k36',
        name: 'Macbook m2'
    },
    {
        id: 3,
        src: 'https://down-vn.img.susercontent.com/file/sg-11134301-7rcc9-ltm7wr6f2v3k36',
        name: 'Macbook m3'
    },
    {
        id: 4,
        src: 'https://down-vn.img.susercontent.com/file/sg-11134301-7rcc9-ltm7wr6f2v3k36',
        name: 'Macbook m4'
    },
    {
        id: 5,
        src: 'https://down-vn.img.susercontent.com/file/sg-11134301-7rcc9-ltm7wr6f2v3k36',
        name: 'Macbook m5'
    },
    {
        id: 6,
        src: 'https://down-vn.img.susercontent.com/file/sg-11134301-7rcc9-ltm7wr6f2v3k36',
        name: 'Macbook m6'
    },
    {
        id: 7,
        src: 'https://down-vn.img.susercontent.com/file/sg-11134301-7rcc9-ltm7wr6f2v3k36',
        name: 'Macbook m7'
    }
]
const product = {
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
    slug: 'macbook-air-13-inch-m2-chinh-hang',
    countComment: 8,
    quantity: 20
}
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
    <div>
        <section class="max-w-screen-xl mx-auto px-4">
            <div class="mt-8">
                <p class="text-sm flex items-center gap-x-2 font-semibold">
                    <span>Home</span>
                    <Icon icon="ph:caret-right" class="size-3" />
                    <span>Laptop</span>
                    <Icon icon="ph:caret-right" class="size-3" />
                    <span class="text-gray-500">Macbook m1</span>
                </p>
            </div>
            <div class="grid grid-cols-12 mt-6">
                <div class="col-span-5 mr-12">
                    <ImgProduct :images="images" />
                    <ReviewsProduct />
                </div>
                <div class="col-span-7 ml-12">
                    <InformationProduct :product="product" />
                    <VariantProduct />
                    <DetailProduct />
                    <DesciptionProduct />
                </div>
            </div>
            <div class="mt-10">
                <p class="font-semibold text-xl uppercase">Có thể bạn sẽ thích</p>
                <div class="grid grid-cols-5 gap-x-2 gap-y-4 mt-6">
                    <ProductItem
                        class="bg-gray-100"
                        v-for="product in products"
                        :key="product.id"
                        :product="product"
                    />
                </div>
            </div>
        </section>
    </div>
</template>
