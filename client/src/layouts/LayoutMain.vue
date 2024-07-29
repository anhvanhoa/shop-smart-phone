<script setup lang="ts">
import { Footer } from './footer'
import { HeaderLayoutMain } from './header'
</script>
<template>
    <div class="font-apple">
        <HeaderLayoutMain />
        <RouterView />
        <Footer></Footer>
    </div>
</template>
