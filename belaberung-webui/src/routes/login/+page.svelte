<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { login } from '$lib/api/auth';
	import Alert from '$lib/components/ErrorNotification.svelte';

	let username = $state('');
	let password = $state('');
	let errorMessage = $state('');
    let showError = $state(false)

	async function handleSubmit() {
		const [doRedirect, errorMessageRes] = await login(username, password);
		errorMessage = errorMessageRes;
        showError = !doRedirect
		if (doRedirect) {
			await goto(resolve('/'));
		}
	}
</script>

<div>
	<h1>Login</h1>
	<form on:submit={async () => await handleSubmit()}>
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

			<button type="submit">Login</button>
            <p>Dont have an account? Create one <a href={resolve("/register")}>here</a></p>

            <Alert isVisible={showError} message={errorMessage} onDismiss={()=> showError = !showError}></Alert>
	</form>
</div>