<script lang="ts">
    import {appStore, type Cart} from "$lib/appStore";

    function groupBySeller(products: Cart['products']): Map<string, Cart['products']> {
        const sellers = new Map<string, Cart['products']>();

        for (const p of products) {
            let sellerProducts = sellers.get(p.product.seller);
            if (!sellerProducts) {
                sellerProducts = [];
            }
            sellerProducts.push(p);
            sellers.set(p.product.seller, sellerProducts);
        }

        return sellers;
    }

    $: grouped = groupBySeller($appStore.cart.products);

</script>

<div class="flex flex-col flex-1 mx-24">
    <h1 class="text-3xl my-5">Cart</h1>
    <div class="flex flex-1">
        <div class="flex flex-col flex-1 items-center">
            {#each grouped as [seller, val], idx}
                <div class="p-4 bg-base-200 rounded-md w-[1000px]">
                    <h1 class="text-2xl">{seller}</h1>
                    <div class="flex flex-col gap-3">
                        {#each val as p}
                            <div class="flex gap-5">
                                <img src={p.product?.preview_url} class="w-[150px] h-[150px]" alt="product preview"/>
                                <div class="flex flex-col gap-2 w-full">
                                    <h2 class="text-2xl">{p.product.name}</h2>
                                    <div class="flex justify-end w-full gap-10">
                                        <div>
                                            <button class="btn btn-square btn-ghost" on:click={() => {p.quantity -= 1}}>-
                                            </button>
                                            <input  type="number" bind:value={p.quantity} class="input w-32">
                                            <button class="btn btn-square btn-ghost" on:click={() => {p.quantity += 1}}>+
                                            </button>
                                        </div>
                                        <div class="text-2xl w-[100px]">{p.product.price * p.quantity} PLN</div>
                                    </div>
                                </div>
                            </div>
                        {/each}
                    </div>
                    <hr class="my-3" />
                    <div class="flex justify-end">
                        <div class="text-2xl w-[100px]">{val.map(v => v.product.price * v.quantity).reduce((p,n) => (p + n), 0)} PLN</div>
                    </div>
                </div>
            {/each}
        </div>
        <div class="bg-base-200 p-8 rounded-md w-[500px]">
            <h2 class="text-2xl">Summary</h2>
            <hr class="my-3" />
            <div class="flex justify-between">
                <div class="text-2xl">Total</div>
                <div class="text-2xl">{($appStore.cart.products.map(p => p.product.price * p.quantity).reduce((p,n) => (p + n), 0)).toFixed(2)} PLN</div>
            </div>
        </div>
    </div>
</div>
