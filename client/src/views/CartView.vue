<script setup lang="ts">
import { Icon } from '@iconify/vue/dist/iconify.js'
import { RouterLink } from 'vue-router'
</script>

<template>
    <div>
        <div class="max-w-screen-xl mx-auto px-8">
            <h1 class="text-2xl font-semibold py-8">Giỏ hàng</h1>
            <div class="grid grid-cols-12 items-start mb-8 gap-x-8">
                <div class="col-span-8">
                    <div class="flex justify-between border p-4 rounded-xl gap-x-32 mb-4">
                        <div class="flex items-center gap-x-2 text-sm">
                            <input type="checkbox" name="" id="" />
                            <span> Chọn tất cả </span>
                        </div>
                        <div>
                            <button
                                class="px-3 py-1 rounded-full text-rose-500 text-sm font-semibold"
                            >
                                Xóa
                            </button>
                        </div>
                    </div>
                    <div class="flex justify-between border p-4 rounded-xl gap-x-32 mb-4">
                        <div class="flex gap-x-6">
                            <label class="flex items-center gap-x-4">
                                <input type="checkbox" />
                                <img
                                    src="http://localhost:5173/src/assets/images/product.png"
                                    alt=""
                                    class="size-24 bg-gray-200/80 rounded-xl"
                                />
                            </label>
                            <div>
                                <RouterLink to="/product/macbook-air-13-inch-m2-chinh-hang">
                                    <h3
                                        class="py-1 font-semibold line-clamp-2 hover:underline text-pretty"
                                    >
                                        MacBook Air 13 inch M2 | Chính hãng
                                    </h3>
                                </RouterLink>
                                <p class="text-sm">Phân loại: Ram 32GB, Màu trắng</p>
                            </div>
                        </div>
                        <div class="ml-12 flex flex-col items-end">
                            <p class="text-gray-500 line-through text-sm">19.000.000đ</p>
                            <p class="text-rose-500 text-lg font-semibold">19.000.000đ</p>
                            <div
                                class="mt-2 inline-flex items-center text-sm border gap-x-2 p-1 rounded-lg w-min"
                            >
                                <span class="p-1 hover:bg-gray-100 rounded-md">
                                    <Icon icon="ph:minus" />
                                </span>
                                <span class="w-6 text-center overflow-hidden">
                                    <input
                                        type="text"
                                        class="w-full outline-none text-center"
                                        value="1"
                                    />
                                </span>
                                <span class="p-1 hover:bg-gray-100 rounded-md">
                                    <Icon icon="ph:plus" />
                                </span>
                            </div>
                        </div>
                    </div>
                    <div class="flex justify-between border p-4 rounded-xl gap-x-32 mb-4">
                        <div class="flex gap-x-6">
                            <label class="flex items-center gap-x-4">
                                <input type="checkbox" />
                                <img
                                    src="http://localhost:5173/src/assets/images/product.png"
                                    alt=""
                                    class="size-24 bg-gray-200/80 rounded-xl"
                                />
                            </label>
                            <div>
                                <RouterLink to="/product/macbook-air-13-inch-m2-chinh-hang">
                                    <h3
                                        class="py-1 font-semibold line-clamp-2 hover:underline text-pretty"
                                    >
                                        MacBook Air 13 inch M2 | Chính hãng
                                    </h3>
                                </RouterLink>
                                <p class="text-sm">Phân loại: Ram 32GB, Màu trắng</p>
                            </div>
                        </div>
                        <div class="ml-12 flex flex-col items-end">
                            <p class="text-gray-500 line-through text-sm">19.000.000đ</p>
                            <p class="text-rose-500 text-lg font-semibold">19.000.000đ</p>
                            <div
                                class="mt-2 inline-flex items-center text-sm border gap-x-2 p-1 rounded-lg w-min"
                            >
                                <span class="p-1 hover:bg-gray-100 rounded-md">
                                    <Icon icon="ph:minus" />
                                </span>
                                <span class="w-6 text-center overflow-hidden">
                                    <input
                                        type="text"
                                        class="w-full outline-none text-center"
                                        value="1"
                                    />
                                </span>
                                <span class="p-1 hover:bg-gray-100 rounded-md">
                                    <Icon icon="ph:plus" />
                                </span>
                            </div>
                        </div>
                    </div>
                    <div class="flex justify-between border p-4 rounded-xl gap-x-32 mb-4">
                        <div class="flex gap-x-6">
                            <label class="flex items-center gap-x-4">
                                <input type="checkbox" />
                                <img
                                    src="http://localhost:5173/src/assets/images/product.png"
                                    alt=""
                                    class="size-24 bg-gray-200/80 rounded-xl"
                                />
                            </label>
                            <div>
                                <RouterLink to="/product/macbook-air-13-inch-m2-chinh-hang">
                                    <h3
                                        class="py-1 font-semibold line-clamp-2 hover:underline text-pretty"
                                    >
                                        MacBook Air 13 inch M2 | Chính hãng
                                    </h3>
                                </RouterLink>
                                <p class="text-sm">Phân loại: Ram 32GB, Màu trắng</p>
                            </div>
                        </div>
                        <div class="ml-12 flex flex-col items-end">
                            <p class="text-gray-500 line-through text-sm">19.000.000đ</p>
                            <p class="text-rose-500 text-lg font-semibold">19.000.000đ</p>
                            <div
                                class="mt-2 inline-flex items-center text-sm border gap-x-2 p-1 rounded-lg w-min"
                            >
                                <span class="p-1 hover:bg-gray-100 rounded-md">
                                    <Icon icon="ph:minus" />
                                </span>
                                <span class="w-6 text-center overflow-hidden">
                                    <input
                                        type="text"
                                        class="w-full outline-none text-center"
                                        value="1"
                                    />
                                </span>
                                <span class="p-1 hover:bg-gray-100 rounded-md">
                                    <Icon icon="ph:plus" />
                                </span>
                            </div>
                        </div>
                    </div>
                    <div class="flex justify-between border p-4 rounded-xl gap-x-32 mb-4">
                        <div class="flex gap-x-6">
                            <label class="flex items-center gap-x-4">
                                <input type="checkbox" />
                                <img
                                    src="http://localhost:5173/src/assets/images/product.png"
                                    alt=""
                                    class="size-24 bg-gray-200/80 rounded-xl"
                                />
                            </label>
                            <div>
                                <RouterLink to="/product/macbook-air-13-inch-m2-chinh-hang">
                                    <h3
                                        class="py-1 font-semibold line-clamp-2 hover:underline text-pretty"
                                    >
                                        MacBook Air 13 inch M2 | Chính hãng
                                    </h3>
                                </RouterLink>
                                <p class="text-sm">Phân loại: Ram 32GB, Màu trắng</p>
                            </div>
                        </div>
                        <div class="ml-12 flex flex-col items-end">
                            <p class="text-gray-500 line-through text-sm">19.000.000đ</p>
                            <p class="text-rose-500 text-lg font-semibold">19.000.000đ</p>
                            <div
                                class="mt-2 inline-flex items-center text-sm border gap-x-2 p-1 rounded-lg w-min"
                            >
                                <span class="p-1 hover:bg-gray-100 rounded-md">
                                    <Icon icon="ph:minus" />
                                </span>
                                <span class="w-6 text-center overflow-hidden">
                                    <input
                                        type="text"
                                        class="w-full outline-none text-center"
                                        value="1"
                                    />
                                </span>
                                <span class="p-1 hover:bg-gray-100 rounded-md">
                                    <Icon icon="ph:plus" />
                                </span>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-span-4 border rounded-xl p-6 sticky top-8">
                    <div>
                        <div class="flex items-center border rounded-xl p-1">
                            <input
                                type="text"
                                class="flex-1 text-sm py-2 outline-none px-2 font-semibold placeholder:font-normal"
                                placeholder="Mã giảm giá"
                            />
                            <button
                                class="bg-slate-950 px-3 py-2 rounded-lg text-white text-sm font-semibold"
                            >
                                Áp dụng
                            </button>
                        </div>
                    </div>
                    <div class="mt-8 px-3 flex flex-col gap-4 text-sm">
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
                    <div class="px-3 mt-6">
                        <button
                            class="bg-primary text-sm font-semibold px-4 py-2.5 rounded-lg w-full"
                        >
                            Thanh toán
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
