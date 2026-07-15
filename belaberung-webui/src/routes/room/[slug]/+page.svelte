<script lang="ts">
	import { api } from "$lib/api/core";
    import { Container, Row, Col } from "@sveltestrap/sveltestrap";
	import { onMount } from "svelte";
    let { data } = $props();
    let id = $derived(data.post.id)
    let users = $state([])

    onMount(async ()=> {
        const tmpUsers = await api.get(`/rooms/${id}/users`)
        users = tmpUsers.data
        console.log(users)
    })
    
</script>

<Container>
    <h1>{data.post.name}</h1>
    <Row>
        <Col>
        </Col>
    </Row>
    <Row>
        <Col>
            <h2>users</h2>
            <h5>admin</h5>
            <ul>
                {#each users as user (user.id)}
                    {#if user.role === "Administrator"}
                        <li class="text-danger">{user.User.username}</li>
                    {/if}
                {/each}
            </ul>
            <h5>moderator</h5>
            <ul>
                {#each users as user (user.id)}
                    {#if user.role === "Moderator"}
                        <li class="text-warning">{user.User.username}</li>
                    {/if}
                {/each}
            </ul>
            <h5>member</h5>
            <ul>
                {#each users as user (user.id)}
                    {#if user.role === "Member"}
                        <li>{user.User.username}</li>
                    {/if}
                {/each}
            </ul>
        </Col>
    </Row>
</Container>
