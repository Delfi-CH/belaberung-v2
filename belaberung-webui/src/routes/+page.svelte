<script lang="ts">
	import { onMount } from "svelte";
    import { getRoomList } from "$lib/lastRoom";
	import { getUsername } from "belaberung-client-libs/auth";
    import { getJoinedRooms } from "belaberung-client-libs";
    import { Container, Row, Col, Card,
		CardHeader,
		CardTitle,
		CardBody,
		CardFooter,
		Button } from "@sveltestrap/sveltestrap";
    import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';

    let lastRooms = $state([])
    let joinedRooms = $state([])
    let username = $state("")

    onMount(()=>{
        lastRooms = getRoomList()
        username = getUsername()
    })

    onMount(async()=>{
        const tmpJoinedRooms = await getJoinedRooms();
		joinedRooms = tmpJoinedRooms;
    })
</script>

<Container>
    <Row>
        <Col>
            <h1>Hello, {username}</h1>
            <h2>Joined rooms</h2>
            <div>
            {#each joinedRooms as room, index (index)}
					<span class="roomList"><Button color="success" href={resolve("/room/"+room.id)}>{room.name}</Button></span>     
				{/each}
                </div>
            <h2>Options</h2>
            <Button href={resolve("/rooms/new")} color="info" size="lg">Create a Room</Button>
            <Button href={resolve("/profile")} color="info" size="lg">Edit your Profile</Button>
            <Button href={resolve("/logout")} color="danger" size="lg">Log out</Button>
        </Col>
        <Col>
            <h2>Previously visited rooms:</h2>
            {#each lastRooms as room, index (index)}
					<Card class="m-1">
						<CardHeader>
							<CardTitle>{room.name}</CardTitle>
						</CardHeader>
						<CardBody>
							{room.description}
						</CardBody>
						<CardFooter>
							<Button
								onclick={() => {
									goto(resolve('/room/' + String(room.id)));
								}}>Goto</Button
							>
						</CardFooter>
					</Card>
				{/each}
        </Col>
    </Row>
</Container>

<style>
    .roomList {
       margin: 1rem !important ;
    }
</style>
