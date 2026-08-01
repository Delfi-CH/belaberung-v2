<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { getUserID, getUsername } from '$lib/api/auth';
	import { api } from '$lib/api/core';
	import ErrorNotification from '$lib/components/ErrorNotification.svelte';
	import { Container, Row, Col, Button, Input, Form, Label } from '@sveltestrap/sveltestrap';
	import { onMount } from 'svelte';

	let user = $state({});
	let newUsername = $state('');
	let oldPassword = $state('');
	let newPassword = $state('');
	let newBiography = $state('');
	let showerr = $state(false);
	let errmsg = $state('');
	let newBioLength = $derived(newBiography.length);
	let newPronouns = $state('');
	let newProLength = $derived(newPronouns.length);
	const uid = getUserID();

	onMount(async () => {
		const res = await api.get('/users/' + uid);
		user = res.data;
		newBiography = user.biography;
		newPronouns = user.pronouns;
	});

	async function saveUsername() {
		try {
			await api.patch('/users/' + uid, {
				type: 'username',
				username: newUsername
			});
			goto(resolve('/logout'));
		} catch (err) {
			showerr = true;
			errmsg = err;
		}
	}

	async function savePassword() {
		try {
			await api.patch('/users/' + uid, {
				type: 'password',
				oldPassword: oldPassword,
				newPassword: newPassword
			});
			goto(resolve('/logout'));
		} catch (err) {
			showerr = true;
			errmsg = err;
		}
	}

	async function saveBiography() {
		try {
			await api.patch('/users/' + uid, {
				type: 'biography',
				biography: newBiography
			});
		} catch (err) {
			showerr = true;
			errmsg = err;
		}
	}

	async function savePronouns() {
		try {
			await api.patch('/users/' + uid, {
				type: 'pronouns',
				pronouns: newPronouns
			});
		} catch (err) {
			showerr = true;
			errmsg = err;
		}
	}
</script>

<Container>
	<Row>
		<h1>Your Profile</h1>
		<Col>
			<p>Username: {getUsername()}</p>
			<Form>
				<Label>Username</Label>
				<Input type="text" bind:value={newUsername} required></Input>
				<Button type="submit" onclick={async () => await saveUsername()}>Change Username</Button>
			</Form>
			<Form>
				<Label>Old Password</Label><Input type="password" bind:value={oldPassword} required></Input>
				<Label>New Password</Label><Input type="password" bind:value={newPassword} required></Input>
				<Button type="submit" onclick={async () => await savePassword()}>Change Password</Button>
			</Form>

			<Form>
				<p>Biography:</p>
				<Input
					type="textarea"
					bind:value={newBiography}
					placeholder="No biograpgy..."
					max="1024"
					rows="5"
					required
				></Input>
				<p><span class={newBioLength > 1024 ? 'text-danger' : ''}>{newBioLength}</span>/{1024}</p>
				<Button
					onclick={async () => await saveBiography()}
					type="submit"
					disabled={newBioLength > 64}>Save biography</Button
				>
			</Form>
			<Form>
				<p>Pronouns:</p>
				<Input
					type="textarea"
					bind:value={newPronouns}
					placeholder="No pronouns..."
					max="64"
					rows="1"
					required
				></Input>
				<p><span class={newProLength > 64 ? 'text-danger' : ''}>{newProLength}</span>/{64}</p>
				<Button onclick={async () => await savePronouns()} disabled={newProLength > 64}
					>Save Pronouns</Button
				>
			</Form>
		</Col>
		<ErrorNotification
			isVisible={showerr}
			message={errmsg}
			onDismiss={() => {
				showerr = false;
			}}
		></ErrorNotification>
	</Row>
</Container>
