<script lang="ts">
  import {browser} from '$app/environment';
  import {page} from "$app/stores";
  import '../app.css';
  import {onMount} from "svelte";
  import {appStore} from "$lib/appStore";

  let showUserInfo = false;
  let showCategories = false;

  let modalRef: HTMLDivElement;

  let theme: 'light' | 'dark' = 'dark';

  $: if (browser) {
    if (theme === 'light') {
      document.documentElement.setAttribute('data-theme', 'light');
    } else {
      document.documentElement.setAttribute('data-theme', 'dark');
    }
  }

  onMount(async () => {
    fetch("http://localhost:8080/products").then(r => r.json()).then(r => $appStore.products = r);
    fetch("http://localhost:8080/users").then(r => r.json()).then(r => $appStore.users = r);
  })

  let search = "";
  let searchCategory = "All categories";
  let searchParams: URLSearchParams = new URLSearchParams();

  $: {
    if (search && search?.trim() !== "") {
      searchParams.set("query", search);
    } else {
      searchParams.delete("query");
    }

    if (searchCategory && searchCategory !== "All categories") {
      searchParams.set("category", searchCategory);
    } else {
      searchParams.delete("category");
    }

    searchParams = searchParams
  }

  $: breaderCrumbs = $page.url.pathname.split("/").filter(p => p !== "");
</script>

<div class="h-screen bg-base-300 flex flex-col">
  <header class="navbar shadow flex flex-col p-4">
    <div class="container mx-auto flex justify-between items-center">
      <div class="flex gap-4 items-center">
        <a href="/" class="btn btn-ghost text-xl">E-COM</a>
        <div class="text-sm breadcrumbs">
          <ul>
            <li><a href="/">Home</a></li>
            {#each breaderCrumbs as crumb}
              <li><a href="/{crumb}">{crumb}</a></li>
            {/each}
          </ul>
        </div>
      </div>
      <div class="join">
        <div>
          <div>
            <input type="search" bind:value={search} class="input input-bordered join-item"
                   placeholder="Search for anything"/>
          </div>
        </div>
        <select bind:value={searchCategory} class="select select-bordered join-item">
          <option selected>All categories</option>
          <option>Sci-fi</option>
          <option>Drama</option>
          <option>Action</option>
        </select>
        <a href={`/products?${searchParams.toString()}`} class="btn join-item">Search</a>
      </div>
      <div class="flex gap-4">
        <div class="dropdown dropdown-end">
          <button tabindex="0" class="btn btn-ghost btn-circle">
            <div class="indicator">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24"
                   stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"/>
              </svg>
              <span class="badge badge-sm indicator-item">8</span>
            </div>
          </button>
          <button tabindex="0" class="mt-3 z-[1] card card-compact dropdown-content w-52 bg-base-100 shadow">
            <div class="card-body">
              <span class="font-bold text-lg">8 Items</span>
              <span class="text-info">Subtotal: $999</span>
              <div class="card-actions">
                <button class="btn btn-primary btn-block">View cart</button>
              </div>
            </div>
          </button>
        </div>
        <div class="dropdown dropdown-end">
          <button tabindex="0" class="btn btn-ghost btn-circle avatar placeholder">
            <div class="bg-neutral-focus text-neutral-content rounded-full w-8">
              <span class="text-sm">K</span>
            </div>
          </button>
          <button tabindex="0">
            <ul class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52">
              <li><a href="/account/products">My prodcuts</a></li>
              <li><a href="/">Logout</a></li>
              <li>
                <button
                  on:click={() => {
								theme = theme === 'dark' ? 'light' : 'dark';
							}}
                  class=""
                >
                  {#if theme === 'light'}
                    Dark theme
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="24"
                      height="24"
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      class="lucide lucide-moon">
                      <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z"/>
                    </svg
                    >
                  {:else}
                    Light theme
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="24"
                      height="24"
                      viewBox="0 0 24 24"
                      fill="none"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      class="lucide lucide-sun-moon"
                    >
                      <circle cx="12" cy="12" r="4"/>
                      <path d="M12 8a2 2 0 1 0 4 4"/>
                      <path
                        d="M12 2v2"
                      />
                      <path d="M12 20v2"/>
                      <path d="m4.93 4.93 1.41 1.41"/>
                      <path
                        d="m17.66 17.66 1.41 1.41"
                      />
                      <path d="M2 12h2"/>
                      <path d="M20 12h2"/>
                      <path d="m6.34 17.66-1.41 1.41"/>
                      <path
                        d="m19.07 4.93-1.41 1.41"
                      />
                    </svg
                    >
                  {/if}
                </button>
              </li>
            </ul>
          </button>
        </div>
      </div>
    </div>

    <div class="flex justify-between px-6 p-4 w-full">
      <div class="relative container mx-auto">
        <button class="btn btn-ghost" on:click={() => (showCategories = !showCategories)}>Categories</button>
        {#if showCategories}
          <div class="absolute top-[100%] left-0 mt-2 bg-base-100 p-6 rounded">
            <div class="w-48">TODO list</div>
          </div>
        {/if}
      </div>
    </div>

  </header>

  <slot/>

  <footer class="p-6">
    <div>footer todo</div>
  </footer>
</div>
