<script lang="ts">
	import type { Message } from 'belaberung-client-libs/message';
	import MiniProfile from './MiniProfile.svelte';

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

	let showMiniProfile = $state(false)
	let miniProfileUserID = $state(0)
	let miniProfileMessageIndex = $state(-1)
</script>

<div class="view">
	{#each messages as message, index (index)}
		<p>
			<span class={determineCSSClass(message)} onclick={()=>{
				showMiniProfile = !showMiniProfile
				miniProfileUserID = message.userId
				miniProfileMessageIndex = index
			}}
			onkeydown={()=>{
				 
			}}
			role="button"
			tabindex=0
			
			>{message.username}</span>
			<span class="date">sent at {message.timestamp.toLocaleString()}</span>
		</p>
		{#if showMiniProfile && miniProfileMessageIndex === index}
		<div style="position: relative;">
			<MiniProfile userID={miniProfileUserID}></MiniProfile>
		</div>
		{/if}
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
