<script lang="ts">
	let { isVisible, message, onDismiss } = $props();
</script>

{#if isVisible}
	<script lang="ts">
	import { TriangleAlertIcon } from '@lucide/svelte';
</script>

<div class="w-full space-y-8">
	<div class="card preset-outlined-error-500 grid grid-cols-1 items-center gap-4 p-4 lg:grid-cols-[auto_1fr_auto]">
		<div>
			<p class="font-bold">Error</p>
			<p class="text-xs opacity-60">{message}</p>
		</div>
		<div class="flex gap-1 justify-end">
			<button class="btn preset-tonal hover:preset-filled" onclick={()=> onDismiss()}>Dismiss</button>
		</div>
	</div>
</div>
{/if}
