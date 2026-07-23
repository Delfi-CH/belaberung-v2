<script lang="ts">
	import { getJoinedRooms, getPublicRooms, joinRoom } from '$lib/api/core';
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
		Button
	} from '@sveltestrap/sveltestrap';
	import { goto } from "$app/navigation";
    import { resolve } from "$app/paths"

	let rooms = $state([]);
	let joinedRooms = $state([]);

	onMount(async () => {
		const tmpRooms = await getPublicRooms();
		rooms = tmpRooms;
		const tmpJoinedRooms = await getJoinedRooms();
		joinedRooms = tmpJoinedRooms;
	});
</script>

<Container>
	<h1>rooms</h1>
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
						{room.description}
					</CardBody>
					<CardFooter>
						<Button onclick={()=>{
							goto(resolve("/room/" + String(room.id)))
						}}>Goto</Button>
					</CardFooter>
				</Card>
			{/each}
			{:else}
				<p>You havent joined any rooms yet...	</p>
			{/if}
		</Col>
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
						<Button onclick={async ()=>{
							const err = await joinRoom(room.id)

							if (err === "joined") {
								goto(resolve("/room/" + String(room.id)))
							} else {
								alert("An unexpected error ocurred!")
							}
							
						}}>Join</Button>
					</CardFooter>
				</Card>
			{/each}
			{:else}
				<p>No rooms found...</p>
			{/if}
		</Col>
	</Row>
</Container>
