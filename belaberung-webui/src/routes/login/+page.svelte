<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { login } from 'belaberung-client-libs/auth';
	import Alert from '$lib/components/ErrorNotification.svelte';
	import { resetRoomList } from '$lib/lastRoom';
	import { Container, Row, Col, Form, Input, Label, Button } from '@sveltestrap/sveltestrap';

	let username = $state('');
	let password = $state('');
	let errorMessage = $state('');
	let showError = $state(false);

	async function handleSubmit() {
		username = username.trim();
		const [doRedirect, errorMessageRes] = await login(username, password);
		errorMessage = errorMessageRes;
		showError = !doRedirect;
		if (doRedirect) {
			resetRoomList()
			await goto(resolve('/'));
		}
	}
</script>

<Container>
	<Row>
	<Col>
	<h1>Register</h1>
	<Form onsubmit={async () => await handleSubmit()}>
		<Label for="username">Username</Label><Input
			id="username"
			required
			bind:value={username}
		/>
		<Label for="password">Password</Label><Input
			type="password"
			id="password"
			required
			bind:value={password}
		/>
		<p style="margin-top: 1em;"><Button type="submit" size="lg" color="primary">Login</Button></p>
		<p>Dont have an account? Create one <a href={resolve('/register')}>here</a>.</p>

		<Alert isVisible={showError} message={errorMessage} onDismiss={() => (showError = !showError)}
		></Alert>
	</Form>
	</Col>
	</Row>
</Container>
