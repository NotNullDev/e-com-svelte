import type { PageLoad } from "./$types";
import type {Product} from "$lib/appStore";
import {redirect} from "@sveltejs/kit";

export const load: PageLoad = async ({url}) => {
    let product: Product | null = null
    const id = url.searchParams.get("id")

    try {
        product = await (await fetch(`http://localhost:8080/products/${id}`)).json() as Product;
    } catch (e) {
        console.error(e);
        throw redirect(301, "/account/products")
    }

    return {
        product
    }
}
