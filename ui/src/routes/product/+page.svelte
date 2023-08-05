<script lang="ts">
  import type { PageData } from "./$types"

  export let data: PageData;
  let amount = 1;
</script>

{#if data.product}

  <div class="mx-auto container">

    <div class="flex justify-between p-4 gap-2">

      <div class="bg-base-200 rounded-md flex-1 flex flex-col p-8 gap-4">
        <div>{data.product.name}</div>
        <div class="flex gap-2">
          <img class="w-[200px]" src={data.product.preview_url} alt="aa" />
          {#each data.product.images as imgSrc}
            <img class="w-[200px]" src={`http://localhost:8080/static/${imgSrc}`} alt="aa" />
          {/each}
        </div>
      </div>

      <div class="bg-base-200 rounded-md flex p-6 flex flex-col justify-between w-[400px]">
        <div class="flex flex-col gap-2">
          <div class="text-2xl">{data.product.name}</div>
          <div>{data.product.price} PLN</div>
        </div>
        <div class="flex justify-between">
          <label>
            <input type="number" class="input input-bordered w-[100px]" bind:value={amount} />
          </label>
          <button class="btn btn-primary">Add to cart</button>
        </div>
      </div>

    </div>

    <div class="bg-base-200 m-4 p-4 rounded-md">
      <div>Real description starts here</div>
      <div class="p-4">{data.product.description}</div>
    </div>

  </div>
{/if}
