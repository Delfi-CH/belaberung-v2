<script lang="ts">
	import { getJoinedRooms, getPublicRooms, joinRoom } from 'belaberung-client-libs';
	import { onMount } from 'svelte';
	import {
		Container,
		Row,
		Col,
		Card,
		CardHeader,
		CardTitle,
		CardBody,
		CardFooter,
		Button,
		Input,
		Form,
		Label
	} from '@sveltestrap/sveltestrap';
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';

	let rooms = $state([]);
	let joinedRooms = $state([]);
	let roomID = $state();
	let roompass = $state('');

	onMount(async () => {
		const tmpRooms = await getPublicRooms();
		rooms = tmpRooms;
		const tmpJoinedRooms = await getJoinedRooms();
		joinedRooms = tmpJoinedRooms;
	});
</script>

<Container>
	<h1>Rooms</h1>
	<Row>
		<h2>Joined</h2>
		<Col>
			{#if joinedRooms.length >= 1}
				{#each joinedRooms as room, index (index)}
					<Card class="m-1">
						<CardHeader>
							<CardTitle>{room.name}</CardTitle>
						</CardHeader>
						<CardBody>
							<p>{room.description}</p>
						</CardBody>
						<CardFooter>
							<p>
								<Button
									onclick={() => {
										if (room.private) {
											goto(resolve('/room/' + String(room.id) + `?password=${roompass}`));
										} else {
											goto(resolve('/room/' + String(room.id)));
										}
									}}>Goto</Button
								>
							</p>
							{#if room.private}
								Password <Input type="password" label="Password" bind:value={roompass} id="password"
								></Input>
							{/if}
						</CardFooter>
					</Card>
				{/each}
			{:else}
				<p>You havent joined any rooms yet...</p>
			{/if}
		</Col>
	</Row>
	<Row>
		<h2>Join a private room</h2>
		<Form
			onsubmit={async () => {
				const err = await joinRoom(roomID, roompass);

				if (err === 'joined') {
					goto(resolve('/room/' + String(roomID) + `/?password=${roompass}`));
				} else {
					alert('An unexpected error ocurred!');
				}
			}}
		>
			<Label>RoomID</Label>
			<Input type="number" bind:value={roomID}></Input>
			<Label>Password</Label>
			<Input type="password" bind:value={roompass}></Input>
			<Button type="submit">Join</Button>
		</Form>
	</Row>
	<Row>
		<h2>Public Rooms</h2>
		<Col>
			{#if rooms.length >= 1}
				{#each rooms as room, index (index)}
					<Card class="m-1">
						<CardHeader>
							<CardTitle>{room.name}</CardTitle>
						</CardHeader>
						<CardBody>
							{room.description}
						</CardBody>
						<CardFooter>
							<Button
								onclick={async () => {
									const err = await joinRoom(room.id);

									if (err === 'joined') {
										goto(resolve('/room/' + String(room.id)));
									} else {
										alert('An unexpected error ocurred!');
									}
								}}>Join</Button
							>
						</CardFooter>
					</Card>
				{/each}
			{:else}
				<p>No rooms found...</p>
			{/if}
		</Col>
	</Row>
</Container>
