<script lang="ts">
  import type {PageData} from "./$types";
  import ProductCard from "./product-card.svelte";

  export let data: PageData;

  $: console.log(data);
</script>

<div class="flex flex-col gap-2 flex-wrap container mx-auto">
  <div class="my-4">
    <h1 class="text-3xl font-bold">Matches for {data.query}</h1>
  </div>
  {#each data?.products ?? [] as item}
    <ProductCard product={item}/>
  {/each}
</div>
