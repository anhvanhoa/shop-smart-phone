<script setup lang="ts">
import { Icon } from '@iconify/vue'
</script>

<template>
    <span class="cursor-pointer absolute top-1/2 transform -translate-y-1/2 right-6">
        <slot>
            <span
                class="block p-2 rounded-full bg-black/30 text-white hover:scale-105 transition-all"
            >
                <Icon icon="hugeicons:arrow-right-02" class="size-6" />
            </span>
        </slot>
    </span>
</template>
