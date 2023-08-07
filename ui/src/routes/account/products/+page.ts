import type {PageLoad} from "./$types";
import type {Product} from "$lib/appStore";

export const load: PageLoad = async ({fetch}) => {
  let myProducts: Product[] = [];

  try {
    myProducts = await (await fetch("http://localhost:8080/products/my")).json() as Product[];
  } catch (e) {
    console.error(e);
  }

  console.log(myProducts)

  return {
    myProducts
  }
}
