<script setup lang="ts">
import { ref } from 'vue'
import NextCarousel from './NextCarousel.vue'
import PrevCarousel from './PrevCarousel.vue'
const positionStyle = ref(0)
const props = defineProps<{
    itemOnRow: 1 | 2 | 3 | 4 | 5
}>()
const cols = {
    1: '*:w-full',
    2: '*:w-1/2',
    3: '*:w-1/3',
    4: '*:w-1/4',
    5: '*:w-1/5'
}
const carousel = ref<HTMLElement | null>(null)
const handleCarousel = (type: 'next' | 'prev') => {
    if (!carousel.value) return
    const total = carousel.value.children.length / props.itemOnRow - 1
    const countCarousel = Math.floor(total)
    const orther = total - countCarousel
    if (type === 'next') {
        if (positionStyle.value === countCarousel && orther) positionStyle.value += orther
        if (positionStyle.value !== countCarousel + orther) {
            positionStyle.value++
        }
    } else {
        if (positionStyle.value !== 0) {
            if (positionStyle.value === orther) positionStyle.value -= orther
            else positionStyle.value--
        }
    }
}
</script>

<template>
    <div class="relative select-none">
        <!-- Carousel -->
        <div class="overflow-hidden rounded-xl">
            <div
                ref="carousel"
                class="flex *:flex-shrink-0 transition-all duration-300 relative"
                :style="{
                    left: `-${positionStyle * 100}%`
                }"
                :class="`${cols[props.itemOnRow]}`"
            >
                <slot></slot>
            </div>
        </div>
        <!-- Navigate -->
        <div>
            <span @click="handleCarousel('prev')">
                <slot name="prev">
                    <PrevCarousel />
                </slot>
            </span>
            <span @click="handleCarousel('next')">
                <slot name="next">
                    <NextCarousel />
                </slot>
            </span>
        </div>
    </div>
</template>
