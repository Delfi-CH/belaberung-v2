<script lang="ts">
	import type { Message } from '$lib/api/message';

	let { messages } = $props();

	function determineCSSClass(message: Message) {
		switch (message.role) {
			case 'Administrator':
				return 'text-danger';
			case 'Moderator':
				return 'text-warning';
			default:
				return '';
		}
	}
</script>

<div class="view">
	{#each messages as message, index (index)}
		<p>
			<span class={determineCSSClass(message)}>{message.username}</span>
			<span class="date">sent at {message.timestamp.toLocaleString()}</span>
		</p>
		<p>{message.content}</p>
	{/each}
</div>

<style>
	.view {
		display: flex;
		flex-direction: column;
		height: 75vh;
		overflow: scroll;
	}

	.date {
		font-size: 65%;
	}
</style>
