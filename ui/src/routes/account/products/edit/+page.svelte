<script lang="ts">
    import type {PageData} from "./$types"
    import ProductCreation from "$lib/components/product-creation.svelte"

    export let data: PageData
</script>

<ProductCreation productData={{
    productPrice: data.product.price,
    productName: data.product.name,
    productDescription: data.product.description,
    productStock: data.product.stock,
    productPreviewUrl: data.product.preview_url,
}}
                 productId={data.product.id}
                 mode="edit"
                 imagesUrls={data.product.images.map(image => `http://localhost:8080/static/${image}`)}
/>
