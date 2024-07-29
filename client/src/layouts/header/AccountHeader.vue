<script setup lang="ts">
import images from '@/assets/images';
import  { RouterLink } from 'vue-router';
</script>
<template>
    <div class="flex items-center">
        <span class="mx-4 h-4 bg-gray-300 rounded-xl w-0.5"></span>
        <RouterLink to="/account">
            <img alt="avatar" :src="images.logo" class="size-7 rounded-full"/>
        </RouterLink>
    </div>
</template>