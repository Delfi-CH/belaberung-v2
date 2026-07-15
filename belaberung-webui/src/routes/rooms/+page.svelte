<script lang="ts">
	import { getJoinedRooms, getPublicRooms } from '$lib/api/core';
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
			{#if joinedRooms.length <= 1}
				{#each joinedRooms as room, index (index)}
					{#if JSON.stringify(room) !== JSON.stringify({})}
						<Card>
							<CardHeader>
								<CardTitle>{room.name}</CardTitle>
							</CardHeader>
							<CardBody>
								{room.description}
							</CardBody>
							<CardFooter>
								<Button>Goto</Button>
							</CardFooter>
						</Card>
					{/if}
				{/each}
			{:else}
				<p>No rooms found...</p>
			{/if}
		</Col>
	</Row>
	<Row>
		<h2>Public Rooms</h2>
		<Col>
			{#if rooms.length <= 1}
				{#each rooms as room, index (index)}
					{#if JSON.stringify(room) !== JSON.stringify({})}
						<Card>
							<CardHeader>
								<CardTitle>{room.name}</CardTitle>
							</CardHeader>
							<CardBody>
								{room.description}
							</CardBody>
							<CardFooter>
								<Button>Join</Button>
							</CardFooter>
						</Card>
					{/if}
				{/each}
			{:else}
				<p>No rooms found...</p>
			{/if}
		</Col>
	</Row>
</Container>
