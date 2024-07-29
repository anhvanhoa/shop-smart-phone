<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
const menus = [
    {
        id: 1,
        name: 'Thông tin tài khoản',
        path: '/account'
    },
    {
        id: 2,
        name: 'Đơn hàng',
        path: '/account/orders'
    },
        // {
        //     id: 3,
        //     name: 'Cài đặt',
        //     path: '/account/settings'
        // }
]
</script>

<template>
    <div class="max-w-screen-xl mx-auto px-8">
        <div class="grid grid-cols-12 mt-8">
            <div class="col-span-3 border rounded-xl p-4 h-max mt-4">
                <RouterLink
                    v-for="menu in menus"
                    :key="menu.id"
                    :to="menu.path"
                    class="block py-1 text-sm"
                >
                    {{ menu.name }}
                </RouterLink>
            </div>
            <div class="col-span-9 ml-16">
                <RouterView />
            </div>
        </div>
    </div>
</template>
