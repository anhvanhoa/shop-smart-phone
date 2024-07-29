<script setup lang="ts">
import { Icon } from '@iconify/vue'
import FilterReviews from './FilterReviews.vue'
const comments = [
    {
        id: 1,
        avatar: 'http://localhost:5173/src/assets/images/logo.png',
        name: 'Jond',
        date: '14/06/2022',
        product: 'phân loại đen, ram 8GB',
        star: 4,
        content:
            'Lorem ipsum dolor sit amet consectetur adipisicing elit. Dolore ad aut exercitationem nisi sit ipsa, fuga dolorum. Unde molestias perspiciatis voluptate impedit omnis perferendis odio, doloremque sed dignissimos voluptas maiores.'
    },
    {
        id: 2,
        avatar: 'http://localhost:5173/src/assets/images/logo.png',
        name: 'Jondny',
        date: '14/06/2022',
        product: 'phân loại đen, ram 8GB',
        star: 3,
        content:
            'Lorem ipsum dolor sit amet consectetur adipisicing elit. Dolore ad aut exercitationem nisi sit ipsa, fuga dolorum. Unde molestias perspiciatis voluptate impedit omnis perferendis odio, doloremque sed dignissimos voluptas maiores.'
    }
]
const stars = [1, 2, 3, 4, 5]
</script>
<template>
    <div class="border rounded-2xl py-3 px-4 mt-8">
        <h4 class="text-xl font-semibold py-2">Đánh giá</h4>
        <FilterReviews />
        <div class="mt-5">
            <div class="flex gap-x-3 mb-4" v-for="comment in comments" :key="comment.id">
                <img :src="comment.avatar" :alt="comment.name" class="size-9 rounded-full" />
                <div class="text-sm flex flex-col gap-y-1">
                    <p class="font-semibold leading-4">{{ comment.name }}</p>
                    <p class="text-xs">{{ comment.date }} | {{ comment.product }}</p>
                    <div class="flex items-center gap-x-1">
                        <span v-for="star in stars" :key="star">
                            <Icon
                                v-if="comment.star >= star"
                                icon="iconoir:star-solid"
                                class="text-orange-500 size-3"
                            />
                            <Icon
                                v-if="comment.star < star"
                                icon="ph:star-light"
                                class="text-orange-500 size-3"
                            />
                        </span>
                    </div>
                    <p>
                        {{ comment.content }}
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>
