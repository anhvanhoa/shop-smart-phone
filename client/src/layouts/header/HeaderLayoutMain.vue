<script setup lang="ts">
import LogoLayout from './LogoLayout.vue'
import ButtonCatgeory from './ButtonCatgeory.vue'
import InputSearch from './InputSearch.vue'
import RightHeaderMain from './RightHeaderMain.vue'
import AccountHeader from './AccountHeader.vue'
</script>
<template>
    <section class="border-b border-gray-200/60 bg-white">
        <header class="max-w-screen-xl mx-auto px-4 py-5 flex items-center gap-x-14">
            <!-- logo -->
            <LogoLayout  />
            <!-- categories -->
            <ButtonCatgeory />
            <!-- search -->
            <InputSearch />
            <!-- right header -->
            <RightHeaderMain>
                <AccountHeader />
            </RightHeaderMain>
        </header>
        
    </section>
</template>
