<template>
    <div class="p-2">
        <h3 class="text-xl font-semibold">Thông tin tài khoản</h3>
        <div class="py-4 grid grid-cols-2 gap-x-8 gap-y-2">
            <div
                class="col-span-2 bg-gray-50 rounded-xl flex items-center px-6 p-4 justify-between mb-4"
            >
                <div>
                    <img
                        class="size-14 object-cover rounded-full"
                        src="http://localhost:5173/src/assets/images/logo.png"
                        alt=""
                    />
                </div>
                <div>
                    <button class="py-1.5 bg-blue-600 text-white rounded-lg px-3">
                        <span class="text-sm font-semibold">Đổi ảnh đại diện</span>
                    </button>
                </div>
            </div>
            <div class="text-sm">
                <p class="py-2">Họ và Tên</p>
                <input type="text" class="border p-2 rounded-lg outline-none w-full" />
            </div>
            <div class="text-sm">
                <p class="py-2">Tên tài khoản</p>
                <input type="text" class="border p-2 rounded-lg outline-none w-full" />
            </div>
            <div class="text-sm">
                <p class="py-2">Số điện thoại</p>
                <input type="text" class="border p-2 rounded-lg outline-none w-full" />
            </div>
            <div class="text-sm">
                <p class="py-2">Ngày sinh</p>
                <input type="date" class="border p-2 rounded-lg outline-none w-full" />
            </div>
            <div class="text-sm col-span-2">
                <p class="py-2">Địa chỉ</p>
                <input type="text" class="border p-2 rounded-lg outline-none w-full" />
            </div>
            <div class="col-span-2 mt-6">
                <button class="py-1.5 bg-primary rounded-lg px-3 ">
                    <span class="text-sm font-semibold">Xác nhận</span>
                </button>
            </div>
        </div>
    </div>
</template>
