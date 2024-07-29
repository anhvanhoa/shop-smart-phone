<script setup lang="ts">
import { ref } from 'vue'
import { computed } from 'vue'

const props = defineProps<{
    data: {
        id: number
        title: string
        value: string
    }[]
}>()
const MAX_ITEM = 8
const itemShow = ref(MAX_ITEM)
const showBtn = computed(() => {
    if(props.data.length === itemShow.value) return false
    return props.data.length > MAX_ITEM
})
const handleShowAll = () => {
    itemShow.value = props.data.length
}
</script>
<template>
    <div>
        <div
            v-for="(item, i) in props.data"
            :key="item.id"
            v-show="i <= itemShow"
            class="flex items-center px-1 gap-x-1.5 my-2.5"
        >
            <input
                type="checkbox"
                :id="`${item.value}-${item.id}`"
                :name="item.title"
                :value="item.value"
            />
            <label :for="`${item.value}-${item.id}`" class="text-sm cursor-pointer">{{
                item.title
            }}</label>
        </div>
        <button @click="handleShowAll" v-if="showBtn" class="hover:text-sky-600 text-xs px-1 text-cente font-semibold">Xem tất cả</button>
    </div>
</template>
