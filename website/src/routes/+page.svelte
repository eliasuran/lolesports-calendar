<script lang="ts">
	import { authURL } from '$lib/auth.js';

	let token = '';
</script>

{#if !token}
	<button
		on:click={async () => {
			const url = await authURL();
			window.location.href = url;
		}}>Login</button
	>
{:else}
	<div class="space-y-12">
		<h1 class="text-4xl">hello world</h1>
		<div class="flex gap-12 flex-wrap">
			<div
				class="h-52 aspect-square bg-primary border border-text rounded-md grid place-items-center font-bold text-2xl"
			>
				xdd
			</div>
		</div>
	</div>
{/if}
