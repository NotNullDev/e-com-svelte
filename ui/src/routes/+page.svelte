<script>
  import {appStore} from "$lib/appStore";
</script>

<div class="h-full flex flex-col items-center mt-40">
  <h1 class="text-3xl">Imagine nice recommended products web page</h1>
</div>

{#each $appStore.products as p}
  <div>
    <h1>{p.name}</h1>
    <p>{p.description}</p>
    <p>{p.price}</p>
  </div>
{/each}
