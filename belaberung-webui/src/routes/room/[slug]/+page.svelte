<script lang="ts">
	import { getUserID, getUsername } from "$lib/api/auth.js";
	import { api, createWebsocket, loadInitialMessages, sendMessage, streamMessages } from "$lib/api/core";
	import { MessageAttachment, MessageAttachmentType, Message } from "$lib/api/message.js";
    import { Container, Row, Col, Button } from "@sveltestrap/sveltestrap";
	import { onMount } from "svelte";
    let { data } = $props();
    let id = $derived(data.post.id)
    let users = $state([])
    let messages: Message[] = $state([])
    let ws: WebSocket

    onMount(async ()=> {
        const tmpUsers = await api.get(`/rooms/${id}/users`)
        users = tmpUsers.data
        const tmpMessages = await loadInitialMessages(id)
        messages = [...messages, ...tmpMessages]
    })

    onMount(()=>{
        ws = createWebsocket()
        streamMessages(ws, (message)=>{
            messages = [...messages, message]
        })

    })
    
</script>

<Container>
    <h1>{data.post.name}</h1>
    <Row>
        <Col>
            {#each messages as message, index (index)}
                <p>{message.username}</p>
                <p>{message.content}</p>
            {/each}
            <Button onclick={()=>sendMessage(ws, "test", getUsername(), Number(getUserID()), id, new MessageAttachment(MessageAttachmentType.None, null))}>test</Button>
        </Col>
    </Row>
    <Row>
        <Col>
            <h2>Users</h2>
            <h5>Administrators</h5>
            <ul>
                {#each users as user, index (index)}
                    {#if user.role === "Administrator"}
                        <li class="text-danger">{user.User.username}</li>
                    {/if}
                {/each}
            </ul>
            <h5>Moderators</h5>
            <ul>
                {#each users as user, index (index)}
                    {#if user.role === "Moderator"}
                        <li class="text-warning">{user.User.username}</li>
                    {/if}
                {/each}
            </ul>
            <h5>Members</h5>
            <ul>
                {#each users as user, index (index)}
                    {#if user.role === "Member"}
                        <li>{user.User.username}</li>
                    {/if}
                {/each}
            </ul>
        </Col>
    </Row>
</Container>
