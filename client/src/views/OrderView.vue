<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
    <div class="p-2">
        <h3 class="text-xl font-semibold">Đơn hàng</h3>

        <div class="relative overflow-x-auto mt-4">
            <table
                class="w-full text-sm text-left rtl:text-right text-gray-500 rounded-lg overflow-hidden"
            >
                <thead class="text-xs text-gray-700 uppercase bg-gray-50">
                    <tr>
                        <th scope="col" class="px-6 py-3">Mã đơn hàng</th>
                        <th scope="col" class="px-6 py-3">Số lượng</th>
                        <th scope="col" class="px-6 py-3">Tổng tiền</th>
                        <th scope="col" class="px-6 py-3">Trạng thái</th>
                    </tr>
                </thead>
                <tbody>
                    <tr class="bg-white border-b">
                        <th
                            scope="row"
                            class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                        >
                            <RouterLink to="orders/1"> Apple MacBook Pro 17" </RouterLink>
                        </th>
                        <td class="px-6 py-4">Silver</td>
                        <td class="px-6 py-4">Laptop</td>
                        <td class="px-6 py-4">$2999</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>
