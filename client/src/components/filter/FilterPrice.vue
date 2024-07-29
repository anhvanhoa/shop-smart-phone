<script setup lang="ts">
import { Icon } from '@iconify/vue';
</script>

<template>
    <div>
        <div class="grid grid-cols-2 gap-x-2">
            <div class="flex items-center my-2.5 gap-x-0.5 text-sm border rounded-lg py-2 px-2">
                <label><Icon class="size-4 mt-0.5" icon="tabler:currency-dong" /></label>
                <input class="outline-none w-full" type="text" placeholder="Từ" />
            </div>
            <div class="flex items-center my-2.5 gap-x-0.5 text-sm border rounded-lg py-2 px-2">
                <label><Icon class="size-4 mt-0.5" icon="tabler:currency-dong" /></label>
                <input class="outline-none w-full" type="text" placeholder="Đến" />
            </div>
        </div>
        <button class="w-full bg-primary rounded-lg py-2 mt-2 text-sm font-semibold text-black">
            Áp dụng
        </button>
    </div>
</template>
