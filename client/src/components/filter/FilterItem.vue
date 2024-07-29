<script setup lang="ts">
import { Icon } from '@iconify/vue/dist/iconify.js'
import { ref } from 'vue'
const props = defineProps<{
    title: string
}>()
const active = ref(true)
const handleClick = () => {
    active.value = !active.value
}
</script>
<template>
    <div class="mb-4 border-b pb-4">
        <p
            @click="handleClick"
            class="flex items-center justify-between px-1 font-semibold cursor-pointer"
        >
            <span class="text-sm">{{ props.title }}</span>
            <span class="p-1 hover:bg-gray-100 rounded-md cursor-pointer select-none">
                <Icon icon="grommet-icons:down" class="size-3 text-gray-500" />
            </span>
        </p>
        <Transition>
            <slot v-if="active"></slot>
        </Transition>
    </div>
</template>

<style scoped>
.v-enter-active,
.v-leave-active {
    transition: opacity 0.1s linear;
}

.v-enter-from,
.v-leave-to {
    opacity: 0;
}
</style>
