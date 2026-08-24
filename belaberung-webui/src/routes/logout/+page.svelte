<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { logout } from 'belaberung-client-libs/auth';
	import { resetRoomList } from '$lib/lastRoom';

	let errorMessage = $state('');

	onMount(async () => {
		let ok = await logout();
		if (ok) {
			resetRoomList();
			goto(resolve('/login'));
		} else {
			errorMessage = 'server error';
		}
	});
</script>

<p>Logging you out...</p>
<p>{errorMessage}</p>
