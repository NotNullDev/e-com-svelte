import type {PageLoad} from "./$types";
import type {Product} from "$lib/appStore";

export const load: PageLoad = async ({url, fetch}) => {
  const id = url.searchParams.get("id");

  try {
    const resp = await fetch(`http://localhost:8080/products/${id}`)

    if (!resp.ok) {
      return {
        status: resp.status,
        product: null
      }
    }

    const product = await resp.json() as Product
    return {
      status: 200,
      product
    }
  } catch (e) {
    console.error(e);
  }

  return {
    status: 405,
    product: null
  }
}
