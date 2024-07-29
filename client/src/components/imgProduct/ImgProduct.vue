<script setup lang="ts">
import { ref } from 'vue'
import { Icon } from '@iconify/vue'
import CarouselCom, { NextCarousel, PrevCarousel } from '@/components/carousel'
import { computed } from 'vue'

type Image = {
    id: number
    src: string
    name: string
}
const props = defineProps<{
    images: Image[]
}>()
const active = ref<number>(1)
const changeThumbnail = (id: number) => {
    active.value = id
}
const thumbnail = computed(() => {
    const image = props.images.find((image) => image.id === active.value)
    return image?.src
})
</script>
<template>
    <div class="bg-gray-100 rounded-3xl relative overflow-hidden">
        <div>
            <img :src="thumbnail" alt="macbook" class="w-full aspect-square object-contain" />
        </div>
    </div>
    <div class="mt-4">
        <CarouselCom :item-on-row="5">
            <div v-for="image in props.images" :key="image.id" class="p-2">
                <img
                    @mouseenter="changeThumbnail(image.id)"
                    :src="image.src"
                    :alt="image.name"
                    class="w-full aspect-square object-contain bg-gray-100 rounded-xl border-2"
                    :class="{
                        'border-rose-500': active === image.id,
                        'border-transparent': active !== image.id
                    }"
                />
            </div>
            <template #next>
                <NextCarousel class="-mr-8 bg-black/30 rounded-full">
                    <Icon icon="mingcute:right-line" class="p-1 size-8 text-white"
                /></NextCarousel>
            </template>
            <template #prev>
                <PrevCarousel class="-ml-8 bg-black/30 rounded-full">
                    <Icon icon="mingcute:left-line" class="p-1 size-8 text-white" />
                </PrevCarousel>
            </template>
        </CarouselCom>
    </div>
</template>
