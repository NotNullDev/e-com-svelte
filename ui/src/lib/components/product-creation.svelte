<script lang="ts">
    import {goto} from "$app/navigation";
    import {browser} from "$app/environment";

    export let productId: number | undefined = undefined;
    export let mode: "create" | "edit" = "create";
    export let productData = {
        productName: "",
        productDescription: "",
        productPrice: 0,
        productStock: "",
        productPreviewUrl: "https://picsum.photos/200/300",
    }
    export let imagesUrls: string[] = [];

    let filesList: FileList;
    let files: File[] = [];

    $: if (browser && filesList) {
        files = Array.from(filesList);
    }
    $: if (browser && imagesUrls) {
        processUrls(imagesUrls)
    }

    async function processUrls(urls: string[]) {
        for (const url of urls) {
            try {
                console.log('fetching', url);
                const response = await fetch(url);
                const blob = await response.blob();
                const file = new File([blob], "image.png", {type: blob.type});
                files.push(file);
            } catch (e) {
                console.error(e)
            }
        }
        files = files
    }

    function textAreaAutoResize(element: HTMLTextAreaElement) {
        element.style.boxSizing = 'border-box';
        const offset = element.offsetHeight - element.clientHeight;
        element.addEventListener('input', function (event) {
            const target = event?.target as HTMLElement
            if (!target) return;
            target.style.height = 'auto';
            target.style.height = target.scrollHeight + offset + 'px';
        });
    }

    async function createProduct() {
        const form = new FormData();

        form.append("productName", productData.productName);
        form.append("productDescription", productData.productDescription);
        form.append("productPrice", productData.productPrice.toString());
        form.append("productStock", productData.productStock);
        form.append("productPreviewUrl", productData.productPreviewUrl);
        for (const f of files) {
            form.append("files", f);
        }
        await fetch("http://localhost:8080/products", {
            method: "POST",
            body: form
        })
        await goto("/products")
    }

    async function updateProduct() {
        const form = new FormData();

        if (!productId) {
            throw new Error("Product id is not defined")
        }

        form.append("productId", productId.toString());
        form.append("productName", productData.productName);
        form.append("productDescription", productData.productDescription);
        form.append("productPrice", productData.productPrice.toString());
        form.append("productStock", productData.productStock);
        form.append("productPreviewUrl", productData.productPreviewUrl);
        for (const f of files) {
            form.append("files", f);
        }
        await fetch("http://localhost:8080/products", {
            method: "PUT",
            body: form
        })
        await goto("/products")
    }

    let selectedFile: File | null = null;
</script>

<div class="mx-auto container relative">
    <h1 class="my-5 text-3xl">Product {mode === "create" ? "creation" : "edition"}</h1>

    <form class="flex flex-col gap-2" on:submit|preventDefault>
        <div class="form-control">
            <label class="label" for="product-name">
                <span class="label-text">Product name</span>
            </label>
            <input bind:value={productData.productName} id="product-name" class="input input-bordered !h-12"
                   placeholder="Product name"/>
        </div>

        <div class="w-full border p-4 border-base-200 flex flex-wrap gap-5">
            {#if files}
                {#each files as f}
                    <div class="relative">
                        <img class="w-[250px] h-[250px] rounded-md" src={URL.createObjectURL(f)}
                             alt="uploaded-file-preview"/>
                        <div class="absolute top-0 right-0">
                            <input type="checkbox" class="checkbox m-4 shadow-xl shadow-primary" readonly value={false}
                                   checked={false}/>
                        </div>
                    </div>
                {/each}
            {/if}
            <label for="file-input"
                   class="w-[250px] h-[250px] flex items-center justify-center bg-base-200 rounded-md cursor-pointer select-none">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none"
                     stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                     class="lucide lucide-plus">
                    <path d="M5 12h14"/>
                    <path d="M12 5v14"/>
                </svg>
            </label>
            <input multiple hidden type="file" bind:files={filesList} id="file-input"/>
        </div>

        <div class="flex gap-2">
            <div class="form-control">
                <label class="label" for="product-price">
                    <span class="label-text">Product price</span>
                </label>
                <input bind:value={productData.productPrice} id="product-price" type="number" placeholder="Price"
                       class="input input-bordered w-[150px]"/>
            </div>

            <div class="form-control">
                <label class="label" for="product-stock">
                    <span class="label-text">Product stock</span>
                </label>
                <input bind:value={productData.productStock} id="product-stock" type="number" placeholder="Stock"
                       class="input input-bordered w-[150px]"/>
            </div>
        </div>

        <div class="form-control">
            <label class="label" for="product-description">
                <span class="label-text">Product description</span>
            </label>

            <textarea
                    id="product-description"
                    placeholder="Product description"
                    bind:value={productData.productDescription}
                    use:textAreaAutoResize
                    class="textarea textarea-bordered"
            />
        </div>

        <button on:click={mode === "create" ? createProduct : updateProduct} class="btn btn-primary mt-2"
                type="submit">{mode === "create" ? "Create" : "Update"}</button>
    </form>

    <!--    <div class="absolute bottom-0 left-1/2 mb-20 -translate-x-1/2 rounded-md p-4">-->
    <!--        <button class="btn !btn-lg btn-primary" >Delete</button>-->
    <!--        <button class="btn !btn-lg btn-primary" >Make primary</button>-->
    <!--    </div>-->

</div>
