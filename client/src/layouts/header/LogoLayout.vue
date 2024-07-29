<script setup lang="ts">
import images from '@/assets/images'
import type { RouterLink } from 'vue-router'
</script>
<template>
    <div>
        <RouterLink to="/">
            <img :src="images.logo" alt="Ánh Văn Hóa Store" class="size-9 rounded-md" />
        </RouterLink>
    </div>
</template>
