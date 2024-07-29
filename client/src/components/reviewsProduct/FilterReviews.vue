<script setup lang="ts">
const filters = [
    {
        id: 2,
        name: '5 Sao'
    },
    {
        id: 3,
        name: '4 Sao'
    },
    {
        id: 4,
        name: '3 Sao'
    },
    {
        id: 5,
        name: '2 Sao'
    },
    {
        id: 6,
        name: '1 Sao'
    }
]
</script>
<template>
    <div class="flex items-center gap-x-2 text-sm">
        <label class="peer select-none" v-for="filter in filters" :key="filter.id">
            <input type="checkbox" name="todo" hidden class="peer" />
            <p
                class="text-xs cursor-pointer rounded-md px-2 py-1 border border-slate-400 text-slate-600 peer-checked:border-orange-500 peer-checked:text-orange-500 font-semibold"
            >
                {{ filter.name }}
            </p>
        </label>
    </div>
</template>
