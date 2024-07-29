import { computed } from "vue"

export const priceFormat = computed(() => {
    return (price: number) => {
        return new Intl.NumberFormat('vi-VN', {
            style: 'currency',
            currency: 'VND'
        }).format(price)
    }
})
