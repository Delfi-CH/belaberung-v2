<script lang="ts">
	import { getUserID, getUsername } from 'belaberung-client-libs/auth';
	import {
		createAPI,
		getBackendUrl,
		createWebsocket,
		loadInitialMessages,
		sendMessage,
		streamMessages
	} from 'belaberung-client-libs';
	import {
		MessageAttachment,
		MessageAttachmentType,
		Message
	} from 'belaberung-client-libs/message';
	import MessageViewer from '$lib/components/MessageViewer.svelte';
	import { Container, Row, Button, Form, Input } from '@sveltestrap/sveltestrap';
	import { onMount } from 'svelte';
	import MiniProfile from '$lib/components/MiniProfile.svelte';
	import { appendToRoomList } from '$lib/lastRoom';

	let { data } = $props();
	let password = $derived(data.password);
	let messageContent = $state('');
	let id = $derived(data.post.id);
	let users = $state([]);
	let messages: Message[] = $state([]);
	let ws: WebSocket;
	let showMiniProfile = $state(false);
	let miniProfileUserID = $state(0);

	onMount(async () => {
		const api = createAPI(getBackendUrl());
		const tmpUsers = await api.get(`/rooms/${id}/users?password=${password}`);
		users = tmpUsers.data;
		const tmpMessages = await loadInitialMessages(id, password ?? '');
		messages = [...messages, ...tmpMessages];
		messages.sort((a, b) => a.timestamp - b.timestamp);
	});

	onMount(() => {
		appendToRoomList(data.post);
		ws = createWebsocket();
		streamMessages(ws, (message) => {
			messages = [...messages, message];
			messages.sort((a, b) => a.timestamp - b.timestamp);
		});
	});
</script>

<Container>
	<h1>{data.post.name}</h1>
	<Row>
		<div class="messages">
			<MessageViewer {messages}></MessageViewer>
			<Form class="d-flex gap-2 align-items-center">
				<Input type="text" bind:value={messageContent} placeholder="Type your message here"></Input>
				<Button
					onclick={(e) => {
						e.preventDefault();
						sendMessage(
							ws,
							messageContent,
							getUsername(),
							Number(getUserID()),
							id,
							new MessageAttachment(MessageAttachmentType.None, null),
							password ?? ''
						);
						messageContent = '';
					}}
					type="submit">test</Button
				>
			</Form>
		</div>
		<div class="users">
			<h2>Users</h2>
			<h5>Administrators</h5>
			<ul>
				{#each users as user, index (index)}
					{#if user.role === 'Administrator'}
						<li
							onclick={() => {
								showMiniProfile = !showMiniProfile;
								miniProfileUserID = user.User.id;
							}}
							onkeydown={(e) => {}}
							role="button"
							tabindex="0"
							class="text-danger"
						>
							{user.User.username}
						</li>
						{#if showMiniProfile && user.role === 'Administrator' && user.User.id === miniProfileUserID}
							<div style="position: relative;">
								<MiniProfile userID={miniProfileUserID}></MiniProfile>
							</div>
						{/if}
					{/if}
				{/each}
			</ul>
			<h5>Moderators</h5>
			<ul>
				{#each users as user, index (index)}
					{#if user.role === 'Moderator'}
						<li
							onclick={() => {
								showMiniProfile = !showMiniProfile;
								miniProfileUserID = user.User.id;
							}}
							onkeydown={(e) => {}}
							role="button"
							tabindex="0"
							class="text-warning"
						>
							{user.User.username}
						</li>
						{#if showMiniProfile && user.role === 'Moderator' && user.User.id === miniProfileUserID}
							<div style="position: relative;">
								<MiniProfile userID={miniProfileUserID}></MiniProfile>
							</div>
						{/if}
					{/if}
				{/each}
			</ul>
			<h5>Members</h5>
			<ul>
				{#each users as user, index (index)}
					{#if user.role === 'Member'}
						<li
							onclick={() => {
								showMiniProfile = !showMiniProfile;
								miniProfileUserID = user.User.id;
							}}
							onkeydown={(e) => {}}
							role="button"
							tabindex="0"
						>
							{user.User.username}
						</li>
						{#if showMiniProfile && user.role === 'Member' && user.User.id === miniProfileUserID}
							<div style="position: relative;">
								<MiniProfile userID={miniProfileUserID}></MiniProfile>
							</div>
						{/if}
					{/if}
				{/each}
			</ul>
		</div>
	</Row>
</Container>

<style>
	.messages {
		width: 100%;
		@media (width >= 768px) {
			width: 70%;
		}
	}

	.users {
		display: none;
		@media (width >= 768px) {
			display: block;
			width: 30%;
		}
	}
</style>
