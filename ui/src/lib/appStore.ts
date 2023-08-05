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
}

export const appStore = writable<AppStoreType>({
    products: [],
    users: []
})
