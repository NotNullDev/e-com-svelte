<script lang="ts">
  import type {Product} from "$lib/appStore";

  export let product: Product;
  export const href = `/product?id=${product.id}`
</script>

<div class="bg-base-200 rounded-md h-48 flex">
  <img src={product.preview_url} class="w-[200px]" width="200" height="200" alt="preview" />
  <div class="p-4 flex w-full">
    <a class="flex flex-col flex-1" href={href}>
      <div class="text-xl">{product.name}</div>
      <div class="text-2xl font-bold">{product.price} PLN</div>
      <div>{product.stock} items in stock</div>
      <div>{product.seller}</div>
    </a>

    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-heart"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>

  </div>
</div>
