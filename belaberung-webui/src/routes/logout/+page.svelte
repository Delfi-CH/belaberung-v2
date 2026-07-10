<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { logout } from '$lib/api/auth';

	let errorMessage = $state('');

	onMount(async () => {
		let ok = await logout();
		if (ok) {
			goto(resolve('/login'));
		} else {
			errorMessage = 'server error';
		}
	});
</script>

<p>Logging you out...</p>
<p>{errorMessage}</p>
