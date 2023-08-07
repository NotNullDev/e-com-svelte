<script lang="ts">
  import type {PageData} from "./$types";
  import AdminProductCard from "./admin-product-card.svelte"

  export let data: PageData
  console.log(data)
</script>

<div class="container mx-auto flex-1 flex flex-col">
  <div class="flex justify-between p-4 my-3">
    <h2 class="text-3xl">My products</h2>
    <a href="/account/products/create" class="btn btn-primary">ADD PRODUCT</a>
  </div>
  <div class="container mx-auto flex flex-col gap-2 flex-1">
    {#each data.myProducts as p}
      <AdminProductCard product={p} />
    {/each}
    {#if data.myProducts.length === 0}
      <div class="flex flex-col gap-4 w-full mt-10 items-center flex-1">
        <div class="text-4xl">You don't have any products yet.</div>
        <a href="/account/products/create" class="btn !btn-lg btn-primary w-[200px]">CREATE NEW PRODUCT</a>
      </div>
    {/if}
  </div>
</div>
