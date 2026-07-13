<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { register } from '$lib/api/auth';
	import Alert from '$lib/components/ErrorNotification.svelte';

	let username = $state('');
	let password = $state('');
	let errorMessage = $state('');
    let showError = $state(false)

	async function handleSubmit() {
		const [doRedirect, errorMessageRes] = await register(username, password);
		errorMessage = errorMessageRes;
        showError = !doRedirect
		if (doRedirect) {
			await goto(resolve('/'));
		}
	}
</script>

<div>
	<h1>Register</h1>
	<form onsubmit={async () => await handleSubmit()}>
			<label for="username">Username</label><input
				type="username"
				id="username"
				required
				bind:value={username}
			/>
			<label for="password">Password</label><input
				type="password"
				id="password"
				required
				bind:value={password}
			/>
			<button type="submit">Sign Up</button>
			<p>Alreay have an account? Log in <a href={resolve("/login")}>here</a></p>

            <Alert isVisible={showError} message={errorMessage} onDismiss={()=> showError = !showError}></Alert>
	</form>
</div>