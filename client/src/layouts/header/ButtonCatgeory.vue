<script setup lang="ts">
import { Icon } from '@iconify/vue'
import { ref, onMounted, onUnmounted } from 'vue'
import { RouterLink } from 'vue-router'
const categories = [
    {
        path: '/dien-thoai',
        name: 'Điện thoại'
    },
    {
        path: '/may-tinh',
        name: 'Máy tính'
    }
]
const active = ref(false)
const box = ref<HTMLUListElement | null>(null)
const toggleActive = () => (active.value = !active.value)
const handleClickOutside = (event: MouseEvent) => {
    if (box.value && !box.value.contains(event.target as Node)) {
        active.value = false
    }
}
onMounted(() => {
    document.addEventListener('click', handleClickOutside)
})
onUnmounted(() => {
    document.removeEventListener('click', handleClickOutside)
})
</script>
<template>
    <div ref="box" class="relative">
        <button @click="toggleActive" class="flex items-center bg-gray-100 px-3 py-2 rounded-md">
            <Icon class="size-5" icon="fluent:layout-column-two-split-left-16-regular" />
            <span class="pl-1 text-sm font-semibold">Danh mục</span>
        </button>
        <Transition>
            <ul
                v-if="active"
                class="min-w-full drop-shadow-md max-w-64 absolute z-10 top-full mt-3 rounded-md max-h-72 overflow-y-auto bg-white py-1"
            >
                <li
                    @click="toggleActive"
                    class="hover:bg-gray-100/80 text-sm whitespace-nowrap"
                    v-for="category in categories"
                    v-bind:key="category.path"
                >
                    <RouterLink class="block py-1.5 px-3" :to="category.path">
                        {{ category.name }}
                    </RouterLink>
                </li>
            </ul>
        </Transition>
    </div>
</template>

<style scoped>
.v-enter-active,
.v-leave-active {
    transition: opacity 0.2s ease;
}

.v-enter-from,
.v-leave-to {
    opacity: 0;
}
</style>
