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

<div class="mx-auto w-full max-w-md space-y-4">
	<h1 class="h1">Register</h1>
	<form onsubmit={async () => await handleSubmit()}>
		<fieldset class="space-y-4">
			<label for="username" class="label">Username</label><input
				type="username"
				id="username"
				required
				bind:value={username}
				class="input"
			/>
			<label for="password" class="label">Password</label><input
				type="password"
				id="password"
				required
				bind:value={password}
				class="input"
			/>
		</fieldset>
		<fieldset class="flex pt-4" >
			<button type="submit" class="btn preset-outlined-surface-300-700 bg-green-400">Sign Up</button>
			<p class="pl-3">Alreay have an account? Log in <a href={resolve("/login")} class="link underline">here</a></p>
		</fieldset>
        <fieldset class="pt-4">
            <Alert isVisible={showError} message={errorMessage} onDismiss={()=> showError = !showError}></Alert>
        </fieldset>
	</form>
</div>