<script lang="ts">
	import { getUserDetails } from 'belaberung-client-libs';
	import { onMount } from 'svelte';
	import { Card, CardBody, CardHeader } from '@sveltestrap/sveltestrap';

	let { userID } = $props();
	let userData = $state({});
	let biography = $state([]);

	onMount(async () => {
		const tmpUserData = await getUserDetails(userID);
		userData = tmpUserData;
		if (tmpUserData.biography !== '') {
			biography = tmpUserData.biography.split('\n');
		} else {
			biography = [];
		}
	});
</script>

<div class="on-top">
	<Card>
		<CardHeader>
			<p>{userData.username}@{userData.domain}</p>
			<p class="pronouns">{userData.pronouns === '' ? 'No pronouns set' : userData.pronouns}</p>
		</CardHeader>
		<CardBody>
			{#if biography.length >= 1}
				{#each biography as line, index (index)}
					<span>{line}</span>
					<br />
				{/each}
			{:else}
				<p>No biography set</p>
			{/if}
		</CardBody>
	</Card>
</div>

<style>
	.on-top {
		top: 0.01rem;
		position: absolute;
		z-index: 1000;
	}
	.pronouns {
		font-size: 85%;
	}
</style>
