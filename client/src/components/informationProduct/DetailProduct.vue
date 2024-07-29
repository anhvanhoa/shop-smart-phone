<script setup lang="ts">
const details = [
    {
        id: 1,
        name: 'Hãng',
        value: 'Apple'
    },
    {
        id: 2,
        name: 'Kích thước',
        value: '17 inch'
    },
    {
        id: 3,
        name: 'Bộ nhớ',
        value: '512GB | 1TB'
    }
]
</script>
<template>
    <div class="mb-4">
        <p class="font-extrabold text-lg">Chi tiết sản phẩm</p>
        <div class="relative overflow-x-auto mt-4">
            <table class="table-fixed w-full text-sm text-left rtl:text-right">
                <tbody>
                    <tr v-for="detail in details" :key="detail.id" :class="{
                        'bg-gray-100': detail.id % 2 !== 0
                    }" >
                        <td scope="row" class="px-3 py-2 font-medium pr-12">
                            {{ detail.name }}
                        </td>
                        <td class="px-3 py-2">{{ detail.value }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>
