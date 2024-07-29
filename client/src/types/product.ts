export type ProductProps = {
    id: number
    name: string
    price: number
    priceOrigin?: number
    thumbnail: string
    sales?: number
    star?: string
    discount?: number
    liked: boolean
    slug: string
}

export type InformationProductProps = {
    countComment: number
    quantity: number
} & Omit<ProductProps, 'id' | 'like' | 'slug'>
