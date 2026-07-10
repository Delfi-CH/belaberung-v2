<script>
	import { getUserID, getUsername } from '$lib/api/auth';
	import { onMount } from 'svelte';

	let username = $state('');
	let userID = $state('');

	onMount(() => {
		username = getUsername();
		userID = getUserID();
	});
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>
<p>username: {username}</p>
<p>id: {userID}</p>
