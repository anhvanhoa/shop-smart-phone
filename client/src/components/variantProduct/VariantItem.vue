<script setup lang="ts">
const props = defineProps<{
    title: string
    variants: { name: string; value: string; id: number }[]
}>()
</script>
<template>
    <div class="flex mb-6">
        <p class="font-semibold mr-6">{{ props.title }}</p>
        <div class="flex items-center flex-wrap gap-3">
            <label class="peer select-none" v-for="variant in props.variants" :key="variant.id">
                <input type="radio" :name="props.title" :value="variant.value" hidden class="peer" />
                <span
                    class="block text-sm cursor-pointer bg-gray-100 rounded-md px-3 py-1.5 peer-checked:bg-black/80 peer-checked:text-white font-semibold"
                >
                    {{ variant.name }}
                </span>
            </label>
        </div>
    </div>
</template>
