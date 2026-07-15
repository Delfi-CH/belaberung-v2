<script lang="ts">
    import { Form, Input, Button, Container, Row, Col, Label } from "@sveltestrap/sveltestrap"
    import { api } from "$lib/api/core"
	import { goto } from "$app/navigation";
    import { resolve } from "$app/paths"
    let name = $state("")
    let description = $state("")
    let isPrivate = $state(false)
    let password = $state("")

    async function handleSubmit() {
        let data;
        if (isPrivate && password != "") {
            data = {
                name: name,
                description: description,
                password: password
            }
            try {
                await api.post("/rooms?private=true", data)
            } catch {
                //
            }
        } else {
            data = {
                name: name,
                description: description,
            }
            try {
                await api.post("/rooms", data)
                goto(resolve("/rooms"))
            } catch {
                //
            }
        }
    }
</script>

<h1>new room</h1>

    <Container>
        <Row>
            <Col>
            <Form on:submit={async()=> await handleSubmit()}>
                <Label>Name</Label>
                <Input type="text" bind:value={name}></Input>
                <Label>Description</Label>
                <Input type="text" bind:value={description}></Input>
                <Label>Private</Label>
                <Input type="checkbox" bind:checked={isPrivate}></Input>
                {#if isPrivate}
                    <Label>Password</Label>
                    <Input type="password" bind:value={password}></Input>
                {/if}
                <Button type="submit">Create</Button>
            </Form>
            
            </Col>
        </Row>
    </Container>
