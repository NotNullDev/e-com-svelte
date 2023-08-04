import type {PageLoad} from "../../../.svelte-kit/types/src/routes/products/$types";
import type {Product} from "$lib/appStore";

export const load: PageLoad = async ({url, fetch}) => {
  const query = url.searchParams.get("query");
  const category = url.searchParams.get("category");

  let products: Product[] = [];
  try {
    products = await (await fetch("http://localhost:8080/products/")).json() as Product[]
  } catch(e) {
    console.error(e);
    return {}
  }

  return {
    products,
    query
  }
}
