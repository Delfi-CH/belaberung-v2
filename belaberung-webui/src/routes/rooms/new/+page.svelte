<script lang="ts">
	import { Form, Input, Button, Container, Row, Col, Label } from '@sveltestrap/sveltestrap';
	import { createAPI, getBackendUrl } from 'belaberung-client-libs';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	let name = $state('');
	let description = $state('');
	let isPrivate = $state(false);
	let password = $state('');

	async function handleSubmit() {
		let data;
		const api = createAPI(getBackendUrl());
		if (isPrivate && password != '') {
			data = {
				name: name,
				description: description,
				password: password
			};
			try {
				await api.post('/rooms?private=true', data);
			} catch {
				//
			}
		} else {
			data = {
				name: name,
				description: description
			};
			try {
				const room = await api.post('/rooms', data);
				//@ts-expect-error womp womp
				goto(resolve('/room/' + room.data.id));
			} catch {
				alert('An unexpected Error ocurred.');
			}
		}
	}
</script>

<Container>
	<Row>
		<Col>
			<h1>Create a room</h1>
			<Form on:submit={async () => await handleSubmit()}>
				<Label>Name</Label>
				<Input type="text" bind:value={name} required></Input>
				<Label>Description</Label>
				<Input type="text" bind:value={description} required></Input>
				<Label>Private</Label>
				<Input type="checkbox" bind:checked={isPrivate}></Input>
				{#if isPrivate}
					<Label>Password</Label>
					<Input type="password" bind:value={password} required></Input>
				{/if}
				<p style="margin-top: 1rem !important;">
					<Button type="submit" color="primary">Create</Button>
				</p>
			</Form>
		</Col>
	</Row>
</Container>
