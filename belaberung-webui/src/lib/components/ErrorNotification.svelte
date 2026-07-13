<script lang="ts">
	let { isVisible, message, onDismiss } = $props();
</script>

{#if isVisible}

<div>
	<p>{message}</p>
	<button onclick={()=> onDismiss()}>Dismiss</button>
</div>
{/if}
