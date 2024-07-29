<script setup lang="ts">
import { BannerHome } from '@/components/banner'
import { FlashSale } from '@/components/flashSale'
import ProductsHome from '@/components/ProductsHome.vue'
</script>

<template>
    <main>
        <BannerHome />
        <FlashSale />
        <ProductsHome title="Sản phẩm của chúng tôi" />
    </main>
</template>
