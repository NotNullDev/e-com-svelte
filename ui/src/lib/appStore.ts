import {writable} from "svelte/store";


export type Product = {
    id: number
    name: string
    price: number
    preview_url: string
    categories: string[]
    images: string[],
    description: string
    stock: string,
    stock_reserved: string,
    seller: string
}

export type User = {
    id: number
    email: string
}

export type AppStoreType = {
    products: Product[]
    users: User[]
    cart: Cart
    saved: Saved
}

export type Saved = {
    products: Product[]
}

export type Cart = {
    products: {
        product: Product,
        quantity: number
    }[]
}

export const appStore = writable<AppStoreType>({
    products: [],
    users: [],
    cart: {
        products: []
    },
    saved: {
        products: []
    }
})
