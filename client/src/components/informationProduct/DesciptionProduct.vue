<script setup lang="ts">
const description = `
        <p>
                Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ratione debitis voluptatum
                rem? Blanditiis odit dolore esse placeat beatae aperiam aliquid id, quisquam
                accusamus velit nisi itaque asperiores tempore? Rerum, tempore!
        </p>
        <img
            src="https://down-vn.img.susercontent.com/file/sg-11134301-7rcc9-ltm7wr6f2v3k36"
            alt=""
            srcset=""
            class="w-96 mx-auto my-6"
        />
        <p>
            Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ratione debitis voluptatum
            rem? Blanditiis odit dolore esse placeat beatae aperiam aliquid id, quisquam
            accusamus velit nisi itaque asperiores tempore? Rerum, tempore!
        </p>
        <img
            src="https://cdn.tgdd.vn/Products/Images/44/289472/apple-macbook-air-m2-2022-16gb-256gb-thumb-600x600.jpg"
            alt=""
            srcset=""
            class="w-96 mx-auto my-6"
        />
    `
</script>
<template>
    <div class="mt-6">
        <p class="font-extrabold text-lg">Thông tin sản phẩm</p>
        <div class="mt-2" v-html="description"></div>
    </div>
</template>
