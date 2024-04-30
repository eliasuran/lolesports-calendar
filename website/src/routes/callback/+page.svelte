<script lang="ts">
	import type { PageData } from './$types.js';

	export let data: PageData;
</script>

{#if data.error}
	<h1 class="text-xl">{data.error}</h1>
	<a href="/">Back</a>
{/if}
