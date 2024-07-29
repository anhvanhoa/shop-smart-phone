<script setup lang="ts">
import { onUnmounted, reactive, computed } from 'vue'
const countDown = reactive({
    hours: 1,
    minutes: 58,
    seconds: 0
})
const timerId = setInterval(() => {
    if (countDown.hours === 0 && countDown.minutes === 0 && countDown.seconds === 0) {
        console.log('Time is up')
    } else {
        if (countDown.seconds == 0) {
            if (countDown.minutes > -1) {
                countDown.minutes -= 1
                countDown.seconds = 59
            }
            if (countDown.minutes == -1) {
                if (countDown.hours > 0) {
                    countDown.hours -= 1
                    countDown.minutes = 59
                }
            }
        } else {
            --countDown.seconds
        }
    }
}, 1000)

const time = computed(() => {
    return {
        hours: String(countDown.hours).padStart(2, '0'),
        minutes: String(countDown.minutes).padStart(2, '0'),
        seconds: String(countDown.seconds).padStart(2, '0')
    }
})

onUnmounted(() => {
    clearInterval(timerId)
})
</script>
<template>
    <div
        class="flex items-center gap-x-3 *:h-8 *:w-9 *:rounded-lg *:bg-black/80 *:text-white *:font-bold *:flex *:items-center *:justify-center"
    >
        <div>{{ time.hours }}</div>
        <div>{{ time.minutes }}</div>
        <div>{{ time.seconds }}</div>
    </div>
</template>
