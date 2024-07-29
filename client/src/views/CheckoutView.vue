<script setup lang="ts">
import { Icon } from '@iconify/vue/dist/iconify.js'
</script>

<template>
    <div>
        <div class="max-w-screen-xl mx-auto px-8">
            <div class="grid grid-cols-12 mt-8 gap-x-20">
                <div class="col-span-7">
                    <h3 class="text-2xl font-semibold">Thanh toán</h3>
                    <div class="py-4 grid grid-cols-2 gap-x-8 gap-y-2">
                        <div class="text-sm">
                            <p class="py-2">Họ và Tên</p>
                            <input type="text" class="border p-2 rounded-lg outline-none w-full" />
                        </div>
                        <div class="text-sm">
                            <p class="py-2">Số điện thoại</p>
                            <input type="text" class="border p-2 rounded-lg outline-none w-full" />
                        </div>
                        <div class="text-sm col-span-2">
                            <p class="py-2">Địa chỉ</p>
                            <input type="text" class="border p-2 rounded-lg outline-none w-full" />
                        </div>
                        <div class="text-sm col-span-2">
                            <p class="py-2">Ghi chú</p>
                            <input type="text" class="border p-2 rounded-lg outline-none w-full" />
                        </div>
                    </div>
                    <div class="mt-4">
                        <p class="font-semibold text-sm">Phương thức thanh toán</p>
                        <div class="mt-3">
                            <button
                                type="button"
                                class="text-gray-900 bg-white hover:bg-gray-100 border border-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center me-2 mb-2"
                            >
                                <Icon icon="guidance:money" class="size-5 mr-2" />
                                Thanh toán khi nhận hàng
                            </button>
                            <button
                                type="button"
                                class="text-gray-900 bg-white hover:bg-gray-100 border border-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center me-2 mb-2"
                            >
                                <Icon icon="mdi:bank-outline" class="size-5 mr-2" />
                                Chuyển khoản
                            </button>
                        </div>
                    </div>
                    <div class="mt-8">
                        <div class="py-6 flex flex-col gap-y-2.5 text-sm border-y">
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
                        <button class="bg-primary w-full py-2.5 rounded-lg mt-8">
                            <span class="text-sm font-semibold">Hoàn tất đơn hàng</span>
                        </button>
                    </div>
                </div>
                <div class="col-span-5">
                    <h3 class="font-semibold my-4 pt-10">Sản phẩm</h3>
                    <div>
                        <div class="my-4 flex justify-between">
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
                        <div class="my-4 flex justify-between">
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
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
