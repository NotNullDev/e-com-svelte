import type { PageLoad } from "./$types";
import type {Product} from "$lib/appStore";

export const load: PageLoad = async ({url, fetch}) => {
  const id = url.searchParams.get("id");

  const product = await (await fetch(`http://localhost:8080/products/${id}`)).json() as Product

  return {
    product
  }
}
