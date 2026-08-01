<script lang="ts">
	import { getUserID, getUsername } from "$lib/api/auth.js";
	import { api, createWebsocket, loadInitialMessages, sendMessage, streamMessages } from "$lib/api/core";
	import { MessageAttachment, MessageAttachmentType, Message } from "$lib/api/message.js";
	import MessageViewer from "$lib/components/MessageViewer.svelte";
    import { Container, Row, Button, Form, Input } from "@sveltestrap/sveltestrap";
	import { onMount } from "svelte";
    let { data } = $props();
    let messageContent = $state("")
    let id = $derived(data.post.id)
    let users = $state([])
    let messages: Message[] = $state([])
    let ws: WebSocket

    onMount(async ()=> {
        const tmpUsers = await api.get(`/rooms/${id}/users`)
        users = tmpUsers.data
        const tmpMessages = await loadInitialMessages(id)
        messages = [...messages, ...tmpMessages]
        messages.sort((a, b) => a.timestamp - b.timestamp)
    })

    onMount(()=>{
        ws = createWebsocket()
        streamMessages(ws, (message)=>{
            messages = [...messages, message]
            messages.sort((a, b) => a.timestamp - b.timestamp)
        })

    })
    
</script>

<Container>
    <h1>{data.post.name}</h1>
    <Row>
        <div class="messages">
            <MessageViewer messages={messages}></MessageViewer>
            <Form class="d-flex gap-2 align-items-center">
                <Input type="text" bind:value={messageContent} placeholder="Type your message here"></Input>
                <Button onclick={()=>sendMessage(ws, messageContent, getUsername(), Number(getUserID()), id, new MessageAttachment(MessageAttachmentType.None, null))} type="submit">test</Button>
            </Form>
            
        </div>
        <div class="users">
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
