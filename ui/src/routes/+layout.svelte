<script lang="ts">
	import { browser } from '$app/environment';
	import { fly } from 'svelte/transition';
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
</script>

<div class="h-screen overflow-hidden bg-base-300">
	<header class="flex flex-col">
		<div class="flex justify-between px-6 p-4 container mx-auto shadow shadow-base-200">
			<a class="btn btn-ghost btn-sm" href="/">E-com</a>
			<div class="join">
				<div>
					<div>
						<input type="search" bind:value={search} class="input input-bordered join-item" placeholder="Search for anything" />
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
			<div class="cursor-pointer relative">
				<button
					class="avatar placeholder"
					on:click={async () => {
						showUserInfo = !showUserInfo;
					}}
				>
					<div class="bg-neutral-focus text-neutral-content rounded-full w-8">
						<span class="text-sm">K</span>
					</div>
				</button>
				{#if showUserInfo}
					<div
						bind:this={modalRef}
						transition:fly
						class="right-0 top-[100%] absolute bg-base-100 mt-2 px-8 mr-2 py-4 rounded-md flex flex-col gap-2 items-center"
					>
						<a href="/login" class="btn btn-sm btn-ghost">login</a>

						<button
							on:click={() => {
								theme = theme === 'dark' ? 'light' : 'dark';
							}}
							class="btn btn-square btn-sm btn-ghost"
						>
							{#if theme === 'light'}
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
									><circle cx="12" cy="12" r="4" /><path d="M12 8a2 2 0 1 0 4 4" /><path
										d="M12 2v2"
									/><path d="M12 20v2" /><path d="m4.93 4.93 1.41 1.41" /><path
										d="m17.66 17.66 1.41 1.41"
									/><path d="M2 12h2" /><path d="M20 12h2" /><path d="m6.34 17.66-1.41 1.41" /><path
										d="m19.07 4.93-1.41 1.41"
									/></svg
								>
							{:else}
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
									class="lucide lucide-moon"><path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z" /></svg
								>
							{/if}
						</button>
					</div>
				{/if}
			</div>
		</div>
		<div class="flex justify-between px-6 p-4 container mx-auto shadow shadow-base-200">
			<div class="relative">
				<button class="btn btn-ghost" on:click={() => (showCategories = !showCategories)}
					>Categories</button
				>
				{#if showCategories}
					<div class="absolute top-[100%] left-0 mt-2 bg-base-100 p-6 rounded">
						<div class="w-48">TODO list</div>
					</div>
				{/if}
			</div>
		</div>
	</header>
	<slot />
</div>
