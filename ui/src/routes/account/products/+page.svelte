<script lang="ts">
  import type {PageData} from "./$types";
  import AdminProductCard from "./admin-product-card.svelte"

  export let data: PageData
  console.log(data)
</script>

<div class="container mx-auto">
  <div class="flex justify-end p-4">
    <a href="/account/products/create" class="btn btn-primary">ADD PRODUCT</a>
  </div>
  <div class="container mx-auto flex flex-col gap-2">
    {#each data.myProducts as p}
      <AdminProductCard product={p} />
    {/each}
    {#if data.myProducts.length === 0}
      <div>No products found</div>
    {/if}
  </div>
  <div>Hah</div>
</div>
