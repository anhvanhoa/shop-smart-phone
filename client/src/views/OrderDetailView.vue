<script setup lang="ts">
import { Icon } from '@iconify/vue/dist/iconify.js'
</script>

<template>
    <div class="p-4">
        <div class="flex items-center justify-between">
            <h3 class="font-semibold text-xl">Đơn hàng</h3>
            <p class="flex items-center gap-x-3">
                <span class="uppercase"> asdha8as8asd90aad6 </span>
                <span class="block h-6 rounded-full w-0.5 bg-gray-200"></span>
                <span class="text-rose-600"> Đang giao hàng </span>
            </p>
        </div>
        <div class="p-2">
            <div class="bg-gray-100 p-3 rounded-lg text-sm my-2">
                <p class="font-semibold">Nguyễn Văn Ánh</p>
                <p>0334376855</p>
                <p>Số 1, Đại Cồ Việt, Hai Bà Trưng, Hà Nội</p>
            </div>
            <div class="my-4 flex justify-between border p-3 rounded-lg">
                <div class="flex gap-x-2.5">
                    <img
                        src="http://localhost:5173/src/assets/images/product.png"
                        alt=""
                        class="size-14 bg-gray-200/80 rounded-lg"
                    />
                    <div>
                        <p class="font-semibold">Tên sản phẩm</p>
                        <p class="text-sm">19.000.000đ</p>
                    </div>
                </div>
                <div class="text-sm">
                    <p class="flex items-center justify-end gap-x-0.5">
                        <Icon icon="material-symbols-light:close" />
                        3
                    </p>
                    <p class="font-semibold">39.000.000đ</p>
                </div>
            </div>
            <div class="my-4 flex justify-between border p-3 rounded-lg">
                <div class="flex gap-x-2.5">
                    <img
                        src="http://localhost:5173/src/assets/images/product.png"
                        alt=""
                        class="size-14 bg-gray-200/80 rounded-lg"
                    />
                    <div>
                        <p class="font-semibold">Tên sản phẩm</p>
                        <p class="text-sm">19.000.000đ</p>
                    </div>
                </div>
                <div class="text-sm">
                    <p class="flex items-center justify-end gap-x-0.5">
                        <Icon icon="material-symbols-light:close" />
                        3
                    </p>
                    <p class="font-semibold">39.000.000đ</p>
                </div>
            </div>
            <div class="mt-8 px-3 flex flex-col gap-4 text-sm">
                <p class="flex items-center justify-between">
                    <span>Phương thức thanh toán:</span>
                    <span>Thanh toán khi nhận hàng</span>
                </p>
                <p class="flex items-center justify-between">
                    <span>Tổng tiền hàng:</span>
                    <span>19.000.000đ</span>
                </p>
                <p class="flex items-center justify-between">
                    <span>Voucher giảm giá:</span>
                    <span>- 19.000.000đ</span>
                </p>
                <p class="flex items-center justify-between">
                    <span>Giảm giá sản phẩm:</span>
                    <span>- 19.000.000đ</span>
                </p>
                <p class="flex items-center justify-between">
                    <span>Tiết kiệm:</span>
                    <span>- 19.000.000đ</span>
                </p>
                <p class="flex items-center justify-between font-semibold text-base">
                    <span>Tổng số tiền:</span>
                    <span>19.000.000đ</span>
                </p>
            </div>
        </div>
    </div>
</template>
